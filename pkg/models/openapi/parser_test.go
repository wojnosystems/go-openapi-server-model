package openapi

import (
	"strings"
	"testing"
)

func TestParseModels(t *testing.T) {
	cases := map[string]struct {
		input string
	} {
		//"empty": {
		//	input: "",
		//},
		"schema structure": {
			input:
`components:
  schemas:
    AccountId:
      $ref: '#/components/schemas/Id'
    Id:
      type: object
      properties:
        type: integer
        Format: int64
        description: An object identifier`,
		},
	}

	for caseName, c := range cases {
		_, _ = ParseModels( strings.NewReader(c.input) )
		_ = caseName
	}
}
