package main

import (
	"flag"
)

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go language", "help")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "php language", "help")
	args := flag.Args()
	if len(args) < 0 {
		return
	}

	switch args[0] {
	case "php":
		_ = phpCmd.Parse(args[1:])
		println("php")
	case "go":
		_ = goCmd.Parse(args[1:])
		println("go")
	default:
		println("Great!")
	}
}
