package main

import (
	"github.com/RickHPotter/upkeeper_site_checker_alura_course/internals"
	"github.com/RickHPotter/upkeeper_site_checker_alura_course/menu"
)

func main() {
	menu.Intro()

	var menuChoice int
	menu.ShowMenu(internals.InitialMenu)
	menu.GetChoice(internals.InitialMenu, &menuChoice)
}
