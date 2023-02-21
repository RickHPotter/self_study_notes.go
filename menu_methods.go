package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// handling nothing relevant
func Intro() {
	fmt.Println("Hello, Sir.", name)
	fmt.Println("App version:", version)
}

// handling menu listing
func ShowMenu(m map[int]string) {
	for i := 1; i <= len(m); i++ {
		fmt.Println(i, m[i])
	}
}

// handling birth defect (aka user)
func GetChoice(menu map[int]string, menuChoice *int) {
	for {
		fmt.Scan(menuChoice)
		_, ok := menu[*menuChoice]

		if ok {
			InitialMenu(*menuChoice)
			break
		}
		fmt.Println("Man, you must be tripping. Choose a right option, would you.")
	}
}

// handling user interaction
func InitialMenu(menuChoice int) {
	switch menuChoice {
	case 1:
		StartMonitoring(delay)
	case 2:
		ShowLogs()
	case 3:
		Exiting()
	default:
		fmt.Println("How did you even get here?")
	}
}

// handling monitoring of many sites
func StartMonitoring(seconds int) {
	fmt.Println("Monitoring...")
	ReadSiteTxt()

	for {
		for i := range sites {
			fmt.Println(i, sites[i])
			HTTPGet(sites[i])
		}
		time.Sleep(delay * time.Second)
	}

}

func ShowLogs() {
	fmt.Println("Showing Logs...")
	exec.Command("explorer", LogPath).Start()
	// IMPLEMENT
}

func Exiting() {
	fmt.Println("Exiting...")
	os.Exit(0)
	// IMPLEMENT
}
