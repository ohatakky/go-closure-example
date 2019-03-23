package main

import "fmt"

func main() {
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
