package dna

import (
	"bufio"
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
	// Объявляем буфер с ссылкой на StringBuilder
	b := bufio.NewWriter(&sourceString)
	// Текущий символ
	var currentChar rune
	// Общий счетчик
	var count byte = 0

	// Проход по символам в строке
	for _, char := range compressed {
		if char != 0 {
			// Для каждой четной итерации записываем символ
			if count%2 == 0 {
				currentChar = char
				b.WriteString(string(char))
				// Сбрасываем счетчик
				count = 0
				// Записываем в StringBuilder
				err = b.Flush()
				if err != nil {
					return err.Error(), err
				}
			} else {
				if !unicode.IsDigit(char) {
					err := fmt.Errorf("unsupported string")
					return err.Error(), err
				}
				count, err := strconv.Atoi(string(char))
				if err != nil {
					err := fmt.Errorf("unsupported string")
					return err.Error(), err
				}
				// Заполняем оставшееся количество символов
				for i := 0; i < count-1; i++ {
					b.WriteString(string(currentChar))
				}
				// Записываем в StringBuilder
				err = b.Flush()
				if err != nil {
					return err.Error(), err
				}
				// Сбрасываем буфер
				b.Reset(&sourceString)
			}
			count++
		}
	}
	return sourceString.String(), nil
}
