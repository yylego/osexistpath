// Package osexistpath: Path existence checking with configurable logging modes
// Provides comprehensive file, DIR, and path existence validation
// Supports multiple detail levels for error handling and logging
// Designed for robust file system operations with enhanced error context
//
// osexistpath: 具有可配置日志模式的路径存在性检查工具
// 提供全面的文件、目录和路径存在性验证功能
// 支持多种详细程度的错误处理和日志记录
// 专为具有增强错误上下文的健壮文件系统操作而设计
package osexistpath

import (
	"os"

	"github.com/pkg/errors"
	"github.com/yylego/erero"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

type CheckMode string

const (
	Noisy CheckMode = "NOISY" // Noisy mode, logs all errors // 比较吵闹的，即无论什么错误都会打印
	Sweet CheckMode = "SWEET" // Sweet mode, logs unexpected errors // 有点适度的，即只打印预料之外的错误
	Quiet CheckMode = "QUIET" // Quiet mode, no logs on errors // 比较安静的，即出错时不打印任何日志
	Might CheckMode = "MIGHT" // Might mode, checks type without raising errors // 表示询问的，即询问文件是何种类型的，在调用函数前没有预期，因此不抛出错误
)

// IsPathExists checks if there is something at the specified path
// Supports configurable detail levels for different logging behaviors
// Returns boolean existence status and possible errors from file system operations
//
// IsPathExists 检查指定路径下是否有内容
// 支持可配置的详细级别以适应不同的日志行为
// 返回布尔存在状态和文件系统操作的潜在错误
func IsPathExists(path string, verb CheckMode) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if verb == Noisy {
				zaplog.LOGS.Skip1.Debug("IS_PATH_EXISTS not exists", zap.String("path", path), zap.Bool("exist", false))
			}
			return false, nil // 路径不存在
		}
		if verb != Quiet {
			zaplog.LOGS.Skip1.Error("IS_PATH_EXISTS", zap.String("path", path), zap.Error(err))
		}
		return false, erero.Wro(err) // 其他的错误
	}
	return true, nil
}

// IsFileExists checks if a file exists at the specified path
// Validates that the path exists AND is a standard file (not a DIR)
// Returns boolean existence status and handles type validation errors
//
// IsFileExists 检查指定路径的文件是否存在
// 验证路径存在且为常规文件（非目录）
// 返回布尔存在状态并处理类型验证错误
func IsFileExists(path string, verb CheckMode) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if verb == Noisy {
				zaplog.LOGS.Skip1.Debug("IS_FILE_EXISTS not exists", zap.String("path", path), zap.Bool("exist", false))
			}
			return false, nil // 路径不存在
		}
		if verb != Quiet {
			zaplog.LOGS.Skip1.Error("IS_FILE_EXISTS wrong stat", zap.String("path", path), zap.Error(err))
		}
		return false, erero.Wro(err) // 其他的错误
	}
	if info.IsDir() {
		//这里不要直接打错误日志，需要根据情况而定
		err := errors.New("PATH-EXIST-BUT-TYPE-IS-NOT-FILE")

		switch verb {
		case Noisy, Sweet:
			zaplog.LOGS.Skip1.Debug("IS_FILE_EXISTS wrong type", zap.String("path", path), zap.String("type", "ROOT"))
			return false, err
		case Quiet:
			return false, err
		case Might:
			// 这里不要返回错误，因为这里的意图是判定是不是这个类型，在逻辑中是if询问性质的，没有明确的预期
			return false, nil
		default:
			// 这里必须返回错误，否则判定不存在接着创建文件就没法创建，而强行写/删也会出问题
			return false, err
		}
	}
	//这里依然要判断类型避免漏判
	return !info.IsDir(), nil
}

// IsRootExists checks if a DIR exists at the specified path
// Validates that the path exists AND is a DIR (not a standard file)
// Returns boolean existence status and handles type validation errors
//
// IsRootExists 检查指定路径的目录是否存在
// 验证路径存在且为目录（非常规文件）
// 返回布尔存在状态并处理类型验证错误
func IsRootExists(path string, verb CheckMode) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			if verb == Noisy {
				zaplog.LOGS.Skip1.Debug("IS_ROOT_EXISTS not exists", zap.String("path", path), zap.Bool("exist", false))
			}
			return false, nil // 路径不存在
		}
		if verb != Quiet {
			zaplog.LOGS.Skip1.Error("IS_ROOT_EXISTS wrong stat", zap.String("path", path), zap.Error(err))
		}
		return false, erero.Wro(err) // 其他的错误
	}
	if !info.IsDir() {
		//这里不要直接打错误日志，需要根据情况而定
		err := errors.New("PATH-EXIST-BUT-TYPE-IS-NOT-ROOT")

		switch verb {
		case Noisy, Sweet:
			zaplog.LOGS.Skip1.Debug("IS_ROOT_EXISTS wrong type", zap.String("path", path), zap.String("type", "FILE"))
			return false, err
		case Quiet:
			return false, err
		case Might:
			// 这里不要返回错误，因为这里的意图是判定是不是这个类型，在逻辑中是if询问性质的，没有明确的预期
			return false, nil
		default:
			// 这里必须返回错误，否则判定不存在接着创建目录就没法创建，而强行建/删也会出问题
			return false, err
		}
	}
	//这里依然要判断类型避免漏判
	return info.IsDir(), nil
}

// IsPathExist checks if there is something at the specified path
// Convenience function using Quiet mode for silent operation
// Returns boolean existence status without logging errors
//
// IsPathExist 检查指定路径下是否有内容
// 使用安静模式的便利函数，用于静默操作
// 返回布尔存在状态，不记录错误日志
func IsPathExist(path string) (bool, error) {
	return IsPathExists(path, Quiet)
}

// IsFileExist checks if a file exists at the specified path
// Convenience function using Quiet mode for silent file validation
// Returns boolean existence status without logging errors
//
// IsFileExist 检查指定路径的文件是否存在
// 使用安静模式的便利函数，用于静默文件验证
// 返回布尔存在状态，不记录错误日志
func IsFileExist(path string) (bool, error) {
	return IsFileExists(path, Quiet)
}

// IsRootExist checks if a DIR exists at the specified path
// Convenience function using Quiet mode for silent DIR validation
// Returns boolean existence status without logging errors
//
// IsRootExist 检查指定路径的目录是否存在
// 使用安静模式的便利函数，用于静默目录验证
// 返回布尔存在状态，不记录错误日志
func IsRootExist(path string) (bool, error) {
	return IsRootExists(path, Quiet)
}

// IsPath checks if there is something at the specified path
// Uses Might mode for query-style checking without error expectations
// Returns boolean existence status with basic error handling
//
// IsPath 检查指定路径下是否有内容
// 使用询问模式进行探查式检查，无错误预期
// 返回布尔存在状态，具有最小的错误处理
func IsPath(path string) (bool, error) {
	return IsPathExists(path, Might)
}

// IsFile checks if a file exists at the specified path
// Uses Might mode for query-style file type checking
// Returns boolean existence status with basic error handling
//
// IsFile 检查指定路径的文件是否存在
// 使用询问模式进行探查式文件类型检查
// 返回布尔存在状态，具有最小的错误处理
func IsFile(path string) (bool, error) {
	return IsFileExists(path, Might)
}

// IsRoot checks if a DIR exists at the specified path
// Uses Might mode for query-style DIR type checking
// Returns boolean existence status with basic error handling
//
// IsRoot 检查指定路径的目录是否存在
// 使用询问模式进行探查式目录类型检查
// 返回布尔存在状态，具有最小的错误处理
func IsRoot(path string) (bool, error) {
	return IsRootExists(path, Might)
}

// PATH validates path existence and returns the path if valid
// Returns the source path string when validation passes
// Returns error when path doesn't exist or validation fails
//
// PATH 验证路径存在性，成功时返回路径
// 验证通过时返回原始路径字符串
// 路径不存在或验证失败时返回错误
func PATH(path string) (string, error) {
	exist, err := IsPathExists(path, Sweet)
	if err != nil {
		zaplog.LOGS.Skip1.Error("path", zap.String("path", path), zap.Error(err))
		return "", erero.Wro(err)
	}
	if !exist {
		zaplog.LOGS.Skip1.Debug("path", zap.String("path", path), zap.Bool("exist", false))
		return "", os.ErrNotExist
	}
	return path, nil
}

// FILE validates file existence and returns the path if valid
// Returns the source path string when file validation passes
// Returns error when file doesn't exist or is not a standard file
//
// FILE 验证文件存在性，成功时返回路径
// 文件验证通过时返回原始路径字符串
// 文件不存在或非常规文件时返回错误
func FILE(path string) (string, error) {
	exist, err := IsFileExists(path, Sweet)
	if err != nil {
		zaplog.LOGS.Skip1.Error("file", zap.String("path", path), zap.Error(err))
		return "", erero.Wro(err)
	}
	if !exist {
		zaplog.LOGS.Skip1.Debug("file", zap.String("path", path), zap.Bool("exist", false))
		return "", os.ErrNotExist
	}
	return path, nil
}

// ROOT validates DIR existence and returns the path if valid
// Returns the source path string when DIR validation passes
// Returns error when DIR doesn't exist or is not a DIR
//
// ROOT 验证目录存在性，成功时返回路径
// 目录验证通过时返回原始路径字符串
// 目录不存在或非目录时返回错误
func ROOT(path string) (string, error) {
	exist, err := IsRootExists(path, Sweet)
	if err != nil {
		zaplog.LOGS.Skip1.Error("root", zap.String("path", path), zap.Error(err))
		return "", erero.Wro(err)
	}
	if !exist {
		zaplog.LOGS.Skip1.Debug("root", zap.String("path", path), zap.Bool("exist", false))
		return "", os.ErrNotExist
	}
	return path, nil
}

// MustPath validates path existence and panics if validation fails
// Uses Noisy mode to log detailed error information before panicking
// Designed for vital path validation where failing should halt execution
//
// MustPath 验证路径存在性，验证失败时触发 panic
// 使用详细模式在 panic 前记录详细错误信息
// 用于关键路径验证，失败时应停止执行
func MustPath(path string) {
	exist, err := IsPathExists(path, Noisy)
	if err != nil {
		zaplog.LOGS.Skip1.Panic("must_path", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.Skip1.Panic("must_path", zap.String("path", path), zap.Bool("exist", false))
	}
}

// MustFile validates file existence and panics if validation fails
// Uses Noisy mode to log detailed error information before panicking
// Designed for vital file validation where failing should halt execution
//
// MustFile 验证文件存在性，验证失败时触发 panic
// 使用详细模式在 panic 前记录详细错误信息
// 用于关键文件验证，失败时应停止执行
func MustFile(path string) {
	exist, err := IsFileExists(path, Noisy)
	if err != nil {
		zaplog.LOGS.Skip1.Panic("must_file", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.Skip1.Panic("must_file", zap.String("path", path), zap.Bool("exist", false))
	}
}

// MustRoot validates DIR existence and panics if validation fails
// Uses Noisy mode to log detailed error information before panicking
// Designed for vital DIR validation where failing should halt execution
//
// MustRoot 验证目录存在性，验证失败时触发 panic
// 使用详细模式在 panic 前记录详细错误信息
// 用于关键目录验证，失败时应停止执行
func MustRoot(path string) {
	exist, err := IsRootExists(path, Noisy)
	if err != nil {
		zaplog.LOGS.Skip1.Panic("must_root", zap.String("path", path), zap.Error(err))
	}
	if !exist {
		zaplog.LOGS.Skip1.Panic("must_root", zap.String("path", path), zap.Bool("exist", false))
	}
}
