package docker

import (
	"bufio"
	"log"
	"os"
)

// Load env var from file
func LoadEnv(filepath string) []string {

	// // load file to byte array
	// content, err := ioutil.ReadFile(filepath)
	// if err != nil {
	// 	panic(err)
	// }
	// // convert []byte -> []string

	// return strings.Split(string(content), "\n")
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}
