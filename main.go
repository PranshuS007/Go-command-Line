// Package main is the entry point for the Filer CLI application.
// Filer is a powerful file management tool that provides commands for listing,
// searching, organizing, and analyzing files and directories.
package main

import (
        "os"

        "github.com/user/filer/cmd"
)

// main is the application entry point. It executes the root command
// and exits with status code 1 if an error occurs.
func main() {
        if err := cmd.Execute(); err != nil {
                os.Exit(1)
        }
}
