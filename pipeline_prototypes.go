package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreatePipelinePrototype(organizationId *string, pipelinePrototype *PipelinePrototype) (*string, error) {
	rb, _ := json.Marshal(pipelinePrototype)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipeline-prototypes", BisonBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetPipelinePrototype(organizationId, pipelinePrototypeId *string) (*PipelinePrototype, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-prototypes/%s", BisonBaseUrl, *organizationId, *pipelinePrototypeId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	vpcPipeline := PipelinePrototype{}
	json.Unmarshal(body, &vpcPipeline)

	return &vpcPipeline, nil
}

func (xc *XilutionClient) GetPipelinePrototypes(organizationId *string, pageSize, pageNumber *int) (*FetchPipelinePrototypesResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-prototypes?pageSize=%d&pageNumber=%d", BisonBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchPipelinePrototypesResponse := FetchPipelinePrototypesResponse{}
	json.Unmarshal(body, &fetchPipelinePrototypesResponse)

	return &fetchPipelinePrototypesResponse, nil
}

func (xc *XilutionClient) UpdatePipelinePrototype(organizationId *string, pipelinePrototype *PipelinePrototype) error {
	rb, _ := json.Marshal(pipelinePrototype)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipeline-prototypes/%s", BisonBaseUrl, *organizationId, pipelinePrototype.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeletePipelinePrototype(organizationId *string, pipelinePrototypeId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipeline-prototypes/%s", BisonBaseUrl, *organizationId, *pipelinePrototypeId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
