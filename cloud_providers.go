package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateCloudProvider(organizationId *string, cloudProvider *CloudProvider) (*string, error) {
	rb, _ := json.Marshal(cloudProvider)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/cloud-providers", RhinoBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetCloudProvider(organizationId *string, cloudProviderId *string) (*CloudProvider, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/cloud-providers/%s", RhinoBaseUrl, *organizationId, *cloudProviderId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	cloudProvider := CloudProvider{}
	json.Unmarshal(body, &cloudProvider)

	return &cloudProvider, nil
}

func (xc *XilutionClient) GetCloudProviders(organizationId *string, pageSize, pageNumber *int) (*FetchCloudProvidersResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/cloud-providers?pageSize=%d&pageNumber=%d", RhinoBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

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

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/cloud-providers/%s", RhinoBaseUrl, *organizationId, cloudProvider.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteCloudProvider(organizationId *string, cloudProviderId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/cloud-providers/%s", RhinoBaseUrl, *organizationId, *cloudProviderId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
