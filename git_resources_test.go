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

func Test__CreateGitAccount__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccount := buildTestGitAccount()
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

	resp, err := xc.CreateGitAccount(&organizationId, &gitAccount)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreateGitAccount__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccount := buildTestGitAccount()

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

	resp, err := xc.CreateGitAccount(&organizationId, &gitAccount)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetGitAccount__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccount := buildTestGitAccount()

	json, _ := json.Marshal(&gitAccount)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetGitAccount(&organizationId, &gitAccount.ID)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &gitAccount, resp)
}

func Test__GetGitAccount__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccount := buildTestGitAccount()

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

	resp, err := xc.GetGitAccount(&organizationId, &gitAccount.ID)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetGitAccounts__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccounts := []GitAccount{
		buildTestGitAccount(),
		buildTestGitAccount(),
		buildTestGitAccount(),
	}
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)
	totalPages := gofakeit.Number(0, 500)
	numberOfElements := gofakeit.Number(0, 500)
	totalElements := gofakeit.Number(0, 500)
	firstPage := gofakeit.Bool()
	lastPage := gofakeit.Bool()
	fetchGitAccountsResponse := FetchGitAccountsResponse{
		Content:          gitAccounts,
		PageSize:         pageSize,
		PageNumber:       pageNumber,
		TotalPages:       totalPages,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		FirstPage:        firstPage,
		LastPage:         lastPage,
	}

	json, _ := json.Marshal(&fetchGitAccountsResponse)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetGitAccounts(&organizationId, &pageSize, &pageNumber)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &fetchGitAccountsResponse, resp)
}

func Test__GetGitAccounts__When_doGetRequest_Fails(t *testing.T) {
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

	resp, err := xc.GetGitAccounts(&organizationId, &pageSize, &pageNumber)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__UpdateGitAccount__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccount := buildTestGitAccount()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateGitAccount(&organizationId, &gitAccount)

	assert.Nil(t, err)
}

func Test__UpdateGitAccount__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccount := buildTestGitAccount()

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

	err := xc.UpdateGitAccount(&organizationId, &gitAccount)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__DeleteGitAccount__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeleteGitAccount(&organizationId, &gitAccountId)

	assert.Nil(t, err)
}

func Test__DeleteGitAccount__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()

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

	err := xc.DeleteGitAccount(&organizationId, &gitAccountId)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__CreateGitRepo__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepo := buildTestGitRepo()
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

	resp, err := xc.CreateGitRepo(&organizationId, &gitAccountId, &gitRepo)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreateGitRepo__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepo := buildTestGitRepo()

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

	resp, err := xc.CreateGitRepo(&organizationId, &gitAccountId, &gitRepo)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetGitRepo__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepo := buildTestGitRepo()

	json, _ := json.Marshal(&gitRepo)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetGitRepo(&organizationId, &gitAccountId, &gitRepo.ID)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &gitRepo, resp)
}

func Test__GetGitRepo__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepo := buildTestGitRepo()

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

	resp, err := xc.GetGitRepo(&organizationId, &gitAccountId, &gitRepo.ID)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__GetGitRepos__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepos := []GitRepo{
		buildTestGitRepo(),
		buildTestGitRepo(),
		buildTestGitRepo(),
	}
	pageSize := gofakeit.Number(0, 100)
	pageNumber := gofakeit.Number(0, 500)
	totalPages := gofakeit.Number(0, 500)
	numberOfElements := gofakeit.Number(0, 500)
	totalElements := gofakeit.Number(0, 500)
	firstPage := gofakeit.Bool()
	lastPage := gofakeit.Bool()
	fetchGitReposResponse := FetchGitReposResponse{
		Content:          gitRepos,
		PageSize:         pageSize,
		PageNumber:       pageNumber,
		TotalPages:       totalPages,
		NumberOfElements: numberOfElements,
		TotalElements:    totalElements,
		FirstPage:        firstPage,
		LastPage:         lastPage,
	}

	json, _ := json.Marshal(&fetchGitReposResponse)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := xc.GetGitRepos(&organizationId, &gitAccountId, &pageSize, &pageNumber)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, &fetchGitReposResponse, resp)
}

func Test__GetGitRepos__When_doGetRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
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

	resp, err := xc.GetGitRepos(&organizationId, &gitAccountId, &pageSize, &pageNumber)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__UpdateGitRepo__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepo := buildTestGitRepo()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.UpdateGitRepo(&organizationId, &gitAccountId, &gitRepo)

	assert.Nil(t, err)
}

func Test__UpdateGitRepo__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepo := buildTestGitRepo()

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

	err := xc.UpdateGitRepo(&organizationId, &gitAccountId, &gitRepo)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}

func Test__DeleteGitRepo__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepoId := buildTestId()

	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusNoContent,
		Body:       r,
	}, nil)

	xc := XilutionClient{
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	err := xc.DeleteGitRepo(&organizationId, &gitAccountId, &gitRepoId)

	assert.Nil(t, err)
}

func Test__DeleteGitRepo__When_doNoContentRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepoId := buildTestId()

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

	err := xc.DeleteGitRepo(&organizationId, &gitAccountId, &gitRepoId)

	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
func Test__CreateGitRepoEvent__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepoId := buildTestId()
	gitRepoEvent := buildTestGitRepoEvent()
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

	resp, err := xc.CreateGitRepoEvent(&organizationId, &gitAccountId, &gitRepoId, &gitRepoEvent)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, location, *resp)
}

func Test__CreateGitRepoEvent__When_doCreateRequest_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	organizationId := buildTestId()
	gitAccountId := buildTestId()
	gitRepoId := buildTestId()
	gitRepoEvent := buildTestGitRepoEvent()

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

	resp, err := xc.CreateGitRepoEvent(&organizationId, &gitAccountId, &gitRepoId, &gitRepoEvent)

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, errMsg, err.Error())
}
