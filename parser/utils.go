package parser

import (
	"strings"
	"time"

	"github.com/charlesbases/reverse/types"
)

const defaultTimeFormatLayou = "2006-01-02 15:04:05"

func today() string {
	return time.Now().Format(defaultTimeFormatLayou)
}

// camelcase aaa_bbb to AaaBbb
func camelcase(s string) string {
	if b, find := types.Abbre[s]; find {
		return b
	} else {
		var bs strings.Builder
		bs.Grow(len(s))
		for _, item := range strings.Split(s, "_") {
			if len(item) > 0 {
				if b, find := types.Abbre[item]; find {
					bs.WriteString(b)
				} else {
					if c := item[0]; isASCIILower(c) {
						bs.WriteByte(c - ('a' - 'A'))
					} else {
						bs.WriteByte(c)
					}
					if len(item) > 1 {
						bs.WriteString(item[1:])
					}
				}
			}
		}
		return bs.String()
	}
}

// encamelcase AaaBbb to aaa_bbb
func encamelcase(s string) string {
	builder := strings.Builder{}
	for index, letter := range []byte(s) {
		if isASCIIUpper(letter) {
			if index != 0 {
				builder.WriteString("_")
			}
			builder.WriteByte(letter + 'a' - 'A')
		} else {
			builder.WriteByte(letter)
		}
	}
	return builder.String()
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
