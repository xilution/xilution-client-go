package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateClient(client Client) (*string, error) {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/clients", ElephantBaseUrl), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (c *XilutionClient) GetClient(clientId string) (*Client, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/clients/%s", ElephantBaseUrl, clientId), nil)

	body, err := c.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	client := Client{}
	json.Unmarshal(body, &client)

	return &client, nil
}

func (xc *XilutionClient) GetClients(pageSize, pageNumber int) (*FetchClientsResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/clients?pageSize=%d&pageNumber=%d", ElephantBaseUrl, pageSize, pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchClientsResponse := FetchClientsResponse{}
	json.Unmarshal(body, &fetchClientsResponse)

	return &fetchClientsResponse, nil
}

func (xc *XilutionClient) UpdateClient(client Client) error {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/clients/%s", ElephantBaseUrl, client.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteClient(clientId string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/clients/%s", ElephantBaseUrl, clientId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
