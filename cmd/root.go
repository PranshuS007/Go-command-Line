package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filer",
	Short: "A powerful file management CLI tool",
	Long: `Filer is a command-line tool for managing files and directories.
It provides functionality to list, search, organize, and analyze files
with various filtering and sorting options.`,
	Version: "1.0.0",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP("format", "f", "table", "output format (table, json, csv)")
}

// Helper function to handle errors consistently
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Helper function to get verbose flag value
func isVerbose() bool {
	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	return verbose
}

// Helper function to get format flag value
func getOutputFormat() string {
	format, _ := rootCmd.PersistentFlags().GetString("format")
	return format
}
