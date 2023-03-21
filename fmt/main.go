package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hallo ", "Print()\n")
	fmt.Println("Hallo Println()")
	var s = "Printf()"
	fmt.Printf("Hallo, %s\n", s)
	fmt.Println()
	// Output:
	// Hallo Print()
	// Hallo Println()
	// Hallo, Printf()

	var nr = 2
	var name = "Lutz"
	fmt.Printf("%03d: Hallo, %s\n", nr, name)
	fmt.Println()
	// Output:
	// 002: Hallo, Lutz

	/* Verben
	%v Die Ausgabe des Wertes einer beliebigen Variablen. Dabei kann jeder Typ über dieses Verb ausgegeben werden
	%+v Gibt bei Strukturen zusätzlich die Namen der einzelnen Felder aus.
	%#v Die Ausgabe erfolgt so, wie wir diese als Go-Code definieren würden.
	%T Gibt den Typ der Variablen aus.
	%b Ausgabe zur Basis 2, also als binäre Zahl. Dieses Verb ist nur für Integer zulässig
	%d Ausgabe zu Basis 10
	%03d oder %03b: Zahlen werden mit 0 aufgefüllt
	%p Gibt die Adresse im Speicher aus, also den Pointer der Variablen
	%s Direkte Ausgabe eines String
	%x Ausgabe eines String als  hexadezimaler Wert
	https://golang.org/pkg/fmt/ */

	s1 := fmt.Sprintf("%03d: Hallo, %s\n", nr, name)
	fmt.Print(s1)
	fmt.Println()
	// Output:
	// 002: Hallo, Lutz

	// Schreiben in Datei:
	file, _ := os.Create("datei.txt")
	fmt.Fprintf(file, "%03d: Hallo, %s\n", nr, name)
	file.Close()

}
