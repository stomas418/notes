package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/models"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	db *sql.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func (h *BaseHandler) getNoteById(id string) (models.Note, error) {
	var note models.Note
	if row, err := h.db.Query("SELECT * FROM notes WHERE note_id = ?", id); err != nil {
		return models.Note{}, err
	} else {
		for row.Next() {
			if err := row.Scan(&note.Id, &note.Title, &note.Content, &note.Author, &note.LastModified, &note.CreationDate); err != nil {
				return models.Note{}, errors.New("database error")
			}
		}
	}
	return note, nil
}

func (h *BaseHandler) getNotesByAuthor(author string) ([]models.Note, error) {
	var notes []models.Note
	if rows, err := h.db.Query("SELECT * FROM notes WHERE author = ?", author); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var note models.Note
			scanError := rows.Scan(&note.Id, &note.Title, &note.Content, &note.Author, &note.LastModified, &note.CreationDate)
			if scanError != nil {
				return nil, scanError
			}
			notes = append(notes, note)
		}
	}
	return notes, nil
}

func (h *BaseHandler) GetNotes(c *gin.Context) {
	author := c.Param("user")
	if notes, err := h.getNotesByAuthor(author); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Couldnt get notes"})
	} else {
		c.IndentedJSON(200, notes)
	}
}

func (h *BaseHandler) GetNoteById(c *gin.Context) {
	id := c.Param("id")
	note, err := h.getNoteById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Note not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, note)
}

func (h *BaseHandler) editNote(id string, newNote models.Note) error {
	_, updateErr := h.db.Exec("UPDATE notes SET title = ?, author = ?, content = ?, lastModified = ? WHERE note_id = ?", newNote.Title, newNote.Author, newNote.Content, newNote.LastModified, id)
	return updateErr
}

func (h *BaseHandler) deleteNote(id string) error {
	_, deleteErr := h.db.Exec("DELETE FROM notes WHERE note_id = ?", id)
	return deleteErr
}

func (h *BaseHandler) EditNote(c *gin.Context) {
	id := c.Param("id")
	var newNote models.Note
	err := json.NewDecoder(c.Request.Body).Decode(&newNote)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	updateErr := h.editNote(id, newNote)
	if updateErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}
	c.IndentedJSON(http.StatusOK, newNote)
}

func (h *BaseHandler) CreateNote(c *gin.Context) {
	var newNote models.Note
	err := json.NewDecoder(c.Request.Body).Decode(&newNote)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}
	_, insertErr := h.db.Exec("INSERT INTO notes VALUES(note_id, ?, ?, ?, ?, ?)", newNote.Title, newNote.Content, newNote.Author, time.Now(), time.Now())
	if insertErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
	}
	c.IndentedJSON(http.StatusCreated, newNote)
}

func (h *BaseHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	deletedNote, err := h.getNoteById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}
	if deleteErr := h.deleteNote(id); deleteErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database error"})
		return
	}
	c.IndentedJSON(http.StatusOK, deletedNote)
}
