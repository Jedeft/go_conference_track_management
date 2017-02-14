package main

import (
	"io/ioutil"
	"os"
)

// laodFile Read data from file
func loadFile(path string) (data []byte, err error) {
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
