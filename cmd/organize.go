package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/user/filer/internal/fileops"
)

var organizeCmd = &cobra.Command{
	Use:   "organize [directory]",
	Short: "Organize files into subdirectories by type",
	Long: `Organize files in the specified directory into subdirectories based on file type.
Files are categorized into: images, videos, audio, documents, archives, and other.
If no directory is specified, the current directory is used.`,
	Aliases: []string{"org", "o"},
	Args:    cobra.MaximumNArgs(1),
	Run:     runOrganize,
}

func init() {
	rootCmd.AddCommand(organizeCmd)
	
	organizeCmd.Flags().BoolP("dry-run", "n", false, "show what would be organized without making changes")
	organizeCmd.Flags().BoolP("confirm", "y", false, "skip confirmation prompt")
}

func runOrganize(cmd *cobra.Command, args []string) {
	// Get directory to organize
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}
	
	// Get flags
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	skipConfirm, _ := cmd.Flags().GetBool("confirm")
	
	if isVerbose() {
		if dryRun {
			fmt.Printf("Dry run: Analyzing organization for '%s'...\n", dir)
		} else {
			fmt.Printf("Organizing files in '%s'...\n", dir)
		}
	}
	
	// Perform organization (dry run first to show preview)
	organized, err := fileops.OrganizeFiles(dir, true)
	checkError(err)
	
	if len(organized) == 0 {
		fmt.Println("No files to organize")
		return
	}
	
	// Show what will be organized
	fmt.Println("Files will be organized as follows:")
	fmt.Println(strings.Repeat("=", 50))
	
	totalFiles := 0
	for category, files := range organized {
		fmt.Printf("\n%s/ (%d files):\n", strings.Title(category), len(files))
		for _, file := range files {
			fmt.Printf("  - %s\n", file)
			totalFiles++
		}
	}
	
	fmt.Printf("\nTotal files to organize: %d\n", totalFiles)
	
	if dryRun {
		fmt.Println("\n(This was a dry run - no files were moved)")
		return
	}
	
	// Confirm unless skip flag is set
	if !skipConfirm {
		fmt.Print("\nProceed with organization? (y/N): ")
		var response string
		fmt.Scanln(&response)
		
		if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
			fmt.Println("Organization cancelled")
			return
		}
	}
	
	// Perform actual organization
	fmt.Println("\nOrganizing files...")
	_, err = fileops.OrganizeFiles(dir, false)
	checkError(err)
	
	fmt.Println("✓ Files organized successfully!")
	
	if isVerbose() {
		fmt.Printf("Organized %d files into %d categories\n", totalFiles, len(organized))
	}
}
