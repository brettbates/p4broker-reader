package main

import (
	"bytes"
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
