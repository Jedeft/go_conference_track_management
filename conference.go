package main

import (
	"errors"
	"fmt"
	"time"
)

// Conference conference struct
type Conference struct {
	MorningSession   Talks
	AfternoonSession Talks
}

// Conferences conference slice
type Conferences []Conference

const (
	dayMaxDuration     = 420
	dayMinDuration     = 360
	sessionMinDuration = 180
	sessionMaxDuration = 240
)

func getScheduleConferences(talks Talks) Conferences {
	totalDuration := talks.getTotalDuration()
	totalDay := totalDuration/dayMaxDuration + 1
	conferences := make(Conferences, totalDay)
	conferences.setMorningSession(talks)
	conferences.setAfternoonSession(talks)
	if !talks.isSchedule() {
		conferences.checkSession(talks)
		if !talks.isSchedule() {
			panic(errors.New("input talks can not compose the conferences"))
		}
	}
	return conferences
}

// setMorningSession set possibale MorningSession
func (conferences Conferences) setMorningSession(talks Talks) {
	for ci := range conferences {
		for ti := range talks {
			var morningSession Talks
			var totalDuration int
			for tj := ti; tj < len(talks); tj++ {
				currentIndex := tj
				if talks[currentIndex].IsSchedule {
					continue
				}
				if talks[currentIndex].Duration > sessionMinDuration ||
					talks[currentIndex].Duration+totalDuration > sessionMinDuration {
					continue
				}
				morningSession = append(morningSession, talks[currentIndex])
				totalDuration += talks[currentIndex].Duration
				if totalDuration == sessionMinDuration ||
					(len(conferences) == 1 && talks.getTotalDuration() < sessionMinDuration) {
					break
				}
			}
			if totalDuration == sessionMinDuration ||
				(len(conferences) == 1 && talks.getTotalDuration() < sessionMinDuration) {
				conferences[ci].MorningSession = morningSession
				talks.setSchedule(morningSession)
				break
			}
		}
	}
}

// setAfternoonSession set possibale AfternoonSession
func (conferences Conferences) setAfternoonSession(talks Talks) {
	for ci := range conferences {
		for ti := range talks {
			var afternoonSession Talks
			var totalDuration int
			for tj := ti; tj < len(talks); tj++ {
				currentIndex := tj
				if talks[currentIndex].IsSchedule {
					continue
				}
				if talks[currentIndex].Duration > sessionMaxDuration ||
					talks[currentIndex].Duration+totalDuration > sessionMaxDuration {
					continue
				}
				afternoonSession = append(afternoonSession, talks[currentIndex])
				totalDuration += talks[currentIndex].Duration
				if totalDuration >= sessionMinDuration && totalDuration <= sessionMaxDuration ||
					(len(conferences) == 1 && talks.getTotalDuration() < dayMinDuration) {
					break
				}
			}
			if totalDuration >= sessionMinDuration && totalDuration <= sessionMaxDuration ||
				(len(conferences) == 1 && talks.getTotalDuration() < dayMinDuration) {
				conferences[ci].AfternoonSession = afternoonSession
				talks.setSchedule(afternoonSession)
				break
			}
		}
	}
}

// check session and append remaining talks
func (conferences Conferences) checkSession(talks Talks) {
	for ci := range conferences {
		for _, talk := range talks {
			if talk.IsSchedule {
				continue
			}
			if conferences[ci].AfternoonSession.getTotalDuration()+talk.Duration <= sessionMaxDuration {
				conferences[ci].AfternoonSession = append(conferences[ci].AfternoonSession, talk)
				talks.setSchedule(Talks{talk})
			}
		}
	}
}

func (conferences Conferences) output() {
	for i, conference := range conferences {
		date, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-02-14 09:00:00", time.Local)
		fmt.Printf("Track %d :\n", i+1)
		for _, talk := range conference.MorningSession {
			fmt.Printf("%sAM %s\n", date.Format("15:04"), talk.Topic)
			date = date.Add(time.Duration(talk.Duration) * time.Minute)
		}
		fmt.Printf("12:00PM Lunch\n")
		date = date.Add(time.Hour)
		for _, talk := range conference.AfternoonSession {
			fmt.Printf("%sPM %s\n", date.Format("15:04"), talk.Topic)
			date = date.Add(time.Duration(talk.Duration) * time.Minute)
		}
		fmt.Printf("05:00PM Networking Event\n\n")
	}
}
