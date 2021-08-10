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

func Test__CreatePipelinePrototype__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestPipelinePrototype()
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

	resp, err := xc.CreatePipelinePrototype(&organizationId, &user)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreatePipelinePrototype__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestPipelinePrototype()

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

	resp, err := xc.CreatePipelinePrototype(&organizationId, &user)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetPipelinePrototype__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestPipelinePrototype()

	json, _ := json.Marshal(&user)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetPipelinePrototype(&organizationId, &user.ID)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &user, resp)
}

func Test__GetPipelinePrototype__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestPipelinePrototype()

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

	resp, err := xc.GetPipelinePrototype(&organizationId, &user.ID)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetPipelinePrototypes__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	users := []PipelinePrototype{
		buildTestPipelinePrototype(),
		buildTestPipelinePrototype(),
		buildTestPipelinePrototype(),
	}
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)
	totalPages := gofakeit.Number(0, 500)
	numberOfElements := gofakeit.Number(0, 500)
	totalElements := gofakeit.Number(0, 500)
	firstPage := gofakeit.Bool()
	lastPage := gofakeit.Bool()
	fetchPipelinePrototypesResponse := FetchPipelinePrototypesResponse{
		Content:          users,
		PageSize:         pageSize,
		PageNumber:       pageNumber,
		TotalPages:       totalPages,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		FirstPage:        firstPage,
		LastPage:         lastPage,
	}

	json, _ := json.Marshal(&fetchPipelinePrototypesResponse)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetPipelinePrototypes(&organizationId, &pageSize, &pageNumber)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &fetchPipelinePrototypesResponse, resp)
}

func Test__GetPipelinePrototypes__When_doGetRequest_Fails(t *testing.T) {
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

	resp, err := xc.GetPipelinePrototypes(&organizationId, &pageSize, &pageNumber)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__UpdatePipelinePrototype__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestPipelinePrototype()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdatePipelinePrototype(&organizationId, &user)

	assert.Nil(t, err)
}

func Test__UpdatePipelinePrototype__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestPipelinePrototype()

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

	err := xc.UpdatePipelinePrototype(&organizationId, &user)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__DeletePipelinePrototype__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	userId := buildTestId()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeletePipelinePrototype(&organizationId, &userId)

	assert.Nil(t, err)
}

func Test__DeletePipelinePrototype__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	userId := buildTestId()

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

	err := xc.DeletePipelinePrototype(&organizationId, &userId)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
