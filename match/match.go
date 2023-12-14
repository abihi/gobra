package match

import (
	"context"
	"gobra/domain"

	"github.com/google/uuid"
)

type Service struct{}

func (s Service) Get(ctx context.Context, id uuid.UUID) (domain.Match, error) {
	return domain.Match{}, nil
}
