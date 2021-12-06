package layout_abcbrasil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type AbcBrasilGeracaoExtratoRequest struct {
	CodCliente  int       `json:"CodCliente"`
	Agencia     string    `json:"Agencia"`
	Conta       string    `json:"Conta"`
	DataInicio  time.Time `json:"DtInicio"`
	DataFim     time.Time `json:"dtFim"`
	UrlCallback string    `json:"UrlCallback"`
}

type AbcBrasilGeracaoExtratoResponse struct {
	Status          bool      `json:"status"`
	Name            string    `json:"name"`
	EnvironmentName string    `json:"environmentName"`
	DateRequest     time.Time `json:"date"`
	Data            struct {
		Protocolo         string `json:"protocolo"`
		CodeStatus        int    `json:"codeStatus"`
		StatusDescription string `json:"statusDescription"`
	} `json:"data"`
	Infos  string `json:"infos"`
	Errors struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
		Title    string `json:"title"`
		Property string `json:"property"`
	} `json:"errors"`
}

type AbcBrasilGeracaoExtratoCallback struct {
	Protocolo                 string    `json:"Protocolo"`
	DataValidade              time.Time `json:"DataValidade"`
	QuantidadeRegistros       int       `json:"QtdRegistros"`
	QuantidadeRegistrosPagina int       `json:"QtdRegistrosPagina"`
	QuantidadePaginas         int       `json:"QtdPaginas"`
}

func (r *AbcBrasilBoletoRequest) GeracaoExtrato(url, token string) (AbcBrasilGeracaoExtratoResponse, error) {
	response := AbcBrasilGeracaoExtratoResponse{}

	pathUrl := `/abcbrasil.openbanking.contacorrente.api/api/v1/extrato/gerar`

	url = fmt.Sprint(url, pathUrl)
	token = fmt.Sprint("Bearer ", token)
	method := "POST"
	client := &http.Client{}

	b, err := json.Marshal(r)
	if err != nil {
		return response, err
	}

	boletoJson := strings.NewReader(string(b))
	req, err := http.NewRequest(method, url, boletoJson)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, err
}
