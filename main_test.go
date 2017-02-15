package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSystem(t *testing.T) {
	// no os.Args
	main()
	// fail os.Args
	os.Args = append(os.Args, "noFile.txt")
	main()
	// success os.Args
	os.Args[1] = "input.txt"
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
