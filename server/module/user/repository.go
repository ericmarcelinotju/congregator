package user

import (
	"context"
	"encoding/json"

	"github.com/go-playground/form/v4"
	"github.com/gofiber/fiber/v2"

	"github.com/ericmarcelinotju/congregator/dto"
)

type Repository interface {
	GetDetail(context.Context, *dto.UserFilter) (*dto.UserDto, error)
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

func (s *repository) GetDetail(ctx context.Context, filter *dto.UserFilter) (*dto.UserDto, error) {
	urlValues, err := s.encoder.Encode(filter)
	if err != nil {
		return nil, err
	}
	agent := s.client.Get("/user?" + urlValues.Encode())

	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var res *dto.UserDto
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
