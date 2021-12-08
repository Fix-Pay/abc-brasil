package layout_abcbrasil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type AbcBrasilBoletoRequest struct {
	CodCliente           string  `json:"codCliente"`
	NumContaHeader       string  `json:"numContaHeader"`
	NumCarteira          string  `json:"numCarteira"`
	NossoNumero          string  `json:"nossoNumero"`
	CodBanco             string  `json:"codBanco"`
	CodModalBancos       string  `json:"codModalBancos"`
	NossoNumeroBancos    string  `json:"nossoNumeroBancos"`
	CodEspecieDoc        string  `json:"codEspecieDoc"`
	ValorNominal         float64 `json:"vlrNominal"`
	ValorAbatimento      float64 `json:"vlrAbatimento"`
	DataEmissao          string  `json:"dataEmissao"`
	DataVencimento       string  `json:"dataVencimento"`
	SeuNumero            string  `json:"seuNumero"`
	Aceite               string  `json:"aceite"`
	CnpjCPF              string  `json:"cnpjCPF"`
	TipoPessoa           string  `json:"tipoPessoa"`
	Nome                 string  `json:"nome"`
	Endereco             string  `json:"endereco"`
	Bairro               string  `json:"bairro"`
	Cidade               string  `json:"cidade"`
	Uf                   string  `json:"uf"`
	Cep                  string  `json:"cep"`
	Email                string  `json:"email"`
	Ddd                  string  `json:"ddd"`
	Telefone             string  `json:"telefone"`
	CampoLivre           string  `json:"campoLivre"`
	TipoPessoaSacador    string  `json:"tipoPessoaSacador"`
	CnpjCPFSacador       string  `json:"cnpjCPFSacador"`
	NomeSacador          string  `json:"nomeSacador"`
	EnderecoSacador      string  `json:"enderecoSacador"`
	BairroSacador        string  `json:"bairroSacador"`
	CepSacador           string  `json:"cepSacador"`
	CidadeSacador        string  `json:"cidadeSacador"`
	UfSacador            string  `json:"ufSacador"`
	Mensagem1            string  `json:"mensagem1"`
	Mensagem2            string  `json:"mensagem2"`
	Mensagem3            string  `json:"mensagem3"`
	Mensagem4            string  `json:"mensagem4"`
	Mensagem5            string  `json:"mensagem5"`
	CodDesconto1         string  `json:"codDesconto1"`
	ValorDesconto1       float64 `json:"vlrDesconto1"`
	TaxaDesconto1        float64 `json:"txDesconto1"`
	DataDesconto1        string  `json:"dataDesconto1"`
	CodDesconto2         string  `json:"codDesconto2"`
	ValorDesconto2       float64 `json:"vlrDesconto2"`
	TaxaDesconto2        float64 `json:"txDesconto2"`
	DataDesconto2        string  `json:"dataDesconto2"`
	CodDesconto3         string  `json:"codDesconto3"`
	ValorDesconto3       float64 `json:"vlrDesconto3"`
	TaxaDesconto3        float64 `json:"txDesconto3"`
	DataDesconto3        string  `json:"dataDesconto3"`
	CodMulta             string  `json:"codMulta"`
	DataMulta            string  `json:"dataMulta"`
	TaxaMulta            float64 `json:"txMulta"`
	ValorMulta           float64 `json:"vlrMulta"`
	CodMora              string  `json:"codMora"`
	DataMora             string  `json:"dataMora"`
	TaxaMora             float64 `json:"txMora"`
	ValorMora            float64 `json:"vlrMora"`
	PossuiAgenda         string  `json:"possuiAgenda"`
	TipoAgendamento      string  `json:"tipoAgendamento"`
	CriterioDias         string  `json:"criterioDias"`
	NumDiasAgenda        int     `json:"numDiasAgenda"`
	CodIndice            string  `json:"codIndice"`
	IndPagtoParcial      string  `json:"indPagtoParcial"`
	QtdPagtosParciais    int     `json:"qtdPagtosParciais"`
	TipoValorPercMinimo  string  `json:"tipoValorPercMinimo"`
	ValorPercMinimo      float64 `json:"vlrPercMinimo"`
	TipoValorPercMaximo  string  `json:"tipoValorPercMaximo"`
	ValorPercMaximo      float64 `json:"vlrPercMaximo"`
	TipoAutRecDivergente string  `json:"tipoAutRecDivergente"`
}

type AbcBrasilBoletoResponse struct {
	Sucesso        bool   `json:"sucesso"`
	Mensagem       string `json:"mensagem"`
	PdfBinario     string `json:"pdfBinario"`
	Imagem         string `json:"imagem"`
	LinhaDigitavel string `json:"linhaDigitavel"`
	CodigoDeBarras string `json:"codigoDeBarras"`
	Desenvolvedor  struct {
		CodigoRetorno string `json:"codigoRetorno"`
		Mensagem      string `json:"mensagem"`
	} `json:"desenvolvedor"`
}

type RetornoErro struct {
	Sucesso        bool        `json:"sucesso"`
	Mensagem       string      `json:"mensagem"`
	PdfBinario     interface{} `json:"pdf_binario"`
	Imagem         interface{} `json:"imagem"`
	LinhaDigitavel interface{} `json:"linha_digitavel"`
	CodigoDeBarras interface{} `json:"codigo_de_barras"`
	Desenvolvedor  struct {
		CodigoRetorno string `json:"codigo_retorno"`
		Mensagem      string `json:"mensagem"`
	} `json:"desenvolvedor"`
}

func (simplificado *AbcBrasilBoletoRequest) GerarBoleto(url, token string, isSimplificado bool) (AbcBrasilBoletoResponse, error) {
	var pathUrl string
	if isSimplificado {
		pathUrl = `/ABCDigital.BoletoOnline/api/v1.0/BoletoSimplificado`
	} else {
		pathUrl = `/ABCDigital.BoletoOnline/api/v1.0/BoletoPDF`
	}

	url = fmt.Sprint(url, pathUrl)
	token = fmt.Sprint("Bearer ", token)
	method := "POST"
	client := &http.Client{}

	returnSuccess := AbcBrasilBoletoResponse{}
	b, err := json.Marshal(simplificado)
	if err != nil {
		return returnSuccess, err
	}

	boletoJson := strings.NewReader(string(b))
	req, err := http.NewRequest(method, url, boletoJson)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return returnSuccess, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return returnSuccess, err
	}

	err = json.Unmarshal(body, &returnSuccess)
	if err != nil {
		return returnSuccess, err
	}
	return returnSuccess, err
}

func CalcularDigitoVerificador(agencia, carteira, nossoNumero string) string {
	digito, _ := digitoVerificador(agencia, carteira, nossoNumero)
	return fmt.Sprint(nossoNumero, digito)
}

func digitoVerificador(agencia, carteira, nossoNumero string) (int, error) {
	peso := [17]int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2}

	if len(agencia) > 4 {
		return -1, errors.New("Campo agência inválido.")
	} else {
		digito, _ := strconv.Atoi(agencia)
		agencia = fmt.Sprintf("%04d", digito)
	}

	if len(carteira) > 3 {
		return -1, errors.New("Campo carteira inválido.")
	} else {
		digito, _ := strconv.Atoi(carteira)
		carteira = fmt.Sprintf("%03d", digito)
	}

	if len(nossoNumero) > 10 {
		return -1, errors.New("Campo nosso número já possui dígito verificador.")
	} else {
		digito, _ := strconv.Atoi(nossoNumero)
		nossoNumero = fmt.Sprintf("%010d", digito)
	}

	array := addWordInArray(agencia, []int{})
	array = addWordInArray(carteira, array)
	array = addWordInArray(nossoNumero, array)

	total := 0
	var i int
	for i = 0; i < len(peso); i++ {
		aux := array[i] * peso[i]
		total += (aux / 10) + (aux % 10)
	}

	total = 10 - (total % 10)

	if total != 10 {
		return total, nil
	} else {
		return 0, nil
	}
}

func addWordInArray(word string, array []int) []int {
	for i, _ := range word {
		i, _ := strconv.Atoi(word[i : i+1])
		array = append(array, i)
	}
	return array
}
