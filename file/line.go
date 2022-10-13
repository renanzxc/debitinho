package file

import (
	"log"
)

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

	if err := parseLine(l.OriginLine, typedLine); err != nil {
		log.Fatal(err)
	}

	return
}

type LineA struct {
	*basicLine
	RegisterCode          string `name:"Código do Registro" json:"codigo_do_registro" posicao:"0"`
	ShippingCode          string `name:"Código de Remessa" json:"codigo_de_remessa" posicao:"1"`
	ConvCode              string `name:"Código do Convênio" json:"codigo_do_convenio" posicao:"2:22"`
	CompanyName           string `name:"Nome da Empresa" json:"nome_da_empresa" posicao:"22:42"`
	BankCode              string `name:"Código do Banco" json:"codigo_do_banco" posicao:"42:45"`
	BankName              string `name:"Nome do Banco" json:"nome_do_banco" posicao:"45:65"`
	GenerationDate        string `name:"Data de Geração" json:"data_de_geracao" posicao:"65:73"`
	NSA                   string `name:"Número Seqüencial do Arquivo (NSA)" json:"nsa" posicao:"73:79"`
	Version               string `name:"Versão do Layout" json:"versao_do_layout" posicao:"79:81"`
	ServiceIdentification string `name:"Identificação do Serviço" json:"identificacao_do_servico" posicao:"81:98"`
	ReserveFuture         string `name:"Reservado para o futuro" json:"reservado_para_o_futuro" posicao:"98:150"`
}

func (line *LineA) Validate() {

}

type LineZ struct {
	*basicLine
	RegisterCode       string `name:"Código do Registro" json:"codigo_do_registro" posicao:"0"`
	TotalRecordsInFile string `name:"Total de registros do arquivo" json:"total_de_registros_do_arquivo" posicao:"1:7"`
	TotalValueInFile   string `name:"Valor total dos registros do arquivo" json:"valor_total_dos_registros_do_arquivo" posicao:"7:24"`
	ReserveFuture      string `name:"Reservado para o futuro" json:"reservado_para_o_futuro" posicao:"24:150"`
}

func (line *LineZ) Validate() {

}
