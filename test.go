package main

import (
	"errors"
	"fmt"
)

func main2() {
	var a = 123
	fmt.Printf("a类型: %T\n", a)
	fmt.Printf("a值: %v\n", a)
	fmt.Printf("a地址: %p\n", &a)
	fmt.Printf("&a地址: %T\n", &a)
	p := &a
	*p = 456
	fmt.Printf("p值: %v\n", p)
	fmt.Printf("&p值: %v\n", &p)
	var p2 *int
	fmt.Println(p2)

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr2)

	fmt.Println(calc(2432, 53))
}

func calc(a, b int) (sum,subtract,multiply int, divide float64, remainder int, err error) {
	sum = a + b
	subtract = a - b
	multiply = a * b
	if b == 0 {
		err = errors.New("divide by zero")
	} else {
		divide = float64(a) / float64(b)
		remainder = a % b
	}
	return
}




