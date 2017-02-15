package main

import "testing"

func TestLoadFile(t *testing.T) {
	if _, err := loadFile("input.txt"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := loadFile("noFile.txt"); err == nil {
		t.Log(err)
		t.Fail()
	}
}
