package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

// mehrere Parameter desselben Typs
func add(a, b, c int) int {
	return a + b + c
}

// mehrere Rückgabewerte
func swap(a, b int) (int, int) {
	return b, a
}

// benannte Rückgabeparameter
func foo(s string) (a string, b string) {
	b = s       // keine Deklaration notwendig
	return a, b // a =''
}

// varadische Funktionen
// weitere Parameter müssen vor dem varadischen Parameter stehen
func addmore(a ...int) int {
	summe := 0
	for _, v := range a {
		summe = summe + v
	}
	return summe
}

// Funktion wird als Typ definiert
type filterFunc func(string) bool

func meinFilter(s []string, filter filterFunc) []string {
	var out []string
	for _, str := range s {
		if filter(str) {
			out = append(out, str)
		}
	}
	return out
}

// defer Anweisungen hinter defer werden erst nach dem beenden der Funktion ausgeführt
func meineFunktion() {
	//anonyme Funktion
	defer func() {
		fmt.Println(1)
		fmt.Println(2)
	}()
	fmt.Println(3)
}

func main() {
	greeting := greet("Alice")
	fmt.Println(greeting) //Output Hallo Alice
	fmt.Println()

	a := 1
	b := 2
	c := 3
	fmt.Printf("%d + %d + %d = %d\n", a, b, c, add(a, b, c)) // 1 + 2 + 3 = 6
	fmt.Println()

	d, e := swap(12, 2)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println()
	// Output:
	// 2
	// 12

	h, g := foo("Hallo Welt!")
	fmt.Println(g) //leer
	fmt.Println(h) // Hallo Welt!
	fmt.Println()

	// varadische Funktion
	fmt.Println(addmore(1, 2, 3, 4))
	fmt.Println(addmore())
	fmt.Println()

	//Verwendung der Funktion als Typ
	f := func(s string) bool {
		if len(s) < 3 {
			return false
		}
		return true
	}
	s := []string{"a", "abcd", "abc", "ab"}
	fmt.Println(meinFilter(s, f))

	fmt.Println()
	meineFunktion()
}
