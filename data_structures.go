package main

import "os"

// INTRO
var name = "Rick"
var version float32 = 1.2

// MONITORING
const delay = 5

// MENU
var initialMenu = map[int]string{
	1: "Start Monitoring",
	2: "Show Logs",
	3: "Exit",
}

// EXPLORER PATH
var homeDir, _ = os.UserHomeDir()
var assetsPath = homeDir + "\\go\\src\\module_un\\assets\\"

// SITES
var sitesPath string = assetsPath + "sites.txt"
var sites []string

// LOGS
var LogPath string = assetsPath + "logs"
