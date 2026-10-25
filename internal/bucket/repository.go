package bucket

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v4"
)

const (
	prefix = "b:"
)

type Repository struct {
	DB *badger.DB
}

func NewRepository(db *badger.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

type Bucket struct {
	Name            string
	AlloweMIMETypes []string
	CreatedAt       string
	Files           []string
}

func (r *Repository) GetBucket(bucketName string) (Bucket, error) {
	var bucket Bucket

	err := r.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(prefix + bucketName))
		if err != nil {
			return err
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
		return Bucket{}, err
	}

	return bucket, nil
}

func (r *Repository) InsertBucket(bucket Bucket) error {
	value, err := json.Marshal(bucket)
	if err != nil {
		return err
	}

	err = r.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(prefix+bucket.Name), value)
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
