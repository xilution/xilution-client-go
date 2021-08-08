package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreateK8sPipeline(organizationId *string, k8sPipeline *K8sPipeline) (*string, error) {
	rb, _ := json.Marshal(k8sPipeline)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", GiraffeBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetK8sPipeline(organizationId *string, k8sPipelineId *string) (*K8sPipeline, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GiraffeBaseUrl, *organizationId, *k8sPipelineId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	k8sPipeline := K8sPipeline{}
	json.Unmarshal(body, &k8sPipeline)

	return &k8sPipeline, nil
}

func (xc *XilutionClient) GetK8sPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchK8sPipelinesResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", GiraffeBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchK8sPipelinesResponse := FetchK8sPipelinesResponse{}
	json.Unmarshal(body, &fetchK8sPipelinesResponse)

	return &fetchK8sPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateK8sPipeline(organizationId *string, k8sPipeline *K8sPipeline) error {
	rb, _ := json.Marshal(k8sPipeline)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GiraffeBaseUrl, *organizationId, k8sPipeline.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteK8sPipeline(organizationId *string, k8sPipelineId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GiraffeBaseUrl, *organizationId, *k8sPipelineId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateK8sPipelineEvent(organizationId *string, pipelineEvent *PipelineEvent) (*string, error) {
	rb, _ := json.Marshal(pipelineEvent)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipeline-events", GazelleBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetK8sPipelineEvent(organizationId *string, pipelineEventId *string) (*PipelineEvent, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events/%s", GazelleBaseUrl, *organizationId, *pipelineEventId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	pipelineEvent := PipelineEvent{}
	json.Unmarshal(body, &pipelineEvent)

	return &pipelineEvent, nil
}

func (xc *XilutionClient) GetK8sPipelineEvents(organizationId *string, pageSize, pageNumber *int) (*FetchPipelineEventsResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events?pageSize=%d&pageNumber=%d", GazelleBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchPipelineEventsResponse := FetchPipelineEventsResponse{}
	json.Unmarshal(body, &fetchPipelineEventsResponse)

	return &fetchPipelineEventsResponse, nil
}
