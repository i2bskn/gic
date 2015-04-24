package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	MetaDir = ".gic"
	TemplateDir = "templates"
)

func templateName(template_path string) string {
	return path.Base(template_path)
}

func exitIfNotInitialized() {
	if requireInitialize() {
		fmt.Println("Require initialize. Please execute `gic init`.")
		os.Exit(1)
	}
}

func getTemplates() []string {
	issues := path.Join(getTemplatePath(), "*.issue")
	templates, err := filepath.Glob(issues)

	if err != nil {
		fmt.Println("Get template list fails.")
		os.Exit(1)
	}

	return templates
}

func requireInitialize() bool {
	template_path := getTemplatePath()
	_, err := os.Stat(template_path)

	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

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

