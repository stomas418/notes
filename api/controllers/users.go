package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *BaseHandler) getUser(username string) (*models.User, error) {
	var user models.User
	if rows, err := h.db.Query("SELECT * FROM users WHERE username = ?", username); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			scanError := rows.Scan(&user.Username, &user.Password)
			if scanError != nil {
				return nil, err
			}
		}
	}
	println(user.Username)
	return &user, nil
}

func (h *BaseHandler) GetUser(c *gin.Context) {
	username := c.Param("user")
	user, err := h.getUser(username)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *BaseHandler) EditUser(c *gin.Context) {
	user, err := h.getUser(c.Param("user"))
	if user.Username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User doesnt exist"})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}
	var editedUser models.EditUser
	json.NewDecoder(c.Request.Body).Decode(&editedUser)
	if nUser, _ := h.getUser(editedUser.Username); nUser.Username != "" && nUser.Username != user.Username {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "username already in use"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(editedUser.OldPassword)); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	if user.Username != editedUser.Username {
		_, updateNotesErr := h.db.Exec("UPDATE notes SET author = ? WHERE author = ?", editedUser.Username, user.Username)
		if updateNotesErr != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
			return
		}
		_, updateUserErr := h.db.Exec("UPDATE users SET username = ? WHERE username = ?", editedUser.Username, user.Username)
		if updateUserErr != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(editedUser.OldPassword)); err == nil {
		newPassword, hashErr := hash(editedUser.NewPassword)
		if hashErr != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		}
		_, updatePasswordErr := h.db.Exec("UPDATE users SET hashed_password = ? WHERE username = ?", newPassword, editedUser.Username)
		if updatePasswordErr != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, editedUser)
}

func (h *BaseHandler) DeleteUser(c *gin.Context) {
	user, err := h.getUser(c.Param("user"))
	if user.Username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User doesnt exist"})
		return
	}
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	if _, notesDeleteErr := h.db.Exec("DELETE FROM notes WHERE author = ?", user.Username); notesDeleteErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error, couldnt delete notes from user"})
		return

	}

	if _, userDeleteErr := h.db.Exec("DELETE FROM users WHERE username = ?", user.Username); userDeleteErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error, couldnt delete user"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func hash(password string) (string, error) {
	println([]byte(password))
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
