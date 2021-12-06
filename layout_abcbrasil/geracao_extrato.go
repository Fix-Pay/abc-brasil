package layout_abcbrasil

import "time"

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
