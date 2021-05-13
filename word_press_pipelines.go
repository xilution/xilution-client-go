package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateWordPressPipeline(organizationId *string, client *WordPressPipeline) (*string, error) {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", PenguinBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetWordPressPipeline(organizationId *string, clientId *string) (*WordPressPipeline, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", PenguinBaseUrl, *organizationId, *clientId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	client := WordPressPipeline{}
	json.Unmarshal(body, &client)

	return &client, nil
}

func (xc *XilutionClient) GetWordPressPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchWordPressPipelinesResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", PenguinBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchWordPressPipelinesResponse := FetchWordPressPipelinesResponse{}
	json.Unmarshal(body, &fetchWordPressPipelinesResponse)

	return &fetchWordPressPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateWordPressPipeline(organizationId *string, client *WordPressPipeline) error {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", PenguinBaseUrl, *organizationId, client.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteWordPressPipeline(organizationId *string, clientId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", PenguinBaseUrl, *organizationId, *clientId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
