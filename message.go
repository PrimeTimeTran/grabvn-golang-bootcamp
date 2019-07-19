package main

import (
	"fmt"
	"strconv"
	"strings"
)

func eval(text string) (value float64) {
	var la = strings.Replace(text, " ", "", -1)

	var expression = strings.Split(la, "")

	var a, _ = strconv.ParseFloat(expression[0], 10)
	var b, _ = strconv.ParseFloat(expression[2], 10)
	var op = expression[1]

	switch op {
	case "+":
		fmt.Println(a, op, b, "=", a+b)
		value = a + b
	case "-":
		fmt.Println(a, op, b, "=", a-b)
		value = a - b
	case "*":
		fmt.Println(a, op, b, "=", a*b)
		value = a * b
	case "/":
		fmt.Println(a, op, b, "=", a/b)
		if a == 0 {
			panic("Can not divide to 0")
		}
		value = a / b
	default:
		fmt.Println("Invalid operator")
	}

	return value
}
