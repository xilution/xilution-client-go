package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateUser(organizationId *string, user *User) (*string, error) {
	rb, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/users", RhinoBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetUser(organizationId *string, userId *string) (*User, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/users/%s", RhinoBaseUrl, *organizationId, *userId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	user := User{}
	json.Unmarshal(body, &user)

	return &user, nil
}

func (xc *XilutionClient) GetUsers(organizationId *string, pageSize, pageNumber *int) (*FetchUsersResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/users?pageSize=%d&pageNumber=%d", RhinoBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchUsersResponse := FetchUsersResponse{}
	json.Unmarshal(body, &fetchUsersResponse)

	return &fetchUsersResponse, nil
}

func (xc *XilutionClient) UpdateUser(organizationId *string, user *User) error {
	rb, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/users/%s", RhinoBaseUrl, *organizationId, user.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteUser(organizationId *string, userId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/users/%s", RhinoBaseUrl, *organizationId, *userId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
