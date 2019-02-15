package main

import (
	"fmt"
)

//Create new type called menu
//slice of stings (same behaviour of []strings)
type menu []string

func (m menu) print() {
	for i, item := range m {
		fmt.Println(i, item)
	}
}
