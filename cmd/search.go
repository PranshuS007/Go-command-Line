package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/user/filer/internal/fileops"
	"github.com/user/filer/internal/models"
)

var searchCmd = &cobra.Command{
	Use:   "search [pattern] [directory]",
	Short: "Search for files matching criteria",
	Long: `Search for files matching the specified pattern and criteria.
The pattern can be a glob pattern or simple substring match.
If no directory is specified, the current directory is used.`,
	Aliases: []string{"find", "f"},
	Args:    cobra.RangeArgs(1, 2),
	Run:     runSearch,
}

func init() {
	rootCmd.AddCommand(searchCmd)
	
	searchCmd.Flags().StringP("extension", "e", "", "filter by file extension")
	searchCmd.Flags().Int64P("min-size", "m", 0, "minimum file size in bytes")
	searchCmd.Flags().Int64P("max-size", "M", 0, "maximum file size in bytes")
	searchCmd.Flags().StringP("modified-since", "s", "", "modified since date (YYYY-MM-DD)")
	searchCmd.Flags().StringP("modified-before", "b", "", "modified before date (YYYY-MM-DD)")
	searchCmd.Flags().BoolP("hidden", "H", false, "include hidden files")
	searchCmd.Flags().StringP("sort", "S", "name", "sort by: name, size, modified, extension")
	searchCmd.Flags().BoolP("reverse", "r", false, "reverse sort order")
	searchCmd.Flags().IntP("limit", "l", 0, "limit number of results (0 = no limit)")
}

func runSearch(cmd *cobra.Command, args []string) {
	pattern := args[0]
	
	// Get directory to search
	dir := "."
	if len(args) > 1 {
		dir = args[1]
	}
	
	// Parse flags
	extension, _ := cmd.Flags().GetString("extension")
	minSize, _ := cmd.Flags().GetInt64("min-size")
	maxSize, _ := cmd.Flags().GetInt64("max-size")
	modifiedSinceStr, _ := cmd.Flags().GetString("modified-since")
	modifiedBeforeStr, _ := cmd.Flags().GetString("modified-before")
	includeHidden, _ := cmd.Flags().GetBool("hidden")
	sortBy, _ := cmd.Flags().GetString("sort")
	reverse, _ := cmd.Flags().GetBool("reverse")
	limit, _ := cmd.Flags().GetInt("limit")
	
	// Parse dates
	var modifiedSince, modifiedBefore time.Time
	var err error
	
	if modifiedSinceStr != "" {
		modifiedSince, err = time.Parse("2006-01-02", modifiedSinceStr)
		checkError(err)
	}
	
	if modifiedBeforeStr != "" {
		modifiedBefore, err = time.Parse("2006-01-02", modifiedBeforeStr)
		checkError(err)
	}
	
	// Create search options
	opts := models.SearchOptions{
		Pattern:        pattern,
		Extension:      extension,
		MinSize:        minSize,
		MaxSize:        maxSize,
		ModifiedSince:  modifiedSince,
		ModifiedBefore: modifiedBefore,
		IncludeHidden:  includeHidden,
		Recursive:      true,
	}
	
	// Perform search
	if isVerbose() {
		fmt.Printf("Searching for '%s' in '%s'...\n", pattern, dir)
	}
	
	files, err := fileops.SearchFiles(dir, opts)
	checkError(err)
	
	// Sort results
	fileops.SortFiles(files, sortBy, reverse)
	
	// Apply limit
	if limit > 0 && len(files) > limit {
		files = files[:limit]
	}
	
	// Output results
	if len(files) == 0 {
		fmt.Println("No files found matching the criteria")
		return
	}
	
	if isVerbose() {
		fmt.Printf("Found %d matching files:\n\n", len(files))
	}
	
	outputFiles(files, getOutputFormat())
	
	if isVerbose() {
		fmt.Printf("\nSearch completed. Found %d files.\n", len(files))
	}
}
