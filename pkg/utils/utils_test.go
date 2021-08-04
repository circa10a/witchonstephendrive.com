package utils

import "testing"

func TestIsDuringQuietTime(t *testing.T) {
	t.Parallel()
	tests := []map[string]interface{}{
		{
			"string": "yes",
			"slice":  []string{"yes", "no"},
			"pass":   true,
		},
		{
			"string": "yes",
			"slice":  []string{"no", "no"},
			"pass":   false,
		},
	}
	for _, test := range tests {
		str := test["string"].(string)
		slice := test["slice"].([]string)
		pass := test["pass"].(bool)
		if StrInSlice(str, slice) != pass {
			t.Fail()
		}
	}
}
