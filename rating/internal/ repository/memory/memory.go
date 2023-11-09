package memory

import (
	"context"
	"github.com/monegim/simple-microservice-in-go/rating/internal/repository"
	"github.com/monegim/simple-microservice-in-go/rating/pkg/model"
)

type Repository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

func New() *Repository {
	return &Repository{map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

func (r *Repository) Get(ctx context.Context,
	recordID model.RecordID,
	recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}
	if ratings, ok := r.data[recordType][recordID]; !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[recordType][recordID], nil

}
