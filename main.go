package main

import (
	"fmt"
	"log"
	"os"

	"bmstu-schedule-photo/api"
	"bmstu-schedule-photo/parse"

	"github.com/benbjohnson/phantomjs"
)

var (
	instructionText = fmt.Sprintf(
		`-all [path to json file with urls] [outdir]
-u [url] [group_name] [outdir]
`,
	)
)

func main() {
	if !(3 <= len(os.Args) && len(os.Args) <= 5) {
		fmt.Printf(instructionText)
		return
	}

	// Start the process once.
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		log.Panic(err)
	}
	defer phantomjs.DefaultProcess.Close()

	// Parse of arguments
	var groups *parse.Groups
	var err error
	var outdir string

	switch os.Args[1] {
	case "-all":
		if groups, err = parse.ParseJsonFile(os.Args[2]); err != nil {
			log.Panic(err)
		}
		outdir = os.Args[3]
	case "-u":
		groups = &parse.Groups{
			&parse.Group{
				URL:  os.Args[2],
				Name: os.Args[3],
			},
		}
		outdir = os.Args[4]
	default:
		fmt.Printf(instructionText)
		return
	}

	if err := api.GetPhotos(groups, outdir); err != nil {
		log.Panic(err)
	}
}
