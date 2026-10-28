package storage

import "time"

type Object struct {
	BucketName   string
	Name         string
	RelativeKey  string
	RelativePath string
	CreatedAt    time.Time
}
