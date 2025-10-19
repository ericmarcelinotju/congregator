package dto

import (
	"time"
)

type ProjectFilter struct{}

type ProjectDto struct {
	ID                                 int                `json:"id"`
	Description                        *string            `json:"description"`
	Name                               string             `json:"name"`
	NameWithNamespace                  string             `json:"name_with_namespace"`
	Path                               string             `json:"path"`
	PathWithNamespace                  string             `json:"path_with_namespace"`
	CreatedAt                          time.Time          `json:"created_at"`
	DefaultBranch                      string             `json:"default_branch"`
	TagList                            []string           `json:"tag_list"`
	Topics                             []string           `json:"topics"`
	SSHURLToRepo                       string             `json:"ssh_url_to_repo"`
	HTTPURLToRepo                      string             `json:"http_url_to_repo"`
	WebURL                             string             `json:"web_url"`
	ReadmeURL                          string             `json:"readme_url"`
	ForksCount                         int                `json:"forks_count"`
	AvatarURL                          *string            `json:"avatar_url"`
	StarCount                          int                `json:"star_count"`
	LastActivityAt                     time.Time          `json:"last_activity_at"`
	Namespace                          NamespaceDto       `json:"namespace"`
	RepositoryStorage                  string             `json:"repository_storage"`
	Links                              LinksDto           `json:"_links"`
	PackagesEnabled                    bool               `json:"packages_enabled"`
	EmptyRepo                          bool               `json:"empty_repo"`
	Archived                           bool               `json:"archived"`
	Visibility                         string             `json:"visibility"`
	ResolveOutdatedDiffs               bool               `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy          ContainerPolicyDto `json:"container_expiration_policy"`
	RepositoryObjectFormat             string             `json:"repository_object_format"`
	IssuesEnabled                      bool               `json:"issues_enabled"`
	MergeRequestsEnabled               bool               `json:"merge_requests_enabled"`
	WikiEnabled                        bool               `json:"wiki_enabled"`
	JobsEnabled                        bool               `json:"jobs_enabled"`
	SnippetsEnabled                    bool               `json:"snippets_enabled"`
	ContainerRegistryEnabled           bool               `json:"container_registry_enabled"`
	ServiceDeskEnabled                 bool               `json:"service_desk_enabled"`
	ServiceDeskAddress                 *string            `json:"service_desk_address"`
	CanCreateMergeRequestIn            bool               `json:"can_create_merge_request_in"`
	IssuesAccessLevel                  string             `json:"issues_access_level"`
	RepositoryAccessLevel              string             `json:"repository_access_level"`
	MergeRequestsAccessLevel           string             `json:"merge_requests_access_level"`
	ForkingAccessLevel                 string             `json:"forking_access_level"`
	WikiAccessLevel                    string             `json:"wiki_access_level"`
	BuildsAccessLevel                  string             `json:"builds_access_level"`
	SnippetsAccessLevel                string             `json:"snippets_access_level"`
	PagesAccessLevel                   string             `json:"pages_access_level"`
	AnalyticsAccessLevel               string             `json:"analytics_access_level"`
	ContainerRegistryAccessLevel       string             `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel   string             `json:"security_and_compliance_access_level"`
	ReleasesAccessLevel                string             `json:"releases_access_level"`
	EnvironmentsAccessLevel            string             `json:"environments_access_level"`
	FeatureFlagsAccessLevel            string             `json:"feature_flags_access_level"`
	InfrastructureAccessLevel          string             `json:"infrastructure_access_level"`
	MonitorAccessLevel                 string             `json:"monitor_access_level"`
	ModelExperimentsAccessLevel        string             `json:"model_experiments_access_level"`
	ModelRegistryAccessLevel           string             `json:"model_registry_access_level"`
	EmailsDisabled                     bool               `json:"emails_disabled"`
	EmailsEnabled                      bool               `json:"emails_enabled"`
	SharedRunnersEnabled               bool               `json:"shared_runners_enabled"`
	LFSEnabled                         bool               `json:"lfs_enabled"`
	CreatorID                          int                `json:"creator_id"`
	ImportURL                          *string            `json:"import_url"`
	ImportType                         *string            `json:"import_type"`
	ImportStatus                       string             `json:"import_status"`
	OpenIssuesCount                    int                `json:"open_issues_count"`
	DescriptionHTML                    string             `json:"description_html"`
	UpdatedAt                          time.Time          `json:"updated_at"`
	CIDefaultGitDepth                  int                `json:"ci_default_git_depth"`
	CIForwardDeploymentEnabled         bool               `json:"ci_forward_deployment_enabled"`
	CIForwardDeploymentRollbackAllowed bool               `json:"ci_forward_deployment_rollback_allowed"`
	CIJobTokenScopeEnabled             bool               `json:"ci_job_token_scope_enabled"`
	CISeparatedCaches                  bool               `json:"ci_separated_caches"`
	CIForkPipelinesInParent            bool               `json:"ci_allow_fork_pipelines_to_run_in_parent_project"`
	CIIDTokenSubClaimComponents        []string           `json:"ci_id_token_sub_claim_components"`
	BuildGitStrategy                   string             `json:"build_git_strategy"`
	KeepLatestArtifact                 bool               `json:"keep_latest_artifact"`
	RestrictUserDefinedVars            bool               `json:"restrict_user_defined_variables"`
	CIPipelineOverrideRole             string             `json:"ci_pipeline_variables_minimum_override_role"`
	RunnersToken                       string             `json:"runners_token"`
	RunnerTokenExpiration              *string            `json:"runner_token_expiration_interval"`
	GroupRunnersEnabled                bool               `json:"group_runners_enabled"`
	AutoCancelPendingPipelines         string             `json:"auto_cancel_pending_pipelines"`
	BuildTimeout                       int                `json:"build_timeout"`
	AutoDevopsEnabled                  bool               `json:"auto_devops_enabled"`
	AutoDevopsStrategy                 string             `json:"auto_devops_deploy_strategy"`
	CIPushRepoJobTokenAllowed          bool               `json:"ci_push_repository_for_job_token_allowed"`
	CIConfigPath                       string             `json:"ci_config_path"`
	PublicJobs                         bool               `json:"public_jobs"`
	SharedWithGroups                   []string           `json:"shared_with_groups"`
	OnlyMergeIfPipelineSucceeds        bool               `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline        *bool              `json:"allow_merge_on_skipped_pipeline"`
	OnlyMergeIfDiscussionsResolved     bool               `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge       bool               `json:"remove_source_branch_after_merge"`
	PrintMergeRequestLinkEnabled       bool               `json:"printing_merge_request_link_enabled"`
	MergeMethod                        string             `json:"merge_method"`
	SquashOption                       string             `json:"squash_option"`
	EnforceAuthChecksOnUploads         bool               `json:"enforce_auth_checks_on_uploads"`
	SuggestionCommitMessage            *string            `json:"suggestion_commit_message"`
	MergeCommitTemplate                *string            `json:"merge_commit_template"`
	SquashCommitTemplate               *string            `json:"squash_commit_template"`
	IssueBranchTemplate                *string            `json:"issue_branch_template"`
	WarnAboutUnwantedChars             bool               `json:"warn_about_potentially_unwanted_characters"`
	AutocloseReferencedIssues          bool               `json:"autoclose_referenced_issues"`
	Permissions                        PermissionsDto     `json:"permissions"`
}

type NamespaceDto struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Kind      string  `json:"kind"`
	FullPath  string  `json:"full_path"`
	ParentID  *int    `json:"parent_id"`
	AvatarURL *string `json:"avatar_url"`
	WebURL    string  `json:"web_url"`
}

type LinksDto struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
	ClusterAgents string `json:"cluster_agents"`
}

type ContainerPolicyDto struct {
	Cadence       string    `json:"cadence"`
	Enabled       bool      `json:"enabled"`
	KeepN         int       `json:"keep_n"`
	OlderThan     string    `json:"older_than"`
	NameRegex     string    `json:"name_regex"`
	NameRegexKeep *string   `json:"name_regex_keep"`
	NextRunAt     time.Time `json:"next_run_at"`
}

type PermissionsDto struct {
	ProjectAccess *AccessDto `json:"project_access"`
	GroupAccess   *AccessDto `json:"group_access"`
}

type AccessDto struct {
	AccessLevel       int `json:"access_level"`
	NotificationLevel int `json:"notification_level"`
}
