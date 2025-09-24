#!/bin/bash

echo "=== Filer CLI Tool Demo ==="
echo

# Create some demo files for testing
echo "Creating demo files..."
mkdir -p demo_files
cd demo_files

# Create various file types
echo "Sample document content" > document.txt
echo "Another document" > report.pdf
echo "#!/bin/bash\necho hello" > script.sh
echo "{\"key\": \"value\"}" > data.json
echo "Sample,CSV,Data" > spreadsheet.csv
echo "Log entry 1" > app.log
touch image.jpg video.mp4 audio.mp3

echo "Demo files created!"
echo

# Go back to filer directory
cd ..

echo "=== Demo 1: List Files ==="
echo "Command: ./filer list demo_files"
./filer list demo_files
echo

echo "=== Demo 2: Search for Text Files ==="
echo "Command: ./filer search \"*.txt\" demo_files"
./filer search "*.txt" demo_files
echo

echo "=== Demo 3: Directory Statistics ==="
echo "Command: ./filer stats demo_files"
./filer stats demo_files
echo

echo "=== Demo 4: File Organization Preview ==="
echo "Command: ./filer organize demo_files --dry-run"
./filer organize demo_files --dry-run
echo

echo "=== Demo 5: JSON Output ==="
echo "Command: ./filer list demo_files --format json"
./filer list demo_files --format json
echo

echo "=== Demo 6: Search with Criteria ==="
echo "Command: ./filer search \"*\" demo_files --extension txt --verbose"
./filer search "*" demo_files --extension txt --verbose
echo

echo "=== Demo Complete! ==="
echo "Clean up demo files with: rm -rf demo_files"
