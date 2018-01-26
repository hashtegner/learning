package phone

import "testing"

var testCases = []struct {
	input    string
	expected string
}{
	{"1234567890", "1234567890"},
	{"123 456 7891", "1234567891"},
	{"(123) 456 7892", "1234567892"},
	{"(123) 456-7893", "1234567893"},
	{"123-456-7894", "1234567894"},
}

func TestNormalizer(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			normalized := Normalize(tc.input)

			if normalized != tc.expected {
				t.Errorf("Expected %s for %s but %s", tc.expected, tc.input, normalized)
			}
		})

	}
}

func BenchmarkNormalizer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Normalize(tc.input)
		}
	}
}

// func BenchmarkNormalizerSpeed(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, tc := range testCases {
// 			NormalizeSpeed(tc.input)
// 		}
// 	}
// }
