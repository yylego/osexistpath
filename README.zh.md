<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/osexistpath/release.yml?branch=main&label=BUILD)](https://github.com/yylego/osexistpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/osexistpath)](https://pkg.go.dev/github.com/yylego/osexistpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/osexistpath/main.svg)](https://coveralls.io/github/yylego/osexistpath?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/osexistpath.svg)](https://github.com/yylego/osexistpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/osexistpath)](https://goreportcard.com/report/github.com/yylego/osexistpath)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# osexistpath

检查路径、文件、目录是否存在，用模式调节日志多少，并决定类型不对算不算错。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

## 主要特性

- 🎯 **路径存在性验证**: 全面检查路径、文件和 DIR 的存在性
- ⚡ **检查模式**: Noisy / Sweet / Quiet 调节日志多少，Might 是在问问题，类型不对只返回 false
- 🔄 **类型安全验证**: 确保路径匹配预期类型（文件 vs DIR）
- 🌍 **错误上下文增强**: 丰富的错误包装，提供详细的位置信息
- 📋 **便捷函数**: 常见用例的简单一行存在性检查

## 安装

```bash
go get github.com/yylego/osexistpath
```

## 使用方法

### 示例 1: 使用 osmustexist 进行路径验证

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

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 示例 2: 使用 osmustexist 检查文件存在性

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

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

## API 参考

### 核心函数 (osexistpath)

#### 路径存在性检查

| 包名        | 函数                                     | 说明                            |
| ----------- | ---------------------------------------- | ------------------------------- |
| osexistpath | `IsPathExists(path, verb) (bool, error)` | 使用可配置日志检查路径是否存在  |
| osexistpath | `IsFileExists(path, verb) (bool, error)` | 检查文件是否存在并进行类型验证  |
| osexistpath | `IsRootExists(path, verb) (bool, error)` | 检查 DIR 是否存在并进行类型验证 |

#### 便捷函数 (Quiet 模式)

| 包名        | 函数                              | 说明                        |
| ----------- | --------------------------------- | --------------------------- |
| osexistpath | `IsPathExist(path) (bool, error)` | 使用安静模式检查路径存在性  |
| osexistpath | `IsFileExist(path) (bool, error)` | 使用安静模式检查文件存在性  |
| osexistpath | `IsRootExist(path) (bool, error)` | 使用安静模式检查 DIR 存在性 |

#### 查询式函数 (Might 模式)

| 包名        | 函数                         | 说明                        |
| ----------- | ---------------------------- | --------------------------- |
| osexistpath | `IsPath(path) (bool, error)` | 使用询问模式查询路径存在性  |
| osexistpath | `IsFile(path) (bool, error)` | 使用询问模式查询文件存在性  |
| osexistpath | `IsRoot(path) (bool, error)` | 使用询问模式查询 DIR 存在性 |

#### 路径验证函数

| 包名        | 函数                         | 说明                                   |
| ----------- | ---------------------------- | -------------------------------------- |
| osexistpath | `PATH(path) (string, error)` | 路径存在时返回路径，否则返回错误       |
| osexistpath | `FILE(path) (string, error)` | 有效文件时返回文件路径，否则返回错误   |
| osexistpath | `ROOT(path) (string, error)` | 有效 DIR 时返回 DIR 路径，否则返回错误 |

#### 必须函数

| 包名        | 函数             | 说明                              |
| ----------- | ---------------- | --------------------------------- |
| osexistpath | `MustPath(path)` | 断言路径存在，不存在则触发 panic  |
| osexistpath | `MustFile(path)` | 断言文件存在，不存在则触发 panic  |
| osexistpath | `MustRoot(path)` | 断言 DIR 存在，不存在则触发 panic |

### osmustexist 包（推荐使用）

`osmustexist` 包通过自动使用 `sure.Must()` 处理错误，提供了简洁的 API。这消除了手动检查错误返回的需要，使代码更加简洁和可读。

#### 存在性检查（无错误返回）

| 包名        | 函数                            | 说明                                        |
| ----------- | ------------------------------- | ------------------------------------------- |
| osmustexist | `IsPathExists(path, verb) bool` | 检查路径存在性并记录日志，错误时 panic      |
| osmustexist | `IsFileExists(path, verb) bool` | 检查文件存在性并进行类型验证，错误时 panic  |
| osmustexist | `IsRootExists(path, verb) bool` | 检查 DIR 存在性并进行类型验证，错误时 panic |

#### 便捷函数 (Quiet 模式)

| 包名        | 函数                     | 说明                                      |
| ----------- | ------------------------ | ----------------------------------------- |
| osmustexist | `IsPathExist(path) bool` | 使用安静模式检查路径存在性，错误时 panic  |
| osmustexist | `IsFileExist(path) bool` | 使用安静模式检查文件存在性，错误时 panic  |
| osmustexist | `IsRootExist(path) bool` | 使用安静模式检查 DIR 存在性，错误时 panic |

#### 查询函数 (Might 模式)

| 包名        | 函数                | 说明                                      |
| ----------- | ------------------- | ----------------------------------------- |
| osmustexist | `IsPath(path) bool` | 使用询问模式查询路径存在性，错误时 panic  |
| osmustexist | `IsFile(path) bool` | 使用询问模式查询文件存在性，错误时 panic  |
| osmustexist | `IsRoot(path) bool` | 使用询问模式查询 DIR 存在性，错误时 panic |

#### 验证函数

| 包名        | 函数                | 说明                                 |
| ----------- | ------------------- | ------------------------------------ |
| osmustexist | `PATH(path) string` | 路径存在时返回路径，否则 panic       |
| osmustexist | `FILE(path) string` | 有效文件时返回文件路径，否则 panic   |
| osmustexist | `ROOT(path) string` | 有效 DIR 时返回 DIR 路径，否则 panic |

#### 必须函数

| 包名        | 函数             | 说明                          |
| ----------- | ---------------- | ----------------------------- |
| osmustexist | `MustPath(path)` | 断言路径存在，不存在则 panic  |
| osmustexist | `MustFile(path)` | 断言文件存在，不存在则 panic  |
| osmustexist | `MustRoot(path)` | 断言 DIR 存在，不存在则 panic |

### CheckMode 类型

| 模式    | 说明                                     |
| ------- | ---------------------------------------- |
| `Noisy` | 记录所有错误和调试信息；类型不对算错     |
| `Sweet` | 仅记录意外错误；类型不对算错             |
| `Quiet` | 错误时不记录日志；类型不对算错           |
| `Might` | 是在问问题：类型不对只返回 false，不算错 |

简单说：`Exist` / `Exists` 系列、`PATH` / `FILE` / `ROOT` 和 `Must*` 都期望路径是它名字里说的那种类型，期望文件却遇到目录就算错（在 osmustexist 里就是 panic）。想在 `if` 里问一问又不担这个风险，用 `IsPath` / `IsFile` / `IsRoot`。

### 使用模式

**环境设置（推荐 - osmustexist）：**

```go
// 使用 osmustexist 进行简洁的设置代码 - 无需错误处理
osmustexist.MustRoot("/project/config")
osmustexist.MustFile("/project/config/app.yaml")
```

**运行时验证（推荐 - osmustexist）：**

```go
// IsFile 是在问：不存在返回 false，那里是个目录也返回 false，不会 panic
if osmustexist.IsFile("config.txt") {
    // 处理配置文件 - 无需处理错误
}
```

**检查操作（推荐 - osmustexist）：**

```go
// 使用 osmustexist 进行简单的探测检查
hasConfig := osmustexist.IsFile("config.json")
hasYamlFile := osmustexist.IsFile("config.yaml")
```

**标准错误处理（当需要时）：**

```go
// 当需要显式错误处理时使用基础包
if exists, err := osexistpath.IsFileExist("config.txt"); err != nil {
    return fmt.Errorf("检查文件失败: %w", err)
} else if exists {
    // 处理配置文件
}
```

## 示例

### 错误处理模式

**使用 osmustexist.PATH 进行简洁的验证（推荐）：**

```go
// 无需错误处理 - 失败时会 panic
configPath := osmustexist.PATH("/app/config")
```

**标准 PATH 函数进行显式错误处理：**

```go
configPath, err := osexistpath.PATH("/app/config")
if err != nil {
    return fmt.Errorf("配置 DIR 缺失: %w", err)
}
```

**使用 osmustexist.ROOT 链式验证（推荐）：**

```go
// 简洁可读 - 无需错误变量
dataPath := osmustexist.ROOT("/data")
logsPath := osmustexist.ROOT(filepath.Join(dataPath, "logs"))
```

### 不同日志模式

**osmustexist Noisy 模式用于调试（推荐）：**

```go
exists := osmustexist.IsFileExists("debug.log", osexistpath.Noisy)
// 记录详细的调试信息，语法简洁
```

**osmustexist Sweet 模式用于生产（推荐）：**

```go
exists := osmustexist.IsRootExists("/var/lib/app", osexistpath.Sweet)
// 仅记录意外错误，无需错误处理
```

### 类型安全验证

**osmustexist：确保路径是文件（推荐）：**

```go
if !osmustexist.IsFile("data.json") {
    return errors.New("data.json 处没有文件")
}
```

**osmustexist：确保路径确实是 DIR（推荐）：**

```go
if !osmustexist.IsRoot("uploads") {
    return errors.New("uploads 处没有 DIR")
}
```

### 基于条件的操作

**osmustexist：处理配置文件（推荐）：**

```go
if osmustexist.IsFile("app-config.yaml") {
    config := loadConfig("app-config.yaml")
    useConfig(config)
}
```

**osmustexist：根据日志 DIR 存在性设置日志（推荐）：**

```go
if osmustexist.IsRoot("/var/log/myapp") {
    setupFileLogging("/var/log/myapp")
} else {
    setupConsoleLogging()
}
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->
