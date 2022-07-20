package models

import "time"

type Note struct {
	Id           int32     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	LastModified time.Time `json:"lastModified"`
	CreationDate time.Time `json:"creationDate"`
}
