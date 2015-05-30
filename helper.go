package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Helper in templates
type Helper struct {
	Env map[string]string
}

// Execute command helper
func (helper Helper) Execute(command string) string {
	prog, args := parseCommand(command)
	cmd := exec.Command(prog, args...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Execute fails: %s\n", prog)
		os.Exit(1)
	}

	return strings.TrimSpace(out.String())
}

func newHelper() *Helper {
	return &Helper{
		Env: getEnvMap(),
	}
}

func parseCommand(command string) (prog string, args []string) {
	elements := strings.Split(command, " ")
	prog = elements[0]
	args = elements[1:]
	return
}
