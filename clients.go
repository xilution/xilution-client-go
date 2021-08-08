package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreateClient(organizationId *string, client *Client) (*string, error) {
	rb, _ := json.Marshal(client)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/clients", HippoBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetClient(organizationId *string, clientId *string) (*Client, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/clients/%s", HippoBaseUrl, *organizationId, *clientId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	client := Client{}
	json.Unmarshal(body, &client)

	return &client, nil
}

func (xc *XilutionClient) GetClients(organizationId *string, pageSize, pageNumber *int) (*FetchClientsResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/clients?pageSize=%d&pageNumber=%d", HippoBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchClientsResponse := FetchClientsResponse{}
	json.Unmarshal(body, &fetchClientsResponse)

	return &fetchClientsResponse, nil
}

func (xc *XilutionClient) UpdateClient(organizationId *string, client *Client) error {
	rb, _ := json.Marshal(client)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/clients/%s", HippoBaseUrl, *organizationId, client.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteClient(organizationId *string, clientId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/clients/%s", HippoBaseUrl, *organizationId, *clientId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
