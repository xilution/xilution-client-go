package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateVpcPipeline(organizationId *string, pipeline *VpcPipeline) (*string, error) {
	rb, _ := json.Marshal(pipeline)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", GazelleBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetVpcPipeline(organizationId *string, pipeline *string) (*VpcPipeline, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GazelleBaseUrl, *organizationId, *pipeline), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	vpcPipeline := VpcPipeline{}
	json.Unmarshal(body, &vpcPipeline)

	return &vpcPipeline, nil
}

func (xc *XilutionClient) GetVpcPipelines(organizationId *string, pageSize, pageNumber *int) (*FetchVpcPipelinesResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines?pageSize=%d&pageNumber=%d", GazelleBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchVpcPipelinesResponse := FetchVpcPipelinesResponse{}
	json.Unmarshal(body, &fetchVpcPipelinesResponse)

	return &fetchVpcPipelinesResponse, nil
}

func (xc *XilutionClient) UpdateVpcPipeline(organizationId *string, vpcPipeline *VpcPipeline) error {
	rb, _ := json.Marshal(vpcPipeline)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GazelleBaseUrl, *organizationId, vpcPipeline.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteVpcPipeline(organizationId *string, pipeline *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", GazelleBaseUrl, *organizationId, *pipeline), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateVpcPipelineEvent(organizationId *string, pipelineEvent *PipelineEvent) (*string, error) {
	rb, _ := json.Marshal(pipelineEvent)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipeline-events", GazelleBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetVpcPipelineEvent(organizationId *string, pipelineEventId *string) (*PipelineEvent, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events/%s", GazelleBaseUrl, *organizationId, *pipelineEventId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	pipelineEvent := PipelineEvent{}
	json.Unmarshal(body, &pipelineEvent)

	return &pipelineEvent, nil
}

func (xc *XilutionClient) GetVpcPipelineEvents(organizationId *string, pageSize, pageNumber *int) (*FetchPipelineEventsResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events?pageSize=%d&pageNumber=%d", GazelleBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchPipelineEventsResponse := FetchPipelineEventsResponse{}
	json.Unmarshal(body, &fetchPipelineEventsResponse)

	return &fetchPipelineEventsResponse, nil
}
