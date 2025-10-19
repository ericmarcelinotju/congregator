package user

import (
	"context"

	"github.com/ericmarcelinotju/congregator/dto"
)

// Service defines user service behavior.
type Service interface {
	ReadDetail(context.Context, *dto.UserFilter) (*dto.UserDto, error)
}

type service struct {
	repo Repository
}

// NewService creates a new service struct
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (svc *service) ReadDetail(ctx context.Context, filter *dto.UserFilter) (*dto.UserDto, error) {
	return svc.repo.GetDetail(ctx, filter)
}
