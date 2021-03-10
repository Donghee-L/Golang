package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Scan(&n1, &n2)
	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)

}

// func Scan(a ...interface{}) (n int, err error): 콘솔에서 공백, 새 줄로 구분하여 입력을 받음
// func Scanln(a ...interface{}) (n int, err error): 콘솔에서 공백으로 구분하여 입력을 받음
// func Scanf(format string, a ...interface{}) (n int, err error): 콘솔에서 형식을 지정하여 입력을 받음
