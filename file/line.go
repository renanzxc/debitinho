package file

import "log"

type LineType rune

const (
	LineTypeA LineType = 'A'
	LineTypeB LineType = 'B'
	LineTypeC LineType = 'C'
	LineTypeD LineType = 'D'
	LineTypeE LineType = 'E'
	LineTypeF LineType = 'F'
	LineTypeH LineType = 'H'
	LineTypeI LineType = 'I'
	LineTypeJ LineType = 'J'
	LineTypeK LineType = 'K'
	LineTypeL LineType = 'L'
	LineTypeT LineType = 'T'
	LineTypeX LineType = 'X'
	LineTypeZ LineType = 'Z'
)

func ParseLine(lineStr string) ILine {
	line := &basicLine{OriginLine: lineStr}
	return line.Parse()
}

type ILine interface {
	Type() LineType
	String() string
	Parse()
	Validate()
}

type basicLine struct {
	MyType     LineType `json:"-"`
	OriginLine string   `json:"-"`
}

func (l *basicLine) Type() LineType {
	return l.MyType
}

func (l *basicLine) String() string {
	return l.OriginLine
}

func (l *basicLine) Parse() (typedLine ILine) {
	basicLineValidation(l.OriginLine)

	switch LineType(l.OriginLine[0]) {
	case LineTypeA:
		typedLine = &LineA{basicLine: l}
		l.MyType = LineTypeA
	case LineTypeZ:
		typedLine = &LineZ{basicLine: l}
		l.MyType = LineTypeZ
	default:
		log.Fatal("Tipo de linha não identificada")
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
	originLine := line.String()

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

type LineZ struct {
	*basicLine
	RegisterCode       string `name:"Código do Registro" json:"codigo_do_registro"`
	TotalRecordsInFile string `name:"Total de registros do arquivo" json:"total_de_registros_do_arquivo"`
	TotalValueInFile   string `name:"Valor total dos registros do arquivo" json:"valor_total_dos_registros_do_arquivo"`
	ReserveFuture      string `name:"Reservado para o futuro" json:"reservado_para_o_futuro"`
}

func (line *LineZ) Parse() {
	originLine := line.String()

	line.RegisterCode = string(originLine[0])
	line.TotalRecordsInFile = string(originLine[1:7])
	line.TotalValueInFile = string(originLine[7:24])
	line.ReserveFuture = string(originLine[24:150])
}

func (line *LineZ) Validate() {

}
