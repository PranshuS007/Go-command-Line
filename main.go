package main

import (
	"os"

	"github.com/user/filer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
