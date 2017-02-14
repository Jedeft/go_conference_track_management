package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	if len(os.Args) < 2 {
		fmt.Println("Please input one fileName")
		os.Exit(1)
	}
	scheduleConference(os.Args[1])
}

func scheduleConference(filePath string) {
	data, err := loadFile(filePath)
	if err != nil {
		panic(err)
	}
	talks, err := fillTalks(data)
	if err != nil {
		panic(err)
	}
	conferences := getScheduleConferences(talks)
	conferences.output()
}
