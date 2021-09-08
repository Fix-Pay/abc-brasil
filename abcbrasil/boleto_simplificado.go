package abcbrasil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"strconv"
	"strings"
)

type BoletoSimplificado struct {
	CodCliente           string  `json:"codCliente"`
	NumContaHeader       string  `json:"numContaHeader"`
	NumCarteira          string  `json:"numCarteira"`
	NossoNumero          string  `json:"nossoNumero"`
	CodModalBancos       string  `json:"codModalBancos"`
	NossoNumeroBanco     string  `json:"nossoNumeroBanco"`
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
	Messagem1            string  `json:"messagem1"`
	Messagem2            string  `json:"messagem2"`
	Messagem3            string  `json:"messagem3"`
	Messagem4            string  `json:"messagem4"`
	Messagem5            string  `json:"messagem5"`
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
	TaxaMora             float64 `json:"txMora"`
	ValorMora            float64 `json:"vlrMora"`
	PossuiAgenda         string  `json:"possuiAgenda"`
	TipoAgendamento      string  `json:"tipoAgendamento"`
	CriterioDias         string  `json:"criterioDias"`
	NumDiasAgenda        float64 `json:"numDiasAgenda"`
	CodIndice            string  `json:"codIndice"`
	IndPagtoParcial      string  `json:"indPagtoParcial"`
	QtdPagtosParciais    int     `json:"qtdPagtosParciais"`
	TipoValorPercMinimo  string  `json:"tipoValorPercMinimo"`
	ValorPercMinimo      float64 `json:"vlrPercMinimo"`
	TipoValorPercMaximo  string  `json:"tipoValorPercMaximo"`
	ValorPercMaximo      float64 `json:"vlrPercMaximo"`
	TipoAutRecDivergente string  `json:"tipoAutRecDivergente"`
}

type RetornoSucesso struct {
	Sucesso        bool   `json:"sucesso"`
	Mensagem       string `json:"mensagem"`
	PdfBinario     string `json:"pdf_binario"`
	Imagem         string `json:"imagem"`
	LinhaDigitavel string `json:"linha_digitavel"`
	CodigoDeBarras string `json:"codigo_de_barras"`
	Desenvolvedor  struct {
		CodigoRetorno string `json:"codigo_retorno"`
		Mensagem      string `json:"mensagem"`
	} `json:"desenvolvedor"`
}

func (simplificado BoletoSimplificado) GerarBoletoSimplificado(url, token string) (RetornoSucesso, error) {
	returnSuccess := RetornoSucesso{}

	url = fmt.Sprint(url, `/ABCDigital.BoletoOnline/api/v1.0/BoletoSimplificado`)
	token = fmt.Sprint("Bearer ", token)
	method := "POST"
	client := &http.Client{}

	boletoJson := strings.NewReader(getBoletoSimplificadoValues(simplificado).Encode())
	req, err := http.NewRequest(method, url, boletoJson)
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
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

func CalculoDigito(agencia, carteira, nossoNumero string) string {
	//agencia := "0019"
	//carteira := "110"
	//nossoNumero := "0062893742"
	digito, _ := calcularDigitoVerificador(agencia, carteira, nossoNumero)
	return fmt.Sprint(nossoNumero, digito)
}

func calcularDigitoVerificador(agencia, carteira, nossoNumero string) (int, error) {
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

func getBoletoSimplificadoValues(boleto BoletoSimplificado) netUrl.Values {
	data := netUrl.Values{}
	data.Set("codCliente", boleto.CodCliente)
	data.Set("numContaHeader", boleto.NumContaHeader)
	data.Set("numCarteira", boleto.NumCarteira)
	data.Set("nossoNumero", boleto.NossoNumero)
	data.Set("codModalBancos", boleto.CodModalBancos)
	data.Set("nossoNumeroBanco", boleto.NossoNumeroBanco)
	data.Set("codEspecieDoc", boleto.CodEspecieDoc)
	data.Set("vlrNominal", fmt.Sprintf("%.2f", boleto.ValorNominal))
	data.Set("vlrAbatimento", fmt.Sprintf("%.2f", boleto.ValorAbatimento))
	data.Set("dataEmissao", boleto.DataEmissao)
	data.Set("dataVencimento", boleto.DataVencimento)
	data.Set("seuNumero", boleto.SeuNumero)
	data.Set("aceite", boleto.Aceite)
	data.Set("cnpjCPF", boleto.CnpjCPF)
	data.Set("tipoPessoa", boleto.TipoPessoa)
	data.Set("nome", boleto.Nome)
	data.Set("endereco", boleto.Endereco)
	data.Set("bairro", boleto.Bairro)
	data.Set("cidade", boleto.Cidade)
	data.Set("uf", boleto.Uf)
	data.Set("cep", boleto.Cep)
	data.Set("email", boleto.Email)
	data.Set("ddd", boleto.Ddd)
	data.Set("telefone", boleto.Telefone)
	data.Set("campoLivre", boleto.CampoLivre)
	data.Set("tipoPessoaSacador", boleto.TipoPessoaSacador)
	data.Set("cnpjCPFSacador", boleto.CnpjCPFSacador)
	data.Set("nomeSacador", boleto.NomeSacador)
	data.Set("enderecoSacador", boleto.EnderecoSacador)
	data.Set("bairroSacador", boleto.BairroSacador)
	data.Set("cepSacador", boleto.CepSacador)
	data.Set("cidadeSacador", boleto.CidadeSacador)
	data.Set("ufSacador", boleto.UfSacador)
	data.Set("messagem1", boleto.Messagem1)
	data.Set("messagem2", boleto.Messagem2)
	data.Set("messagem3", boleto.Messagem3)
	data.Set("messagem4", boleto.Messagem4)
	data.Set("messagem5", boleto.Messagem5)
	data.Set("codDesconto1", boleto.CodDesconto1)
	data.Set("vlrDesconto1", fmt.Sprintf("%.2f", boleto.ValorDesconto1))
	data.Set("txDesconto1", fmt.Sprintf("%.2f", boleto.TaxaDesconto1))
	data.Set("dataDesconto1", boleto.DataDesconto1)
	data.Set("codDesconto2", boleto.CodDesconto2)
	data.Set("vlrDesconto2", fmt.Sprintf("%.2f", boleto.ValorDesconto2))
	data.Set("txDesconto2", fmt.Sprintf("%.2f", boleto.TaxaDesconto2))
	data.Set("dataDesconto2", boleto.DataDesconto2)
	data.Set("codDesconto3", boleto.CodDesconto3)
	data.Set("vlrDesconto3", fmt.Sprintf("%.2f", boleto.ValorDesconto3))
	data.Set("txDesconto3", fmt.Sprintf("%.2f", boleto.TaxaDesconto3))
	data.Set("dataDesconto3", boleto.DataDesconto3)
	data.Set("codMulta", boleto.CodMulta)
	data.Set("dataMulta", boleto.DataMulta)
	data.Set("txMulta", fmt.Sprintf("%.2f", boleto.TaxaMulta))
	data.Set("vlrMulta", fmt.Sprintf("%.2f", boleto.ValorMulta))
	data.Set("codMora", boleto.CodMora)
	data.Set("txMora", fmt.Sprintf("%.2f", boleto.TaxaMora))
	data.Set("vlrMora", fmt.Sprintf("%.2f", boleto.ValorMora))
	data.Set("possuiAgenda", boleto.PossuiAgenda)
	data.Set("tipoAgendamento", boleto.TipoAgendamento)
	data.Set("criterioDias", boleto.CriterioDias)
	data.Set("numDiasAgenda", fmt.Sprint(boleto.NumDiasAgenda))
	data.Set("codIndice", boleto.CodIndice)
	data.Set("indPagtoParcial", boleto.IndPagtoParcial)
	data.Set("qtdPagtosParciais", fmt.Sprint(boleto.QtdPagtosParciais))
	data.Set("tipoVlrPercMinimo", boleto.TipoValorPercMinimo)
	data.Set("vlrPercMinimo", fmt.Sprintf("%.2f", boleto.ValorPercMinimo))
	data.Set("tipoVlrPercMaximo", boleto.TipoValorPercMaximo)
	data.Set("vlrPercMaximo", fmt.Sprintf("%.2f", boleto.ValorPercMaximo))
	data.Set("tipoAutRecDivergente", boleto.TipoValorPercMaximo)
	return data
}
