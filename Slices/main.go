package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Ayo"

	userNames = append(userNames, "fm")
	userNames = append(userNames, "samuel")
	fmt.Println(userNames)
}
