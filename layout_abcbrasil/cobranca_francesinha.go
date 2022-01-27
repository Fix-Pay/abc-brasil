package layout_abcbrasil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ConsultaFrancesinhaRequest struct {
	DataInicio  string `json:"DataInicio"`
	DataFim     string `json:"DataFim"`
	PaginaAtual int    `json:"PaginaAtual"`
}

type ConsultaFrancesinhaResponse struct {
	Status          bool      `json:"status"`
	Name            string    `json:"name"`
	EnvironmentName string    `json:"environmentName"`
	DateRequest     time.Time `json:"date"`
	Data            struct {
		Header ConsultaFrancesinhaHeader `json:"header"`
		Body   []ConsultaFrancesinhaBody `json:"body"`
	} `json:"data"`
	Infos  string `json:"infos"`
	Errors struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
		Title    string `json:"title"`
		Property string `json:"property"`
	} `json:"errors"`
}

type ConsultaFrancesinhaHeader struct {
	DataInicio                    string    `json:"DataInicio"`
	DataFim                       string    `json:"DataFim"`
	NomeCliente                   string    `json:"nomeCliente"`
	DataGeracaoRelatorio          time.Time `json:"dataGeracaoRelatorio"`
	TotalPaginas                  int       `json:"totalPaginas"`
	TotalRegistrosConsulta        int       `json:"totalRegistrosConsulta"`
	TotalRegistrosPagina          int       `json:"totalRegistrosPagina"`
	TotalMaximoRegistrosPorPagina int       `json:"TotalMaximoRegistrosPorPagina"`
	QuantidadeDiasPagina          int       `json:"quantidadeDiasPagina"`
	SucessoExecucao               bool      `json:"sucessoExecucao"`
	Mensagem                      string    `json:"mensagem"`
	PaginaAtual                   string    `json:"paginaAtual"`
}

type ConsultaFrancesinhaBody struct {
	DataMovimento          string                                    `json:"dataMovimento"`
	NomeProduto            string                                    `json:"nomeProduto"`
	TotalRegistrosData     int                                       `json:"totalRegistrosData"`
	Dados                  []ConsultaFrancesinhaDados                `json:"dados"`
	Ropade                 ConsultaFrancesinhaRodape                 `json:"ropade"`
	ResumoLancamentosConta ConsultaFrancesinhaResumoLancamentosConta `json:"resumoLancamentosConta"`
}

type ConsultaFrancesinhaDados struct {
	NomeProduto               string  `json:"nomeProduto"`
	NumeroOperacao            string  `json:"numeroOperacao"`
	SeuNumero                 string  `json:"seuNumero"`
	NossoNumero               string  `json:"nossoNumero"`
	NomeSacado                string  `json:"nomeSacado"`
	CodigoHistorico           string  `json:"codigoHistorico"`
	DescricaoHistorico        string  `json:"descricaoHistorico"`
	DataVencimento            string  `json:"dataVencimento"`
	DataReferencia            string  `json:"dataReferencia"`
	DataContabilLancar        string  `json:"dataContabilLancar"`
	ValorTitulo               float64 `json:"valorTitulo"`
	ValorDescontos            float64 `json:"valorDescontos"`
	ValorAbatimento           float64 `json:"valorAbatimento"`
	ValorEncargos             float64 `json:"valorEncargos"`
	ValorMora                 float64 `json:"valorMora"`
	ValorMulta                float64 `json:"valorMulta"`
	ValorAtualizacaoMonetaria float64 `json:"valorAtualizacaoMonetaria"`
	ValorIOF                  float64 `json:"valorIOF"`
	ValorTarifa               float64 `json:"valorTarifa"`
	ValorLancarConta          float64 `json:"valorLancarConta"`
	CodigoAgenciaLancar       string  `json:"codigoAgenciaLancar"`
	NumeroContaLancar         string  `json:"numeroContaLancar"`
}

type ConsultaFrancesinhaRodape struct {
	SaldoAnteriorProduto                    float64 `json:"saldoAnteriorProduto"`
	QuantidadeTitulosAbertosAnteriorProduto int     `json:"quantidadeTitulosAbertosAnteriorProduto"`
	ValorTotalEntrada                       float64 `json:"valorTotalEntrada"`
	QuantidadeEntrada                       int     `json:"quantidadeEntrada"`
	ValorTotalSaida                         float64 `json:"valorTotalSaida"`
	QuantidadeSaida                         int     `json:"quantidadeSaida"`
	SaldoAtual                              float64 `json:"saldoAtual"`
	QuantidadeAtual                         int     `json:"quantidadeAtual"`
}

type ConsultaFrancesinhaResumoLancamentosConta struct {
	DataReferencia string  `json:"dataReferencia"`
	Agencia        string  `json:"agencia"`
	ContaCorrente  string  `json:"contaCorrente"`
	TotalDebito    float64 `json:"totalDebito"`
	TotalCredito   float64 `json:"totalCredito"`
}

type HistoricoFrancesinha struct {
	CodigoHistorico    string `json:"codigoHistorico"`
	DescricaoHistorico string `json:"descricaoHistorico"`
}

func (e *ConsultaFrancesinhaRequest) ConsultaFrancesinha(url, token string) (ConsultaFrancesinhaResponse, error) {
	pathUrl := fmt.Sprint(`/abcbrasil.openbanking.cobranca.api/api/v1/francesinha/operation?DataInicio=`, e.DataInicio, `&DataFim=`,
		e.DataFim, `&PaginaAtual=`, e.PaginaAtual)
	url = fmt.Sprint(url, pathUrl)
	token = fmt.Sprint("Bearer ", token)
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()

	response := ConsultaFrancesinhaResponse{}
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

var historicoFrancesinhaEnum = []HistoricoFrancesinha{
	{"21", "PAGO"},
	{"22", "PAGO APÓS A BAIXA"},
	{"23", "PAGO NO CORRESPONDENTE"},
	{"25", "PAGO EM CARTÓRIO"},
	{"26", "PAGO NO CAIXA"},
	{"27", "PAGO VIA COMPENSAÇÃO"},
	{"28", "PAGO EM CARTÓRIO (CORRESPONDENTE)"},
	{"29", "PAGO VIA DOC/SPB"},
	{"42", "BAIXA PARA ACERTOS"},
	{"43", "BAIXADO POR TER SIDO PROTESTADO"},
	{"44", "BAIXADO PARA DEVOLUÇÃO"},
	{"48", "BAIXADO POR TER SIDO PAGO DIRETO AO CEDENTE"},
	{"49", "BAIXADO PARA SUBSTITUIÇÃO"},
	{"50", "BAIXADO POR FALTA DE SOLUÇÃO"},
	{"51", "BAIXADO A PEDIDO DO CEDENTE"},
	{"52", "BAIXA POR PROTESTO, NOSSO BANCO"},
	{"53", "BXA POR PROTESTO DO CORRESPONDENTE"},
	{"0101", "ENTRADA"},
	{"0102", "ENTRADA P/ TRANSF. DE CARTEIRA"},
	{"0103", "ENTRADA DE COBRANÇA EXPRESSA NA BXA"},
	{"0110", "BAIXA P/ TRANSF. DE CARTEIRA"},
	{"0125", "PERDÃO DE DÍVIDA"},
	{"0130", "TÍTULO ENVIADO AO CARTÓRIO"},
	{"0133", "T.C.O."},
	{"0134", "ADIANTAMENTO S/ OPERAÇÃO T.D."},
	{"0144", "AGENDAMENTO DE CARTÓRIO"},
	{"0145", "ALTERAÇÃO DE VENCIMENTO"},
	{"0147", "CANCEL AGENDAMENTO DE DEVOLUÇÃO"},
	{"0148", "CONCESSÃO DE ABATIMENTO"},
	{"0149", "CANCEL ABATIMENTO"},
	{"0150", "SUSTAÇÃO DE PROTESTO"},
	{"0151", "SUSTAR PROTESTO E ALTERAR VENCTO"},
	{"0154", "CONCESSÃO DE DESCONTO"},
	{"0155", "CANCELAMENTO DE DESCONTO"},
	{"0190", "ESTORNO DE LIQUIDAÇÃO"},
	{"0191", "ESTORNO DE BAIXA"},
	{"0192", "ESTORNO DE LIQUIDAÇÃO PARCIAL"},
	{"0193", "ESTORNO DE PRORROGAÇÃO DE DESCONTO"},
	{"0198", "REPASSE DO VALOR DE CUSTAS CARTÓRIO"},
	{"0199", "MANUTENÇÃO DE TÍTULOS EM CARTEIRA"},
	{"0201", "SUSTAR PROTESTO E BAIXAR O TÍTULO"},
	{"0230", "PROTESTO ENVIADO A CARTÓRIO/CORRESP TARIFA"},
	{"0250", "TARIFA SUSTAÇÃO PROTESTO NO CORRESP"},
	{"0253", "BAIXA POR PROTESTO APÓS BAIXA NORMAL"},
	{"0255", "EMISSÃO DE CARTA DE ANUÊNCIA"},
	{"0261", "TAR. MENSAL - PROT/SUST MÊS ANTERIOR"},
	{"0262", "TAR. MENSAL - ENTRADAS MÊS ANTERIOR"},
	{"0263", "TAR. MENSAL - BAIXAS MÊS ANTERIOR"},
	{"0264", "TAR. MENSAL - INSTRUÇÕES MÊS ANTERIOR"},
}

func GetHistoricoFrancesinhaEnum() []HistoricoFrancesinha {
	return historicoFrancesinhaEnum
}
