package main

import (
	"fmt"
)

func main() {
	// maps, speical type of slice
	m := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`batman`:     []string{`Bruce Wayne`, `Billionaire`, `Vigilante`},
		`ironman`:    []string{`Tony Stark`, `Pepper`, `AI`},
	}
	for k, v := range m {
		fmt.Printf("key(index): %s\nvalue: %v\n", k, v)
	}
	fmt.Println()

	// add to map
	m["goku"] = []string{`dragon ball`, `super seiyan`, `monkey`}
	for k, v := range m {
		fmt.Printf("key(index): %s\nvalue: %v\n", k, v)
	}
	fmt.Println()

	// remove from map
	delete(m, "bond_james")
	for k, v := range m {
		fmt.Printf("key(index): %s\nvalue: %v\n", k, v)
	}
	fmt.Println()
}
