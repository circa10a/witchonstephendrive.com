package sounds

import "testing"

func TestIsDuringQuietTime(t *testing.T) {
	t.Parallel()
	tests := []map[string]interface{}{
		// 10pm (start time)
		{
			"now":   22,
			"start": 22,
			"end":   7,
			"pass":  true,
		},
		// 9pm
		{
			"now":   21,
			"start": 22,
			"end":   7,
			"pass":  false,
		},
		// 12am
		{
			"now":   00,
			"start": 22,
			"end":   7,
			"pass":  true,
		},
		// 7am (end time)
		{
			"now":   7,
			"start": 22,
			"end":   7,
			"pass":  false,
		},
		{
			"now":   8,
			"start": 8,
			"end":   9,
			"pass":  true,
		},
	}
	for _, test := range tests {
		now := test["now"].(int)
		start := test["start"].(int)
		end := test["end"].(int)
		pass := test["pass"].(bool)
		if IsDuringQuietTime(now, start, end) != pass {
			t.Fail()
		}
	}
}
