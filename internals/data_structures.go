package internals

import "os"

// INTRO
var Name = "Rick"
var Version float32 = 1.2

// MONITORING
const Delay = 5

// MENU
var InitialMenu = map[int]string{
	1: "Start Monitoring",
	2: "Show Logs Folder",
	3: "Display Today Logs",
	4: "Exit",
}

// EXPLORER PATH
var HomeDir, _ = os.UserHomeDir()
var AssetsPath = HomeDir + "\\go\\src\\github\\upkeeper_site_checker_alura_course\\assets\\"

// SITES
var SitesPath string = AssetsPath + "sites.txt"
var Sites []string

// LOGS
var LogPath string = AssetsPath + "logs\\"
