package dto

import "time"

type UserFilter struct {
	AuthorUsername *string    `uri:"author_username"`
	CreatedAfter   *time.Time `uri:"created_after"`
	CreatedBefore  *time.Time `uri:"created_before"`
}

type UserDto struct {
	Id                  int       `json:"id"`
	Username            string    `json:"username"`
	Name                string    `json:"name"`
	State               string    `json:"state"`
	Locked              bool      `json:"locked"`
	AvatarURL           string    `json:"avatarUrl"`
	WebURL              string    `json:"webUrl"`
	CreatedAt           time.Time `json:"createdAt"`
	Bio                 string    `json:"bio"`
	Location            string    `json:"location"`
	PublicEmail         *string   `json:"publicEmail"`
	Skype               string    `json:"skype"`
	Linkedin            string    `json:"linkedin"`
	Twitter             string    `json:"twitter"`
	Discord             string    `json:"discord"`
	WebsiteURL          string    `json:"websiteUrl"`
	Organization        string    `json:"organization"`
	JobTitle            string    `json:"jobTitle"`
	Pronouns            *string   `json:"pronouns"`
	Bot                 bool      `json:"bot"`
	WorkInformation     *string   `json:"workInformation"`
	LocalTime           *string   `json:"localTime"`
	LastSignInAt        time.Time `json:"lastSignInAt"`
	ConfirmedAt         time.Time `json:"confirmedAt"`
	LastActivityOn      string    `json:"lastActivityOn"`
	Email               string    `json:"email"`
	ThemeId             int       `json:"themeId"`
	ColorSchemeId       int       `json:"colorSchemeId"`
	ProjectsLimit       int       `json:"projectsLimit"`
	CurrentSignInAt     time.Time `json:"currentSignInAt"`
	Identities          []any     `json:"identities"`
	CanCreateGroup      bool      `json:"canCreateGroup"`
	CanCreateProject    bool      `json:"canCreateProject"`
	TwoFactorEnabled    bool      `json:"twoFactorEnabled"`
	External            bool      `json:"external"`
	PrivateProfile      bool      `json:"privateProfile"`
	CommitEmail         string    `json:"commitEmail"`
	IsAdmin             bool      `json:"isAdmin"`
	Note                string    `json:"note"`
	NamespaceId         int       `json:"namespaceId"`
	CreatedBy           *UserDto  `json:"createdBy"`
	EmailResetOfferedAt *string   `json:"emailResetOfferedAt"`
}
