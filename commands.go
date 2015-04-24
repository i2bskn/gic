package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandInit,
	commandList,
	commandEdit,
	commandUpload,
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

var commandUpload = cli.Command{
	Name:  "upload",
	Usage: "",
	Description: ``,
	Action: doUpload,
}

func doInit(c *cli.Context) {
	if requireInitialize() {
		template_path := getTemplatePath()
		os.MkdirAll(template_path, 0777)
		fmt.Printf("Created %s\n", template_path)
	}
}

func doList(c *cli.Context) {
	exitIfNotInitialized()
	templates := getTemplates()

	for _, template := range templates {
		fmt.Println(templateName(template))
	}
}

func doEdit(c *cli.Context) {
}

func doUpload(c *cli.Context) {
}

