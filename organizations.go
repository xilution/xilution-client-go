package xilution

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func (xc *XilutionClient) CreateOrganization(organization *Organization) (*string, error) {
	rb, _ := json.Marshal(organization)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations", ElephantBaseUrl), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetOrganization(organizationId *string) (*Organization, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s", ElephantBaseUrl, *organizationId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	organization := Organization{}
	json.Unmarshal(body, &organization)

	return &organization, nil
}

func (xc *XilutionClient) GetOrganizations(pageSize, pageNumber *int) (*FetchOrganizationsResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations?pageSize=%d&pageNumber=%d", ElephantBaseUrl, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchOrganizationsResponse := FetchOrganizationsResponse{}
	json.Unmarshal(body, &fetchOrganizationsResponse)

	return &fetchOrganizationsResponse, nil
}

func (xc *XilutionClient) UpdateOrganization(organization *Organization) error {
	rb, _ := json.Marshal(organization)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s", ElephantBaseUrl, organization.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteOrganization(organizationId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s", ElephantBaseUrl, *organizationId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateSubOrganization(organizationId *string, subOrganization *Organization) (*string, error) {
	rb, _ := json.Marshal(subOrganization)

	req, _ := retryablehttp.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/sub-organizations", ElephantBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetSubOrganization(organizationId, subOrganizationId *string) (*Organization, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/sub-organizations/%s", ElephantBaseUrl, *organizationId, *subOrganizationId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	organization := Organization{}
	json.Unmarshal(body, &organization)

	return &organization, nil
}

func (xc *XilutionClient) GetSubOrganizations(organizationId *string, pageSize, pageNumber *int) (*FetchOrganizationsResponse, error) {
	req, _ := retryablehttp.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/sub-organizations?pageSize=%d&pageNumber=%d", ElephantBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchOrganizationsResponse := FetchOrganizationsResponse{}
	json.Unmarshal(body, &fetchOrganizationsResponse)

	return &fetchOrganizationsResponse, nil
}

func (xc *XilutionClient) UpdateSubOrganization(organizationId *string, subOrganization *Organization) error {
	rb, _ := json.Marshal(subOrganization)

	req, _ := retryablehttp.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/sub-organizations/%s", ElephantBaseUrl, *organizationId, subOrganization.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteSubOrganization(organizationId, subOrganizationId *string) error {
	req, _ := retryablehttp.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/sub-organizations/%s", ElephantBaseUrl, *organizationId, *subOrganizationId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}
