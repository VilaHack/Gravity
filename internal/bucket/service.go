// Package bucket contains bucket service and repository definition and implementation
package bucket

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) ListFiles(bucketName string) ([]string, error) {
	bucket, err := s.Repository.GetBucket(bucketName)
	if err != nil {
		return nil, err
	}

	return append([]string{}, bucket.Files...), nil
}

func (s *Service) UpsertBucket(bucket Bucket)
