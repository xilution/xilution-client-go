package xilution

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test__CreateOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organization := buildTestOrganization()
	location := gofakeit.URL()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusCreated,
		Body:       r,
		Header:     map[string][]string{"Location": {location}},
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.CreateOrganization(organization)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreateOrganization__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organization := buildTestOrganization()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.CreateOrganization(organization)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organization := buildTestOrganization()

	json, _ := json.Marshal(&organization)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetOrganization(organization.ID)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &organization, resp)
}

func Test__GetOrganization__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organization := buildTestOrganization()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetOrganization(organization.ID)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetOrganizations__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizations := []Organization{
		buildTestOrganization(),
		buildTestOrganization(),
		buildTestOrganization(),
	}
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)
	totalPages := gofakeit.Number(0, 500)
	numberOfElements := gofakeit.Number(0, 500)
	totalElements := gofakeit.Number(0, 500)
	firstPage := gofakeit.Bool()
	lastPage := gofakeit.Bool()
	fetchOrganizationsResponse := FetchOrganizationsResponse{
		Content:          organizations,
		PageSize:         pageSize,
		PageNumber:       pageNumber,
		TotalPages:       totalPages,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		FirstPage:        firstPage,
		LastPage:         lastPage,
	}

	json, _ := json.Marshal(&fetchOrganizationsResponse)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetOrganizations(pageSize, pageNumber)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &fetchOrganizationsResponse, resp)
}

func Test__GetOrganizations__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetOrganizations(pageSize, pageNumber)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__UpdateOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organization := buildTestOrganization()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateOrganization(organization)

	assert.Nil(t, err)
}

func Test__UpdateOrganization__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organization := buildTestOrganization()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateOrganization(organization)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__DeleteOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeleteOrganization(organizationId)

	assert.Nil(t, err)
}

func Test__DeleteOrganization__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeleteOrganization(organizationId)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__CreateSubOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganization := buildTestOrganization()
	location := gofakeit.URL()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusCreated,
		Body:       r,
		Header:     map[string][]string{"Location": {location}},
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.CreateSubOrganization(organizationId, subOrganization)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreateSubOrganization__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganization := buildTestOrganization()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.CreateSubOrganization(organizationId, subOrganization)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetSubOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganization := buildTestOrganization()

	json, _ := json.Marshal(&subOrganization)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetSubOrganization(organizationId, subOrganization.ID)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &subOrganization, resp)
}

func Test__GetSubOrganization__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganization := buildTestOrganization()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetSubOrganization(organizationId, subOrganization.ID)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetSubOrganizations__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganizations := []Organization{
		buildTestOrganization(),
		buildTestOrganization(),
		buildTestOrganization(),
	}
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)
	totalPages := gofakeit.Number(0, 500)
	numberOfElements := gofakeit.Number(0, 500)
	totalElements := gofakeit.Number(0, 500)
	firstPage := gofakeit.Bool()
	lastPage := gofakeit.Bool()
	fetchOrganizationsResponse := FetchOrganizationsResponse{
		Content:          subOrganizations,
		PageSize:         pageSize,
		PageNumber:       pageNumber,
		TotalPages:       totalPages,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		FirstPage:        firstPage,
		LastPage:         lastPage,
	}

	json, _ := json.Marshal(&fetchOrganizationsResponse)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetSubOrganizations(organizationId, pageSize, pageNumber)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &fetchOrganizationsResponse, resp)
}

func Test__GetSubOrganizations__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetSubOrganizations(organizationId, pageSize, pageNumber)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__UpdateSubOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganization := buildTestOrganization()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateSubOrganization(organizationId, subOrganization)

	assert.Nil(t, err)
}

func Test__UpdateSubOrganization__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganization := buildTestOrganization()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateSubOrganization(organizationId, subOrganization)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__DeleteSubOrganization__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganizationId := buildTestId()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeleteSubOrganization(organizationId, subOrganizationId)

	assert.Nil(t, err)
}

func Test__DeleteSubOrganization__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	subOrganizationId := buildTestId()

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeleteSubOrganization(organizationId, subOrganizationId)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
