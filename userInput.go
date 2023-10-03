package main

import "fmt"

func Input(prompt string) (input string) {
	println(prompt + " ")
	fmt.Scan(&input)

	return input
}
