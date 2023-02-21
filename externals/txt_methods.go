package externals

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RickHPotter/upkeeper_site_checker_alura_course/internals"
)

func WriteToFile(filename string, text []string) {
	// CREATE THE FILE IF SUCH DOESN'T EXIST
	fn, e := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	// panic
	if e != nil {
		fmt.Println("Seems as if this writing-to-file didn't quite work, I'm afraid.")
		panic(e)
	}

	for _, line := range text {
		fn.WriteString(line)
	}
	fn.Close()
}

func ReadFromFile(filename string) []string {
	var text []string

	fn, e := os.OpenFile(filename, os.O_RDONLY, 0400)

	if e != nil {
		fmt.Println("Seems as if this reading-from-file didn't quite work, I'm afraid.")
		panic(e)
	}

	fileScanner := bufio.NewScanner(fn)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())
		text = append(text, line)
	}

	defer fn.Close()
	return text
}

// PROJECT'S SPECIFICS

func ReadSiteTxt() {
	text := ReadFromFile(internals.SitesPath)
	internals.Sites = append(internals.Sites, text...)
}

func WriteToLog(chunk []string) {
	// CREATE THE FILE IF SUCH DOESN'T EXIST
	log_fn := internals.LogPath + chunk[0][1:11] + ".txt"
	WriteToFile(log_fn, chunk)
}
