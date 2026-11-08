package storage

import (
	"fmt"
	"os"
	"uuid"
)

const (
	FilePermissions = 0600
)

type FileSystem struct {
	rootPath string
}

func NewFileSystem(rootPath string) *FileSystem {
	return &FileSystem{
		rootPath: rootPath,
	}
}

func getObjectPath(id uuid.UUID) string {
	idStr := id.String()
	return fmt.Sprintf("%s/%s.blob", idStr[0:2], idStr)
}

func getBucketPath(rootPath string, id uuid.UUID) string {
	idStr := id.String()
	return fmt.Sprintf("%s/%s/%s/", rootPath, idStr[0:2], idStr)
}

func createBucketPath(rootPath string, id uuid.UUID) (*string, error) {
	path := getBucketPath(rootPath, id)

	err := os.Mkdir(path, FilePermissions)
	if err != nil {
		return nil, fmt.Errorf("[FileSystem]: failed to create path %s. Err: %v", path, err)
	}

	return &path, nil
}

func createObjectPath(rootPath string, bucket, id uuid.UUID) (*string, error) {
	bucketPath, err := createBucketPath(rootPath, bucket)
	if err != nil {
		return nil, err
	}

	idStr := id.String()
	fullPath := fmt.Sprintf("%s%s", bucketPath, idStr[0:2])

	err = os.Mkdir(fullPath, FilePermissions)
	if err != nil {
		return nil, fmt.Errorf("[FileSystem]: failed to create path %s. Err: %v", fullPath, err)
	}

	return &fullPath, nil
}

func (fileSystem *FileSystem) UpsertObject(meta Object, data []byte) error {
	path, err := createObjectPath(fileSystem.rootPath, meta.ID, meta.ID)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s.blob", path, meta.ID), data, FilePermissions)
	if err != nil {
		fmt.Errorf("[FileSystem]: failed to upsert object %s. Err: %v", meta.Name, err)
	}

	return nil
}

func (fileSystem *FileSystem) DeleteObject(meta Object) error {
	err := os.Remove(getBucketPath(fileSystem.rootPath, meta.ID) + getObjectPath(meta.ID))

	if err != nil {
		fmt.Errorf("[FileSystem]: failed to delete object %s. Err: %v", meta.Name, err)
	}

	return nil
}
