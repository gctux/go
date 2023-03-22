# Konstanten

- Deklaration wie Variablen
```
// einfache Deklaration
const a = 10

// Deklaration als Gruppe 
const (
    maxBreite = 100
    maxLaenge = 100
)
```

## Automatischer Zähler ` iota `

```
const (
    Montag = iota       // 0
    Dienstag = iota     // 1
    Mittwoch = iota     // 2
    Donnerstag          // 3
    Freitag             // 4
    Samstag             // 5
    Sonntag             // 6
)

// Beginnt hier wider bei 0
const (
    _ = iota            // 0 ausgelassen
    eins
    _                   // 2 ausgelassen
    drei
)

```