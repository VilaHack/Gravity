package file

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v4"
)

const (
	Prefix = "f:"
)

type Repository struct {
	DB *badger.DB
}

type File struct {
	Name         string
	RelativePath string
	CreatedAt    string
}

func NewRepository(db *badger.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetFile(bucketName string, fileName string) (File, error) {
	var file File

	err := r.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(Prefix + bucketName + ":" + fileName))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, &file)
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return File{}, nil
	}

	return file, nil
}

func (r *Repository) InsertFile(bucketName string, file File) error {
	data, err := json.Marshal(file)
	if err != nil {
		return err
	}

	err = r.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(Prefix+bucketName+":"+file.Name), data)
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
