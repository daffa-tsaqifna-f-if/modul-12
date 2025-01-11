package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y)
	for x >= y {
		x = x - y
		z++
	}
	fmt.Print(z)
}
