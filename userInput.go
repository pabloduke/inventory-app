package main

import "fmt"

func Input(prompt string) (input string) {
	print(prompt + " ")
	fmt.Scan(&input)

	return input
}
