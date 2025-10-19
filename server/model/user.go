package model

import (
	"time"
)

type User struct {
	ID                  int       `json:"id"`
	Username            string    `json:"username"`
	Name                string    `json:"name"`
	State               string    `json:"state"`
	Locked              bool      `json:"locked"`
	AvatarURL           string    `json:"avatar_url"`
	WebURL              string    `json:"web_url"`
	CreatedAt           time.Time `json:"created_at"`
	Bio                 string    `json:"bio"`
	Location            string    `json:"location"`
	PublicEmail         *string   `json:"public_email"`
	Skype               string    `json:"skype"`
	Linkedin            string    `json:"linkedin"`
	Twitter             string    `json:"twitter"`
	Discord             string    `json:"discord"`
	WebsiteURL          string    `json:"website_url"`
	Organization        string    `json:"organization"`
	JobTitle            string    `json:"job_title"`
	Pronouns            *string   `json:"pronouns"`
	Bot                 bool      `json:"bot"`
	WorkInformation     *string   `json:"work_information"`
	LocalTime           *string   `json:"local_time"`
	LastSignInAt        time.Time `json:"last_sign_in_at"`
	ConfirmedAt         time.Time `json:"confirmed_at"`
	LastActivityOn      string    `json:"last_activity_on"`
	Email               string    `json:"email"`
	ThemeID             int       `json:"theme_id"`
	ColorSchemeID       int       `json:"color_scheme_id"`
	ProjectsLimit       int       `json:"projects_limit"`
	CurrentSignInAt     time.Time `json:"current_sign_in_at"`
	Identities          []any     `json:"identities"`
	CanCreateGroup      bool      `json:"can_create_group"`
	CanCreateProject    bool      `json:"can_create_project"`
	TwoFactorEnabled    bool      `json:"two_factor_enabled"`
	External            bool      `json:"external"`
	PrivateProfile      bool      `json:"private_profile"`
	CommitEmail         string    `json:"commit_email"`
	IsAdmin             bool      `json:"is_admin"`
	Note                string    `json:"note"`
	NamespaceID         int       `json:"namespace_id"`
	CreatedBy           *User     `json:"created_by"`
	EmailResetOfferedAt *string   `json:"email_reset_offered_at"`
}
