package utils

import "strings"

func PlaceHolders(n int) string {
	var b strings.Builder
	for i := 0; i < n-1; i++ {
		b.WriteString("?,")
	}
	if n > 0 {
		b.WriteString("?")
	}
	return b.String()
}