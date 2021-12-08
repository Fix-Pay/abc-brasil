package layout_abcbrasil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AbcBrasilExtractDetailsResponse struct {
	Status          bool      `json:"status"`
	Name            string    `json:"name"`
	EnvironmentName string    `json:"environmentName"`
	Date            time.Time `json:"date"`
	Data            struct {
		Protocolo   string `json:"protocolo"`
		UrlCallback string `json:"urlCallback"`
		Saldo       struct {
			DataSaldo       time.Time `json:"dataSaldo"`
			Limite          float64   `json:"limite"`
			SaldoAliberar   float64   `json:"saldoAliberar"`
			SaldoAtualCC    float64   `json:"saldoAtualCC"`
			SaldoAplicado   float64   `json:"saldoAplicado"`
			SaldoDisponivel float64   `json:"saldoDisponivel"`
			SaldoBloqueado  float64   `json:"saldoBloqueado"`
		} `json:"saldo"`
		Extrato struct {
			Lancamentos []Lancamento `json:"lancamentos"`
		} `json:"extrato"`
	} `json:"data"`
	Infos  string `json:"infos"`
	Errors struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
		Title    string `json:"title"`
		Property string `json:"property"`
	} `json:"errors"`
}

type Lancamento struct {
	DataMovto    time.Time `json:"dataMovto"`
	Descricao    string    `json:"descricao"`
	Categoria    string    `json:"categoria"`
	NumDocumento string    `json:"numDocumento"`
	Natureza     string    `json:"natureza"`
	Fidelidade   string    `json:"fidelidade"`
	Valor        float64   `json:"valor"`
	SaldoAnt     float64   `json:"saldoAnt"`
	SaldoMovto   float64   `json:"saldoMovto"`
}

func GetExtratoDetalhado(url, token, protocolo string, codCliente, numPagina int) (AbcBrasilExtractDetailsResponse, error) {
	pathUrl := fmt.Sprint(`/abcbrasil.openbanking.contacorrente.api/api/v1/extrato/detalhado/`, protocolo, `/`, codCliente, `/`, numPagina)
	url = fmt.Sprint(url, pathUrl)
	token = fmt.Sprint("Bearer ", token)
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()

	response := AbcBrasilExtractDetailsResponse{}
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
