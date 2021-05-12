package xilution

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
}

// FetchOrganizationsResponse -
type FetchOrganizationsResponse struct {
	Content          []Organization
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
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
}

// FetchClientsResponse -
type FetchClientsResponse struct {
	Content          []Client
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
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
}

// FetchUsersResponse -
type FetchUsersResponse struct {
	Content          []User
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
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
}

// FetchGitAccountsResponse -
type FetchGitAccountsResponse struct {
	Content          []GitAccount
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
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
}

// FetchGitReposResponse -
type FetchGitReposResponse struct {
	Content          []GitRepo
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
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
}

// FetchGitRepoEventsResponse -
type FetchGitRepoEventsResponse struct {
	Content          []GitRepoEvent
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
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
}

// FetchCloudProvidersResponse -
type FetchCloudProvidersResponse struct {
	Content          []CloudProvider
	PageSize         int  `json:"pageSize"`
	PageNumber       int  `json:"pageNumber"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	TotalElements    int  `json:"totalElements"`
	FirstPage        bool `json:"firstPage"`
	LastPage         bool `json:"lastPage"`
}
