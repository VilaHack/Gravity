package storage

import (
	"time"
	"uuid"
)

type Object struct {
	ID           uuid.UUID
	BucketName   string
	Name         string
	RelativePath string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
