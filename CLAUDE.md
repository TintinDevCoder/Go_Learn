# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository is a Go language learning and practice codebase organized into three independent Go modules, each serving a distinct purpose:
- **CNJC**: Fundamental Go syntax and core concepts (functions, data structures, concurrency)
- **SGG**: Structured learning path from basics to advanced topics (pointers, structs, goroutines, testing)
- **leetcode**: LeetCode algorithm problem solutions with completion tracking

Each module has its own `go.mod` file and operates independently with no shared packages. Go version 1.25 is used throughout.

## Common Commands

### Running Programs
```bash
# Run the simple example in CNJC
cd CNJC
go run main.go

# Run the SGG module entry point
cd SGG
go run main.go

# Run a specific LeetCode solution
cd leetcode
go run t1.go
```

### Building Executables
```bash
# Build a specific file (Windows .exe output)
cd CNJC
go build -o program.exe Array.go

# Build all files in a module
cd SGG
go build ./...
```

### Testing
Only one test file exists (`SGG/advanced/union_test/u1_test.go`), demonstrating Go's testing framework.
```bash
# Run all tests in the SGG module
cd SGG
go test -v ./...

# Run a specific test function
cd SGG
go test -v ./advanced/union_test/ -test.run TestAdd
```

### Code Formatting
```bash
# Format all Go files in a module
cd leetcode
go fmt ./...
```

## Project Structure Quick Reference

```
Go_Learn/
├── CNJC/                      # Core Go concepts (individual examples)
│   ├── *.go                   # Files named after concepts: function.go, slice.go, etc.
│   └── genericExmaple/        # Generic type examples (Stack.go, Map.go)
├── SGG/                       # Structured learning
│   ├── basic/                 # Fundamentals: data types, variables, operators
│   ├── advanced/              # Advanced topics: pointers, structs, goroutines, testing
│   ├── packageRules/          # Module structure examples
│   ├── test/                  # Practice exercises (CustomerRelationshipManagement.go)
│   └── README.md              # Learning outline and file descriptions
└── leetcode/                  # LeetCode solutions
    ├── t1.go                  # Main algorithm solutions file
    ├── status.md              # Problem completion tracker (over, review, dp categories)
    ├── TreeNode/              # Binary tree node types
    ├── Codec/                 # Tree serialization/deserialization
    └── LRU.go                 # LRU cache implementation
```

## Module Information
- **Root Module**: `module Go_Learn` (go 1.25)
- **CNJC Module**: `module CNJC` (go 1.25)
- **SGG Module**: `module SGG` (go 1.25)
- **LeetCode Module**: `module leetcode` (go 1.25)

## Development Notes
- **IDE**: GoLand project files (`.idea/`) are present; the code includes IDE tip comments.
- **Comments**: Many files contain detailed Chinese comments explaining concepts.
- **Executables**: Compiled Windows `.exe` files exist in some directories; they can be regenerated with `go build`.
- **Testing**: The single test file serves as an example; new tests should follow the same pattern (`*_test.go`, `TestXxx` functions).
- **Dependencies**: No external dependencies; all modules use only the Go standard library.
