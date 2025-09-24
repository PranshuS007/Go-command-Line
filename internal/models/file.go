package models

import (
        "fmt"
        "os"
        "path/filepath"
        "time"
)

// FileInfo represents enhanced file information
type FileInfo struct {
        Name       string    `json:"name"`
        Path       string    `json:"path"`
        Size       int64     `json:"size"`
        SizeHuman  string    `json:"size_human"`
        ModTime    time.Time `json:"mod_time"`
        Mode       string    `json:"mode"`
        IsDir      bool      `json:"is_dir"`
        Extension  string    `json:"extension"`
        MimeType   string    `json:"mime_type,omitempty"`
        Hidden     bool      `json:"hidden"`
}

// DirectoryStats represents directory statistics
type DirectoryStats struct {
        Path           string            `json:"path"`
        TotalFiles     int               `json:"total_files"`
        TotalDirs      int               `json:"total_dirs"`
        TotalSize      int64             `json:"total_size"`
        TotalSizeHuman string            `json:"total_size_human"`
        FileTypes      map[string]int    `json:"file_types"`
        LargestFile    *FileInfo         `json:"largest_file,omitempty"`
        OldestFile     *FileInfo         `json:"oldest_file,omitempty"`
        NewestFile     *FileInfo         `json:"newest_file,omitempty"`
        Extensions     map[string]int    `json:"extensions"`
}

// SearchOptions represents search criteria
type SearchOptions struct {
        Pattern    string
        Extension  string
        MinSize    int64
        MaxSize    int64
        ModifiedSince time.Time
        ModifiedBefore time.Time
        IncludeHidden bool
        Recursive     bool
}

// NewFileInfo creates a FileInfo from os.FileInfo
func NewFileInfo(path string, info os.FileInfo) *FileInfo {
        ext := filepath.Ext(info.Name())
        if ext != "" && len(ext) > 1 {
                ext = ext[1:] // Remove the dot
        }
        
        return &FileInfo{
                Name:      info.Name(),
                Path:      path,
                Size:      info.Size(),
                SizeHuman: formatBytes(info.Size()),
                ModTime:   info.ModTime(),
                Mode:      info.Mode().String(),
                IsDir:     info.IsDir(),
                Extension: ext,
                Hidden:    info.Name()[0] == '.',
        }
}

// formatBytes converts bytes to human readable format
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
