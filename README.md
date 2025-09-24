# Filer - Advanced File Management CLI Tool

A powerful, fast, and intuitive command-line tool for managing files and directories. Built with Go for cross-platform compatibility and optimal performance.

## Features

- **List Files**: Display files with advanced filtering, sorting, and formatting options
- **Search Files**: Find files using patterns, extensions, size, and date criteria  
- **Organize Files**: Automatically categorize files into type-based subdirectories
- **Directory Stats**: Comprehensive analysis of directory contents and statistics
- **Multiple Output Formats**: Table, JSON, and CSV output support
- **Cross-Platform**: Works on Linux, macOS, and Windows

## Installation

### From Source
```bash
git clone <repository-url>
cd filer
go build -o filer
```

### Direct Download
Download the binary for your platform from the releases page.

## Usage

### Basic Commands

#### List Files
```bash
# List current directory
./filer list

# List with all files (including hidden)
./filer list -a

# List files recursively
./filer list -r /path/to/directory

# Sort by size, show only files
./filer list -s size -F

# Filter by extension
./filer list -e go

# JSON output
./filer list --format json
```

#### Search Files
```bash
# Search by pattern
./filer search "*.pdf"

# Search with multiple criteria
./filer search "config" --extension json --min-size 1024

# Search by date range
./filer search "*" --modified-since 2024-01-01

# Limit results
./filer search "*.log" --limit 10
```

#### Organize Files
```bash
# Preview organization (dry run)
./filer organize --dry-run

# Organize files by type
./filer organize

# Skip confirmation
./filer organize -y
```

#### Directory Statistics
```bash
# Show comprehensive stats
./filer stats

# Focus on top 5 extensions
./filer stats --top 5

# JSON format for processing
./filer stats --format json
```

### Advanced Usage

#### Global Flags
- `-v, --verbose`: Enable verbose output
- `-f, --format`: Output format (table, json, csv)
- `-h, --help`: Show help information

#### List Command Options
- `-r, --recursive`: List files recursively
- `-a, --all`: Include hidden files
- `-s, --sort`: Sort by name, size, modified, extension
- `-R, --reverse`: Reverse sort order
- `-d, --dirs-only`: Show directories only
- `-F, --files-only`: Show files only
- `-e, --extension`: Filter by file extension
- `-m, --min-size`: Minimum file size filter
- `-M, --max-size`: Maximum file size filter

#### Search Command Options
- `-e, --extension`: Filter by extension
- `-m, --min-size`: Minimum file size
- `-M, --max-size`: Maximum file size
- `-s, --modified-since`: Files modified after date (YYYY-MM-DD)
- `-b, --modified-before`: Files modified before date (YYYY-MM-DD)
- `-H, --hidden`: Include hidden files
- `-S, --sort`: Sort results
- `-r, --reverse`: Reverse sort order
- `-l, --limit`: Limit number of results

#### Organize Command Options
- `-n, --dry-run`: Preview without making changes
- `-y, --confirm`: Skip confirmation prompt

#### Stats Command Options
- `-e, --extensions`: Show extensions breakdown
- `-t, --top`: Show top N extensions

## File Organization Categories

Files are automatically categorized into:

- **Images**: jpg, jpeg, png, gif, bmp
- **Videos**: mp4, avi, mov, mkv, wmv  
- **Audio**: mp3, wav, flac, aac, ogg
- **Documents**: pdf, doc, docx, txt, xls, xlsx, ppt, pptx
- **Archives**: zip, rar, tar, gz, 7z
- **Other**: All other file types

## Examples

### Find Large Files
```bash
./filer search "*" --min-size 10485760  # Files > 10MB
./filer list -s size -R --files-only    # Largest files first
```

### Clean Up Downloads Folder
```bash
./filer stats ~/Downloads                    # Analyze first
./filer organize ~/Downloads --dry-run       # Preview organization
./filer organize ~/Downloads                 # Execute organization
```

### Find Recent Files
```bash
./filer search "*" --modified-since 2024-09-01 --sort modified
```

### Export File Inventory
```bash
./filer list -r --format csv > file_inventory.csv
./filer stats --format json > directory_stats.json
```

### Development Cleanup
```bash
./filer search "node_modules" -F        # Find node_modules directories
./filer search "*.log" --min-size 1024  # Find large log files
```

## Performance

- **Fast**: Optimized Go implementation with concurrent processing
- **Memory Efficient**: Streaming file processing for large directories
- **Cross-Platform**: Single binary deployment

## Output Formats

### Table (Default)
Clean, human-readable tabular output with proper alignment.

### JSON
Structured data perfect for integration with other tools:
```json
[
  {
    "name": "example.txt",
    "path": "/full/path/example.txt", 
    "size": 1024,
    "size_human": "1.0 KB",
    "mod_time": "2024-09-24T10:30:00Z",
    "is_dir": false,
    "extension": "txt"
  }
]
```

### CSV  
Spreadsheet-compatible output for analysis and reporting.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

MIT License - see LICENSE file for details.

## Version

Current version: 1.0.0

---

**Filer** - Making file management simple, powerful, and efficient.
