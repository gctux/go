package main

import "fmt"

// Attribute der Klasse als struct
type rechteck struct {
	laenge int
	breite int
}

func (r rechteck) flaeche() int {
	return r.laenge * r.breite
}

// * um werte zu ändern
func (r *rechteck) setBreite(b int) {
	r.breite = b
}

func main() {
	r := rechteck{laenge: 10, breite: 5}
	r.setBreite(100)
	fmt.Println(r.flaeche())
}
