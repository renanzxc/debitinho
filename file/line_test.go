package file

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	t.Run(`Teste Parse Linha E`, func(t *testing.T) {
		var (
			mockLineEStr    = "E                        11234             12077010100000000000011103                                                          1X2000000999999999    0"
			mockParsedLineE = LineE{
				basicLine:          &basicLine{OriginLine: mockLineEStr},
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
			}
		)

		t.Run(`Deve fazer o parse da linha com sucesso`, func(t *testing.T) {
			var lineE = LineE{
				basicLine: &basicLine{OriginLine: mockLineEStr},
			}
			err := parseLine(mockLineEStr, &lineE)
			require.Nil(t, err)

			require.Equal(t, mockParsedLineE, lineE)
		})
	})
}
