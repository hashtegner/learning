package normalizer

import (
	"bytes"
)

func Normalize(phone string) string {
	var normalized bytes.Buffer

	for _, char := range phone {
		if char >= '0' && char <= '9' {
			normalized.WriteRune(char)
		}
	}

	return normalized.String()
}
