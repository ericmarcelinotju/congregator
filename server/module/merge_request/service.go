package mergerequest

import (
	"context"

	"github.com/ericmarcelinotju/congregator/dto"
	"github.com/ericmarcelinotju/congregator/dto/concern"
)

// Service defines user service behavior.
type Service interface {
	Get(context.Context, *dto.MergeRequestFilter) (*concern.PaginationResp[dto.MergeRequestDto], error)
}

type service struct {
	repo Repository
}

// NewService creates a new service struct
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (svc *service) Get(ctx context.Context, filter *dto.MergeRequestFilter) (*concern.PaginationResp[dto.MergeRequestDto], error) {
	return svc.repo.Get(ctx, filter)
}
