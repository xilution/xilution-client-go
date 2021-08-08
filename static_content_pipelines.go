package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreateStaticContentPipeline(organizationId *string, staticContent *StaticContentPipeline) (*string, error) {
	rb, _ := json.Marshal(staticContent)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", CoyoteBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetStaticContentPipeline(organizationId *string, staticContentId *string) (*StaticContentPipeline, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", CoyoteBaseUrl, *organizationId, *staticContentId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	staticContent := StaticContentPipeline{}
	json.Unmarshal(body, &staticContent)

	return &staticContent, nil
}

func (xc *XilutionClient) GetStaticContentPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchStaticContentPipelinesResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", CoyoteBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchStaticContentPipelinesResponse := FetchStaticContentPipelinesResponse{}
	json.Unmarshal(body, &fetchStaticContentPipelinesResponse)

	return &fetchStaticContentPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateStaticContentPipeline(organizationId *string, staticContent *StaticContentPipeline) error {
	rb, _ := json.Marshal(staticContent)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", CoyoteBaseUrl, *organizationId, staticContent.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteStaticContentPipeline(organizationId *string, staticContentId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", CoyoteBaseUrl, *organizationId, *staticContentId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateStaticContentPipelineEvent(organizationId *string, pipelineEvent *PipelineEvent) (*string, error) {
	rb, _ := json.Marshal(pipelineEvent)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipeline-events", GazelleBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetStaticContentPipelineEvent(organizationId *string, pipelineEventId *string) (*PipelineEvent, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events/%s", GazelleBaseUrl, *organizationId, *pipelineEventId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	pipelineEvent := PipelineEvent{}
	json.Unmarshal(body, &pipelineEvent)

	return &pipelineEvent, nil
}

func (xc *XilutionClient) GetStaticContentPipelineEvents(organizationId *string, pageSize, pageNumber *int) (*FetchPipelineEventsResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events?pageSize=%d&pageNumber=%d", GazelleBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchPipelineEventsResponse := FetchPipelineEventsResponse{}
	json.Unmarshal(body, &fetchPipelineEventsResponse)

	return &fetchPipelineEventsResponse, nil
}
