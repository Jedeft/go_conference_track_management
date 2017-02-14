package main

import (
	"strconv"
	"testing"
)

func TestFillTalks(t *testing.T) {
	data, err := loadFile("input.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	talks, err := fillTalks(data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if len(talks) == 0 {
		t.Fail()
	}
}

func TestGetTotalDuration(t *testing.T) {
	var totalDuration int
	talks := make(Talks, 4)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
		totalDuration += talks[ti].Duration
	}
	if totalDuration != talks.getTotalDuration() {
		t.Fail()
	}
}

func TestSetSchedule(t *testing.T) {
	talks := make(Talks, 2)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
	}
	talks.setSchedule(Talks{Talk{Topic: "test topic 1 10min"}})
	if !talks[0].IsSchedule {
		t.Fail()
	}
}

func TestIsSchedule(t *testing.T) {
	talks := make(Talks, 4)
	for ti := range talks {
		talks[ti].Topic = "test topic " + strconv.Itoa(ti+1) + " " + strconv.Itoa((ti+1)*10) + "min"
		talks[ti].Duration = (ti + 1) * 10
		talks[ti].IsSchedule = true
	}
	if !talks.isSchedule() {
		t.Fail()
	}
}
