package storage

import "fmt"

const (
	bucketPrefix = "b:"
	objectPrefix = "o:"
)

func bucketKey(name string) []byte {
	return []byte(bucketPrefix + name)
}

func objectKey(bucket, name string, directories []string) []byte {
	directoryKeys := ""

	for _, directory := range directories {
		directoryKeys = fmt.Sprintf("%s:%s", directoryKeys, directory)
	}
	return []byte(objectPrefix + bucket + directoryKeys + ":" + name)
}
