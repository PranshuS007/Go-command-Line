package fileops

import (
        "fmt"
        "io/fs"
        "os"
        "path/filepath"
        "sort"
        "strings"

        "github.com/user/filer/internal/models"
)

// ListFiles lists files in a directory with optional filtering
func ListFiles(dir string, recursive bool, showHidden bool) ([]*models.FileInfo, error) {
        var files []*models.FileInfo
        
        if recursive {
                err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
                        if err != nil {
                                return err
                        }
                        
                        // Skip hidden files if not requested
                        if !showHidden && strings.HasPrefix(d.Name(), ".") && path != dir {
                                if d.IsDir() {
                                        return filepath.SkipDir
                                }
                                return nil
                        }
                        
                        info, err := d.Info()
                        if err != nil {
                                return err
                        }
                        
                        files = append(files, models.NewFileInfo(path, info))
                        return nil
                })
                return files, err
        }
        
        // Non-recursive listing
        entries, err := os.ReadDir(dir)
        if err != nil {
                return nil, err
        }
        
        for _, entry := range entries {
                if !showHidden && strings.HasPrefix(entry.Name(), ".") {
                        continue
                }
                
                info, err := entry.Info()
                if err != nil {
                        continue
                }
                
                fullPath := filepath.Join(dir, entry.Name())
                files = append(files, models.NewFileInfo(fullPath, info))
        }
        
        return files, nil
}

// SearchFiles searches for files matching criteria
func SearchFiles(dir string, opts models.SearchOptions) ([]*models.FileInfo, error) {
        var matches []*models.FileInfo
        
        err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
                if err != nil {
                        return err
                }
                
                // Skip hidden files if not requested
                if !opts.IncludeHidden && strings.HasPrefix(d.Name(), ".") && path != dir {
                        if d.IsDir() {
                                return filepath.SkipDir
                        }
                        return nil
                }
                
                info, err := d.Info()
                if err != nil {
                        return err
                }
                
                fileInfo := models.NewFileInfo(path, info)
                
                // Apply filters
                if !matchesPattern(fileInfo, opts) {
                        return nil
                }
                
                matches = append(matches, fileInfo)
                return nil
        })
        
        return matches, err
}

// OrganizeFiles organizes files into subdirectories by type
func OrganizeFiles(dir string, dryRun bool) (map[string][]string, error) {
        organized := make(map[string][]string)
        
        files, err := ListFiles(dir, false, false)
        if err != nil {
                return nil, err
        }
        
        // Define file type mappings
        typeMap := map[string]string{
                "jpg": "images", "jpeg": "images", "png": "images", "gif": "images", "bmp": "images",
                "mp4": "videos", "avi": "videos", "mov": "videos", "mkv": "videos", "wmv": "videos",
                "mp3": "audio", "wav": "audio", "flac": "audio", "aac": "audio", "ogg": "audio",
                "pdf": "documents", "doc": "documents", "docx": "documents", "txt": "documents",
                "xls": "documents", "xlsx": "documents", "ppt": "documents", "pptx": "documents",
                "zip": "archives", "rar": "archives", "tar": "archives", "gz": "archives", "7z": "archives",
        }
        
        for _, file := range files {
                if file.IsDir {
                        continue
                }
                
                ext := strings.ToLower(file.Extension)
                category, exists := typeMap[ext]
                if !exists {
                        category = "other"
                }
                
                targetDir := filepath.Join(dir, category)
                targetPath := filepath.Join(targetDir, file.Name)
                
                organized[category] = append(organized[category], file.Name)
                
                if !dryRun {
                        // Create directory if it doesn't exist
                        if err := os.MkdirAll(targetDir, 0755); err != nil {
                                return organized, err
                        }
                        
                        // Move the file
                        if err := os.Rename(file.Path, targetPath); err != nil {
                                return organized, err
                        }
                }
        }
        
        return organized, nil
}

// GetDirectoryStats calculates comprehensive directory statistics
func GetDirectoryStats(dir string) (*models.DirectoryStats, error) {
        stats := &models.DirectoryStats{
                Path:       dir,
                FileTypes:  make(map[string]int),
                Extensions: make(map[string]int),
        }
        
        var largestFile, oldestFile, newestFile *models.FileInfo
        
        err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
                if err != nil {
                        return err
                }
                
                info, err := d.Info()
                if err != nil {
                        return err
                }
                
                fileInfo := models.NewFileInfo(path, info)
                
                if info.IsDir() {
                        stats.TotalDirs++
                } else {
                        stats.TotalFiles++
                        stats.TotalSize += info.Size()
                        
                        // Track largest file
                        if largestFile == nil || info.Size() > largestFile.Size {
                                largestFile = fileInfo
                        }
                        
                        // Track oldest file
                        if oldestFile == nil || info.ModTime().Before(oldestFile.ModTime) {
                                oldestFile = fileInfo
                        }
                        
                        // Track newest file
                        if newestFile == nil || info.ModTime().After(newestFile.ModTime) {
                                newestFile = fileInfo
                        }
                        
                        // Count extensions
                        if fileInfo.Extension != "" {
                                stats.Extensions[fileInfo.Extension]++
                        }
                }
                
                return nil
        })
        
        if err != nil {
                return nil, err
        }
        
        stats.TotalSizeHuman = formatBytes(stats.TotalSize)
        stats.LargestFile = largestFile
        stats.OldestFile = oldestFile
        stats.NewestFile = newestFile
        
        return stats, nil
}

// SortFiles sorts files by various criteria
func SortFiles(files []*models.FileInfo, sortBy string, reverse bool) {
        switch sortBy {
        case "name":
                sort.Slice(files, func(i, j int) bool {
                        result := strings.Compare(files[i].Name, files[j].Name) < 0
                        if reverse {
                                return !result
                        }
                        return result
                })
        case "size":
                sort.Slice(files, func(i, j int) bool {
                        result := files[i].Size < files[j].Size
                        if reverse {
                                return !result
                        }
                        return result
                })
        case "modified":
                sort.Slice(files, func(i, j int) bool {
                        result := files[i].ModTime.Before(files[j].ModTime)
                        if reverse {
                                return !result
                        }
                        return result
                })
        case "extension":
                sort.Slice(files, func(i, j int) bool {
                        result := strings.Compare(files[i].Extension, files[j].Extension) < 0
                        if reverse {
                                return !result
                        }
                        return result
                })
        }
}

// Helper functions

func matchesPattern(file *models.FileInfo, opts models.SearchOptions) bool {
        // Pattern matching
        if opts.Pattern != "" {
                matched, err := filepath.Match(opts.Pattern, file.Name)
                if err != nil || !matched {
                        // Try substring match if glob pattern fails
                        if !strings.Contains(strings.ToLower(file.Name), strings.ToLower(opts.Pattern)) {
                                return false
                        }
                }
        }
        
        // Extension filter
        if opts.Extension != "" && strings.ToLower(file.Extension) != strings.ToLower(opts.Extension) {
                return false
        }
        
        // Size filters
        if opts.MinSize > 0 && file.Size < opts.MinSize {
                return false
        }
        if opts.MaxSize > 0 && file.Size > opts.MaxSize {
                return false
        }
        
        // Date filters
        if !opts.ModifiedSince.IsZero() && file.ModTime.Before(opts.ModifiedSince) {
                return false
        }
        if !opts.ModifiedBefore.IsZero() && file.ModTime.After(opts.ModifiedBefore) {
                return false
        }
        
        return true
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
