package transformations

import (
	"io/ioutil"
	"os"
)

var JSScripts = struct {
	Queue []string
}{
	Queue: make([]string, 0),
}

func LoadTransformations(pathToScripts string) {

	files, err := ioutil.ReadDir(pathToScripts)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		scriptFile, _ := os.Open(pathToScripts + file.Name())
		script, _ := ioutil.ReadAll(scriptFile)
		JSScripts.Queue = append(JSScripts.Queue, string(script))
	}
}
