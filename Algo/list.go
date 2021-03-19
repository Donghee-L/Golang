package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i]) // Fscan은 space로 나누어서 읽음
	}

	// sol1) sort.Slice를 활용하는 방법
	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] < arr[j]
	// })

	// sol2) for문으로 min max if 사용?

	max_value := -1000000
	min_value := 1000000

	for i := 0; i < n; i++ {
		if arr[i] > max_value {
			max_value = arr[i]
		}

		if arr[i] < min_value {
			min_value = arr[i]
		}
	}

	fmt.Println(min_value, max_value)

}
