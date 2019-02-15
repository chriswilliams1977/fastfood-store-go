package main

func main() {

	burgers := menu{"Slim Jim", newBurger()}
	burgers = append(burgers, "Average Joe")

	//burgers is type menu
	//menu has func print
	burgers.print()

}

func newBurger() string {
	return "Big Devon"
}

func getBurger(menuSelection int) string {
	burger := ""
	switch menuSelection {
	case 1:
		burger = "Slim Jim"
	case 2:
		burger = "Average Joe"
	case 3:
		burger = "Big Devon"
	default:
		burger = "Big Devon"
	}
	return burger
}
