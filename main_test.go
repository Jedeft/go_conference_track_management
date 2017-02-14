package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestScheduleConference(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
			t.Fail()
		}
	}()
	scheduleConference("input.txt")
}
