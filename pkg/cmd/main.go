package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/bejelith/docker_entrypoint/pkg/render"
)

var (
	templates StringValue
	//commandLine *flag.FlagSet
)

func init() {
	standardUsageFunc := flag.Usage
	flag.Usage = func() {
		standardUsageFunc()
		_, _ = fmt.Fprintf(os.Stderr, "Command line arguments are executed after template generation\n")
	}
	flag.Var(&templates, "template", "Templates to render, can be more than one")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if err := render.ExecTemplates(templates...); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) > 0 {
		if err := syscall.Exec(args[0], args, os.Environ()); err != nil {
			fmt.Printf("Error running Exec for %s: %s\n", args[0], err)
			os.Exit(2)
		}
	} else {
		fmt.Println("Nothing to execute, exiting")
	}
}
