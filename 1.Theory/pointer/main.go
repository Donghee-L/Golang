package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10
	a = 6
	*b = 29
	fmt.Println(a, *b)

}

// 변수앞에 &를 붙일시 메모리 주소를 확인할 수 있다.  b := &a -> b는 a의 주소를 바라몬다.
// "*" -> 메모리 주소의 값을 확인
