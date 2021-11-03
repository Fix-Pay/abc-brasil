package abcbrasil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"strings"
)

type AccessTokenParameters struct {
	Username      string
	Password      string
	CodigoEmpresa string
	ClientID      string
	ClientSecret  string
	GrantType     string
	Url           string
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Error        string
}

func (params AccessTokenParameters) GenerateToken() AccessToken {
	acToken := AccessToken{}

	url := fmt.Sprint(params.Url, `/api/oauth/token/openbanking`)
	contentTypeValue := "application/x-www-form-urlencoded"

	data := netUrl.Values{}
	data.Set("username", params.Username)
	data.Set("password", params.Password)
	data.Set("codigoEmpresa", params.CodigoEmpresa)
	data.Set("client_id", params.ClientID)
	data.Set("client_secret", params.ClientSecret)
	data.Set("grant_type", params.GrantType)

	reader := strings.NewReader(data.Encode())
	res, err := http.Post(url, contentTypeValue, reader)

	defer res.Body.Close()
	if err != nil {
		acToken.Error = err.Error()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		acToken.Error = err.Error()
	}

	err = json.Unmarshal(body, &acToken)
	if err != nil {
		acToken.Error = err.Error()
	}
	return acToken
}
