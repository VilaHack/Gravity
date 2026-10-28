package storage

const (
	bucketPrefix = "b:"
	objectPrefix = "o:"
)

func bucketKey(name string) []byte {
	return []byte(bucketPrefix + name)
}

func objectKey(bucket string, relativeKey string) []byte {
	return []byte(objectPrefix + bucket + ":" + relativeKey)
}
