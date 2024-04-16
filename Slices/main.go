package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) outPut() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Ayo"

	userNames = append(userNames, "fm")
	userNames = append(userNames, "samuel")
	fmt.Println(userNames)

	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["angular"] = 4.8

	courseRating.outPut()

	// fmt.Println(courseRating)

	for index, value := range userNames {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	for key, value := range courseRating {
		fmt.Println("value:", value)
		fmt.Println("key:", key)
	}
}
