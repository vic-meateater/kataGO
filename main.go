package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabicNums = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabicToRomanNums = map[int]string{
	1:   "I",
	4:   "IV",
	5:   "V",
	9:   "IX",
	10:  "X",
	40:  "XL",
	50:  "L",
	90:  "XC",
	100: "C",
}

func main() {
	greetings()
	fmt.Println(getNums())
}

func greetings() {
	fmt.Println("Приветствую путник!")
	fmt.Println("Простой калькулятр, поддерживаются операторы '+' '-' '*' '/', числа от 1 до 10 / I - X")
}

func getNums() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите пример, напр: 1 + 1 (через пробел) и нажмите Enter")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	result := performCalculate(text)
	return result

}

func performCalculate(input string) string {
	tokens := strings.Split(input, " ")

	if len(tokens) != 3 {
		return "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	}

	leftNumber := strings.ToUpper(tokens[0])
	rightNumber := strings.ToUpper(tokens[2])
	operator := tokens[1]

	isRomanNumbs := isRomanNumber(leftNumber) && isRomanNumber(rightNumber)

	if isRomanNumbs {
		leftNumberInt := romanToArabicNums[leftNumber]
		rightNumberInt := romanToArabicNums[rightNumber]

		result := calculate(leftNumberInt, rightNumberInt, operator)
		return arabicToRoman(result)
	}

	leftNumberInt, err1 := strconv.Atoi(leftNumber)
	rightNumberInt, err2 := strconv.Atoi(rightNumber)

	isArabicNums := err1 == nil || err2 == nil

	if !isRomanNumbs && isArabicNums {
		return "Вывод ошибки, так как используются одновременно разные системы счисления."
	}

	if !isArabicNums {
		return "Калькулятор умеет работать только с арабскими или римскими цифрами одновременно"
	}

	if !inValidRange(leftNumberInt, 1, 10) || !inValidRange(rightNumberInt, 1, 10) {
		return "Калькулятор должен принимать на вход числа от 1 до 10 включительно"
	}

	result := calculate(leftNumberInt, rightNumberInt, operator)

	return strconv.Itoa(result)
}

func isRomanNumber(number string) bool {
	_, exists := romanToArabicNums[number]
	return exists
}

func calculate(leftNum, rightNum int, operator string) int {
	switch operator {
	case "+":
		return leftNum + rightNum
	case "-":
		return leftNum - rightNum
	case "*":
		return leftNum * rightNum
	case "/":
		if rightNum == 0 {
			fmt.Println("Деление на 0")
			return 0
		}
		return leftNum / rightNum

	default:
		fmt.Println("Пока что неподдерживаемый оператор")
		return 0
	}
}

func arabicToRoman(number int) string {
	if number <= 0 {
		return "Результатом работы калькулятора с римскими числами могут быть только положительные числа"
	}

	result, exist := arabicToRomanNums[number]
	if exist {
		return result
	}

	var romanNum string

	sortedArabics := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, arabic := range sortedArabics {
		for number >= arabic {
			romanNum += arabicToRomanNums[arabic]
			number -= arabic
		}
	}
	return romanNum
}

func inValidRange(number int, start int, end int) bool {
	result := number >= start && number <= end
	return result
}
