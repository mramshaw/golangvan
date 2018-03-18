package main

import "fmt"

func main() {
	str := "Hello world!"

	for _, c := range str {
		fmt.Printf("%[01]s %[01]v %[01]T\n", c)
	}
}
