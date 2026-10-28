package storage

import "time"

type Bucket struct {
	Name            string
	AlloweMIMETypes []string
	CreatedAt       time.Time
}
