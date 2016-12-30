package main

import (
	"io/ioutil"

	"github.com/andreasf/cf-mysql-plugin/cfmysql/interactivity"
)

func main() {
	interactivityDetector := interactivity.FileStatDetector{}

	if interactivityDetector.IsInteractive() {
		ioutil.WriteFile("./output.txt", []byte("interactive mode"), 0777)
	} else {
		ioutil.WriteFile("./output.txt", []byte("pipe mode"), 0777)
	}
}
