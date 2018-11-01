package api

import (
	"bmstu-schedule-photo/parse"
	"fmt"

	"github.com/benbjohnson/phantomjs"
)

const dirForImages = "./"

// GetPhoto gets params (name of group and URL) and renders Web page with BMSTU Schedule. Then it saves it in .PNG format.
func GetPhoto(url, groupName string) error {
	p := phantomjs.NewProcess()

	page, err := p.CreateWebPage()
	if err != nil {
		return err
	}
	defer page.Close()

	// Open a URL.
	if err = page.Open(url); err != nil {
		return err
	}

	// Setup the viewport and render the results view.
	if err = page.SetViewportSize(1024, 600); err != nil {
		return err
	}

	// Set up photo options.
	options := phantomjs.Rect{
		Top:    150,
		Left:   20,
		Width:  1500,
		Height: 1370,
	}

	if err = page.SetClipRect(options); err != nil {
		return err
	}

	// Render a photo.
	name := fmt.Sprintf("%s%s.png", dirForImages, groupName)
	if err = page.Render(name, "png", 100); err != nil {
		return err
	}

	return nil
}

func GetAllPhoto(pathToJSON string) error {
	groups, err := parse.ParseJsonFile(pathToJSON)
	if err != nil {
		return err
	}

	for _, group := range *groups {
		err = GetPhoto(group.URL, group.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
