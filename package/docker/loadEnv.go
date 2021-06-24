package docker

import (
	"io/ioutil"
	"strings"
)

// Load env var from file
func LoadEnv(filepath string) []string {

	// load file to byte array
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// convert []byte -> []string
	return strings.Split(string(content), "\n")
}
