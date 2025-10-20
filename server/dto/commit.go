package dto

import (
	"time"
)

type CommitFilter struct {
	ProjectID   *int
	AuthorEmail *string
	Since       *time.Time
	Until       *time.Time
}

type CommitDto struct {
	ID             string            `json:"id"`
	ShortID        string            `json:"short_id"`
	CreatedAt      time.Time         `json:"created_at"`
	ParentIDs      []string          `json:"parent_ids"`
	Title          string            `json:"title"`
	Message        string            `json:"message"`
	AuthorName     string            `json:"author_name"`
	AuthorEmail    string            `json:"author_email"`
	AuthoredDate   time.Time         `json:"authored_date"`
	CommitterName  string            `json:"committer_name"`
	CommitterEmail string            `json:"committer_email"`
	CommittedDate  time.Time         `json:"committed_date"`
	Trailers       map[string]string `json:"trailers"`
	WebURL         string            `json:"web_url"`
}
