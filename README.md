
# Filer - A Powerful File Management CLI Tool

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://i.ytimg.com/vi/4cgpu9L2AE8/maxresdefault.jpg)

Filer is a modern, feature-rich command-line tool for managing files and directories. It provides powerful functionality to list, search, organize, and analyze files with various filtering and sorting options.

## Features

- 📁 **List Files** - Display files with customizable sorting and filtering
- 🔍 **Search** - Find files using patterns, extensions, size, and date criteria
- 📊 **Statistics** - Analyze directory contents with comprehensive stats
- 🗂️ **Organize** - Automatically organize files by type into subdirectories
- 🎨 **Multiple Output Formats** - Support for table, JSON, and CSV output
- ⚡ **Fast & Efficient** - Built with Go for optimal performance
- 🎯 **Flexible Filtering** - Filter by size, extension, date, and more

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/PranshuS007/Go-command-Line.git
cd Go-command-Line

# Build the binary
go build -o filer

# Optional: Install to your PATH
sudo mv filer /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/PranshuS007/Go-command-Line@latest
```

## Quick Start

```bash
# List files in current directory
filer list

# List files recursively with details
filer list -r -v

# Search for all .go files
filer search "*.go"

# Get directory statistics
filer stats

# Organize files by type (dry run)
filer organize --dry-run

# List files in JSON format
filer list --format json
```

## Usage

### List Command

List files and directories with various options:

```bash
# Basic listing
filer list [directory]

# List recursively
filer list -r

# Include hidden files
filer list -a

# Sort by size (descending)
filer list -s size -R

# Filter by extension
filer list -e go

# Filter by size range
filer list --min-size 1024 --max-size 1048576

# List only directories
filer list -d

# List only files
filer list -F
```

**Available Flags:**
- `-r, --recursive` - List files recursively
- `-a, --all` - Include hidden files
- `-s, --sort` - Sort by: name, size, modified, extension (default: name)
- `-R, --reverse` - Reverse sort order
- `-d, --dirs-only` - List directories only
- `-F, --files-only` - List files only
- `-e, --extension` - Filter by file extension
- `-m, --min-size` - Minimum file size in bytes
- `-M, --max-size` - Maximum file size in bytes

### Search Command

Search for files matching specific criteria:

```bash
# Search by pattern
filer search "*.txt"

# Search with extension filter
filer search "config" -e json

# Search by size range
filer search "*" --min-size 1024 --max-size 10240

# Search by modification date
filer search "*" --modified-since 2024-01-01

# Limit results
filer search "*.go" -l 10

# Include hidden files
filer search "*" -H
```

**Available Flags:**
- `-e, --extension` - Filter by file extension
- `-m, --min-size` - Minimum file size in bytes
- `-M, --max-size` - Maximum file size in bytes
- `-s, --modified-since` - Modified since date (YYYY-MM-DD)
- `-b, --modified-before` - Modified before date (YYYY-MM-DD)
- `-H, --hidden` - Include hidden files
- `-S, --sort` - Sort by: name, size, modified, extension
- `-r, --reverse` - Reverse sort order
- `-l, --limit` - Limit number of results

### Stats Command

Display comprehensive directory statistics:

```bash
# Show statistics for current directory
filer stats

# Show statistics for specific directory
filer stats /path/to/directory

# Output in JSON format
filer stats --format json

# Show top 20 extensions
filer stats -t 20
```

**Statistics Include:**
- Total files and directories count
- Total size (human-readable and bytes)
- Largest file
- Oldest and newest files
- File extension breakdown
- File type distribution

### Organize Command

Organize files into subdirectories by type:

```bash
# Preview organization (dry run)
filer organize --dry-run

# Organize files (with confirmation)
filer organize

# Organize without confirmation
filer organize -y

# Organize specific directory
filer organize /path/to/directory
```

**File Categories:**
- **images** - jpg, jpeg, png, gif, bmp
- **videos** - mp4, avi, mov, mkv, wmv
- **audio** - mp3, wav, flac, aac, ogg
- **documents** - pdf, doc, docx, txt, xls, xlsx, ppt, pptx
- **archives** - zip, rar, tar, gz, 7z
- **other** - All other file types

### Global Flags

These flags work with all commands:

- `-v, --verbose` - Verbose output with additional details
- `-f, --format` - Output format: table, json, csv (default: table)
- `--help` - Show help for any command
- `--version` - Show version information

## Output Formats

### Table Format (Default)

```
MODE       SIZE         MODIFIED             NAME
--------------------------------------------------------------------------------
-rw-r--r-- 1.2 KB       2024-01-15 10:30:45  config.json
-rw-r--r-- 3.4 MB       2024-01-14 09:15:22  data.csv
drwxr-xr-x 4.0 KB       2024-01-13 14:20:10  documents

Total: 3 items
```

### JSON Format

```json
[
  {
    "name": "config.json",
    "path": "/path/to/config.json",
    "size": 1234,
    "size_human": "1.2 KB",
    "mod_time": "2024-01-15T10:30:45Z",
    "mode": "-rw-r--r--",
    "is_dir": false,
    "extension": "json"
  }
]
```

### CSV Format

```csv
name,path,size,size_human,modified,mode,is_dir,extension
"config.json","/path/to/config.json",1234,"1.2 KB","2024-01-15 10:30:45","-rw-r--r--",false,"json"
```

## Examples

### Example 1: Find Large Files

```bash
# Find files larger than 10MB
filer search "*" --min-size 10485760 -s size -r
```

### Example 2: Clean Up Old Files

```bash
# Find files modified before 2023
filer search "*" --modified-before 2023-01-01 -v
```

### Example 3: Analyze Project Structure

```bash
# Get statistics for a project directory
filer stats ~/projects/myapp --format json > project-stats.json
```

### Example 4: Export File List

```bash
# Export all Go files to CSV
filer search "*.go" --format csv > go-files.csv
```

### Example 5: Organize Downloads

```bash
# Preview organization of downloads folder
filer organize ~/Downloads --dry-run

# Actually organize if preview looks good
filer organize ~/Downloads -y
```

## Project Structure

```
.
├── cmd/                    # Command implementations
│   ├── root.go            # Root command and global flags
│   ├── list.go            # List command
│   ├── search.go          # Search command
│   ├── stats.go           # Statistics command
│   └── organize.go        # Organize command
├── internal/
│   ├── fileops/           # File operations
│   │   └── operations.go  # Core file operations
│   └── models/            # Data models
│       └── file.go        # File information models
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── README.md              # This file
```

## Development

### Prerequisites

- Go 1.19 or higher
- Git

### Building from Source

```bash
# Clone the repository
git clone https://github.com/PranshuS007/Go-command-Line.git
cd Go-command-Line

# Download dependencies
go mod download

# Build
go build -o filer

# Run tests (when available)
go test ./...
```

### Running the Demo

```bash
# Make the demo script executable
chmod +x demo.sh

# Run the demo
./demo.sh
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development Guidelines

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Roadmap

- [ ] Add unit tests
- [ ] Add file content search functionality
- [ ] Support for file operations (copy, move, delete)
- [ ] Add file comparison features
- [ ] Implement file watching/monitoring
- [ ] Add duplicate file detection
- [ ] Support for custom file type categories
- [ ] Add configuration file support
- [ ] Implement plugin system

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) - A powerful CLI framework for Go
- Inspired by traditional Unix tools like `ls`, `find`, and `du`

## Support

If you encounter any issues or have questions:

- Open an issue on [GitHub Issues](https://github.com/PranshuS007/Go-command-Line/issues)
- Check existing issues for solutions

## Author

**Pranshu Sharma** - [PranshuS007](https://github.com/PranshuS007)

---

⭐ If you find this project useful, please consider giving it a star on GitHub!
