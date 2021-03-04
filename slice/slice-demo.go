package main

import (
	"fmt"
)

func main() {
	// array composite literal
	// default len and cap of 5
	arr := [5]int{0, 1, 3, 5, 9}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println(len(arr), cap(arr))
	fmt.Printf("type: %T\n\n", arr)

	// slice composite literal
	// default len and cap of 10
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Print("slice: \n")
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
	fmt.Printf("type: %T\n\n", arr)

	// slicing
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Printf("\n\n")

	// append to modify
	fmt.Println("Append")
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x, y)
	fmt.Print("\n\n")

	// delete by re-slicing to preserve order
	fmt.Println("Delete")
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	b := append(a[:3], a[6:]...)
	fmt.Println("mutated original: ", a)
	fmt.Println("deleted result: ", b)
	// delete by swapping each elements to be removed, order not preserved
	a = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a[2] = a[len(a)-1]
	b = a[:len(a)-1]
	fmt.Println(b)
	fmt.Print("\n\n")

	// make to initialize length and capacity; avoid frequent allocation
	states := make([]string, 50, 100)
	for i, v := range []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`,
		` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`,
		` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`,
		` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`,
		` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`,
		` West Virginia`, ` Wisconsin`, ` Wyoming`,
	} {
		states[i] = v
	}

	for i := 0; i < len(states); i++ {
		fmt.Printf("%s, ", states[i])
	}
	fmt.Println("\nlen: ", len(states))
	fmt.Println("cap: ", cap(states))
	fmt.Print("\n\n")

	// multi dimensional slice
	matrix := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(matrix)
	for i, xs := range matrix {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

	// maps, speical type of slice
	m := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`batman`:     []string{`Bruce Wayne`, `Billionaire`, `Vigilante`},
		`ironman`:    []string{`Tony Stark`, `Pepper`, `AI`},
	}
	for k, v := range m {
		fmt.Printf("key(index): %s\nvalue: %v\n", k, v)
	}
}
