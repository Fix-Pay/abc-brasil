package v1

import (
	"encoding/json"
	"fmt"
	"github.com/Fix-Pay/models"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"strings"
)

type TokenAcesso struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
}

func GerarToken() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		var params models.ABCBrasilParameters
		params.FindParameters()

		url := fmt.Sprint(params.UrlHomologacao, `/api/oauth/token/openbanking`)
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
			ctx.JSON(fiber.Map{"error": "Falha na requisicao."})
			return err
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			ctx.JSON(fiber.Map{"error": "Falha na requisicao."})
			return err
		}

		accessToken := TokenAcesso{}
		err = json.Unmarshal(body, &accessToken)
		if err != nil {
			ctx.Status(fiber.StatusUnauthorized)
			ctx.JSON(fiber.Map{"error": "Token Inválido! :("})
			return err
		}
		ctx.JSON(fiber.Map{"data": accessToken})
		return err
	}
}
