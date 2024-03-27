package main

import (
	"fmt"
	"strings"
)

func dividir(dividendo, divisor int, carry string) string {
	if len(carry) > 10 {
		return carry
	}

	var intResult, residuo, cursor int

	// get num divisor in dividendo
	for i := 0; cursor <= dividendo; i++ {
		residuo = dividendo - cursor
		intResult = i
		cursor += divisor
	}

	result := fmt.Sprintf("%s%d", carry, intResult)
	if residuo == 0 {
		return result
	}

	if !strings.Contains(result, ".") {
		result = fmt.Sprintf("%s.", result)
	}

	// residuo x 10
	_dividendo := (residuo << 3) + (residuo << 1)
	result = dividir(_dividendo, divisor, result)

	return result
}

func main() {
	var dividendo int
	var divisor int

	println("Ingresar dividendo: ")
	fmt.Scanln(&dividendo)
	println("Ingresar dividendo: ")
	fmt.Scanln(&divisor)

	fmt.Printf("result: %s", dividir(dividendo, divisor, ""))
}
