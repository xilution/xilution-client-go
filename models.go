package xilution

// Organization -
type Organization struct {
	Type           string `json:"@type"`
	ID             string `json:"id"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt"`
	ModifiedAt     string `json:"modifiedAt"`
	Name           string `json:"name"`
	Logo           string `json:"logo"`
	Domain         string `json:"domain"`
	IamClientId    string `json:"iamClientId"`
	OrganizationId string `json:"organizationId"`
	Active         bool   `json:"active"`
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
	ID             string   `json:"id"`
	OwningUserId   string   `json:"owningUserId"`
	CreatedAt      string   `json:"createdAt"`
	ModifiedAt     string   `json:"modifiedAt"`
	Name           string   `json:"name"`
	Grants         []string `json:"grants"`
	RedirectUris   []string `json:"redirectUris"`
	ClientUserId   string   `json:"clientUserId"`
	OrganizationId string   `json:"organizationId"`
	Active         bool     `json:"active"`
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
	ID             string `json:"id"`
	OwningUserId   string `json:"owningUserId"`
	CreatedAt      string `json:"createdAt"`
	ModifiedAt     string `json:"modifiedAt"`
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
