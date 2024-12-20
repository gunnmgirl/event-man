package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}
	fmt.Println("websites-", websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["Dzana"] = "gunnmgirl.com"
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
	// struct vs map
	// maps can have any type for a key (integer, struct, string)
	// struct we use when we have a clearly defined data
	// map we use when we have a collection of values with custom labels (keys)
}
