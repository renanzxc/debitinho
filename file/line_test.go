package file

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	var (
		originLineE = "E                        11234             12077010100000000000011103                                                          1X2000000999999999    0"
		originLineT = "T00000000000000000000000                                                                                                                              "

		dataTest = map[string]ILine{
			originLineE: &LineE{
				basicLine: &basicLine{
					OriginLine: originLineE,
					MyType:     LineTypeE,
				},
				RegisterCode:       "E",
				CustomerCompanyID:  "                        1",
				Agency:             "1234",
				CustomerBankID:     "             1",
				DueDate:            "20770101",
				DebitValue:         "000000000000111",
				CoinCode:           "03",
				CompanyUse:         "                                                          1X",
				IdentificationType: "2",
				Identification:     "000000999999999",
				ReserveFuture:      "    ",
				MovingCode:         "0",
			},
			originLineT: &LineT{
				basicLine: &basicLine{
					OriginLine: originLineT,
					MyType:     LineTypeT,
				},
				RegisterCode:             "T",
				TotalDebitRegisters:      "000000",
				TotalDebitRegistersValue: "00000000000000000",
				ReserveFuture:            "                                                                                                                              ",
			},
		}
	)

	for originLineTest, expectedParsedLine := range dataTest {
		t.Run(`Teste Parse Linha `+string(expectedParsedLine.Type()), func(t *testing.T) {
			t.Run(`Deve fazer o parse da linha com sucesso`, func(t *testing.T) {
				resultParsedLine := ParseLine(originLineTest)
				require.Equal(t, expectedParsedLine, resultParsedLine)
			})
		})
	}
}
