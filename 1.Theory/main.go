package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	// something.SayHello()

	// const name string = "donghee"  -> const임으로 재할당 불가능
	// or
	name := "donghee"
	fmt.Println(name)

	var nick string = "dongryong"
	nick = "ddong"
	fmt.Println(nick)

	fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper(name)
	// fmt.Println(totalLength, upperName)

	totalLength, _ := lenAndUpper(name) // 두개의 리턴 값중 하나를 무시하겠다.
	fmt.Println(totalLength)

	repeatMe("dong", "hee", "lee")

	_ = superAdd(1, 2, 3, 4, 5, 6, 7)
	// fmt.Println(total)
}

// 변수의 타입과 리턴의 타입을 정해주어야함
func multiply(a, b int) int {
	return a * b
}

// go는 한개이상의 값을 return할  수 있다

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// 여러개의 arguments를 받는법 '...'
func repeatMe(words ...string) {
	fmt.Println(words)
}

//naked return
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done!") // defer return이 된후 실행한다. ex) 파일을 닫는다든거 등에서 활용 가능

	// 위에서 만들어저있기 때문에 length := 는 사용하지 않는다
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// for loop and range
func superAdd(numbers ...int) int {
	fmt.Println(numbers)

	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
