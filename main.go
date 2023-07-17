package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	greetings()
	fmt.Println(getNums())
}

func greetings() {
	fmt.Println("Приветствую путник!")
	fmt.Println("Простой калькулятр, поддерживаются операторы '+' '-' '*' '/' ")
}

func getNums() float64 {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите пример, напр: 1 + 1 (через пробел) и нажмите Enter")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	result := calculate(text)
	return result

}

func calculate(input string) float64 {
	tokens := strings.Split(input, " ")

	if len(tokens) != 3 {
		fmt.Println("Не верный формат ввода.")
		return 0
	}

	firstNum, err1 := strconv.ParseFloat(tokens[0], 64)
	secondNum, err2 := strconv.ParseFloat(tokens[2], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка преобразования чисел")
	}

	operator := tokens[1]
	switch operator {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		if secondNum == 0 {
			fmt.Println("Деление на 0")
			return 0
		}
		return firstNum / secondNum

	default:
		fmt.Println("Пока что неподдерживаемый оператор")
		return 0
	}
}
