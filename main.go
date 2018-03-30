package main

import "fmt"

func main() {
	str := "Hello world!"

	for _, c := range str {

        // Technically, c is a rune, so 'go tool vet' will complain
        // So lets use appropriate format verbs for a rune
        // Specifically hexadecimal (%x) and Unicode (%U)

		//fmt.Printf("%[01]s %[01]v %[01]T\n", c)
		fmt.Printf("%[01]x %[01]U %[01]v %[01]T\n", c)
	}
}
