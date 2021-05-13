package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateApiPipeline(organizationId *string, client *ApiPipeline) (*string, error) {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", FoxBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetApiPipeline(organizationId *string, clientId *string) (*ApiPipeline, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", FoxBaseUrl, *organizationId, *clientId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	client := ApiPipeline{}
	json.Unmarshal(body, &client)

	return &client, nil
}

func (xc *XilutionClient) GetApiPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchApiPipelinesResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", FoxBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchApiPipelinesResponse := FetchApiPipelinesResponse{}
	json.Unmarshal(body, &fetchApiPipelinesResponse)

	return &fetchApiPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateApiPipeline(organizationId *string, client *ApiPipeline) error {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", FoxBaseUrl, *organizationId, client.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteApiPipeline(organizationId *string, clientId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", FoxBaseUrl, *organizationId, *clientId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
