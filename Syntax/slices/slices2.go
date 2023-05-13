package main

import "fmt"

func printSlice(name string, slice []int) {
	fmt.Println(name, ":")

	for _, s := range slice {
		fmt.Printf("%p - len: %d cap: %d - %#v\n", s, len(s), cap(s), s)
	}
}
