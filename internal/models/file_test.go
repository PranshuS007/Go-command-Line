
package models

import (
	"os"
	"testing"
	"time"
)

// mockFileInfo implements os.FileInfo for testing
type mockFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (m mockFileInfo) Name() string       { return m.name }
func (m mockFileInfo) Size() int64        { return m.size }
func (m mockFileInfo) Mode() os.FileMode  { return m.mode }
func (m mockFileInfo) ModTime() time.Time { return m.modTime }
func (m mockFileInfo) IsDir() bool        { return m.isDir }
func (m mockFileInfo) Sys() interface{}   { return nil }

func TestNewFileInfo(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		mockInfo      mockFileInfo
		wantName      string
		wantExtension string
		wantHidden    bool
		wantIsDir     bool
	}{
		{
			name: "regular file with extension",
			path: "/path/to/file.txt",
			mockInfo: mockFileInfo{
				name:    "file.txt",
				size:    1024,
				mode:    0644,
				modTime: time.Now(),
				isDir:   false,
			},
			wantName:      "file.txt",
			wantExtension: "txt",
			wantHidden:    false,
			wantIsDir:     false,
		},
		{
			name: "hidden file",
			path: "/path/to/.hidden",
			mockInfo: mockFileInfo{
				name:    ".hidden",
				size:    512,
				mode:    0644,
				modTime: time.Now(),
				isDir:   false,
			},
			wantName:      ".hidden",
			wantExtension: "",
			wantHidden:    true,
			wantIsDir:     false,
		},
		{
			name: "directory",
			path: "/path/to/dir",
			mockInfo: mockFileInfo{
				name:    "dir",
				size:    4096,
				mode:    os.ModeDir | 0755,
				modTime: time.Now(),
				isDir:   true,
			},
			wantName:      "dir",
			wantExtension: "",
			wantHidden:    false,
			wantIsDir:     true,
		},
		{
			name: "file with multiple dots",
			path: "/path/to/archive.tar.gz",
			mockInfo: mockFileInfo{
				name:    "archive.tar.gz",
				size:    2048,
				mode:    0644,
				modTime: time.Now(),
				isDir:   false,
			},
			wantName:      "archive.tar.gz",
			wantExtension: "gz",
			wantHidden:    false,
			wantIsDir:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInfo := NewFileInfo(tt.path, tt.mockInfo)

			if fileInfo.Name != tt.wantName {
				t.Errorf("Name = %v, want %v", fileInfo.Name, tt.wantName)
			}
			if fileInfo.Extension != tt.wantExtension {
				t.Errorf("Extension = %v, want %v", fileInfo.Extension, tt.wantExtension)
			}
			if fileInfo.Hidden != tt.wantHidden {
				t.Errorf("Hidden = %v, want %v", fileInfo.Hidden, tt.wantHidden)
			}
			if fileInfo.IsDir != tt.wantIsDir {
				t.Errorf("IsDir = %v, want %v", fileInfo.IsDir, tt.wantIsDir)
			}
			if fileInfo.Path != tt.path {
				t.Errorf("Path = %v, want %v", fileInfo.Path, tt.path)
			}
		})
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes int64
		want  string
	}{
		{"zero bytes", 0, "0 B"},
		{"bytes", 512, "512 B"},
		{"kilobytes", 1024, "1.0 KB"},
		{"megabytes", 1048576, "1.0 MB"},
		{"gigabytes", 1073741824, "1.0 GB"},
		{"terabytes", 1099511627776, "1.0 TB"},
		{"mixed KB", 1536, "1.5 KB"},
		{"mixed MB", 2621440, "2.5 MB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatBytes(tt.bytes)
			if got != tt.want {
				t.Errorf("formatBytes(%d) = %v, want %v", tt.bytes, got, tt.want)
			}
		})
	}
}
