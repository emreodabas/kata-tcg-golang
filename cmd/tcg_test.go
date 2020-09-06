package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainFunctionCmd(t *testing.T) {
	var err error
	cmd := exec.Command("../tcg", "some", "bad", "args")
	out, err := cmd.CombinedOutput()
	sout := string(out)
	if err != nil && !strings.Contains(sout, "accepts at most 1 arg") {
		t.Errorf("%v", err)
	}
}

func TestMainFunctionParameter(t *testing.T) {
	var err error
	cmd := exec.Command("../tcg", "PlayerName")
	out, err := cmd.CombinedOutput()
	sout := string(out)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !strings.Contains(sout, "PlayerName") {
		t.Errorf("%v", err)
	}
}
