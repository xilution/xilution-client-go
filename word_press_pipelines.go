package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateWordPressPipeline(organizationId *string, wordPress *WordPressPipeline) (*string, error) {
	rb, _ := json.Marshal(wordPress)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipelines", PenguinBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetWordPressPipeline(organizationId *string, wordPressId *string) (*WordPressPipeline, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipelines/%s", PenguinBaseUrl, *organizationId, *wordPressId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	wordPress := WordPressPipeline{}
	json.Unmarshal(body, &wordPress)

	return &wordPress, nil
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

func (xc *XilutionClient) UpdateWordPressPipeline(organizationId *string, wordPress *WordPressPipeline) error {
	rb, _ := json.Marshal(wordPress)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/pipelines/%s", PenguinBaseUrl, *organizationId, wordPress.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteWordPressPipeline(organizationId *string, wordPressId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/pipelines/%s", PenguinBaseUrl, *organizationId, *wordPressId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateWordPressPipelineEvent(organizationId *string, pipelineEvent *PipelineEvent) (*string, error) {
	rb, _ := json.Marshal(pipelineEvent)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/pipeline-events", GazelleBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetWordPressPipelineEvent(organizationId *string, pipelineEventId *string) (*PipelineEvent, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events/%s", GazelleBaseUrl, *organizationId, *pipelineEventId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	pipelineEvent := PipelineEvent{}
	json.Unmarshal(body, &pipelineEvent)

	return &pipelineEvent, nil
}

func (xc *XilutionClient) GetWordPressPipelineEvents(organizationId *string, pageSize, pageNumber *int) (*FetchPipelineEventsResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/pipeline-events?pageSize=%d&pageNumber=%d", GazelleBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchPipelineEventsResponse := FetchPipelineEventsResponse{}
	json.Unmarshal(body, &fetchPipelineEventsResponse)

	return &fetchPipelineEventsResponse, nil
}
