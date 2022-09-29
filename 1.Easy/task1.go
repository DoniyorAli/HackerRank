package main

import "fmt"

func main() {

	var num1, num2 uint32
	fmt.Scanf("%d\n%d", &num1, &num2)
	total := add(num1,num2)
	fmt.Printf("Total: %d\n", total)
	
}

func add(a, b uint32) uint32 {
	return a + b
}