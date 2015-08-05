package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func listPr(c *cli.Context) {
	fmt.Println("Listing all pull requests")
}

func setRemote(c *cli.Context) {
	fmt.Println("Setting remote context")
}

func main() {
	app := cli.NewApp()
	app.Name = "probe"
	app.Usage = "Test pull requests"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "set",
			Usage: "Set project context",
			Subcommands: []cli.Command{
				{
					Name:   "remote",
					Usage:  "Set remote context for current project",
					Action: setRemote,
				},
				{
					Name:   "token",
					Usage:  "Set access token for current user",
					Action: setRemote,
				},
			},
		},
		{
			Name:   "view",
			Usage:  "View specific pull request",
			Action: setRemote,
		},
		{
			Name:   "apply",
			Usage:  "Apply specific pull request",
			Action: setRemote,
		},
		{
			Name:   "sync",
			Usage:  "Sync pull requests",
			Action: setRemote,
		},
		{
			Name:   "approve",
			Usage:  "Approve current pull request",
			Action: setRemote,
		},
		{
			Name:   "reject",
			Usage:  "Reject current pull request",
			Action: setRemote,
		},
		{
			Name:   "comment",
			Usage:  "Comment on current pull request",
			Action: setRemote,
		},
		{
			Name:  "list",
			Usage: "view pull requests",
			Subcommands: []cli.Command{
				{
					Name:   "remote",
					Usage:  "View all remote",
					Action: listPr,
				},
				{
					Name:   "local",
					Usage:  "View all local",
					Action: listPr,
				},
			},
		},
	}
	//figure out Aliases

	app.Run(os.Args)
}
