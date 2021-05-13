package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateK8sPipeline(organizationId *string, client *K8sPipeline) (*string, error) {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", GiraffeBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetK8sPipeline(organizationId *string, clientId *string) (*K8sPipeline, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GiraffeBaseUrl, *organizationId, *clientId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	client := K8sPipeline{}
	json.Unmarshal(body, &client)

	return &client, nil
}

func (xc *XilutionClient) GetK8sPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchK8sPipelinesResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", GiraffeBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchK8sPipelinesResponse := FetchK8sPipelinesResponse{}
	json.Unmarshal(body, &fetchK8sPipelinesResponse)

	return &fetchK8sPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateK8sPipeline(organizationId *string, client *K8sPipeline) error {
	rb, _ := json.Marshal(client)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GiraffeBaseUrl, *organizationId, client.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteK8sPipeline(organizationId *string, clientId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GiraffeBaseUrl, *organizationId, *clientId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
