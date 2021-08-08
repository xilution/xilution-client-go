package xilution

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
)

var (
	organizationId string
	grantType      string
	scope          string
	clientId       string
	clientSecret   string
	username       string
	password       string
	token          string
)

func init() {
	organizationId = buildTestId()
	grantType = "password"
	scope = "read write"
	clientId = buildTestId()
	clientSecret = buildTestId()
	username = gofakeit.Username()
	password = gofakeit.Password(true, true, true, true, false, 8)
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}

func Test__NewXilutionClientWithToken__Happy_Path(t *testing.T) {
	resp, err := NewXilutionClientWithToken(&token)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, token, resp.Token)
}

func Test__NewXilutionClient__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	resp, err := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, token, resp.Token)
}

func Test__NewClient__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg))

	resp, err := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__NewClient__When_Auth_Response_Code_Is_Not_OK(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	resp, err := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doGetRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	json2 := `{"foo": "bar"}`
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, json2, resp)
}

func Test__doGetRequest__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doGetRequest__When_Response_Code_Is_Not_OK(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(10)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doCreateRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte("")))
	location := gofakeit.URL()
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusCreated,
		Body:       r2,
		Header:     map[string][]string{"Location": {location}},
	}, nil).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("POST", gofakeit.URL(), nil)

	resp, err := c.doCreateRequest(req)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, location, *resp)
}

func Test__doCreateRequest__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("POST", gofakeit.URL(), nil)

	resp, err := c.doCreateRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doCreateRequest__When_Response_Code_Is_Not_Created(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(10)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("POST", gofakeit.URL(), nil)

	resp, err := c.doCreateRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doNoContentRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("PUT", gofakeit.URL(), nil)

	err := c.doNoContentRequest(req)

	assert.Nil(t, err)
}

func Test__doNoContentRequest__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("PUT", gofakeit.URL(), nil)

	err := c.doNoContentRequest(req)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doNoContentRequest__When_Response_Code_Is_Not_Created(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)
	IHttpClientImpl = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(10)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewXilutionClient(&organizationId, &grantType, &scope, &clientId, &clientSecret, &username, &password)

	req, _ := retryablehttp.NewRequest("PUT", gofakeit.URL(), nil)

	err := c.doNoContentRequest(req)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
