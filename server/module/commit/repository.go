package commit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-playground/form/v4"
	"github.com/gofiber/fiber/v2"

	"github.com/ericmarcelinotju/congregator/dto"
	"github.com/ericmarcelinotju/congregator/dto/concern"
)

// Repository provides an abstraction on top of the user data source
type Repository interface {
	Get(context.Context, *dto.CommitFilter) (*concern.PaginationResp[dto.CommitDto], error)
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

func (s *repository) Get(ctx context.Context, filter *dto.CommitFilter) (*concern.PaginationResp[dto.CommitDto], error) {
	urlValues, err := s.encoder.Encode(filter)
	if err != nil {
		return nil, err
	}

	agent := s.client.Get(fmt.Sprintf("/projects/%d/repository/commits?%s", filter.ProjectID, urlValues.Encode()))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var res []dto.CommitDto
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &concern.PaginationResp[dto.CommitDto]{
		Data:  res,
		Total: int64(len(res)),
	}, nil
}
