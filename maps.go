package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}
	fmt.Println("websites-", websites)
	fmt.Println(websites["Amazon Web Services"])
}