package main

import (
	"git-stack/cmd"
	"os"
)

func main() {

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}