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
	DataMovimento          string                                      `json:"dataMovimento"`
	NomeProduto            string                                      `json:"nomeProduto"`
	TotalRegistrosData     int                                         `json:"totalRegistrosData"`
	Dados                  []ConsultaFrancesinhaDados                  `json:"dados"`
	Ropade                 ConsultaFrancesinhaRodape                   `json:"ropade"`
	ResumoLancamentosConta []ConsultaFrancesinhaResumoLancamentosConta `json:"resumoLancamentosConta"`
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
	Codigo    string `json:"codigoHistorico"`
	Descricao string `json:"descricaoHistorico"`
	Situacao  string
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
	{"21", "PAGO", "PAGO"},
	{"22", "PAGO AP??S A BAIXA", "PAGO"},
	{"23", "PAGO NO CORRESPONDENTE", "PAGO"},
	{"25", "PAGO EM CART??RIO", "PAGO"},
	{"26", "PAGO NO CAIXA", "PAGO"},
	{"27", "PAGO VIA COMPENSA????O", "PAGO"},
	{"28", "PAGO EM CART??RIO (CORRESPONDENTE)", "PAGO"},
	{"29", "PAGO VIA DOC/SPB", "PAGO"},
	{"42", "BAIXA PARA ACERTOS", "BAIXA"},
	{"43", "BAIXADO POR TER SIDO PROTESTADO", "BAIXA"},
	{"44", "BAIXADO PARA DEVOLU????O", "BAIXA"},
	{"48", "BAIXADO POR TER SIDO PAGO DIRETO AO CEDENTE", "BAIXA"},
	{"49", "BAIXADO PARA SUBSTITUI????O", "BAIXA"},
	{"50", "BAIXADO POR FALTA DE SOLU????O", "BAIXA"},
	{"51", "BAIXADO A PEDIDO DO CEDENTE", "BAIXA"},
	{"52", "BAIXA POR PROTESTO, NOSSO BANCO", "BAIXA"},
	{"53", "BXA POR PROTESTO DO CORRESPONDENTE", "BAIXA"},
	{"0101", "ENTRADA", "ENTRADA"},
	{"0102", "ENTRADA P/ TRANSF. DE CARTEIRA", "ENTRADA"},
	{"0103", "ENTRADA DE COBRAN??A EXPRESSA NA BXA", "ENTRADA"},
	{"0110", "BAIXA P/ TRANSF. DE CARTEIRA", ""},
	{"0125", "PERD??O DE D??VIDA", ""},
	{"0130", "T??TULO ENVIADO AO CART??RIO", ""},
	{"0133", "T.C.O.", ""},
	{"0134", "ADIANTAMENTO S/ OPERA????O T.D.", ""},
	{"0144", "AGENDAMENTO DE CART??RIO", ""},
	{"0145", "ALTERA????O DE VENCIMENTO", ""},
	{"0147", "CANCEL AGENDAMENTO DE DEVOLU????O", ""},
	{"0148", "CONCESS??O DE ABATIMENTO", ""},
	{"0149", "CANCEL ABATIMENTO", ""},
	{"0150", "SUSTA????O DE PROTESTO", ""},
	{"0151", "SUSTAR PROTESTO E ALTERAR VENCTO", ""},
	{"0154", "CONCESS??O DE DESCONTO", ""},
	{"0155", "CANCELAMENTO DE DESCONTO", ""},
	{"0190", "ESTORNO DE LIQUIDA????O", ""},
	{"0191", "ESTORNO DE BAIXA", ""},
	{"0192", "ESTORNO DE LIQUIDA????O PARCIAL", ""},
	{"0193", "ESTORNO DE PRORROGA????O DE DESCONTO", ""},
	{"0198", "REPASSE DO VALOR DE CUSTAS CART??RIO", ""},
	{"0199", "MANUTEN????O DE T??TULOS EM CARTEIRA", ""},
	{"0201", "SUSTAR PROTESTO E BAIXAR O T??TULO", ""},
	{"0230", "PROTESTO ENVIADO A CART??RIO/CORRESP TARIFA", ""},
	{"0250", "TARIFA SUSTA????O PROTESTO NO CORRESP", ""},
	{"0253", "BAIXA POR PROTESTO AP??S BAIXA NORMAL", ""},
	{"0255", "EMISS??O DE CARTA DE ANU??NCIA", ""},
	{"0261", "TAR. MENSAL - PROT/SUST M??S ANTERIOR", ""},
	{"0262", "TAR. MENSAL - ENTRADAS M??S ANTERIOR", ""},
	{"0263", "TAR. MENSAL - BAIXAS M??S ANTERIOR", ""},
	{"0264", "TAR. MENSAL - INSTRU????ES M??S ANTERIOR", ""},
}

func GetHistoricoFrancesinhaEnum() []HistoricoFrancesinha {
	return historicoFrancesinhaEnum
}
