package main

import log "github.com/Sirupsen/logrus"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	scheduleConference("input.txt")
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
	getScheduleConferences(talks)
}
