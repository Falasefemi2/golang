package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	sum := sumUp(1, 10, 15, -89)
	anotherSum := sumUp(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
