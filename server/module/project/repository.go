package user

import (
	"context"
	"encoding/json"

	"github.com/go-playground/form/v4"
	"github.com/gofiber/fiber/v2"

	"github.com/ericmarcelinotju/congregator/dto"
	"github.com/ericmarcelinotju/congregator/dto/concern"
)

// Repository provides an abstraction on top of the user data source
type Repository interface {
	Get(context.Context, *dto.ProjectFilter) (*concern.PaginationResp[dto.ProjectDto], error)
}

type repository struct {
	client  *fiber.Client
	encoder *form.Encoder
}

func NewRepository(
	client *fiber.Client,
) Repository {
	return &repository{
		client,
		form.NewEncoder(),
	}
}

func (s *repository) Get(ctx context.Context, filter *dto.ProjectFilter) (*concern.PaginationResp[dto.ProjectDto], error) {
	urlValues, err := s.encoder.Encode(filter)
	if err != nil {
		return nil, err
	}
	agent := s.client.Get("/projects?" + urlValues.Encode())

	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var res []dto.ProjectDto
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &concern.PaginationResp[dto.ProjectDto]{
		Data:  res,
		Total: int64(len(res)),
	}, nil
}
