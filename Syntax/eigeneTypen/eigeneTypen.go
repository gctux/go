package main

import "fmt"

type (
	meter      int
	zentimeter int
)

func MeterToZentimeter(m meter) zentimeter {
	return zentimeter(m * 100)
}

func main() {
	var m meter
	var z zentimeter
	m = 3
	z = MeterToZentimeter(m)
	fmt.Printf("%v Meter sind %v Zentimeter.\n", m, z)
}
