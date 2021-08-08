package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreateApiPipeline(organizationId *string, api *ApiPipeline) (*string, error) {
	rb, _ := json.Marshal(api)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", FoxBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetApiPipeline(organizationId *string, apiId *string) (*ApiPipeline, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", FoxBaseUrl, *organizationId, *apiId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	api := ApiPipeline{}
	json.Unmarshal(body, &api)

	return &api, nil
}

func (xc *XilutionClient) GetApiPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchApiPipelinesResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", FoxBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchApiPipelinesResponse := FetchApiPipelinesResponse{}
	json.Unmarshal(body, &fetchApiPipelinesResponse)

	return &fetchApiPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateApiPipeline(organizationId *string, api *ApiPipeline) error {
	rb, _ := json.Marshal(api)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", FoxBaseUrl, *organizationId, api.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteApiPipeline(organizationId *string, apiId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", FoxBaseUrl, *organizationId, *apiId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateApiPipelineEvent(organizationId *string, pipelineEvent *PipelineEvent) (*string, error) {
	rb, _ := json.Marshal(pipelineEvent)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipeline-events", GazelleBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetApiPipelineEvent(organizationId *string, pipelineEventId *string) (*PipelineEvent, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events/%s", GazelleBaseUrl, *organizationId, *pipelineEventId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	pipelineEvent := PipelineEvent{}
	json.Unmarshal(body, &pipelineEvent)

	return &pipelineEvent, nil
}

func (xc *XilutionClient) GetApiPipelineEvents(organizationId *string, pageSize, pageNumber *int) (*FetchPipelineEventsResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events?pageSize=%d&pageNumber=%d", GazelleBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchPipelineEventsResponse := FetchPipelineEventsResponse{}
	json.Unmarshal(body, &fetchPipelineEventsResponse)

	return &fetchPipelineEventsResponse, nil
}
