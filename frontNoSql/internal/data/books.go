package data

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int32     `json:"year,omitempty"`
	Size      Size      `json:"-"`
	Pages     int       `json:"pages"`
	Genres    []string  `json:"genres,omitempty"`
	Version   uuid.UUID `json:"version"`
}
