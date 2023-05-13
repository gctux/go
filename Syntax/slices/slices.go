package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%p - len: %d cap: %d %#v\n", s, len(s), cap(s), s)
}

func main() {
	var a []int
	printSlice(a)
	a = append(a, 1)
	printSlice(a)
	a = append(a, 2)
	printSlice(a)
	a = append(a, 3)
	printSlice(a)

	// Composite Literal
	b := []int{}
	printSlice(b)
	c := []int{1, 2, 3}
	printSlice(c)

	d := make([]int, 2, 8)
	printSlice(d)

	s := []string{"null", "eins", "zwei"}

	for i, v := range s {
		fmt.Printf("%02d: %s\n", i, v)
	}

}
