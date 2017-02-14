package main

import (
	"strconv"
	"testing"
)

func TestGetScheduleConferences(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Fail()
		}
	}()
	talks := make(Talks, 6)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
	}
	getScheduleConferences(talks)
}

func TestSetMorningSession(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Fail()
		}
	}()
	talks := make(Talks, 3)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
	}
	totalDuration := talks.getTotalDuration()
	totalDay := totalDuration/dayMaxDuration + 1
	conferences := make(Conferences, totalDay)
	conferences.setMorningSession(talks)
	for _, conference := range conferences {
		if len(conference.MorningSession) == 0 {
			t.Fail()
		}
	}
}

func TestSetAfternoonSession(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Fail()
		}
	}()
	talks := make(Talks, 4)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
	}
	totalDuration := talks.getTotalDuration()
	totalDay := totalDuration/dayMaxDuration + 1
	conferences := make(Conferences, totalDay)
	conferences.setAfternoonSession(talks)
	for _, conference := range conferences {
		if len(conference.AfternoonSession) == 0 {
			t.Fail()
		}
	}
}
