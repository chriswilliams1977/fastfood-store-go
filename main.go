package main

import (
	"fmt"
)

func main() {

	//get burger type
	//fmt.Println(getBurger(2))

	//get menu items by tyep
	order := getMenuItems("")
	for _, orderItem := range order {
		fmt.Println(orderItem)
	}
}
