package xilution

//go:generate mockgen -source=$GOFILE -destination=client_mock.go -package=xilution

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ProductUrl string

const (
	Elephant ProductUrl = "https://elephant.basics.api.xilution.com"
	Zebra    ProductUrl = "https://zebra.basics.api.xilution.com"
)

// HTTPClient interface
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	IHttpClientImpl IHttpClient
)

func init() {
	IHttpClientImpl = &http.Client{Timeout: 10 * time.Second}
}

// Client -
type Client struct {
	ProductUrl     ProductUrl
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

// NewClient -
func NewClient(host ProductUrl, clientId, organizationId, username, password *string) (*Client, error) {
	c := Client{
		HttpClient:     IHttpClientImpl,
		ProductUrl:     host,
		ClientId:       clientId,
		OrganizationId: organizationId,
	}

	if organizationId != nil && username != nil && password != nil {
		// form request body
		data := url.Values{}
		data.Set("grant_type", "password")
		data.Set("username", *username)
		data.Set("password", *password)
		data.Set("client_id", *clientId)
		data.Set("scope", "read write")

		// authenticate
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/oauth/token", *organizationId, Zebra), strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		res, err := c.HttpClient.Do(req)
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

		c.Token = ar.Token
	}

	return &c, nil
}

func handleErrorResponse(body []byte) error {
	er := ErrorResponse{}
	json.Unmarshal(body, &er)
	return fmt.Errorf(er.Message)
}

func (c *Client) doGetRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)
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

func (c *Client) doCreateRequest(req *http.Request) (*string, error) {
	req.Header.Set("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)
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

func (c *Client) doNoContentRequest(req *http.Request) (*string, error) {
	req.Header.Set("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusNoContent {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return nil, handleErrorResponse(body)
	}

	return nil, err
}
