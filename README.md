<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/osexistpath/release.yml?branch=main&label=BUILD)](https://github.com/yylego/osexistpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/osexistpath)](https://pkg.go.dev/github.com/yylego/osexistpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/osexistpath/main.svg)](https://coveralls.io/github/yylego/osexistpath?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/osexistpath.svg)](https://github.com/yylego/osexistpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/osexistpath)](https://goreportcard.com/report/github.com/yylego/osexistpath)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# osexistpath

Check if a path exists, and if it is a file / DIR, with modes that tune the logging and decide if a wrong type counts as a fault.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

## Main Features

- 🎯 **Path Existence Validation**: Comprehensive checking for paths, files, and directories
- ⚡ **Check Modes**: Noisy / Sweet / Quiet tune the logging, while Might asks a question and treats a wrong type as a plain false
- 🔄 **Type-Safe Validation**: Ensures paths match expected types (file vs DIR)
- 🌍 **Error Context Enhancement**: Rich error wrapping with detailed location information
- 📋 **Convenience Functions**: Simple one-line existence checks for common use cases

## Installation

```bash
go get github.com/yylego/osexistpath
```

## Usage

### Example 1: Using osmustexist for Path Validation

```go
package main

import (
	"os"
	"path/filepath"

	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

func main() {
	tempDIR := rese.V1(os.MkdirTemp("", "osexistpath-demo-*"))
	defer func() {
		must.Done(os.RemoveAll(tempDIR))
	}()

	zaplog.SUG.Debug("working in:", tempDIR)

	root := osmustexist.ROOT(tempDIR)
	zaplog.LOG.Debug("must exist DIR:", zap.String("root", root))

	path := filepath.Join(tempDIR, "test.txt")
	must.Done(os.WriteFile(path, []byte("hello world"), 0644))

	osmustexist.FILE(path)
	zaplog.LOG.Debug("must exist FILE:", zap.String("path", path))

	data := rese.A1(os.ReadFile(path))
	zaplog.LOG.Debug("read data:", zap.ByteString("data", data))
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Example 2: Checking File Existence with osmustexist

```go
package main

import (
	"os"
	"path/filepath"

	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/rese"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

func main() {
	tempDIR := osmustexist.ROOT(rese.V1(os.MkdirTemp("", "osexistpath-demo-*")))
	defer func() {
		must.Done(os.RemoveAll(tempDIR))
	}()

	zaplog.SUG.Debug("working in:", tempDIR)

	path := filepath.Join(tempDIR, "test.txt")
	must.False(osmustexist.IsPath(path))
	zaplog.LOG.Debug("must not exist FILE:", zap.String("path", path))

	must.Done(os.WriteFile(path, []byte("hello world"), 0644))
	must.True(osmustexist.IsFile(path))
	zaplog.LOG.Debug("must exist FILE:", zap.String("path", path))

	data := rese.A1(os.ReadFile(osmustexist.FILE(path)))
	zaplog.LOG.Info("read data:", zap.ByteString("data", data))
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## API Reference

### Core Functions (osexistpath)

#### Path Existence Checking

| Package     | Function                                 | Description                                    |
| ----------- | ---------------------------------------- | ---------------------------------------------- |
| osexistpath | `IsPathExists(path, verb) (bool, error)` | Check if path exists with configurable logging |
| osexistpath | `IsFileExists(path, verb) (bool, error)` | Check if file exists with type validation      |
| osexistpath | `IsRootExists(path, verb) (bool, error)` | Check if DIR exists with type validation       |

#### Convenience Functions (Quiet Mode)

| Package     | Function                          | Description                        |
| ----------- | --------------------------------- | ---------------------------------- |
| osexistpath | `IsPathExist(path) (bool, error)` | Check path existence in Quiet mode |
| osexistpath | `IsFileExist(path) (bool, error)` | Check file existence in Quiet mode |
| osexistpath | `IsRootExist(path) (bool, error)` | Check DIR existence in Quiet mode  |

#### Check-Style Functions (Might Mode)

| Package     | Function                     | Description                           |
| ----------- | ---------------------------- | ------------------------------------- |
| osexistpath | `IsPath(path) (bool, error)` | Check path existence using Might mode |
| osexistpath | `IsFile(path) (bool, error)` | Check file existence using Might mode |
| osexistpath | `IsRoot(path) (bool, error)` | Check DIR existence using Might mode  |

#### Path Validation Functions

| Package     | Function                     | Description                                     |
| ----------- | ---------------------------- | ----------------------------------------------- |
| osexistpath | `PATH(path) (string, error)` | Return path if exists, error otherwise          |
| osexistpath | `FILE(path) (string, error)` | Return file path if valid file, error otherwise |
| osexistpath | `ROOT(path) (string, error)` | Return DIR path if valid DIR, error otherwise   |

#### Must Functions

| Package     | Function         | Description                      |
| ----------- | ---------------- | -------------------------------- |
| osexistpath | `MustPath(path)` | Assert path exists, panic if not |
| osexistpath | `MustFile(path)` | Assert file exists, panic if not |
| osexistpath | `MustRoot(path)` | Assert DIR exists, panic if not  |

### osmustexist Package (Recommended)

The `osmustexist` package provides a clean API through auto handling errors with `sure.Must()`. This eliminates the need to check error returns, making code more concise and readable.

#### Existence Checking (No Error Returns)

| Package     | Function                        | Description                                               |
| ----------- | ------------------------------- | --------------------------------------------------------- |
| osmustexist | `IsPathExists(path, verb) bool` | Check path existence with logging, panic on error         |
| osmustexist | `IsFileExists(path, verb) bool` | Check file existence with type validation, panic on error |
| osmustexist | `IsRootExists(path, verb) bool` | Check DIR existence with type validation, panic on error  |

#### Convenience Functions (Quiet Mode)

| Package     | Function                 | Description                                        |
| ----------- | ------------------------ | -------------------------------------------------- |
| osmustexist | `IsPathExist(path) bool` | Check path existence in Quiet mode, panic on error |
| osmustexist | `IsFileExist(path) bool` | Check file existence in Quiet mode, panic on error |
| osmustexist | `IsRootExist(path) bool` | Check DIR existence in Quiet mode, panic on error  |

#### Check Functions (Might Mode)

| Package     | Function            | Description                                           |
| ----------- | ------------------- | ----------------------------------------------------- |
| osmustexist | `IsPath(path) bool` | Check path existence using Might mode, panic on error |
| osmustexist | `IsFile(path) bool` | Check file existence using Might mode, panic on error |
| osmustexist | `IsRoot(path) bool` | Check DIR existence using Might mode, panic on error  |

#### Validation Functions

| Package     | Function            | Description                                     |
| ----------- | ------------------- | ----------------------------------------------- |
| osmustexist | `PATH(path) string` | Return path if exists, panic otherwise          |
| osmustexist | `FILE(path) string` | Return file path if valid file, panic otherwise |
| osmustexist | `ROOT(path) string` | Return DIR path if valid DIR, panic otherwise   |

#### Must Functions

| Package     | Function         | Description                      |
| ----------- | ---------------- | -------------------------------- |
| osmustexist | `MustPath(path)` | Assert path exists, panic if not |
| osmustexist | `MustFile(path)` | Assert file exists, panic if not |
| osmustexist | `MustRoot(path)` | Assert DIR exists, panic if not  |

### CheckMode Types

| Mode    | Description                                                    |
| ------- | -------------------------------------------------------------- |
| `Noisy` | Logs each fault and debug information; a wrong type is a fault |
| `Sweet` | Logs just unexpected errors; a wrong type is a fault           |
| `Quiet` | No logging on errors; a wrong type is a fault                  |
| `Might` | Asks a question: a wrong type gives false with no fault        |

In short: the `Exist` / `Exists` functions, `PATH` / `FILE` / `ROOT` and `Must*` expect the type in the function name, so a DIR where a file is expected is a fault (a panic in osmustexist). To ask in an `if` without that risk, use `IsPath` / `IsFile` / `IsRoot`.

### Usage Patterns

**Environment Setup (Recommended - osmustexist):**

```go
// Use osmustexist for clean setup code - no error handling needed
osmustexist.MustRoot("/project/config")
osmustexist.MustFile("/project/config/app.yaml")
```

**Runtime Validation (Recommended - osmustexist):**

```go
// IsFile asks: false when missing, and false when a DIR sits there, with no panic
if osmustexist.IsFile("config.txt") {
    // Process config file - no need to handle error
}
```

**Check Operations (Recommended - osmustexist):**

```go
// Use osmustexist for simple probe checks
hasConfig := osmustexist.IsFile("config.json")
hasYamlFile := osmustexist.IsFile("config.yaml")
```

**Standard Error Handling (When needed):**

```go
// Use base package when you need explicit error handling
if exists, err := osexistpath.IsFileExist("config.txt"); err != nil {
    return fmt.Errorf("failed to check file: %w", err)
} else if exists {
    // Process config file
}
```

## Examples

### Error Handling Patterns

**Using osmustexist.PATH for clean validation (Recommended):**

```go
// No error handling needed - panics on failure
configPath := osmustexist.PATH("/app/config")
```

**Standard PATH function with explicit error handling:**

```go
configPath, err := osexistpath.PATH("/app/config")
if err != nil {
    return fmt.Errorf("config DIR missing: %w", err)
}
```

**Chaining validations with osmustexist.ROOT (Recommended):**

```go
// Clean and readable - no error variables needed
dataPath := osmustexist.ROOT("/data")
logsPath := osmustexist.ROOT(filepath.Join(dataPath, "logs"))
```

### Different Logging Modes

**osmustexist Noisy mode for debugging (Recommended):**

```go
exists := osmustexist.IsFileExists("debug.log", osexistpath.Noisy)
// Logs detailed debug information, clean syntax
```

**osmustexist Sweet mode for production (Recommended):**

```go
exists := osmustexist.IsRootExists("/var/lib/app", osexistpath.Sweet)
// Logs just unexpected errors, no error handling needed
```

### Type Safety Validation

**osmustexist: Ensure path is a file (Recommended):**

```go
if !osmustexist.IsFile("data.json") {
    return errors.New("found no file at data.json")
}
```

**osmustexist: Ensure path is a DIR (Recommended):**

```go
if !osmustexist.IsRoot("uploads") {
    return errors.New("found no DIR at uploads")
}
```

### Condition-Based Operations

**osmustexist: Process config files (Recommended):**

```go
if osmustexist.IsFile("app-config.yaml") {
    config := loadConfig("app-config.yaml")
    useConfig(config)
}
```

**osmustexist: Setup logging DIR if exists (Recommended):**

```go
if osmustexist.IsRoot("/var/log/myapp") {
    setupFileLogging("/var/log/myapp")
} else {
    setupConsoleLogging()
}
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
