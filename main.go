package main

import (
	"fmt"
	"os"

	"bmstu-schedule-photo/api"
	"bmstu-schedule-photo/config"
	"bmstu-schedule-photo/parse"
	"bmstu-schedule-photo/transformations"

	"github.com/benbjohnson/phantomjs"
	log "github.com/kataras/golog"
)

var (
	instructionText = fmt.Sprintf(
		`
-c [path to config]
[ -all [path to json file with urls] [outdir] ], [ -u [url] [group_name] [outdir]]`)
)

func main() {
	if !(5 <= len(os.Args) && len(os.Args) <= 7) {
		fmt.Printf(instructionText)
		return
	}

	if os.Args[1] != "-c" {
		fmt.Printf(instructionText)
		return
	}

	cfg := config.LoadConfig(os.Args[2])
	transformations.LoadTransformations(cfg.TransformationsPath)
	// Start the process once.
	if err := phantomjs.DefaultProcess.Open(); err != nil {
		panic(err)
	}
	defer phantomjs.DefaultProcess.Close()
	log.Info("PROCESS IS STARTED...")

	// Parse of arguments
	var groups *parse.Groups
	var err error
	var outdir string

	switch os.Args[3] {
	case "-all":
		if groups, err = parse.ParseJsonFile(os.Args[4]); err != nil {
			panic(err)
		}
		outdir = os.Args[5]
	case "-u":
		groups = &parse.Groups{
			&parse.Group{
				URL:       os.Args[4],
				GroupName: os.Args[5],
			},
		}
		outdir = os.Args[6]
	default:
		fmt.Printf(instructionText)
		return
	}

	if err := api.GetPhotos(groups, outdir); err != nil {
		panic(err)
	}
}
