package storage

import "uuid"

const (
	bucketPrefix = "b:"
	objectPrefix = "o:"
)

func bucketKey(name string) []byte {
	return []byte(bucketPrefix + name)
}

func objectKey(bucket string, id uuid.UUID) []byte {
	return []byte(objectPrefix + bucket + ":" + id.String())
}
