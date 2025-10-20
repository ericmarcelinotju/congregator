package dto

import (
	"time"
)

type MergeRequestFilter struct {
	ProjectID *int
	State     *string
}

type MergeRequestDto struct {
	ID                        int        `json:"id"`
	IID                       int        `json:"iid"`
	ProjectID                 int        `json:"project_id"`
	Title                     string     `json:"title"`
	Description               string     `json:"description"`
	State                     string     `json:"state"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
	MergedBy                  *UserDto   `json:"merged_by,omitempty"`
	MergedAt                  *time.Time `json:"merged_at,omitempty"`
	ClosedBy                  *UserDto   `json:"closed_by,omitempty"`
	ClosedAt                  *time.Time `json:"closed_at,omitempty"`
	TargetBranch              string     `json:"target_branch"`
	SourceBranch              string     `json:"source_branch"`
	Upvotes                   int        `json:"upvotes"`
	Downvotes                 int        `json:"downvotes"`
	Author                    UserDto    `json:"author"`
	Assignee                  *UserDto   `json:"assignee,omitempty"`
	Reviewers                 []UserDto  `json:"reviewers,omitempty"`
	SourceProjectID           int        `json:"source_project_id"`
	TargetProjectID           int        `json:"target_project_id"`
	Labels                    []string   `json:"labels"`
	WorkInProgress            bool       `json:"work_in_progress"` // deprecated, replaced by Draft attribute below
	Draft                     bool       `json:"draft"`            // true if MR is draft (alternative to WIP)
	MergeWhenPipelineSucceeds bool       `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string     `json:"merge_status"` // e.g., "can_be_merged"
	SHA                       string     `json:"sha"`          // head commit SHA
	MergeCommitSHA            string     `json:"merge_commit_sha,omitempty"`
	SquashCommitSHA           string     `json:"squash_commit_sha,omitempty"`
	DiscussionLocked          bool       `json:"discussion_locked"`
	ShouldRemoveSourceBranch  bool       `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool       `json:"force_remove_source_branch"`
	WebURL                    string     `json:"web_url"`
	TimeStats                 struct {
		TimeEstimate        int    `json:"time_estimate"`
		TotalTimeSpent      int    `json:"total_time_spent"`
		HumanTimeEstimate   string `json:"human_time_estimate"`
		HumanTotalTimeSpent string `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool `json:"has_conflicts"`
	BlockingDiscussionsResolved bool `json:"blocking_discussions_resolved"`
}
