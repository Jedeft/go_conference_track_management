package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSystem(t *testing.T) {
	os.Args[1] = "noFile.txt"
	main()
	os.Args = make([]string, 1)
	main()
}

func TestScheduleConference(t *testing.T) {
	if err := scheduleConference("input.txt"); err != nil {
		t.Fail()
	}
	if err := scheduleConference("noFile.txt"); err == nil {
		t.Fail()
	}
	if err := scheduleConference("input_test.txt"); err == nil {
		t.Fail()
	}
}
