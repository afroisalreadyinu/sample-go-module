// A module with an entry point that prints the path to the binary.
//
// This module is for demo purposes. It does not do anything useful.
// You can read the blog post at http://okigiveup.net.
package main

import (
	"github.com/fatih/color"
	"myprinter/pathfinder"
)

func main() {
	color.Blue(pathfinder.Find())
}
