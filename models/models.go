package models

type Cnab struct {

	//Header
	TipoOperacaoRemessa    string `json: "tipoOperacaoRemessa"`
	LiteralRemessa         string `json: "literalRemessa"`
	CodigoServico          string `json: "codigoServico"`
	LiteralServico         string `json: "literalServico"`
	ISPBParticipante       string `json: "iSPBParticipante"`
	TipoPessoaRecebedor    string `json: "tipoPessoaRecebedor"`
	CPForCNPJ              string `json: "cPForCNPJ"`
	Agencia                string `json: "agencia"`
	Conta                  string `json: "conta"`
	TipoConta              string `json: "tipoConta"`
	ChavePix               string `json: "chavePix"`
	DataGeracao            string `json: "dataGeracao"`
	CodigoConvenio         string `json: "codigoConvenio"`
	ExclusivoPSPRecebedor  string `json: "exclusivoPSPRecebedor"`
	NomeRecebedor          string `json: "nomeRecebedor"`
	Bancos                 string `json: "bancos"`
	NumeroSequenciaRemessa string `json: "numeroSequenciaRemessa"`
	VersaoArquivo          string `json: "versaoArquivo"`

	//Detalhe
	TipoRegistro           string `json: "tipoRegistro"`
	Identificador          string `json: "identificador"`
	Tipo                   string `json: "tipo"`
	TipoCobranca           string `json: "tipoCobranca"`
	CodOcorrencia          string `json: "codOcorrencia"`
	TimesTampExpiracao     string `json: "timesTampExpiracao"`
	DataVencimento         string `json: "dataVencimento"`
	ValidadeAposVencimento string `json: "validadeAposVencimento"`
	ValorOriginal          string `json: "valorOriginal"`
	TipoPessoaDevedor      string `json: "tipoPessoaDevedor"`
	CPForCNPJDevedor       string `json: "cPForCNPJDevedor"`
	NomeDevedor            string `json: "nomeDevedor"`
	CampoTextoAoPagador    string `json: "campoTextoAoPagador"`

	//Trailer
	ValorTotal       string `json: "tipoRegistro"`
	QTdRegistros     string `json: "identificador"`
	NumeroSequencial string `json: "tipo"`

	//Status
	StatusProjeto string `json: "tipo"`
}
