package main

import "fmt"

func main()  {
	var ch byte = 'A'
	var xch byte = '\x41'
	var uch byte = '\u0041'

	fmt.Println(ch)
	fmt.Printf("%c\n", ch)

	fmt.Printf("%d - %d - %d\n", ch, xch, uch) // integer
	fmt.Printf("%c - %c - %c\n", ch, xch, uch) // character
	fmt.Printf("%X - %X - %X\n", ch, xch, uch) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, xch, uch) // UTF-8 code point
}