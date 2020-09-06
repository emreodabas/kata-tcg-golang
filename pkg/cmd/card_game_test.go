package cmd

import (
	"bytes"
	"github.com/fatih/color"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewGame(t *testing.T) {
	cmd := NewGame()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"-h"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "TCG player_name") {
		t.Fatalf("expected \"%s\" got \"%s\"", "TCG player_name", string(out))
	}
}
func TestStartGame(t *testing.T) {
	rb := new(bytes.Buffer)
	color.Output = rb
	color.NoColor = false
	cmd := NewGame()
	cmd.SetArgs([]string{"Emre"})
	cmd.Execute()
	line, _ := ioutil.ReadAll(rb)
	if !strings.Contains(string(line), "Congratulations") {
		t.Fatalf("expected \"%s\" got \"%s\"", "Congratulations", string(line))
	}
}
