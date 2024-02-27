package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Compress(source string) (compressed string, err error) {
	// Проверка на пустую строку
	if len(source) == 0 {
		return "", nil
	}

	// Объявляем string builder
	var builder strings.Builder
	// Предыдущий символ
	var previousChar rune
	// Количество символов
	var countChar int

	// Проход по символам в строке
	for _, char := range source {
		if char != previousChar {
			// Проеверяем, что он не первый и записываем в билдер символ и количество
			if previousChar != 0 {
				builder.WriteString(string(previousChar))
				builder.WriteString(strconv.Itoa(countChar)) // Конвертриуем int в string
			}
			previousChar = char
			countChar = 1
		} else {
			// Если новый символ совпадает с предыдущим - увеличиваем количество
			countChar++
		}
	}

	// Добавляем последний символ и его количество
	builder.WriteString(string(previousChar))
	builder.WriteString(strconv.Itoa(countChar))

	return builder.String(), nil
}

func Decompress(compressed string) (source string, err error) {
	// Объявляем String Builder
	var sourceString strings.Builder
	// Текущий символ
	var currentChar rune
	// Общий счетчик
	var count int = 0

	// Проход по символам в строке
	for _, char := range compressed {
		if char != 0 {
			// Для каждой четной итерации записываем символ
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
				// Заполняем оставшееся количество символов
				for i := 0; i < count-1; i++ {
					sourceString.WriteString(string(currentChar))
				}
			}
			count++
		}
	}
	return sourceString.String(), nil
}

func main() {
	fmt.Println(Compress("aaaabbсaa"))
	fmt.Println(Decompress("a4b2с1a2"))
}
