package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReadBasic(t *testing.T) {
	// A stdin replica
	var stdin bytes.Buffer
	stdin.Write([]byte("argCount: 1\nArg0: find me\n"))

	res, err := Read(&stdin)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if res["Arg0"] != "find me" {
		t.Errorf("Expected Arg0 = find me, got %s", res["Arg0"])
	}

}

func TestReadFull(t *testing.T) {
	var stdin bytes.Buffer
	// Read in a previously captured p4broker out
	data, err := ioutil.ReadFile("./tst-in")
	stdin.Write(data)

	res, err := Read(&stdin)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if res["argCount"] != "2" {
		t.Errorf("Expected argCount = 2, got %s", res["argCount"])
	}

	if res["Arg0"] != "read" {
		t.Errorf("Expected Arg0 = read, got %s", res["Arg0"])
	}

	if res["Arg1"] != "//path/to/file" {
		t.Errorf("Expected Arg1 = //path/to/file, got %s", res["Arg1"])
	}
}
