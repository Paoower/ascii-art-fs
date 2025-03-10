# ASCII Art FS (File System)

A command-line utility that generates ASCII art with enhanced file system operations.

## Description

This program extends the basic ASCII art functionality by improving how it interacts with the file system. It provides better handling of banner files and can output the ASCII art to files.

## Features

- All base functionality of the ascii-art program
- Improved file system interactions
- Error handling for file operations
- Support for output redirection

## Usage

```bash
go run . [STRING] [BANNER]
go run . [STRING] [BANNER] --output=<filename>
```

Examples:
```bash
go run . "Hello World" standard
go run . "Hello World" shadow --output=result.txt
```

## Implementation Details

- Properly handles file not found errors
- Creates output files if they don't exist
- Validates banner files before processing
- Safe file system operations

## Requirements

- Go 1.13 or higher
