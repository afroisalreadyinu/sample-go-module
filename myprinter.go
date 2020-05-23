// A module with an entry point that prints the path to the binary.
//
// This module is for demo purposes. It does not do anything useful.
// You can read the blog post at http://okigiveup.net.
package main

import (
	"myprinter/pathfinder"

	"github.com/afroisalreadyinu/gowsay"
	"github.com/fatih/color"
)

func main() {
	path := pathfinder.Find()
	message, err := gowsay.MakeCow(path, gowsay.Mooptions{})
	if err != nil {
		message = path
	}
	color.Blue(message)
}
