package externals

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// handling status code of each site
func HTTPGet(site string) string {
	resp, e := http.Get(site)

	if e != nil {
		panic(e)
	}

	if resp.StatusCode == 200 {
		fmt.Println(site, "is working fine.")

	} else {
		fmt.Println(site, "is having problems. SC:", resp.StatusCode)
	}

	// writing to log in chunks of len(sites.txt) instead of one by one
	date := time.Now().Format("2006-01-02 15:04:05")
	line := "[" + date + "] " + strconv.Itoa(resp.StatusCode) + " - " + site + "\n"
	return line

}
