package main

import "fmt"

func main() {
	m := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}

	for key, value := range m {
		fmt.Println(value)
	}
}
