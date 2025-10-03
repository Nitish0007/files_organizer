# Personal Files Organizer

A powerful CLI tool written in Go that automatically organizes your files by type into categorized folders. Perfect for keeping your Downloads and Documents directories clean and organized.

## 🚀 Features

- **Automatic File Categorization**: Organizes files into 7+ categories (Images, Documents, Videos, Audio, Archives, Code, Executables, Others(Rest goes here))
- **Concurrent Processing**: Uses worker pools for efficient file processing
- **Interactive CLI**: User-friendly command-line interface with prompts
- **Safe Operations**: Uses atomic file moves to prevent data loss
- **Customizable Directories**: Choose which directory to organize

## 📁 Supported File Types

### Images
- `.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`, `.svg`, `.webp`

### Documents
- `.pdf`, `.doc`, `.docx`, `.txt`, `.rtf`, `.odt`, `.xls`, `.csv`

### Videos
- `.mp4`, `.avi`, `.mov`, `.wmv`, `.flv`, `.mkv`, `.webm`

### Audio
- `.mp3`, `.wav`, `.flac`, `.aac`, `.ogg`

### Archives
- `.zip`, `.rar`, `.7z`, `.tar`, `.gz`

### Code Files
- `.go`, `.js`, `.py`, `.java`, `.cpp`, `.c`, `.html`, `.css`, `.php`, `.rb`

### Executables
- `.exe`, `.app`, `.dmg`, `.pkg`

## 🛠️ Installation

### Prerequisites
- Go 1.23.5 or higher

### Build from Source
```bash
# Clone the repository
git clone https://github.com/Nitish0007/files_organizer.git
cd files_organizer

# Build the executable
go build -o file-organizer ./cmd

# Run the application
./file-organizer
```

### Cross-Platform Builds
```bash
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o file-organizer.exe ./cmd

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o file-organizer-linux ./cmd

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o file-organizer-macos ./cmd
```

## 🎯 Usage

### Basic Usage
```bash
# Run the application
./file-organizer

# Or if you built it with a different name
./personal-file-organizer
```