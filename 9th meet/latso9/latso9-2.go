package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for x != 0 {
		fmt.Println(x % 10)
		x = x / 10
	}
}
