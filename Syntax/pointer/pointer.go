package main

import "fmt"

func foo() *int {
	bar := 123
	return &bar
}

func main() {
	var a int
	var b *int // Typ Pointer auf int
	a = 123
	b = &a // b erhält die Speicheradresse von a
	fmt.Println(b, *b)
	*b = 100 // Dereferenzierung
	fmt.Println(a)

	c := foo()
	fmt.Println(c, *c)
}
