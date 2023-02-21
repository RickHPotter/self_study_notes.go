package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadSiteTxt() {
	filename, e := os.Open(sitesPath)

	if e != nil {
		fmt.Println("Seems as if this didn't quite work, I'm afraid.")
		fmt.Println(e)
		return
	}

	fileScanner := bufio.NewScanner(filename)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		sites = append(sites, line)
	}

	filename.Close()
}

func WriteToLog(site string, statusCode int, date string) {

}
