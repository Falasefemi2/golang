package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":             "https://google.com",
		"Amazon web service": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon web service"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
