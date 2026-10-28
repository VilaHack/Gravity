package storage

import (
	"encoding/json"
	"errors"

	"github.com/dgraph-io/badger/v4"
)

var (
	ErrBucketNotFound      = errors.New("bucket not found")
	ErrObjectNotFound      = errors.New("object not found")
	ErrObjectAlreadyExists = errors.New("object already exists")
)

type Repository struct {
	DB *badger.DB
}

func NewRepository(db *badger.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetBucket(bucketName string) (*Bucket, error) {
	var bucket Bucket

	err := r.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(bucketKey(bucketName))
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

func (r *Repository) InsertBucket(b Bucket) error {
	value, err := json.Marshal(b)
	if err != nil {
		return err
	}

	err = r.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(bucketKey(b.Name)), value)
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

func (r *Repository) InsertObject(o Object) error {
	value, err := json.Marshal(o)
	if err != nil {
		return err
	}

	err = r.DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(objectKey(o.BucketName, o.RelativeKey)))
		if err == nil {
			return ErrObjectAlreadyExists
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = r.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(objectKey(o.BucketName, o.RelativeKey)), value)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
