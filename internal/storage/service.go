// Package storage contains bucket and object related logic
package storage

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) UpsertObject(bucket, name, contentType string, data []byte) ([]byte, error) {
	return nil, nil // TODO
}
