package main

import "fmt"

func main() {
	var x int
	defer func() {
		fmt.Println(x)
	}()
	x = 1
}
