package coster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git-rbi.jatismobile.com/jatis_electrolux/electrolux-crm/models"
	"github.com/labstack/echo/v4"
)

type CosterClient interface {
	FindMessageByTemplate(ctx echo.Context, templateId string) (*Template, error)
}

type costerClient struct {
	host string
}

func NewCosterClient() CosterClient {
	host := "http://192.168.217.15:8000"

	return &costerClient{
		host: host,
	}
}

func (cc *costerClient) FindMessageByTemplate(ctx echo.Context, templateId string) (*Template, error) {

	userInfo := ctx.Get("auth_token").(*models.AuthSSO)
	token := userInfo.AccessToken

	endpoint := "/message-template/detail-by-template/"
	parameterQuery := fmt.Sprintf("%s%s", endpoint, templateId)

	url := fmt.Sprintf("%s%s", cc.host, parameterQuery)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+token)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("token_bearer", "user"+token)

	httpClient := http.Client{}

	response, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var resp Template
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
