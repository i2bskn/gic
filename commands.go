package main

import (
	"bytes"
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
	Permission = 0777
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
		os.MkdirAll(template_dir, Permission)
		fmt.Printf("Created %s\n", template_dir)
	}
}

func doList(c *cli.Context) {
	exitIfNotInitialized()

	for _, template := range getTemplates() {
		fmt.Println(getTemplateName(template))
	}
}

func doEdit(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))

	editor := getEditor()
	template_path := getTemplatePath(c.Args().First())
	cmd := exec.Command(editor, template_path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func doPreview(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))

	tmpl := template.Must(template.ParseFiles(getTemplatePath(c.Args().First())))
	helper := newHelper()
	err := tmpl.Execute(os.Stdout, *helper)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func doApply(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))
}

func exitIfNotInitialized() {
	if requireInitialize() {
		fail("Require initialize. Please execute `gic init`.")
	}
}

func exitIfNotSpecifiedTemplate(arg_size int) {
	if arg_size < 1 {
		fail("Require template name.")
	}
}

func requireInitialize() bool {
	_, err := os.Stat(getTemplateDir())

	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

func getTemplates() (templates []string) {
	templates, err := filepath.Glob(path.Join(getTemplateDir(), "*"))

	if err != nil {
		fail("Get template list fails.")
	}
	return
}

func getTemplateName(template_path string) string {
	return path.Base(template_path)
}

func getTemplatePath(template_name string) string {
	return path.Join(getTemplateDir(), template_name)
}

func getTemplateDir() string {
	return path.Join(getMetaPath(), TemplateDir)
}

func getMetaPath() string {
	out, err := getProjectRoot()

	if err != nil {
		fail(out)
	}
	return path.Join(out, MetaDir)
}

func getProjectRoot() (out string, err error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	out = strings.TrimSpace(out.String())
	return
}

func getEditor() (editor string) {
	envs := getEnvMap()
	editor = envs["EDITOR"]

	if len(editor) == 0 {
		editor = DefaultEditor
	}
	return
}

func getGitConfig(key string) string {
	cmd := exec.Command("git", "config", key)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	out = strings.TrimSpace(out.String())
	return
}

func getEnvMap() (envs map[string]string) {
	envs = make(map[string]string)

	for _, env := range os.Environ() {
		key_and_value := strings.SplitN(env, "=", 2)
		envs[key_and_value[0]] = key_and_value[1]
	}
	return
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}

