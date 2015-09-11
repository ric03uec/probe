package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var DEFAULT_REMOTE = "origin"
var ACCESS_TOKEN = ""

func SetRemote(c *cli.Context) {
	if len(c.Args()) > 0 {
		DEFAULT_REMOTE = c.Args()[0]
	}
	fmt.Println("Setting remote: ", DEFAULT_REMOTE)
}

func SetToken(c *cli.Context) {
	if len(c.Args()) > 0 {
		ACCESS_TOKEN = c.Args()[0]
		fmt.Println("Setting token: ", DEFAULT_REMOTE)
	} else {
		fmt.Println("Error, no access token specified")
		os.Exit(1)
	}
}

func SetViewPR(c *cli.Context) {
	if len(c.Args()) > 0 {
		prNumber := c.Args()[0]
		fmt.Println("Showing PR : ", prNumber)
	} else {
		fmt.Println("Error, no PR number specified")
	}
}

func SetApplyPR(c *cli.Context) {
	if len(c.Args()) > 0 {
		prNumber := c.Args()[0]
		fmt.Println("Patching PR : ", prNumber)
	} else {
		fmt.Println("Error, no PR number specified")
	}
}

func SetSync(c *cli.Context) {
	fmt.Println("Syncing PR's")
}

func SetApprovePR(c *cli.Context) {
	if len(c.Args()) > 0 {
		prNumber := c.Args()[0]
		fmt.Println("Approving PR : ", prNumber)
	} else {
		fmt.Println("Error, no PR number specified")
	}
}

func SetRejectPR(c *cli.Context) {
	if len(c.Args()) > 0 {
		prNumber := c.Args()[0]
		fmt.Println("Rejecting PR : ", prNumber)
	} else {
		fmt.Println("Error, no PR number specified")
	}
}

func SetPRComment(c *cli.Context) {
	// comment on this pr
	if len(c.Args()) > 0 {
		prNumber := c.Args()[0]
		fmt.Println("Commenting on PR : ", prNumber)
	} else {
		fmt.Println("Error, no comment specified")
	}
}

func ListPRLocal(c *cli.Context) {
	fmt.Println("Listing local PR's")
}

func ListPRRemote(c *cli.Context) {
	fmt.Println("Listing remote PR's")
}
