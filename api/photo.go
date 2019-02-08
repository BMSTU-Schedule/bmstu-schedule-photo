package api

import (
	"fmt"
	"os/exec"

	"github.com/benbjohnson/phantomjs"
	log "github.com/kataras/golog"

	"bmstu-schedule-photo/parse"
	transform "bmstu-schedule-photo/transformations"
)

// getPhoto gets params (name of group, URL and outdir) and renders
// Web page with BMSTU Schedule. Then it saves it in .PNG format.
func getPhoto(url, groupName, outdir string) {
	p := phantomjs.NewProcess()

	page, err := p.CreateWebPage()
	if err != nil {
		log.Error(err)
		return
	}

	defer func() {
		p.Close()
		page.Close()
	}()

	// Open a URL.
	if err = page.Open(url); err != nil {
		log.Error(err)
		return
	}

	// Run JS Scripts
	for _, script := range transform.JSScripts.Queue {
		log.Infof("Running: %s", script)
		page.Evaluate(script)
	}

	text, _ := page.PlainText()

	// Setup the viewport and render the results view.
	if err = page.SetViewportSize(1248, len(text)); err != nil {
		log.Error(err)
		return
	}

	// Set up photo options.
	options := phantomjs.Rect{
		Top:    150,
		Left:   0,
		Width:  1230,
		Height: 880 + int(float64(len(text))*0.24),
	}
	if err = page.SetClipRect(options); err != nil {
		log.Error(err)
		return
	}

	// Render a photo.
	name := fmt.Sprintf("%s/tmp.png", outdir)
	if err = page.Render(name, "png", 70); err != nil {
		log.Error(err)
		return
	}

	finalName := fmt.Sprintf("%s/%s.png", outdir, groupName)
	if err := exec.Command("mv", name, finalName).Run(); err != nil {
		log.Error(err)
		return
	}

	log.Infof("%s IS DOWNLOADED...", groupName)
	return
}

// GetPhotos gets Groups info and call
// getPhoto method for each, which renders Web pages with BMSTU
// Schedule. Then it saves it in .PNG format.
func GetPhotos(groups *parse.Groups, outdir string) error {
	for _, group := range *groups {
		log.Info(group.GroupName)
		getPhoto(group.URL, group.GroupName, outdir)
	}
	return nil
}
