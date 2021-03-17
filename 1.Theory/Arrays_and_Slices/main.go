package main

import "fmt"

func main() {
	// names := [5]string{"dong", "dal"}
	// names[2] = "lyn"
	// names[3] = "james"
	names := []string{}
	names = append(names, "dong")
	fmt.Println(names)

}

// slice : 길이가 없는 array -> []string{}
