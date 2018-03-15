package main

import "fmt"

func main() {
	str := "Hello world!"

	for _, c := range str {
		fmt.Printf("%s %v %T\n", c, c, c)
	}
}
