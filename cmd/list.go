package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/user/filer/internal/fileops"
	"github.com/user/filer/internal/models"
)

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List files and directories",
	Long: `List files and directories in the specified path with various filtering and sorting options.
If no directory is specified, the current directory is used.`,
	Aliases: []string{"ls", "l"},
	Args:    cobra.MaximumNArgs(1),
	Run:     runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	
	listCmd.Flags().BoolP("recursive", "r", false, "list files recursively")
	listCmd.Flags().BoolP("all", "a", false, "include hidden files")
	listCmd.Flags().StringP("sort", "s", "name", "sort by: name, size, modified, extension")
	listCmd.Flags().BoolP("reverse", "R", false, "reverse sort order")
	listCmd.Flags().BoolP("dirs-only", "d", false, "list directories only")
	listCmd.Flags().BoolP("files-only", "F", false, "list files only")
	listCmd.Flags().StringP("extension", "e", "", "filter by file extension")
	listCmd.Flags().Int64P("min-size", "m", 0, "minimum file size in bytes")
	listCmd.Flags().Int64P("max-size", "M", 0, "maximum file size in bytes")
}

func runList(cmd *cobra.Command, args []string) {
	// Get directory to list
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}
	
	// Get flags
	recursive, _ := cmd.Flags().GetBool("recursive")
	showHidden, _ := cmd.Flags().GetBool("all")
	sortBy, _ := cmd.Flags().GetString("sort")
	reverse, _ := cmd.Flags().GetBool("reverse")
	dirsOnly, _ := cmd.Flags().GetBool("dirs-only")
	filesOnly, _ := cmd.Flags().GetBool("files-only")
	extension, _ := cmd.Flags().GetString("extension")
	minSize, _ := cmd.Flags().GetInt64("min-size")
	maxSize, _ := cmd.Flags().GetInt64("max-size")
	
	// List files
	files, err := fileops.ListFiles(dir, recursive, showHidden)
	checkError(err)
	
	// Apply filters
	filteredFiles := filterFiles(files, extension, minSize, maxSize, dirsOnly, filesOnly)
	
	// Sort files
	fileops.SortFiles(filteredFiles, sortBy, reverse)
	
	// Output results
	outputFiles(filteredFiles, getOutputFormat())
}

func filterFiles(files []*models.FileInfo, extension string, minSize, maxSize int64, dirsOnly, filesOnly bool) []*models.FileInfo {
	var filtered []*models.FileInfo
	
	for _, file := range files {
		// Directory/file filter
		if dirsOnly && !file.IsDir {
			continue
		}
		if filesOnly && file.IsDir {
			continue
		}
		
		// Extension filter
		if extension != "" && strings.ToLower(file.Extension) != strings.ToLower(extension) {
			continue
		}
		
		// Size filters (only for files)
		if !file.IsDir {
			if minSize > 0 && file.Size < minSize {
				continue
			}
			if maxSize > 0 && file.Size > maxSize {
				continue
			}
		}
		
		filtered = append(filtered, file)
	}
	
	return filtered
}

func outputFiles(files []*models.FileInfo, format string) {
	switch format {
	case "json":
		outputJSON(files)
	case "csv":
		outputCSV(files)
	default:
		outputTable(files)
	}
}

func outputTable(files []*models.FileInfo) {
	if len(files) == 0 {
		fmt.Println("No files found")
		return
	}
	
	// Print header
	fmt.Printf("%-10s %-12s %-20s %s\n", "MODE", "SIZE", "MODIFIED", "NAME")
	fmt.Println(strings.Repeat("-", 80))
	
	for _, file := range files {
		mode := file.Mode
		if file.IsDir {
			mode = "d" + mode[1:]
		}
		
		fmt.Printf("%-10s %-12s %-20s %s\n",
			mode[:10],
			file.SizeHuman,
			file.ModTime.Format("2006-01-02 15:04:05"),
			file.Name)
	}
	
	fmt.Printf("\nTotal: %d items\n", len(files))
}

func outputJSON(files []*models.FileInfo) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	checkError(encoder.Encode(files))
}

func outputCSV(files []*models.FileInfo) {
	fmt.Println("name,path,size,size_human,modified,mode,is_dir,extension")
	
	for _, file := range files {
		fmt.Printf("%q,%q,%d,%q,%q,%q,%t,%q\n",
			file.Name,
			file.Path,
			file.Size,
			file.SizeHuman,
			file.ModTime.Format("2006-01-02 15:04:05"),
			file.Mode,
			file.IsDir,
			file.Extension)
	}
}
