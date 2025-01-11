package main

import "fmt"

func main() {
	var use, pass string
	var x int
	var A bool
	for A == false {
		fmt.Scan(&use, &pass)
		if use == "Admin" && pass == "Admin" {
			A = true
		} else {
			x += 1
		}
	}
	fmt.Print(x, " percobaan gagal login")
}
