package commit

import (
	"context"

	"github.com/ericmarcelinotju/congregator/dto"
	"github.com/ericmarcelinotju/congregator/dto/concern"
)

// Service defines user service behavior.
type Service interface {
	Get(context.Context, *dto.CommitFilter) (*concern.PaginationResp[dto.CommitDto], error)
}

type service struct {
	repo Repository
}

// NewService creates a new service struct
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (svc *service) Get(ctx context.Context, filter *dto.CommitFilter) (*concern.PaginationResp[dto.CommitDto], error) {
	return svc.repo.Get(ctx, filter)
}
