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

func getBucketPath(rootPath, bucket string) string {
	return fmt.Sprintf("%s/%s/", rootPath, bucket)
}

func createPath(rootPath, bucket string, id uuid.UUID) (*string, error) {
	idStr := id.String()
	path := fmt.Sprintf("%s/%s/%s/", rootPath, bucket, idStr[0:2])
	err := os.Mkdir(path, FilePermissions)
	if err != nil {
		return nil, fmt.Errorf("[FileSystem]: failed to create path %s. Err: %v", path, err)
	}

	return &path, nil
}

func (fs *FileSystem) UpsertObject(meta Object, data []byte) error {
	path, err := createPath(fs.rootPath, meta.BucketName, meta.ID)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s.blob", path, meta.ID), data, FilePermissions)
	if err != nil {
		fmt.Errorf("[FileSystem]: failed to upsert object %s. Err: %v", meta.Name, err)
	}

	return nil
}

func (fs *FileSystem) DeleteObject(meta Object) error {
	err := os.Remove(getBucketPath(fs.rootPath, meta.BucketName) + getObjectPath(meta.ID))

	if err != nil {
		fmt.Errorf("[FileSystem]: failed to delete object %s. Err: %v", meta.Name, err)
	}

	return nil
}
