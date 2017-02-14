package main

import "testing"

func TestLoadFile(t *testing.T) {
	_, err := loadFile("input.txt")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
