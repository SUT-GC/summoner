package utils

import (
	"encoding/base64"
	"fmt"
	"strings"
)

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

func ChangeIntListToStringList(numbers []int64) []string {
	sl := make([]string, 0)
	for _, n := range numbers {
		sl = append(sl, fmt.Sprint(n))
	}

	return sl
}

func DecodeWithBase64(str string) string {
	sDec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("decode base64 error")
		return ""
	}

	return string(sDec)
}

func EncodeWithBase64(str string) string {
	b := []byte(str)

	sEnc := base64.StdEncoding.EncodeToString(b)

	return sEnc

}
