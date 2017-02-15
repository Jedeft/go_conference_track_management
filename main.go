package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input one fileName")
		return
	}
	if err := scheduleConference(os.Args[1]); err != nil {
		log.Print(err)
	}
}

func scheduleConference(filePath string) error {
	data, err := loadFile(filePath)
	if err != nil {
		return err
	}
	talks, err := fillTalks(data)
	if err != nil {
		return err
	}
	conferences, err := getScheduleConferences(talks)
	conferences.output()
	return err
}
