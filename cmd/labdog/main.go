// Labdog is a collection of small tools for home labs, shipped as one binary.
package main

import (
	"fmt"
	"os"
)

var version = "dev"

const usage = `Labdog - small watchdogs for your lab

Usage:
  labdog <command> [flags]

Commands:
  watchdog   shut down gracefully when a peer sharing your power goes silent
  version    print the version

Run "labdog <command> -h" for command flags.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "watchdog":
		os.Exit(runWatchdog(os.Args[2:]))
	case "version", "-version", "--version":
		fmt.Println(version)
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "labdog: unknown command %q\n\n%s", os.Args[1], usage)
		os.Exit(2)
	}
}
