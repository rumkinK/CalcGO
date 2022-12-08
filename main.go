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
	digits := strings.Fields(Scan1()) // создаем массив из введенной строки
	switch {                          // проверяем ошибки ввода
	case len(digits) == 1:
		fmt.Println("Ошибка ввода! Cтрока не является математической операцией")
		return
	case len(digits) == 2:
		fmt.Println("Ошибка ввода! Возможно отсутствует знак математической операции")
		return
	case len(digits) > 3:
		fmt.Println("Ошибка ввода! Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}
	var action string = digits[1] // в этой переменной хранится арифметическое действие
	var a, b int
	var romandigits bool // идентификатор показывающий, что вычисления производятся в Римской системе
	switch {
	case romanToInt(digits[0]) != 0 && romanToInt(digits[2]) != 0:
		a = romanToInt(digits[0])
		b = romanToInt(digits[2])
		romandigits = true
	case (romanToInt(digits[0]) == 0 && romanToInt(digits[2]) != 0) || (romanToInt(digits[0]) != 0 && romanToInt(digits[2]) == 0):
		fmt.Println("Ошибка, используются одновременно разные системы счисления")
		return
	case romanToInt(digits[0]) == 0 && romanToInt(digits[2]) == 0:
		a, _ = strconv.Atoi(digits[0])
		b, _ = strconv.Atoi(digits[2])
	}

	if a > 10 || a < 1 || b > 10 || b < 1 {
		fmt.Println("Ошибка, операции производятся с целыми Римскими или Арабскими числами от 1 до 10 включительно!")
		return
	}
	res, err := Calc(a, action, b) // производим вычисление и выводим результат на консоль
	if err == nil {
		if romandigits {
			if res < 0 || res == 0 {
				fmt.Println("Ошибка, в Римской системе нет нуля и отрицательных чисел!")
				return
			}
			fmt.Println(Roman(res))
		} else {
			fmt.Println(res)
		}
	} else {
		fmt.Println(err)
	}
}

func Calc(a int, action string, b int) (int, error) {
	switch action {
	case "*":
		return a * b, nil
	case "-":
		return a - b, nil
	case "+":
		return a + b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("Ошибка ввода - недопустимая операция!")
	}
}

// Функция считывает введенные данные с консоли
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return strings.ToUpper(in.Text())
}

// Функция преобразования Арабских чисел в Римские найдена в Интернете
func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

// Функция преобразования Римских чисел в арабские найдена в Интернете
func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}
