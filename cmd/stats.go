package cmd

import (
        "encoding/json"
        "fmt"
        "os"
        "sort"
        "strings"

        "github.com/spf13/cobra"
        "github.com/user/filer/internal/fileops"
        "github.com/user/filer/internal/models"
)

var statsCmd = &cobra.Command{
        Use:   "stats [directory]",
        Short: "Show directory statistics and analysis",
        Long: `Display comprehensive statistics about files and directories including:
- File and directory counts
- Total size and size distribution
- File type breakdown
- Largest, oldest, and newest files
- Extension analysis

If no directory is specified, the current directory is used.`,
        Aliases: []string{"stat", "info", "analyze"},
        Args:    cobra.MaximumNArgs(1),
        Run:     runStats,
}

func init() {
        rootCmd.AddCommand(statsCmd)
        
        statsCmd.Flags().BoolP("extensions", "e", true, "show file extensions breakdown")
        statsCmd.Flags().IntP("top", "t", 10, "show top N extensions (0 = all)")
}

func runStats(cmd *cobra.Command, args []string) {
        // Get directory to analyze
        dir := "."
        if len(args) > 0 {
                dir = args[0]
        }
        
        // Get flags
        showExtensions, _ := cmd.Flags().GetBool("extensions")
        topN, _ := cmd.Flags().GetInt("top")
        
        if isVerbose() {
                fmt.Printf("Analyzing directory: %s\n", dir)
        }
        
        // Calculate statistics
        stats, err := fileops.GetDirectoryStats(dir)
        checkError(err)
        
        // Output based on format
        format := getOutputFormat()
        switch format {
        case "json":
                outputStatsJSON(stats)
        default:
                outputStatsTable(stats, showExtensions, topN)
        }
}

func outputStatsJSON(stats *models.DirectoryStats) {
        encoder := json.NewEncoder(os.Stdout)
        encoder.SetIndent("", "  ")
        checkError(encoder.Encode(stats))
}

func outputStatsTable(stats *models.DirectoryStats, showExtensions bool, topN int) {
        fmt.Printf("Directory Statistics for: %s\n", stats.Path)
        fmt.Println(strings.Repeat("=", 60))
        
        // Basic counts
        fmt.Printf("Files:       %d\n", stats.TotalFiles)
        fmt.Printf("Directories: %d\n", stats.TotalDirs)
        fmt.Printf("Total Size:  %s (%d bytes)\n", stats.TotalSizeHuman, stats.TotalSize)
        
        if stats.TotalFiles > 0 {
                avgSize := stats.TotalSize / int64(stats.TotalFiles)
                fmt.Printf("Average Size: %s per file\n", formatBytes(avgSize))
        }
        
        // Notable files
        fmt.Printf("\nNotable Files:\n")
        fmt.Println(strings.Repeat("-", 30))
        
        if stats.LargestFile != nil {
                fmt.Printf("Largest:  %s (%s)\n", stats.LargestFile.Name, stats.LargestFile.SizeHuman)
        }
        
        if stats.OldestFile != nil {
                fmt.Printf("Oldest:   %s (%s)\n", 
                        stats.OldestFile.Name, 
                        stats.OldestFile.ModTime.Format("2006-01-02"))
        }
        
        if stats.NewestFile != nil {
                fmt.Printf("Newest:   %s (%s)\n", 
                        stats.NewestFile.Name, 
                        stats.NewestFile.ModTime.Format("2006-01-02"))
        }
        
        // Extensions breakdown
        if showExtensions && len(stats.Extensions) > 0 {
                fmt.Printf("\nFile Extensions:\n")
                fmt.Println(strings.Repeat("-", 30))
                
                // Sort extensions by count
                type extCount struct {
                        ext   string
                        count int
                }
                
                var extensions []extCount
                for ext, count := range stats.Extensions {
                        extensions = append(extensions, extCount{ext, count})
                }
                
                sort.Slice(extensions, func(i, j int) bool {
                        return extensions[i].count > extensions[j].count
                })
                
                // Show top N or all
                maxShow := len(extensions)
                if topN > 0 && topN < maxShow {
                        maxShow = topN
                }
                
                for i := 0; i < maxShow; i++ {
                        ext := extensions[i]
                        percentage := float64(ext.count) / float64(stats.TotalFiles) * 100
                        fmt.Printf("%-8s %4d files (%.1f%%)\n", "."+ext.ext, ext.count, percentage)
                }
                
                if topN > 0 && len(extensions) > topN {
                        remaining := 0
                        for i := topN; i < len(extensions); i++ {
                                remaining += extensions[i].count
                        }
                        fmt.Printf("%-8s %4d files (others)\n", "...", remaining)
                }
        }
        
        // Summary
        fmt.Printf("\nSummary: %d items (%d files, %d dirs) totaling %s\n",
                stats.TotalFiles+stats.TotalDirs,
                stats.TotalFiles,
                stats.TotalDirs,
                stats.TotalSizeHuman)
}

func formatBytes(bytes int64) string {
        const unit = 1024
        if bytes < unit {
                return fmt.Sprintf("%d B", bytes)
        }
        div, exp := int64(unit), 0
        for n := bytes / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
