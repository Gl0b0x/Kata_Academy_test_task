package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		result string
		err    error
	)
	for err == nil {
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		result, err = Calculation(text)
		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}
}
func ToRoman(number int) string {
	cifr := map[int]string{100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	keys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for number != 0 {
		for _, key := range keys {
			if number-key >= 0 {
				result += cifr[key]
				number -= key
				break
			}
		}
	}
	return result
}
func ToDecimal(number string) (int, error) {
	cifr := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	value, ok := cifr[number]
	if ok {
		return value, nil
	} else {
		return 0, errors.New("некорректная запись римского числа")
	}
}
func IsRomanNumber(number string) bool {
	for _, value := range number {
		if value != 'V' && value != 'I' && value != 'X' {
			return false
		}
	}
	return true
}

func ToStringResult(number int, flag int) (string, error) {
	if flag == 0 {
		return strconv.Itoa(number), nil
	} else {
		if number > 0 {
			return ToRoman(number), nil
		}
		return "", errors.New("в римской системе есть только положительные числа")
	}
}

func Calculation(text string) (string, error) {
	errIncorectString := errors.New("некорректная строка")
	errZeroDivision := errors.New("на ноль делить нельзя")
	errPresentation := errors.New("строка не является математической операцией")
	words := strings.Fields(text)
	if len(words) != 3 {
		return "", errIncorectString
	}
	var (
		a, b, flag int // Флаг показывает в каком формате числа; flag == 0 - арабские числа; flag == 1, римские числа
		err        error
	)
	if !(IsRomanNumber(words[0])) && !(IsRomanNumber(words[2])) {
		a, err = strconv.Atoi(words[0])
		b, err = strconv.Atoi(words[2])
	} else {
		a, err = ToDecimal(words[0])
		b, err = ToDecimal(words[2])
		flag = 1
	}
	if err != nil {
		return "", errIncorectString
	}
	switch words[1] {
	case "+":
		return ToStringResult(a+b, flag)
	case "-":
		return ToStringResult(a-b, flag)
	case "*":
		return ToStringResult(a*b, flag)
	case "/":
		if b == 0 {
			return "", errZeroDivision
		}
		return ToStringResult(a/b, flag)
	default:
		return "", errPresentation
	}
}
