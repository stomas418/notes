package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/jwt"
	"golang.org/x/crypto/bcrypt"
)

func correctPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *BaseHandler) SignIn(c *gin.Context) {
	//validate user
	username := c.Query("username")
	requestUser, _ := h.getUser("")
	user, _ := h.getUser(username)
	incorrectCredentials := gin.H{"message": "username or password are incorrect"}
	if user.Username == "" {
		c.IndentedJSON(http.StatusBadRequest, incorrectCredentials)
		return
	}
	decodeErr := json.NewDecoder(c.Request.Body).Decode(&requestUser)
	if decodeErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	if !correctPassword(requestUser.Password, user.Password) {
		c.IndentedJSON(http.StatusBadRequest, incorrectCredentials)
		return
	}
	//make token
	token, err := jwt.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "There was an error logging in, please try again"})
		return
	}
	//set token as cookie
	c.SetCookie("token", token, 60*60*24, "/", "", true, false)
	//return token
	c.IndentedJSON(http.StatusAccepted, token)
}

func (h *BaseHandler) SignUp(c *gin.Context) {
	user, _ := h.getUser(c.Query("username"))
	if user.Username != "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}
	decodeErr := json.NewDecoder(c.Request.Body).Decode(&user)
	if decodeErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	hashedPassword, hashErr := hash(user.Password)
	if hashErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	_, insertError := h.db.Exec("INSERT INTO users VALUES(?,?)", user.Username, hashedPassword)
	if insertError != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}
