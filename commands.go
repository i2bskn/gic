package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandList,
	commandEdit,
	commandUpload,
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

func doList(c *cli.Context) {
}

func doEdit(c *cli.Context) {
}

func doUpload(c *cli.Context) {
}

