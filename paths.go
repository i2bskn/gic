package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	MetaDir = ".gic"
	TemplateDir = "templates"
)

func getTemplatePath() string {
	meta_path := getMetaPath()
	return path.Join(meta_path, TemplateDir)
}

func getMetaPath() string {
	out, err := getProjectRoot()
	if err != nil {
		fmt.Println(out)
		os.Exit(1)
	}
	return path.Join(out, MetaDir)
}

func getProjectRoot() (string, error) {
	result, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	out := strings.Trim(string(result), "\n")
	return out, err
}

