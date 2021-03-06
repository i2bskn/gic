package main

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const (
	// MetaDir is saved templates and settings
	MetaDir = ".gic"
	// TemplateDir is saved templates path from MetaDir
	TemplateDir = "templates"
	// DefaultEditor is used when there is no EDITOR environment variable
	DefaultEditor = "vi"
	// Permission of MetaDir
	Permission = 0777
	// PersonalAccessTokenKey in .gitconfig
	PersonalAccessTokenKey = "github.token"
)

// Commands of CLI
var Commands = []cli.Command{
	commandInit,
	commandList,
	commandEdit,
	commandPreview,
	commandApply,
}

var commandInit = cli.Command{
	Name:   "init",
	Usage:  "Initialize gic settings of project.",
	Action: doInit,
}

var commandList = cli.Command{
	Name:   "list",
	Usage:  "Display a list of templates.",
	Action: doList,
}

var commandEdit = cli.Command{
	Name:   "edit",
	Usage:  "Edit template.",
	Action: doEdit,
}

var commandPreview = cli.Command{
	Name:   "preview",
	Usage:  "Display a preview of template.",
	Action: doPreview,
}

var commandApply = cli.Command{
	Name:   "apply",
	Usage:  "Create Issue with given template.",
	Action: doApply,
}

func doInit(c *cli.Context) {
	if requireInitialize() {
		templateDir := getTemplateDir()
		os.MkdirAll(templateDir, Permission)
		fmt.Printf("Created %s\n", templateDir)
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
	templatePath := getTemplatePath(c.Args().First())
	cmd := exec.Command(editor, templatePath)
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
		fail("Render template fails.")
	}
}

func doApply(c *cli.Context) {
	exitIfNotInitialized()
	exitIfNotSpecifiedTemplate(len(c.Args()))

	title := createTitle()

	tmpl := template.Must(template.ParseFiles(getTemplatePath(c.Args().First())))
	var body bytes.Buffer
	helper := newHelper()
	err := tmpl.Execute(&body, *helper)
	if err != nil {
		fail("Render template fails")
	}

	owner, repo := parseOriginURL()

	token, err := getGitConfig(PersonalAccessTokenKey)
	if err != nil {
		fail("Must be token settings to .gitconfig")
	}

	createIssue(title, body.String(), owner, repo, token)
}

func exitIfNotInitialized() {
	if requireInitialize() {
		fail("Require initialize. Please execute `gic init`.")
	}
}

func exitIfNotSpecifiedTemplate(argSize int) {
	if argSize < 1 {
		fail("Require template name.")
	}
}

func createTitle() string {
	now := time.Now().Format("20060102150405")
	return "Post from gic " + now
}

func requireInitialize() bool {
	_, err := os.Stat(getTemplateDir())

	if os.IsNotExist(err) {
		return true
	}
	return false
}

func getTemplates() (templates []string) {
	templates, err := filepath.Glob(path.Join(getTemplateDir(), "*"))

	if err != nil {
		fail("Get template list fails.")
	}
	return
}

func getTemplateName(templatePath string) string {
	return path.Base(templatePath)
}

func getTemplatePath(templateName string) string {
	return path.Join(getTemplateDir(), templateName)
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
	var result bytes.Buffer
	cmd.Stdout = &result

	err = cmd.Run()
	out = strings.TrimSpace(result.String())
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

func parseOriginURL() (owner, repo string) {
	originURL, err := getGitConfig("remote.origin.url")
	if err != nil {
		fail("Origin URI not found.")
	}

	re := regexp.MustCompile(`^(?:git@github\.com:|https://github\.com/)([^/]+)/([^/]+?)(?:\.git)$`)
	submatch := re.FindSubmatch([]byte(originURL))
	if len(submatch) != 3 {
		fail("Origin URL parse error.")
	}

	return string(submatch[1]), string(submatch[2])
}

func getGitConfig(key string) (out string, err error) {
	cmd := exec.Command("git", "config", key)
	var result bytes.Buffer
	cmd.Stdout = &result

	err = cmd.Run()
	out = strings.TrimSpace(result.String())
	return
}

func getEnvMap() (envs map[string]string) {
	envs = make(map[string]string)

	for _, env := range os.Environ() {
		keyAndValue := strings.SplitN(env, "=", 2)
		envs[key_and_value[0]] = keyAndValue[1]
	}
	return
}

func fail(message string) {
	fmt.Println(message)
	os.Exit(1)
}
