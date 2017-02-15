package main

import (
	"sort"
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
	sort.Sort(talks)
	if _, err := getScheduleConferences(talks); err != nil {
		t.Log(err)
		t.Fail()
	}

	failTalks := make(Talks, 100)
	for ti := range failTalks {
		failTalks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		failTalks[ti].Duration = (ti + 1) * 10
	}
	if _, err := getScheduleConferences(failTalks); err == nil {
		t.Fail()
	}
}

func TestSetMorningSession(t *testing.T) {
	talks := make(Talks, 20)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
	}
	sort.Sort(talks)
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
	talks := make(Talks, 20)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
	}
	sort.Sort(talks)
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
