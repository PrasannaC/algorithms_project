package FileUtils

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	stringData := string(data[:])
	return strings.Split(stringData, "\n")
}
