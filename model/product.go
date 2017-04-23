package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	ImageURLs []string  `json:"imgUrls"`
}
