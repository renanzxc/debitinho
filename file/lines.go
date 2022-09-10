package file

type LineType rune

const (
	LineTypeA LineType = 'A'
	LineTypeB LineType = 'B'
)

func NewLine(lineStr string) ILine {
	line := &basicLine{OLine: lineStr}
	return line.Parse()
}

type ILine interface {
	Type() LineType
	OriginLine() string
	Parse()
	Validate()
}

type basicLine struct {
	MyType LineType `json:"-"`
	OLine  string   `json:"-"`
}

func (l *basicLine) Type() LineType {
	return l.MyType
}

func (l *basicLine) OriginLine() string {
	return l.OLine
}

func (l *basicLine) DefaultValidation() {
	defaultLineValidation(l.OLine)
}

func (l *basicLine) Parse() (typedLine ILine) {
	l.DefaultValidation()

	switch LineType(l.OLine[0]) {
	case LineTypeA:
		typedLine = &LineA{basicLine: l}
		l.MyType = LineTypeA
	}

	typedLine.Parse()

	return
}

type LineA struct {
	*basicLine
	RegisterCode          string `name:"Código do Registro" json:"codigo_do_registro"`
	ShippingCode          string `name:"Código de Remessa" json:"codigo_de_remessa"`
	ConvCode              string `name:"Código do Convênio" json:"codigo_do_convenio"`
	CompanyName           string `name:"Nome da Empresa" json:"nome_da_empresa"`
	BankCode              string `name:"Código do Banco" json:"codigo_do_banco"`
	BankName              string `name:"Nome do Banco" json:"nome_do_banco"`
	GenerationDate        string `name:"Data de Geração" json:"data_de_geracao"`
	NSA                   string `name:"Número Seqüencial do Arquivo (NSA)" json:"nsa"`
	Version               string `name:"Versão do Layout" json:"versao_do_layout"`
	ServiceIdentification string `name:"Identificação do Serviço" json:"identificacao_do_servico"`
	ReserveFuture         string `name:"Reservado para o futuro" json:"reservado_para_o_futuro"`
}

func (line *LineA) Parse() {
	originLine := line.OriginLine()

	line.RegisterCode = string(originLine[0])
	line.ShippingCode = string(originLine[1])
	line.ConvCode = string(originLine[2:22])
	line.CompanyName = string(originLine[22:42])
	line.BankCode = string(originLine[42:45])
	line.BankName = string(originLine[45:65])
	line.GenerationDate = string(originLine[65:73])
	line.NSA = string(originLine[73:79])
	line.Version = string(originLine[79:81])
	line.ServiceIdentification = string(originLine[81:98])
	line.ReserveFuture = string(originLine[98:150])
}

func (line *LineA) Validate() {

}
