package xilution

//go:generate mockgen -source=$GOFILE -destination=client_mock.go -package=xilution

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ProductUrl string

const (
	Elephant ProductUrl = "https://elephant.basics.api.xilution.com"
	Zebra    ProductUrl = "https://zebra.basics.api.xilution.com"
)

// HTTPClient interface
type SomeHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	SomeClient SomeHTTPClient
)

func init() {
	SomeClient = &http.Client{Timeout: 10 * time.Second}
}

// Client -
type Client struct {
	ProductUrl ProductUrl
	HTTPClient SomeHTTPClient
	Token      string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// NewClient -
func NewClient(host ProductUrl, organizationId, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: SomeClient,
		ProductUrl: host,
	}

	if organizationId != nil && username != nil && password != nil {
		// form request body
		// TODO - fix this
		rb, _ := json.Marshal(AuthStruct{
			Username: *username,
			Password: *password,
		})

		// authenticate
		req, _ := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/oauth/token", *organizationId, Zebra), strings.NewReader(string(rb)))

		res, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		if res.StatusCode != http.StatusOK {
			er := ErrorResponse{}
			json.Unmarshal(body, &er)
			return nil, fmt.Errorf(er.Message)
		}

		// parse response body
		ar := AuthResponse{}
		json.Unmarshal(body, &ar)

		c.Token = ar.Token
	}

	return &c, nil
}

func (c *Client) doGetRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		er := ErrorResponse{}
		json.Unmarshal(body, &er)
		return nil, fmt.Errorf(er.Message)
	}

	return body, err
}
