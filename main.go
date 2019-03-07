package main

import (
	"fmt"
	"github.com/modmuss50/CAV2"
	"os"
)

var (
	runDir = "."
)

func main() {

	cav2.UseCurseMeta()

	args := os.Args[1:]

	if len(args) == 0 {
		help()
	}

	if args[0] == "update" {
		update()
	}

	// update			asks if each mod with an update wants to be updated
	// update -all 		updates all mods no questions asked

	//install <slug/id> does what it says

	//search			cli interface for finding mods
}

func help() {
	fmt.Println("Twitch Mod Manager")
	fmt.Println("update, updates mods in the current dir")
}
