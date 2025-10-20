package dto

import "time"

type NoteFilter struct {
	ProjectID      *string
	MergeRequestID *string
}

type NoteDto struct {
	ID           int          `json:"id"`
	Type         string       `json:"type,omitempty"`       // e.g. "DiffNote", "DiscussionNote" (optional)
	Body         string       `json:"body"`                 // comment text (sometimes `note` in older API)
	Attachment   string       `json:"attachment,omitempty"` // if comment includes attachment
	Author       UserDto      `json:"author"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	System       bool         `json:"system"`                  // true if system-generated note (like status changes)
	NoteableID   int          `json:"noteable_id"`             // e.g., MR or issue ID this note belongs to
	NoteableType string       `json:"noteable_type"`           // "MergeRequest", "Issue", "Snippet", etc.
	Resolved     bool         `json:"resolved,omitempty"`      // true if discussion/note is resolved
	Resolvable   bool         `json:"resolvable,omitempty"`    // true if discussion/note is resolvable
	DiscussionID string       `json:"discussion_id,omitempty"` // discussion thread identifier
	Position     *PositionDto `json:"position,omitempty"`      // for diff notes, file position info
	Confidential bool         `json:"confidential,omitempty"`
	URL          string       `json:"url,omitempty"` // web URL of the comment
}

type PositionDto struct {
	BaseSHA      string `json:"base_sha"`
	StartSHA     string `json:"start_sha"`
	HeadSHA      string `json:"head_sha"`
	OldPath      string `json:"old_path"`
	NewPath      string `json:"new_path"`
	PositionType string `json:"position_type"` // "text" or "image"
	OldLine      int    `json:"old_line,omitempty"`
	NewLine      int    `json:"new_line,omitempty"`
}
