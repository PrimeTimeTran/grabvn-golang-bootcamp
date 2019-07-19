package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		text := scanner.Text()
		value := eval(text)

		fmt.Println(value)
		fmt.Print("> ")
	}

	return
}
