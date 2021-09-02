package v1

type BoletoSimplificado struct {
	Amount               string  `json:"codCliente"`
	RefundDateTime       string  `json:"numContaHeader"`
	RefundID             string  `json:"numCarteira"`
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