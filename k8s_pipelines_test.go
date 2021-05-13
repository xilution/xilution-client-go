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

func Test__CreateK8sPipeline__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestK8sPipeline()
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

	resp, err := xc.CreateK8sPipeline(&organizationId, &user)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreateK8sPipeline__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestK8sPipeline()

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

	resp, err := xc.CreateK8sPipeline(&organizationId, &user)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetK8sPipeline__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestK8sPipeline()

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

	resp, err := xc.GetK8sPipeline(&organizationId, &user.ID)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &user, resp)
}

func Test__GetK8sPipeline__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestK8sPipeline()

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

	resp, err := xc.GetK8sPipeline(&organizationId, &user.ID)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetK8sPipelines__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	users := []K8sPipeline{
		buildTestK8sPipeline(),
		buildTestK8sPipeline(),
		buildTestK8sPipeline(),
	}
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)
	totalPages := gofakeit.Number(0, 500)
	numberOfElements := gofakeit.Number(0, 500)
	totalElements := gofakeit.Number(0, 500)
	firstPage := gofakeit.Bool()
	lastPage := gofakeit.Bool()
	fetchK8sPipelinesResponse := FetchK8sPipelinesResponse{
		Content:          users,
		PageSize:         pageSize,
		PageNumber:       pageNumber,
		TotalPages:       totalPages,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		FirstPage:        firstPage,
		LastPage:         lastPage,
	}

	json, _ := json.Marshal(&fetchK8sPipelinesResponse)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetK8sPipelines(&organizationId, &pageSize, &pageNumber)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &fetchK8sPipelinesResponse, resp)
}

func Test__GetK8sPipelines__When_doGetRequest_Fails(t *testing.T) {
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

	resp, err := xc.GetK8sPipelines(&organizationId, &pageSize, &pageNumber)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__UpdateK8sPipeline__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestK8sPipeline()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateK8sPipeline(&organizationId, &user)

	assert.Nil(t, err)
}

func Test__UpdateK8sPipeline__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	user := buildTestK8sPipeline()

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

	err := xc.UpdateK8sPipeline(&organizationId, &user)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__DeleteK8sPipeline__Happy_Path(t *testing.T) {
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

	err := xc.DeleteK8sPipeline(&organizationId, &userId)

	assert.Nil(t, err)
}

func Test__DeleteK8sPipeline__When_doNoContentRequest_Fails(t *testing.T) {
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

	err := xc.DeleteK8sPipeline(&organizationId, &userId)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
