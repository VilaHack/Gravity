package storage

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

var (
	ErrBucketNotFound      = errors.New("bucket not found")
	ErrBucketAlreadyExists = errors.New("bucket already exists")
	ErrObjectNotFound      = errors.New("object not found")
	ErrObjectAlreadyExists = errors.New("object already exists")
)

type Repository struct {
	Database *badger.DB
}

func NewRepository(database *badger.DB) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repository *Repository) GetBucket(bucketName string) (*Bucket, error) {
	var bucket Bucket

	err := repository.Database.View(func(transaction *badger.Txn) error {
		item, err := transaction.Get(bucketKey(bucketName))
		if err != nil {
			return ErrBucketNotFound
		}

		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, &bucket)
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &bucket, nil
}

func (repository *Repository) InsertBucket(bucket *Bucket) error {
	value, err := json.Marshal(bucket)
	if err != nil {
		return err
	}

	err = repository.Database.Update(func(transaction *badger.Txn) error {
		err := transaction.Set(bucketKey(bucket.Name), value)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) GetObject(
	bucket, name string,
	directories []string,
) (*Object, error) {
	var object Object

	err := repository.Database.View(func(transaction *badger.Txn) error {
		item, err := transaction.Get(objectKey(bucket, name, directories))
		if err != nil {
			return ErrObjectNotFound
		}

		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, &object)
		})
		if err != nil {
			return fmt.Errorf("[Repository]: Failed to Unmarshal Item: %v", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &object, nil
}

func (repository *Repository) InsertObject(object *Object) error {
	value, err := json.Marshal(object)
	if err != nil {
		return err
	}

	objectKey := objectKey(object.BucketName, object.Name, object.Directories)

	err = repository.Database.View(func(transaction *badger.Txn) error {
		_, err := transaction.Get(objectKey)
		if err == nil {
			return ErrObjectAlreadyExists
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = repository.Database.Update(func(transaction *badger.Txn) error {
		err := transaction.Set(objectKey, value)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
