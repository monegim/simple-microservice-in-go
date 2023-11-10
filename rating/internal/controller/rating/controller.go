package rating

import (
	"context"
	"errors"
	"github.com/monegim/simple-microservice-in-go/rating/internal/repository"
	"github.com/monegim/simple-microservice-in-go/rating/pkg/model"
)

var ErrNotFound = errors.New("rating not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

type Controller struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GetAggregateRating(ctx context.Context,
	recordID model.RecordID, recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordID, recordType)
	if err != nil && err == repository.ErrNotFound {
		return 0, err
	}
	sum := float64(0)
	for _, r := range ratings {
		sum += float64(r.Value)
	}
	return sum / float64(len(ratings)), nil
}