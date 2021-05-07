package xilution

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test__GetCoffees__Happy_Path(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockIHttpClient(ctrl)

	coffees := []Coffee{
		{
			ID: buildTestId(),
		},
		{
			ID: buildTestId(),
		},
	}
	json, _ := json.Marshal(&coffees)
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	m.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil)

	c := Client{
		ProductUrl: ProductUrl(gofakeit.URL()),
		HttpClient: m,
		Token:      buildJwtToken(),
	}

	resp, err := c.GetCoffees()

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	// assert.EqualValues(t, Elephant, resp.ProductUrl)
}
