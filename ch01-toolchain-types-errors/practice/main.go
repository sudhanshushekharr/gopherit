package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil

}

func half(n int) (int, error) {
	if n%2 != 0 {
		return 0, errors.New("n is odd")
	}
	return n / 2, nil

}

func doubleIt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("doubleIt: %w", err)
	}
	return n * 2, nil
}

func main() {
	//var n int = nil
	//     a, b := 1, 2
	// c, b := 3, 4
	// a, b = 5, 6
	// a, b = 5, 6
	//     fmt.Println(a,b,c)

	// res, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)

	// res, err := half(7)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)

	res, err := doubleIt("oops")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
