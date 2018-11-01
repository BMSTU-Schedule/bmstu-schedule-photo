package main

import (
	"bmstu-schedule-photo/api"
	"fmt"
	"log"
	"os"

	"github.com/benbjohnson/phantomjs"
)

var (
	instructionText = fmt.Sprintf("-all [path to json file with urls]\n-u   [url] [group_name]\n")
)

func main() {
	if !(len(os.Args) == 3 || len(os.Args) == 4) {
		fmt.Printf(instructionText)
		return
	}

	// Start the process once.
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer phantomjs.DefaultProcess.Close()

	// Parse of arguments
	switch os.Args[1] {
	case "-all":
		if err := api.GetAllPhoto(os.Args[2]); err != nil {
			log.Panic(err)
		}
	case "-u":
		if err := api.GetPhoto(os.Args[2], os.Args[3]); err != nil {
			log.Panic(err)
		}
	default:
		fmt.Printf(instructionText)
		return
	}
}
