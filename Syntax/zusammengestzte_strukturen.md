# Zusammengestzte Strukturen

- Schlüsselwort `struct`
- wird als neuer typ definiert

**Deklarationsmöglichkeiten:**

```
type adresse struct {
    strasse string
    stadt string
}
```

```
// Deklaration mit Feldnamen
a := adresse {
    strasse: "Musterstr. 1",
    stadt "0000 Musterstadt",
}
```

```
// Deklaration mit allen Feldern
b := adresse("Musterstr. 1", "0000 Musterstadt")
```

- bereits definierte Strukturen können auch in anderen Strukturen verwendet werden
```
type user struct {
    name string
    adresse1 adresse
    adresse2 adresse
}
```

- Deklaration mit adresse als **anonymer Struktur**
```
type user struct {
    name string
    adresse
}
```

- das anonyme Feld bekommt den Typnamen als bezeichner automatisch zugewiesen
- mögliche Deklaration

```
a := user {
    name : "Max",
    adresse: adresse {
        strasse: "Musterstr. 1",
        stadt: "0000 Musterstadt",
    },
}
```

## Anonyme Strukturen

- Strukturen können auch einfach bei der Zuweisung der Variablen definiert werden.
- Anwendungsfall: wenn die Struktur nur einmal verwendet wird