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
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	organizationId string
	username       string
	password       string
)

func init() {
	organizationId = strings.Join(strings.Split(gofakeit.UUID(), "-"), "")
	username = gofakeit.Username()
	password = gofakeit.Password(true, true, true, true, false, 8)
}

func Test__NewClient__Happy_Path_2(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := gofakeit.UUID()
	json := fmt.Sprintf(`{"token": "%s"}`, token)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil)

	resp, err := NewClient(Elephant, &organizationId, &username, &password)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, Elephant, resp.ProductUrl)
	assert.EqualValues(t, token, resp.Token)
}

func Test__NewClient__When_Username_Is_Nil(t *testing.T) {
	resp, err := NewClient(Elephant, &organizationId, nil, &password)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, Elephant, resp.ProductUrl)
	assert.EqualValues(t, "", resp.Token)
}

func Test__NewClient__When_Password_Is_Nil(t *testing.T) {
	resp, err := NewClient(Elephant, &organizationId, &username, nil)

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

	resp, err := NewClient(Elephant, &organizationId, &username, &password)

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

	resp, err := NewClient(Elephant, &organizationId, &username, &password)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__doGetRequest__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockSomeHTTPClient(ctrl)
	SomeClient = m

	token := gofakeit.UUID()
	json1 := fmt.Sprintf(`{"token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       r1,
	}, nil)
	json2 := `{"foo": "bar"}`
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewClient(Elephant, &organizationId, &username, &password)

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

	token := gofakeit.UUID()
	json1 := fmt.Sprintf(`{"token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       r1,
	}, nil)
	errMsg := gofakeit.Sentence(5)
	m.EXPECT().Do(gomock.Any()).Return(nil, errors.New(errMsg)).After(first)

	c, _ := NewClient(Elephant, &organizationId, &username, &password)

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

	token := gofakeit.UUID()
	json1 := fmt.Sprintf(`{"token": "%s"}`, token)
	r1 := ioutil.NopCloser(bytes.NewReader([]byte(json1)))
	errMsg := gofakeit.Sentence(10)
	first := m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 200,
		Body:       r1,
	}, nil)
	json2 := fmt.Sprintf(`{"message": "%s"}`, errMsg)
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json2)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: 500,
		Body:       r2,
	}, nil).After(first)

	c, _ := NewClient(Elephant, &organizationId, &username, &password)

	req, _ := http.NewRequest("GET", gofakeit.URL(), nil)

	resp, err := c.doGetRequest(req)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
