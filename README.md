# GoSh - POSIX Compliant Shell Implementation in Go

[![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org/doc/go1.22)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CodeCrafters](https://img.shields.io/badge/CodeCrafters-Shell%20Challenge-orange.svg)](https://app.codecrafters.io/courses/shell/overview)

A custom POSIX-compliant shell implementation written in Go as part of the [CodeCrafters "Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview). This shell provides essential command-line functionality including built-in commands, external program execution, and advanced features like quote handling.

## Features

### Built-in Commands
- **`echo`** - Display text with support for single and double quotes
- **`exit`** - Exit the shell with optional exit codes
- **`type`** - Display information about command types (builtin vs external)
- **`pwd`** - Print current working directory
- **`cd`** - Change directory with support for `~` (home directory)

### Advanced Features
- **External Program Execution** - Run any executable found in your PATH
- **Quote Handling** - Proper parsing of single (`'`) and double (`"`) quoted strings
- **PATH Resolution** - Automatic discovery and execution of external commands
- **Error Handling** - Comprehensive error reporting for invalid commands and arguments
- **REPL Interface** - Interactive command prompt with continuous input processing

## Quick Start

### Prerequisites
- Go 1.22 or later
- Unix-like operating system (Linux, macOS, WSL)

### Installation & Usage

1. **Clone the repository:**
   ```bash
   git clone https://github.com/EvgeniiKlepilin/gosh.git
   cd gosh
   ```

2. **Compile and run manually:**
   ```bash
   go build -o gosh cmd/myshell/*.go
   ./gosh
   ```

### Example Usage

```bash
$ echo "Hello, World!"
Hello, World!

$ pwd
/home/user/projects/GoSh

$ cd ~
$ pwd
/home/user

$ type echo
echo is a shell builtin

$ type ls
ls is /usr/bin/ls

$ ls -la
# Lists directory contents

$ exit 0
```

## Project Structure

```
├── cmd/myshell/
│   └── main.go              # Main shell implementation
├── .codecrafters/
│   ├── compile.sh           # Build script for CodeCrafters
│   └── run.sh              # Run script for CodeCrafters
├── your_program.sh         # Local development script
├── codecrafters.yml        # CodeCrafters configuration
├── go.mod                  # Go module definition
└── README.md               # This file
```

## Implementation Details

### Command Processing Flow
1. **Input Reading** - Uses `bufio.Scanner` for line-by-line input
2. **Quote Parsing** - Custom `HandleQuotes()` function processes quoted strings
3. **Command Classification** - Determines if command is builtin or external
4. **Execution** - Routes to appropriate handler function
5. **Error Handling** - Provides meaningful error messages

### Built-in Command Implementation
Each built-in command has its own dedicated function:
- `ExitCommand()` - Handles exit codes and graceful shutdown
- `EchoCommand()` - Processes various quote scenarios
- `TypeCommand()` - Searches builtins and PATH for commands
- `PwdCommand()` - Returns current working directory
- `CdCommand()` - Changes directory with error handling

### External Command Execution
The shell searches through PATH directories to find and execute external programs using Go's `os/exec` package.

## Development

### Running Tests
This project is designed to work with the CodeCrafters testing platform:

```bash
git commit -am "your changes"
git push origin master
```

### Local Development
Use the provided `your_program.sh` script for local testing and development.

## Contributing

This is a learning project as part of the CodeCrafters challenge. While it's primarily educational, improvements and bug fixes are welcome.

## License

This project is part of the CodeCrafters educational platform. Please refer to CodeCrafters' terms of service for usage guidelines.

## Acknowledgments

- [CodeCrafters](https://codecrafters.io) for the excellent shell implementation challenge
- The Go community for comprehensive documentation and examples
- POSIX standards for shell behavior specifications
