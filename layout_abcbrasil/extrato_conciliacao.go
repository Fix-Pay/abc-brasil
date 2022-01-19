package layout_abcbrasil

import "time"

type AbcBrasilExtratoConciliacaoResponse struct {
	Status          bool      `json:"status"`
	Name            string    `json:"name"`
	EnvironmentName string    `json:"environmentName"`
	Date            time.Time `json:"date"`
	Data            struct {
		Protocolo string `json:"protocolo"`
		Header    struct {
			Controle struct {
				Banco    int `json:"banco"`
				Lote     int `json:"lote"`
				Registro int `json:"registro"`
			} `json:"controle"`
			Servico struct {
				Operacao        string `json:"operacao"`
				TipoServico     int    `json:"tipoServico"`
				FormaLancamento int    `json:"formaLancamento"`
				LayoutLote      string `json:"layoutLote"`
			} `json:"servico"`
			Empresa struct {
				NomeEmpresa     string `json:"nomeEmpresa"`
				TipoInscricao   int    `json:"tipoInscricao"`
				NumeroInscricao string `json:"numeroInscricao"`
				Convenio        string `json:"convenio"`
				Agencia         string `json:"agencia"`
				DVAgencia       string `json:"DVAgencia"`
				Conta           string `json:"conta"`
				DVConta         string `json:"DVConta"`
				DVAgenciaConta  string `json:"DVAgenciaConta"`
			} `json:"empresa"`
			SaldoInicial struct {
				Data       time.Time `json:"data"`
				Valor      float64   `json:"valor"`
				Situacao   string    `json:"situacao"`
				Status     string    `json:"status"`
				TipoMoeda  string    `json:"tipoMoeda"`
				SeqExtrato int       `json:"seqExtrato"`
			}
		} `json:"header"`
		Detalhes    []ExtratoConciliacaoDetalhes `json:"detalhes"`
		TrailerLote struct {
			Controle struct {
				Banco    int `json:"banco"`
				Lote     int `json:"lote"`
				Registro int `json:"registro"`
			} `json:"controle"`
			Empresa struct {
				TipoInscricao   int    `json:"tipoInscricao"`
				NumeroInscricao string `json:"numeroInscricao"`
				Convenio        string `json:"convenio"`
				Agencia         string `json:"agencia"`
				DVAgencia       string `json:"DVAgencia"`
				Conta           string `json:"conta"`
				DVConta         string `json:"DVConta"`
				DVAgenciaConta  string `json:"DVAgenciaConta"`
			} `json:"empresa"`
			Valores struct {
				BloqueadoAcima24 float64 `json:"bloqueadoAcima24"`
				Limite           float64 `json:"limite"`
				BloqueadoAte24   float64 `json:"bloqueadoAte24"`
			} `json:"valores"`
			SaldoFinal struct {
				Data     time.Time `json:"data"`
				Valor    float64   `json:"valor"`
				Situacao string    `json:"situacao"`
				Status   string    `json:"status"`
			} `json:"saldoFinal"`
			Totais struct {
				Registros    int     `json:"registros"`
				ValorDebitos float64 `json:"valorDebitos"`
				ValorCredito float64 `json:"valorCredito"`
			} `json:"totais"`
		} `json:"trailerLote"`
	} `json:"data"`
	Infos  interface{} `json:"infos"`
	Errors struct {
		Code     string `json:"code"`
		Message  string `json:"message"`
		Title    string `json:"title"`
		Property string `json:"property"`
	} `json:"errors"`
}

type ExtratoConciliacaoDetalhes struct {
	Banco               int       `json:"banco"`
	Lote                int       `json:"lote"`
	Registro            int       `json:"registro"`
	NumRegistro         int       `json:"numRegistro"`
	Segmento            string    `json:"segmento"`
	TipoInscricao       int       `json:"tipoInscricao"`
	NumeroInscricao     string    `json:"numeroInscricao"`
	Convenio            string    `json:"convenio"`
	Agencia             string    `json:"agencia"`
	DVAgencia           string    `json:"DVAgencia"`
	Conta               string    `json:"conta"`
	DVConta             string    `json:"DVConta"`
	DVAgenciaConta      string    `json:"DVAgenciaConta"`
	NomeEmpresa         string    `json:"nomeEmpresa"`
	Natureza            string    `json:"natureza"`
	TipoComplemento     int       `json:"tipoComplemento"`
	Complemento         string    `json:"complemento"`
	Cpmf                string    `json:"cpmf"`
	DataContabil        time.Time `json:"dataContabil"`
	DataLancamento      time.Time `json:"dataLancamento"`
	ValorLancamento     float64   `json:"valorLancamento"`
	TipoLancamento      string    `json:"tipoLancamento"`
	CategoriaLancamento string    `json:"categoriaLancamento"`
	CodigoHistorico     string    `json:"codigoHistorico"`
	Historico           string    `json:"historico"`
	NumDocumento        string    `json:"numDocumento"`
}

type AbcBrasilConsultaSituacaoExtratoReponse struct {
	Status          bool   `json:"status"`
	Name            string `json:"name"`
	EnvironmentName string `json:"environmentName"`
	Date            string `json:"date"`
	Data            struct {
		Codigo             int    `json:"codigo"`
		Protocolo          string `json:"protocolo"`
		Descricao          string `json:"descricao"`
		DataValidade       string `json:"dataValidade"`
		QtdRegistros       int    `json:"qtdRegistros"`
		QtdRegistrosPagina int    `json:"qtdRegistrosPagina"`
		QtdPaginas         int    `json:"qtdPaginas"`
	} `json:"data"`
	Infos  interface{} `json:"infos"`
	Errors interface{} `json:"errors"`
}
