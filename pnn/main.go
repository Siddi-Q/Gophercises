package main

import "bytes"

func main() {

}

func normalizePhoneNumber(pn string) string {
	var buf bytes.Buffer
	for _, rune := range pn {
		if rune >= '0' && rune <= '9' {
			buf.WriteRune(rune)
		}
	}
	return buf.String()
}
