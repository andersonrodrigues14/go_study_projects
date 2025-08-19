package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	courseRatings.output()

	fmt.Println(courseRatings)

	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
