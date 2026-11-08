package storage

import (
	"time"
	"uuid"
)

type Object struct {
	ID          uuid.UUID
	Name        string
	Directories []string
	MIMEType    string
	BucketName  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewObject(name, MIMEType, bucketName string, directories []string) *Object {
	currentTime := time.Now()
	return &Object{
		ID:          uuid.NewV4(),
		Name:        name,
		Directories: directories,
		MIMEType:    MIMEType,
		BucketName:  bucketName,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
}
