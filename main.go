package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	data, err := loadFile("input.txt")
	if err != nil {
		log.Errorf("load File fail : %s", err)
		os.Exit(1)
	}
	talks, err := fillTalks(data)
	if err != nil {
		log.Errorf("fill talks fail : %s", err)
	}
	fmt.Println(talks)
}
