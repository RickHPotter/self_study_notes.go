package menu

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/RickHPotter/upkeeper_site_checker_alura_course/externals"
	"github.com/RickHPotter/upkeeper_site_checker_alura_course/internals"
)

// handling nothing relevant
func Intro() {
	fmt.Println("Hello, Sir.", internals.Name)
	fmt.Println("App version:", internals.Version)
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
		StartMonitoring(internals.Delay)
	case 2:
		ShowLogs()
	case 3:
		DisplayTodaysLogs()
	case 4:
		Exiting()
	default:
		fmt.Println("How did you even get here?")
	}
}

// handling monitoring of many sites
func StartMonitoring(seconds int) {
	fmt.Println("Monitoring...")
	externals.ReadSiteTxt()

	for {
		var line []string
		for _, site := range internals.Sites {
			line = append(line, externals.HTTPGet(site))
		}
		time.Sleep(time.Duration(seconds) * time.Second)
		externals.WriteToLog(line)
	}

}

// for now, just open file explores with all logs
func ShowLogs() {
	fmt.Println("Showing Logs...")
	exec.Command("explorer", internals.LogPath).Start()
}

func DisplayTodaysLogs() {
	dateToday := time.Now().Format("2006-01-02")
	fn := internals.LogPath + dateToday + ".txt"

	var text []string = externals.ReadFromFile(fn)
	for _, line := range text {
		fmt.Println(line)
	}

}

// pretty straight-forward
func Exiting() {
	fmt.Println("Exiting...")
	os.Exit(0)
}
