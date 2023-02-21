package main

func main() {
	Intro()

	var menuChoice int
	ShowMenu(initialMenu)
	GetChoice(initialMenu, &menuChoice)
}
