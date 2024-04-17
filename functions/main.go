package main

import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)
// 	fmt.Println(doubled)
// 	fmt.Println(tripled)
// }

// func transformNumbers(numbers *[]int, transform func(int) int) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}
// 	return dNumbers
// }

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David"}

	greetedNames := transformNames(&names, addGreeting)
	fmt.Println(greetedNames)

}

func transformNames(names *[]string, t func(string) string) []string {
	dWords := []string{}
	for _, v := range *names {
		dWords = append(dWords, t(v))
	}
	return dWords
}

func addGreeting(name string) string {
	return "Hello, " + name
}
