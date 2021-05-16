package xilution

// Link Label -
type LinkLabel struct {
	Default string `json:"default"`
	En      string `json:"en,omitempty"`
}

// Links -
type Link struct {
	Href  string    `json:"href"`
	Rel   string    `json:"rel,omitempty"`
	Label LinkLabel `json:"label,omitempty"`
	Type  string    `json:"type,omitempty"`
}

// Organization -
type Organization struct {
	Type           string `json:"@type"`
	ID             string `json:"id,omitempty"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt,omitempty"`
	ModifiedAt     string `json:"modifiedAt,omitempty"`
	Name           string `json:"name"`
	Logo           string `json:"logo,omitempty"`
	Domain         string `json:"domain,omitempty"`
	AuthClientId   string `json:"authClientId,omitempty"`
	OrganizationId string `json:"organizationId"`
	Active         bool   `json:"active"`
	Url            string `json:"url,omitempty"`
	AutoAuth       bool   `json:"autoAuth,omitempty"`
	ShowSignUp     bool   `json:"showSignUp,omitempty"`
	Links          []Link `json:"links,omitempty"`
}

// Fetch Organizations Response -
type FetchOrganizationsResponse struct {
	Content          []Organization
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Client -
type Client struct {
	Type           string   `json:"@type"`
	ID             string   `json:"id,omitempty"`
	OwningUserId   string   `json:"owningUserId"`
	CreatedAt      string   `json:"createdAt,omitempty"`
	ModifiedAt     string   `json:"modifiedAt,omitempty"`
	Name           string   `json:"name"`
	Grants         []string `json:"grants"`
	RedirectUris   []string `json:"redirectUris"`
	ClientUserId   string   `json:"clientUserId"`
	OrganizationId string   `json:"organizationId"`
	Active         bool     `json:"active"`
	Secret         string   `json:"secret,omitempty"`
	Links          []Link   `json:"links,omitempty"`
}

// Fetch Clients Response -
type FetchClientsResponse struct {
	Content          []Client
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// User -
type User struct {
	Type           string `json:"@type"`
	ID             string `json:"id,omitempty"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt,omitempty"`
	ModifiedAt     string `json:"modifiedAt,omitempty"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	OrganizationId string `json:"organizationId"`
	Active         bool   `json:"active"`
	Links          []Link `json:"links,omitempty"`
}

// Fetch Users Response -
type FetchUsersResponse struct {
	Content          []User
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Git Account -
type GitAccount struct {
	Type           string `json:"@type"`
	ID             string `json:"id,omitempty"`
	Provider       string `json:"provider"`
	Name           string `json:"name"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt,omitempty"`
	ModifiedAt     string `json:"modifiedAt,omitempty"`
	OrganizationId string `json:"organizationId"`
	Status         string `json:"status,omitempty"`
	Links          []Link `json:"links,omitempty"`
}

// Fetch Git Accounts Response -
type FetchGitAccountsResponse struct {
	Content          []GitAccount
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Git Repo -
type GitRepo struct {
	Type           string `json:"@type"`
	ID             string `json:"id,omitempty"`
	GitAccountId   string `json:"gitAccountId"`
	Name           string `json:"name"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt,omitempty"`
	ModifiedAt     string `json:"modifiedAt,omitempty"`
	OrganizationId string `json:"organizationId"`
	Status         string `json:"status,omitempty"`
	Links          []Link `json:"links,omitempty"`
}

// Fetch Git Repos Response -
type FetchGitReposResponse struct {
	Content          []GitRepo
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Git Repo Event -
type GitRepoEvent struct {
	Type           string                 `json:"@type"`
	ID             string                 `json:"id,omitempty"`
	GitAccountId   string                 `json:"gitAccountId"`
	GitRepoId      string                 `json:"gitRepoId"`
	EventType      string                 `json:"eventType"`
	Parameters     map[string]interface{} `json:"parameters"`
	OwningUserId   string                 `json:"owningUserId"`
	CreatedAt      string                 `json:"createdAt,omitempty"`
	ModifiedAt     string                 `json:"modifiedAt,omitempty"`
	OrganizationId string                 `json:"organizationId"`
	Links          []Link                 `json:"links,omitempty"`
}

// Fetch Git Repo Events Response -
type FetchGitRepoEventsResponse struct {
	Content          []GitRepoEvent
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Cloud Provider -
type CloudProvider struct {
	Type           string `json:"@type"`
	ID             string `json:"id,omitempty"`
	Provider       string `json:"provider"`
	Name           string `json:"name"`
	AccountId      string `json:"accountId"`
	Region         string `json:"region"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt,omitempty"`
	ModifiedAt     string `json:"modifiedAt,omitempty"`
	OrganizationId string `json:"organizationId"`
	Status         string `json:"status,omitempty"`
	Links          []Link `json:"links,omitempty"`
}

// Fetch Cloud Providers Response -
type FetchCloudProvidersResponse struct {
	Content          []CloudProvider
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Continuous Integration Status -
type ContinuousIntegrationStatus struct {
	LatestDownExecutionStatus string `json:"latestDownExecutionStatus,omitempty"`
	LatestUpExecutionStatus   string `json:"latestUpExecutionStatus,omitempty"`
}

// Pipeline Status -
type PipelineStatus struct {
	ContinuousIntegrationStatus ContinuousIntegrationStatus `json:"continuousIntegrationStatus,omitempty"`
	InfrastructureStatus        string                      `json:"infrastructureStatus,omitempty"`
}

// VPC Pipelines -
type VpcPipeline struct {
	Type            string         `json:"@type"`
	ID              string         `json:"id,omitempty"`
	Name            string         `json:"name"`
	PipelineType    string         `json:"pipelineType"`
	CloudProviderId string         `json:"cloudProviderId"`
	OwningUserId    string         `json:"owningUserId"`
	CreatedAt       string         `json:"createdAt,omitempty"`
	ModifiedAt      string         `json:"modifiedAt,omitempty"`
	OrganizationId  string         `json:"organizationId"`
	Status          PipelineStatus `json:"status,omitempty"`
	Links           []Link         `json:"links,omitempty"`
}

// Fetch VPC Pipelines Response -
type FetchVpcPipelinesResponse struct {
	Content          []VpcPipeline
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// K8s Pipelines -
type K8sPipeline struct {
	Type           string         `json:"@type"`
	ID             string         `json:"id,omitempty"`
	Name           string         `json:"name"`
	PipelineType   string         `json:"pipelineType"`
	VpcPipelineId  string         `json:"gazellePipelineId"`
	OwningUserId   string         `json:"owningUserId"`
	CreatedAt      string         `json:"createdAt,omitempty"`
	ModifiedAt     string         `json:"modifiedAt,omitempty"`
	OrganizationId string         `json:"organizationId"`
	Status         PipelineStatus `json:"status,omitempty"`
	Links          []Link         `json:"links,omitempty"`
}

// Fetch K8s Pipelines Response -
type FetchK8sPipelinesResponse struct {
	Content          []K8sPipeline
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// WordPress Stage -
type WordPressStage struct {
	Name    string `json:"name"`
	Links   []Link `json:"links,omitempty"`
	SiteUrl string `json:"siteUrl,omitempty"`
}

// WordPress Pipelines -
type WordPressPipeline struct {
	Type           string           `json:"@type"`
	ID             string           `json:"id,omitempty"`
	Name           string           `json:"name"`
	PipelineType   string           `json:"pipelineType"`
	K8sPipelineId  string           `json:"gazellePipelineId"`
	GitRepoId      string           `json:"gitRepoId"`
	Branch         string           `json:"branch"`
	Stages         []WordPressStage `json:"stages"`
	OwningUserId   string           `json:"owningUserId"`
	CreatedAt      string           `json:"createdAt,omitempty"`
	ModifiedAt     string           `json:"modifiedAt,omitempty"`
	OrganizationId string           `json:"organizationId"`
	Status         PipelineStatus   `json:"status,omitempty"`
	Links          []Link           `json:"links,omitempty"`
}

// Fetch WordPress Pipelines Response -
type FetchWordPressPipelinesResponse struct {
	Content          []WordPressPipeline
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// Static Content Stage -
type StaticContentStage struct {
	Name    string `json:"name"`
	Links   []Link `json:"links,omitempty"`
	SiteUrl string `json:"siteUrl,omitempty"`
}

// Static Content Pipelines -
type StaticContentPipeline struct {
	Type            string               `json:"@type"`
	ID              string               `json:"id,omitempty"`
	Name            string               `json:"name"`
	PipelineType    string               `json:"pipelineType"`
	CloudProviderId string               `json:"cloudProviderId"`
	GitRepoId       string               `json:"gitRepoId"`
	Branch          string               `json:"branch"`
	Stages          []StaticContentStage `json:"stages"`
	OwningUserId    string               `json:"owningUserId"`
	CreatedAt       string               `json:"createdAt,omitempty"`
	ModifiedAt      string               `json:"modifiedAt,omitempty"`
	OrganizationId  string               `json:"organizationId"`
	Status          PipelineStatus       `json:"status,omitempty"`
	Links           []Link               `json:"links,omitempty"`
}

// Fetch Static Content Pipelines Response -
type FetchStaticContentPipelinesResponse struct {
	Content          []StaticContentPipeline
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}

// API Stage -
type ApiStage struct {
	Name       string `json:"name"`
	Links      []Link `json:"links,omitempty"`
	ApiBaseUrl string `json:"apiBaseUrl,omitempty"`
}

// API Pipelines -
type ApiPipeline struct {
	Type           string         `json:"@type"`
	ID             string         `json:"id,omitempty"`
	Name           string         `json:"name"`
	PipelineType   string         `json:"pipelineType"`
	VpcPipelineId  string         `json:"gazellePipelineId"`
	GitRepoId      string         `json:"gitRepoId"`
	Branch         string         `json:"branch"`
	Stages         []ApiStage     `json:"stages"`
	OwningUserId   string         `json:"owningUserId"`
	CreatedAt      string         `json:"createdAt,omitempty"`
	ModifiedAt     string         `json:"modifiedAt,omitempty"`
	OrganizationId string         `json:"organizationId"`
	Status         PipelineStatus `json:"status,omitempty"`
	Links          []Link         `json:"links,omitempty"`
}

// Fetch API Pipelines Response -
type FetchApiPipelinesResponse struct {
	Content          []ApiPipeline
	PageSize         int    `json:"pageSize,omitempty"`
	PageNumber       int    `json:"pageNumber,omitempty"`
	TotalPages       int    `json:"totalPages"`
	NumberOfElements int    `json:"numberOfElements,omitempty"`
	TotalElements    int    `json:"totalElements"`
	FirstPage        bool   `json:"firstPage"`
	LastPage         bool   `json:"lastPage"`
	ID               Link   `json:"link"`
	Links            []Link `json:"links,omitempty"`
}
