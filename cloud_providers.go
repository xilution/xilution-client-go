package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreateCloudProvider(organizationId *string, cloudProvider *CloudProvider) (*string, error) {
	rb, _ := json.Marshal(cloudProvider)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/cloud-providers", KangarooBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetCloudProvider(organizationId *string, cloudProviderId *string) (*CloudProvider, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/cloud-providers/%s", KangarooBaseUrl, *organizationId, *cloudProviderId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	cloudProvider := CloudProvider{}
	json.Unmarshal(body, &cloudProvider)

	return &cloudProvider, nil
}

func (xc *XilutionClient) GetCloudProviders(organizationId *string, pageSize, pageNumber *int) (*FetchCloudProvidersResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/cloud-providers?pageSize=%d&pageNumber=%d", KangarooBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchCloudProvidersResponse := FetchCloudProvidersResponse{}
	json.Unmarshal(body, &fetchCloudProvidersResponse)

	return &fetchCloudProvidersResponse, nil
}

func (xc *XilutionClient) UpdateCloudProvider(organizationId *string, cloudProvider *CloudProvider) error {
	rb, _ := json.Marshal(cloudProvider)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/cloud-providers/%s", KangarooBaseUrl, *organizationId, cloudProvider.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteCloudProvider(organizationId *string, cloudProviderId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/cloud-providers/%s", KangarooBaseUrl, *organizationId, *cloudProviderId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
