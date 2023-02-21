package main

import (
	"fmt"
	"net/http"
	"time"
)

// handling status code of each site
func HTTPGet(site string) {
	fmt.Println("[HTTP]", site)
	resp, e := http.Get(site)

	if e != nil {
		fmt.Println(e)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "is working fine.")
		WriteToLog(site, resp.StatusCode, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println(site, "is having problems. SC:", resp.StatusCode)
	}
}
