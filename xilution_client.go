package xilution

//go:generate mockgen -source=$GOFILE -destination=xilution_client_mock.go -package=xilution

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ProductUrl string

const (
	ElephantBaseUrl ProductUrl = "https://elephant.basics.api.xilution.com"
	RhinoBaseUrl    ProductUrl = "https://rhino.basics.api.xilution.com"
	HippoBaseUrl    ProductUrl = "https://hippo.basics.api.xilution.com"
	ZebraBaseUrl    ProductUrl = "https://zebra.basics.api.xilution.com"
	SwanBaseUrl     ProductUrl = "https://swan.basics.api.xilution.com"
	KangarooBaseUrl ProductUrl = "https://kangaroo.basics.api.xilution.com"
	GazelleBaseUrl  ProductUrl = "https://gazelle.basics.api.xilution.com"
	GiraffeBaseUrl  ProductUrl = "https://giraffe.basics.api.xilution.com"
	PenguinBaseUrl  ProductUrl = "https://penguin.cms.api.xilution.com"
	CoyoteBaseUrl   ProductUrl = "https://coyote.content-delivery.api.xilution.com"
	FoxBaseUrl      ProductUrl = "https://fox.integration.api.xilution.com"
)

// HTTPClient interface
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	IHttpClientImpl IHttpClient
)

func init() {
	IHttpClientImpl = &http.Client{Timeout: 30 * time.Second}
}

// XilutionClient -
type XilutionClient struct {
	ClientId       *string
	OrganizationId *string
	HttpClient     IHttpClient
	Token          string
}

// AuthResponse -
type AuthResponse struct {
	Token string `json:"access_token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// NewXilutionClient -
func NewXilutionClient(clientId, organizationId, username, password *string) (*XilutionClient, error) {
	xc := XilutionClient{
		HttpClient:     IHttpClientImpl,
		ClientId:       clientId,
		OrganizationId: organizationId,
	}

	grantType := "password"
	scope := "read write"

	if clientId != nil && organizationId != nil && username != nil && password != nil {
		data := fmt.Sprintf("grant_type=%s&client_id=%s&password=%s&username=%s&scope=%s", grantType, *clientId, *password, *username, scope)

		// authenticate
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/oauth/token", ZebraBaseUrl, *organizationId), strings.NewReader(data))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data)))

		res, err := xc.HttpClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		if res.StatusCode != http.StatusOK {
			return nil, handleErrorResponse(body)
		}

		// parse response body
		ar := AuthResponse{}
		json.Unmarshal(body, &ar)

		xc.Token = ar.Token
	}

	return &xc, nil
}

func handleErrorResponse(body []byte) error {
	er := ErrorResponse{}
	json.Unmarshal(body, &er)
	return fmt.Errorf(er.Message)
}

func (xc *XilutionClient) doGetRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", xc.Token))
	req.Header.Add("Content-Type", "application/json")

	res, err := xc.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, handleErrorResponse(body)
	}

	return body, err
}

func (xc *XilutionClient) doCreateRequest(req *http.Request) (*string, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", xc.Token))
	req.Header.Add("Content-Type", "application/json")

	res, err := xc.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	location := res.Header.Get("Location")

	if res.StatusCode != http.StatusCreated {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return nil, handleErrorResponse(body)
	}

	return &location, err
}

func (xc *XilutionClient) doNoContentRequest(req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", xc.Token))
	req.Header.Add("Content-Type", "application/json")

	res, err := xc.HttpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return handleErrorResponse(body)
	}

	return nil
}
