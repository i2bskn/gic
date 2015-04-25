package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
)

const (
	MetaDir = ".gic"
	TemplateDir = "templates"
	DefaultEditor = "vi"
)

var Commands = []cli.Command{
	commandInit,
	commandList,
	commandEdit,
	commandPreview,
	commandApply,
}

var commandInit = cli.Command{
	Name:  "init",
	Usage: "",
	Description: ``,
	Action: doInit,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "",
	Description: ``,
	Action: doList,
}

var commandEdit = cli.Command{
	Name:  "edit",
	Usage: "",
	Description: ``,
	Action: doEdit,
}

var commandPreview = cli.Command{
	Name:  "preview",
	Usage: "",
	Description: ``,
	Action: doPreview,
}

var commandApply = cli.Command{
	Name:  "apply",
	Usage: "",
	Description: ``,
	Action: doApply,
}

func doInit(c *cli.Context) {
	if requireInitialize() {
		template_dir := getTemplateDir()
		os.MkdirAll(template_dir, 0777)
		fmt.Printf("Created %s\n", template_dir)
	}
}

func doList(c *cli.Context) {
	exitIfNotInitialized()
	templates := getTemplates()

	for _, template := range templates {
		fmt.Println(getTemplateName(template))
	}
}

func doEdit(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))
	editTemplateWithEditor(c.Args().First())
}

func doPreview(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))
	template_path := getTemplatePath(c.Args().First())
	tpl := template.Must(template.ParseFiles(template_path))
	helper := Helper{}
	err := tpl.Execute(os.Stdout, helper)
	if err != nil {
		fmt.Println(err)
	}
}

func doApply(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))
}

func getTemplateName(template_path string) string {
	return path.Base(template_path)
}

func getTemplatePath(template_name string) string {
	return path.Join(getTemplateDir(), template_name)
}

func exitIfNotInitialized() {
	if requireInitialize() {
		fmt.Println("Require initialize. Please execute `gic init`.")
		os.Exit(1)
	}
}

func exitIfNotSpecifiedTemplate(i int) {
	if i == 0 {
		fmt.Println("Require template name.")
		os.Exit(1)
	}
}

func getTemplates() (templates []string) {
	pattern := path.Join(getTemplateDir(), "*")
	templates, err := filepath.Glob(pattern)

	if err != nil {
		fmt.Println("Get template list fails.")
		os.Exit(1)
	}
	return
}

func requireInitialize() bool {
	template_path := getTemplateDir()
	_, err := os.Stat(template_path)

	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func getTemplateDir() string {
	return path.Join(getMetaPath(), TemplateDir)
}

func getMetaPath() string {
	out, err := getProjectRoot()

	if err != nil {
		fmt.Println(out)
		os.Exit(1)
	}

	return path.Join(out, MetaDir)
}

func getProjectRoot() (out string, err error) {
	result, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	out = strings.TrimSpace(string(result))
	return
}

func editTemplateWithEditor(template string) {
	editor := getEditor()
	template_path := getTemplatePath(template)
	cmd := exec.Command(editor, template_path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func getEditor() (editor string) {
	editor = getEnviron("EDITOR")
	if len(editor) == 0 {
		editor = DefaultEditor
	}
	return
}

func getEnviron(key string) (value string) {
	envs := os.Environ()
	for _, env := range envs {
		key_and_value := strings.SplitN(env, "=", 2)
		if key == key_and_value[0] {
			value = key_and_value[1]
		}
	}
	return
}

