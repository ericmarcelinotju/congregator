package note

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
	Get(context.Context, *dto.NoteFilter) (*concern.PaginationResp[dto.NoteDto], error)
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

func (s *repository) Get(ctx context.Context, filter *dto.NoteFilter) (*concern.PaginationResp[dto.NoteDto], error) {

	agent := s.client.Get(fmt.Sprintf("/projects/%d/merge_requests/%d/notes", filter.ProjectID, filter.MergeRequestID))
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var res []dto.NoteDto
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &concern.PaginationResp[dto.NoteDto]{
		Data:  res,
		Total: int64(len(res)),
	}, nil
}
