package note

import (
	"context"

	"github.com/ericmarcelinotju/congregator/dto"
	"github.com/ericmarcelinotju/congregator/dto/concern"
)

// Service defines user service behavior.
type Service interface {
	Get(context.Context, *dto.NoteFilter) (*concern.PaginationResp[dto.NoteDto], error)
}

type service struct {
	repo Repository
}

// NewService creates a new service struct
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (svc *service) Get(ctx context.Context, filter *dto.NoteFilter) (*concern.PaginationResp[dto.NoteDto], error) {
	return svc.repo.Get(ctx, filter)
}
