package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateUser(user *User) (*string, error) {
	rb, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/users", ElephantBaseUrl), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetUser(userId *string) (*User, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", ElephantBaseUrl, *userId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	user := User{}
	json.Unmarshal(body, &user)

	return &user, nil
}

func (xc *XilutionClient) GetUsers(pageSize, pageNumber *int) (*FetchUsersResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/users?pageSize=%d&pageNumber=%d", ElephantBaseUrl, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchUsersResponse := FetchUsersResponse{}
	json.Unmarshal(body, &fetchUsersResponse)

	return &fetchUsersResponse, nil
}

func (xc *XilutionClient) UpdateUser(user *User) error {
	rb, _ := json.Marshal(user)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s", ElephantBaseUrl, user.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteUser(userId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%s", ElephantBaseUrl, *userId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
