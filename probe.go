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
					Action: SetRemote,
				},
				{
					Name:   "token",
					Usage:  "Set access token for current user",
					Action: SetToken,
				},
			},
		},
		{
			Name:   "view",
			Usage:  "View specific pull request",
			Action: SetViewPR,
		},
		{
			Name:   "apply",
			Usage:  "Apply specific pull request",
			Action: SetApplyPR,
		},
		{
			Name:   "sync",
			Usage:  "Sync pull requests",
			Action: SetSync,
		},
		{
			Name:   "approve",
			Usage:  "Approve current pull request",
			Action: SetApprovePR,
		},
		{
			Name:   "reject",
			Usage:  "Reject current pull request",
			Action: SetRejectPR,
		},
		{
			Name:   "comment",
			Usage:  "Comment on current pull request",
			Action: SetPRComment,
		},
		{
			Name:  "list",
			Usage: "view pull requests",
			Subcommands: []cli.Command{
				{
					Name:   "remote",
					Usage:  "View all remote",
					Action: ListPRRemote,
				},
				{
					Name:   "local",
					Usage:  "View all local",
					Action: ListPRLocal,
				},
			},
		},
	}
	//figure out Aliases

	app.Run(os.Args)
}
