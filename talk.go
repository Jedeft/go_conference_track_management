package main

import (
	"strconv"
	"strings"
)

// Talk talk struct
type Talk struct {
	Name     string // topic and time
	Duration int
	Status   bool // isSchedule
}

// Talks talk slice
type Talks []Talk

const (
	lightning = 5
)

func fillTalks(data []byte) (talks Talks, err error) {
	lines := strings.Split(string(data), "\n")
	talks = make(Talks, len(lines)-1)
	for i := range talks {
		talks[i] = Talk{Name: lines[i]}
		timeStr := lines[i][strings.LastIndex(lines[i], " ")+1:]
		if strings.Contains(timeStr, "min") {
			talks[i].Duration, err = strconv.Atoi(timeStr[:len(timeStr)-3])
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(timeStr, "lightning") {
			talks[i].Duration = lightning
		}
	}
	return talks, nil
}
