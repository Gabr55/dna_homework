package dna

import (
	"strconv"
	"strings"
	"unicode"
)

func Compress(source string) (compressed string, err error) {
	if len(source) == 0 {
		return "", nil
	}

	var builder strings.Builder
	var previousChar rune
	var countChar int

	for _, char := range source {
		if char != previousChar {
			if previousChar != 0 {
				builder.WriteString(string(previousChar))
				builder.WriteString(strconv.Itoa(countChar))
			}
			previousChar = char
			countChar = 1
		} else {
			countChar++
		}
	}
	builder.WriteString(string(previousChar))
	builder.WriteString(strconv.Itoa(countChar))

	return builder.String(), nil
}

func Decompress(compressed string) (source string, err error) {
	var sourceString strings.Builder
	var currentChar rune
	var count int = 0

	for _, char := range compressed {
		if char != 0 {
			if count%2 == 0 {
				currentChar = char
				sourceString.WriteString(string(char))
			} else {
				if !unicode.IsDigit(char) {
					panic("Неподдерживаемая строка")
				}
				count, err := strconv.Atoi(string(char))
				if err != nil {
					panic("Неподдерживаемая строка")
				}
				for i := 0; i < count-1; i++ {
					sourceString.WriteString(string(currentChar))
				}
			}
			count++
		}
	}
	return sourceString.String(), nil
}
