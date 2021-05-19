package xilution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (xc *XilutionClient) CreateGitAccount(organizationId *string, gitAccount *GitAccount) (*string, error) {
	rb, _ := json.Marshal(gitAccount)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/git-accounts", SwanBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetGitAccount(organizationId *string, gitAccountId *string) (*GitAccount, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/git-accounts/%s", SwanBaseUrl, *organizationId, *gitAccountId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	gitaccount := GitAccount{}
	json.Unmarshal(body, &gitaccount)

	return &gitaccount, nil
}

func (xc *XilutionClient) GetGitAccounts(organizationId *string, pageSize, pageNumber *int) (*FetchGitAccountsResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/git-accounts?pageSize=%d&pageNumber=%d", SwanBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchGitAccountsResponse := FetchGitAccountsResponse{}
	json.Unmarshal(body, &fetchGitAccountsResponse)

	return &fetchGitAccountsResponse, nil
}

func (xc *XilutionClient) UpdateGitAccount(organizationId *string, gitAccount *GitAccount) error {
	rb, _ := json.Marshal(gitAccount)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/git-accounts/%s", SwanBaseUrl, *organizationId, gitAccount.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteGitAccount(organizationId *string, gitAccountId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/git-accounts/%s", SwanBaseUrl, *organizationId, *gitAccountId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateGitRepo(organizationId *string, gitRepo *GitRepo) (*string, error) {
	rb, _ := json.Marshal(gitRepo)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/git-repos", SwanBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetGitRepo(organizationId, gitRepoId *string) (*GitRepo, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/git-repos/%s", SwanBaseUrl, *organizationId, *gitRepoId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	gitRepo := GitRepo{}
	json.Unmarshal(body, &gitRepo)

	return &gitRepo, nil
}

func (xc *XilutionClient) GetGitRepos(organizationId *string, pageSize, pageNumber *int) (*FetchGitReposResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/git-repos?pageSize=%d&pageNumber=%d", SwanBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchGitReposResponse := FetchGitReposResponse{}
	json.Unmarshal(body, &fetchGitReposResponse)

	return &fetchGitReposResponse, nil
}

func (xc *XilutionClient) UpdateGitRepo(organizationId *string, gitRepo *GitRepo) error {
	rb, _ := json.Marshal(gitRepo)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/organizations/%s/git-repos/%s", SwanBaseUrl, *organizationId, gitRepo.ID), strings.NewReader(string(rb)))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) DeleteGitRepo(organizationId, gitRepoId *string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/organizations/%s/git-repos/%s", SwanBaseUrl, *organizationId, *gitRepoId), strings.NewReader(string("")))

	err := xc.doNoContentRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XilutionClient) CreateGitRepoEvent(organizationId *string, gitRepoEvent *GitRepoEvent) (*string, error) {
	rb, _ := json.Marshal(gitRepoEvent)

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/git-repo-events", SwanBaseUrl, *organizationId), strings.NewReader(string(rb)))

	location, err := xc.doCreateRequest(req)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (xc *XilutionClient) GetGitRepoEvent(organizationId, eventId *string) (*GitRepoEvent, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/git-repo-events/%s", SwanBaseUrl, *organizationId, *eventId), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	gitRepoEvent := GitRepoEvent{}
	json.Unmarshal(body, &gitRepoEvent)

	return &gitRepoEvent, nil
}

func (xc *XilutionClient) GetGitRepoEvents(organizationId *string, pageSize, pageNumber *int) (*FetchGitRepoEventsResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%s/git-repo-events?pageSize=%d&pageNumber=%d", SwanBaseUrl, *organizationId, *pageSize, *pageNumber), nil)

	body, err := xc.doGetRequest(req)
	if err != nil {
		return nil, err
	}

	fetchGitRepoEventsResponse := FetchGitRepoEventsResponse{}
	json.Unmarshal(body, &fetchGitRepoEventsResponse)

	return &fetchGitRepoEventsResponse, nil
}
