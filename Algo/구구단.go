package test

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i < 10; i++ {
		// result := i * n
		// fmt.Println(n, "*", i, "=", result)

		var result string
		result = strconv.Itoa(n) + " * " + strconv.Itoa(i) + " = " + strconv.Itoa(n*i)
		fmt.Println(result)
	}

}
