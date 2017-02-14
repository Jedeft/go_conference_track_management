package main

import (
	"strconv"
	"strings"
)

// Talk talk struct
type Talk struct {
	Topic      string // topic and time
	Duration   int
	IsSchedule bool
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
		talks[i] = Talk{Topic: lines[i]}
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

func (talks Talks) getTotalDuration() int {
	var totalDuration int
	for _, talk := range talks {
		totalDuration += talk.Duration
	}
	return totalDuration
}

func (talks Talks) setSchedule(scheduleTalks Talks) {
	scheduleMap := make(map[string]bool, len(scheduleTalks))
	for _, scheduleTalk := range scheduleTalks {
		scheduleMap[scheduleTalk.Topic] = true
	}
	for ti := range talks {
		if _, exist := scheduleMap[talks[ti].Topic]; exist {
			talks[ti].IsSchedule = true
		}
	}
}

// isSchedule all talk isSchedule
func (talks Talks) isSchedule() bool {
	for _, talk := range talks {
		if !talk.IsSchedule {
			return false
		}
	}
	return true
}
