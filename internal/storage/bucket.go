package storage

import (
	"time"
	"uuid"
)

type Bucket struct {
	ID              uuid.UUID
	Name            string
	AlloweMIMETypes []string
	CreatedAt       time.Time
}

func NewDefaultBucket(name string) *Bucket {
	return &Bucket{
		ID:              uuid.NewV4(),
		Name:            name,
		AlloweMIMETypes: []string{"*"},
		CreatedAt:       time.Now(),
	}
}
