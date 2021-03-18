package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"ramen", "kimchi"}
	donghee := person{"donghee", 27, favFood}
	dong := person{name: "dong", age: 27, favFood: favFood}
	fmt.Println(donghee)
	fmt.Println(dong)
}
