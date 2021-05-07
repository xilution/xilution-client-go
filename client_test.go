package xilution

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	clientId       string
	organizationId string
	username       string
	password       string
)

func init() {
	clientId = buildThingId()
	organizationId = buildThingId()
	username = gofakeit.Username()
	password = gofakeit.Password(true, true, true, true, false, 8)
}

func buildThingId() string {
	return strings.Join(strings.Split(gofakeit.UUID(), "-"), "")
}

func buildJwtToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})
	tokenString, _ := token.SignedString([]byte("AllYourBase"))

	return tokenString
}

func Test__NewClient__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	resp, err := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, Elephant, resp.ProductUrl)
	assert.EqualValues(t, token, resp.Token)
}

func Test__NewClient__When_Username_Is_Nil(t *testing.T) {
	resp, err := NewClient(Elephant, &clientId, &organizationId, nil, &password)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, Elephant, resp.ProductUrl)
	assert.EqualValues(t, "", resp.Token)
}

func Test__NewClient__When_Password_Is_Nil(t *testing.T) {
	resp, err := NewClient(Elephant, &clientId, &organizationId, &username, nil)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, Elephant, resp.ProductUrl)
	assert.EqualValues(t, "", resp.Token)
}

func Test__NewClient__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg))

	resp, err := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__NewClient__When_Auth_Response_Code_Is_Not_OK(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	errMsg := gofakeit.Sentence(10)
	json := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r,
	}, nil)

	resp, err := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doGetRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

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

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, json2, resp)
}

func Test__doGetRequest__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doGetRequest__When_Response_Code_Is_Not_OK(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	errMsg := gofakeit.Sentence(10)
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doCreateRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

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

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("POST", gofakeit.URL(), nil)

	resp, err := c.doCreateRequest(req)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, location, *resp)
}

func Test__doCreateRequest__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("POST", gofakeit.URL(), nil)

	resp, err := c.doCreateRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doCreateRequest__When_Response_Code_Is_Not_Created(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	errMsg := gofakeit.Sentence(10)
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("POST", gofakeit.URL(), nil)

	resp, err := c.doCreateRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doNoContentRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

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

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("PUT", gofakeit.URL(), nil)

	resp, err := c.doNoContentRequest(req)

	assert.Nil(t, err)
	assert.Nil(t, resp)
}

func Test__doNoContentRequest__When_Auth_Request_Returns_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("PUT", gofakeit.URL(), nil)

	resp, err := c.doNoContentRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doNoContentRequest__When_Response_Code_Is_Not_Created(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := buildJwtToken()
	json1 := fmt.Sprintf(`{"access_token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	errMsg := gofakeit.Sentence(10)
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r1,
	}, nil)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewClient(Elephant, &clientId, &organizationId, &username, &password)

	req, _ := http.NewRequest("PUT", gofakeit.URL(), nil)

	resp, err := c.doNoContentRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
