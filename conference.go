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

	if len(talks) != 0 {
		conferences.checkSession(talks)
		if len(talks) != 0 {
			panic(errors.New("input talks can not compose the conferences"))
		}
	}
	return conferences
}

// TODO: set possibale MorningSession
func (conferences Conferences) setMorningSession(talks Talks) error {
	return nil
}

// TODO: set possibale AfternoonSession
func (conferences Conferences) setAfternoonSession(talks Talks) error {
	return nil
}

// TODO: check session and append remaining talks
func (conferences Conferences) checkSession(talks Talks) {

}
