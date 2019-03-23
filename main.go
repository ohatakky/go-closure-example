package main

import "fmt"

func print_x_closure() func() {
	x := 0
	return func() {
		x = x + 1
		fmt.Println(x)
	}
}

func main() {
	pxc := print_x_closure()
	pxc()
	pxc()
	pxc()

	var x int
	defer func() {
		var x int
		fmt.Println(x)
	}()
	defer func() {
		fmt.Println(x)
	}()
	x = 1
}
