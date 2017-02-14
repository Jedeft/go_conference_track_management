package main

import "errors"

// Conference conference struct
type Conference struct {
	MorningSession   Talks
	AfternoonSession Talks
}

// Conferences conference slice
type Conferences []Conference

const (
	dayMaxDuration     = 420
	sessionMinDuration = 180
	sessionMaxDuration = 240
)

func getScheduleConferences(talks Talks) Conferences {
	totalDuration := talks.getTotalDuration()
	totalDay := totalDuration/dayMaxDuration + 1
	conferences := make(Conferences, totalDay)
	for i := 0; i < totalDay; i++ {
		conferences.setMorningSession(talks)
		conferences.setAfternoonSession(talks)
	}
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
					talks[currentIndex].Duration+totalDuration > sessionMaxDuration {
					continue
				}
				morningSession = append(morningSession, talks[currentIndex])
				totalDuration += talks[currentIndex].Duration
				if totalDuration == sessionMinDuration {
					break
				}
			}
			if totalDuration == morningSession.getTotalDuration() {
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
				if talks[currentIndex].Duration > sessionMinDuration ||
					talks[currentIndex].Duration+totalDuration > sessionMaxDuration {
					continue
				}
				afternoonSession = append(afternoonSession, talks[currentIndex])
				totalDuration += talks[currentIndex].Duration
				if totalDuration >= sessionMinDuration && totalDuration <= sessionMaxDuration {
					break
				}
			}
			if totalDuration >= sessionMinDuration && totalDuration <= sessionMaxDuration {
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
			if conferences[ci].AfternoonSession.getTotalDuration()+talk.Duration <= sessionMaxDuration {
				conferences[ci].AfternoonSession = append(conferences[ci].AfternoonSession, talk)
				talks.setSchedule(Talks{talk})
			}
		}
	}
}
