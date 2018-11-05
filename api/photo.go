package api

import (
	"bmstu-schedule-photo/parse"
	"fmt"
	"log"
	"os/exec"

	"github.com/benbjohnson/phantomjs"
)

var (
	p = phantomjs.NewProcess()
)

// getPhoto gets params (name of group, URL and outdir) and renders 
// Web page with BMSTU Schedule. Then it saves it in .PNG format.
func getPhoto(url, groupName, outdir string) {
	page, err := p.CreateWebPage()
	if err != nil {
		log.Print(err)
		return
	}

	// Open a URL.
	if err = page.Open(url); err != nil {
		log.Print(err)
		return
	}

	// Setup the viewport and render the results view.
	if err = page.SetViewportSize(1250, 1250); err != nil {
		log.Print(err)
		return
	}

	// Set up photo options.
	options := phantomjs.Rect{
		Top:    150,
		Left:   20,
		Width:  1220,
		Height: 1300,
	}
	if err = page.SetClipRect(options); err != nil {
		log.Print(err)
		return
	}

	// Render a photo.
	name := fmt.Sprintf("%s/tmp.png", outdir)
	if err = page.Render(name, "png", 100); err != nil {
		log.Print(err)
		return
	}

	finalName := fmt.Sprintf("%s/%s.png", outdir, groupName)
	if err := exec.Command("mv", name, finalName).Run(); err != nil {
		log.Print(err)
		return
	}

	return
}

// GetPhotos gets Groups info and call 
// getPhoto method for each, which renders Web pages with BMSTU 
// Schedule. Then it saves it in .PNG format.
func GetPhotos(groups *parse.Groups, outdir string) error {
	for _, group := range *groups {
		getPhoto(group.URL, group.Name, outdir)
	}
	return nil
}
