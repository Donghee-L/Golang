package main

import "fmt"

func main() {
	dong := map[string]string{"name": "dong", "age": "12"}
	fmt.Println(dong)
	for key, value := range dong {
		fmt.Println(key, value)
	}
}
