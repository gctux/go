Definition (Typ nach Namen, bei Wertzuweisung kein Typ nötig)

var nummer int                      //Definition ohne Wert
var name string = "Obiwan Kenobi"   //Definition mit Wert und Typ
var anzahl = 10                     //Definition ohne Typ


Standardwerte bei der Definition

var nummer int                      // 0
var txt string                      // ""
var checked bool                    // false
var meinUser *meinUser              // nil - Pointer
var liste []string                  // nil - liste


Gruppierung von Variablen

var (
    home = os.Getenv("HOME")
    user = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)


Kurzdeklaration

nummer := 10                        // int
name := getName()                   // string


Mehrere Variablen in einer Anweisung

a, b := 12, 34
a, b = b, a                         // Werte tauschen


Deklarierte variablen müssen im Code verewendet werden!


Variablennamen
- keine Leerzeichen
- Buchstaben, Ziffern, Unterstrich
- dürfen nicht mit Ziffer beginnen
- öffentliche Bezeichner beginnen mit Großbuchstaben, private Bezeichner mit Kleinbuchstaben
- Länge der Variablen sollte sich nach Verwendung richten
- lange Variablen: Kamelhöckerschreibweise 
- Abkürzungen Großbuchstaben z.B. HTTPServer

