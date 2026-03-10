// Package osexistpath_test: Comprehensive test suite for path existence checking
// Tests all main functions across different detail modes and edge cases
// Validates file, DIR, and path existence detection with type checks
//
// osexistpath_test: 路径存在性检查的综合测试套件
// 测试所有主要函数在不同详细模式和边界情况下的表现
// 验证文件、目录和路径存在性检测的类型安全性
package osexistpath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/osexistpath"
	"github.com/yylego/runpath"
	"github.com/yylego/runpath/runtestpath"
)

// TestIsPathExists validates path existence checking with Sweet detail mode
// Tests positive cases with current and parent paths
// Tests negative cases with non-existent file paths
//
// TestIsPathExists 使用 Sweet 详细模式验证路径存在性检查
// 测试当前和父路径的正向情况
// 测试不存在文件路径的负向情况
func TestIsPathExists(t *testing.T) {
	done.VBE(osexistpath.IsPathExists(runpath.CurrentPath(), osexistpath.Sweet)).TRUE()
	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Path(), osexistpath.Sweet)).TRUE()

	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Join("abc.txt"), osexistpath.Sweet)).FALSE()
	done.VBE(osexistpath.IsPathExists(runpath.PARENT.Join("xyz.uvw"), osexistpath.Sweet)).FALSE()
}

// TestIsFileExists validates file existence checking with type validation
// Tests positive cases with real source files
// Tests negative cases and type validation errors for directories
//
// TestIsFileExists 验证带有类型验证的文件存在性检查
// 测试实际源文件的正向情况
// 测试目录的负向情况和类型验证错误
func TestIsFileExists(t *testing.T) {
	done.VBE(osexistpath.IsFileExists(runtestpath.SrcPath(t), osexistpath.Sweet)).TRUE()

	done.VBE(osexistpath.IsFileExists(runpath.PARENT.Join("abc.txt"), osexistpath.Sweet)).FALSE()
	done.VBE(osexistpath.IsFileExists(runpath.PARENT.Join("xyz.uvw"), osexistpath.Sweet)).FALSE()

	require.Error(t, done.VBE(osexistpath.IsFileExists(runpath.PARENT.Path(), osexistpath.Sweet)).E)
}

// TestIsRootExists validates DIR existence checking with type validation
// Tests positive cases with real DIR paths
// Tests negative cases and type validation errors for files
//
// TestIsRootExists 验证带有类型验证的目录存在性检查
// 测试实际目录路径的正向情况
// 测试文件的负向情况和类型验证错误
func TestIsRootExists(t *testing.T) {
	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Path(), osexistpath.Sweet)).TRUE()

	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Join("abc.txt"), osexistpath.Sweet)).FALSE()
	done.VBE(osexistpath.IsRootExists(runpath.PARENT.Join("xyz.uvw"), osexistpath.Sweet)).FALSE()

	require.Error(t, done.VBE(osexistpath.IsRootExists(runtestpath.SrcPath(t), osexistpath.Sweet)).E)
}

// TestPath validates query-style path checking using Might mode
// Tests path existence without error expectations
// Verifies basic error handling actions
//
// TestPath 使用 Might 模式验证探查式路径检查
// 测试无错误预期的路径存在性
// 验证最小错误处理行为
func TestPath(t *testing.T) {
	done.VBE(osexistpath.IsPath(runpath.CurrentPath())).TRUE()
	done.VBE(osexistpath.IsPath(runpath.PARENT.Path())).TRUE()

	done.VBE(osexistpath.IsPath(runpath.PARENT.Join("abc.txt"))).FALSE()
	done.VBE(osexistpath.IsPath(runpath.PARENT.Join("xyz.uvw"))).FALSE()
}

// TestFile validates query-style file checking using Might mode
// Tests file existence with basic error handling
// Verifies type checking actions without exceptions
//
// TestFile 使用 Might 模式验证探查式文件检查
// 测试具有最小错误处理的文件存在性
// 验证无异常的类型检查行为
func TestFile(t *testing.T) {
	done.VBE(osexistpath.IsFile(runtestpath.SrcPath(t))).TRUE()
}

// TestRoot validates query-style DIR checking using Might mode
// Tests DIR existence with basic error handling
// Verifies type checking actions without exceptions
//
// TestRoot 使用 Might 模式验证探查式目录检查
// 测试具有最小错误处理的目录存在性
// 验证无异常的类型检查行为
func TestRoot(t *testing.T) {
	done.VBE(osexistpath.IsRoot(runpath.PARENT.Path())).TRUE()
}

// TestMustPath validates panic-on-failure path validation
// Tests vital path validation that should halt execution on failing
// Verifies valid validation with existing paths
//
// TestMustPath 验证失败时触发 panic 的路径验证
// 测试失败时应停止执行的关键路径验证
// 验证现有路径的成功验证
func TestMustPath(t *testing.T) {
	osexistpath.MustPath(runpath.CurrentPath())
	osexistpath.MustPath(runpath.PARENT.Path())
}

// TestMustFile validates panic-on-failure file validation
// Tests vital file validation that should halt execution on failing
// Verifies valid validation with existing source files
//
// TestMustFile 验证失败时触发 panic 的文件验证
// 测试失败时应停止执行的关键文件验证
// 验证现有源文件的成功验证
func TestMustFile(t *testing.T) {
	osexistpath.MustFile(runtestpath.SrcPath(t))
}

// TestMustRoot validates panic-on-failure DIR validation
// Tests vital DIR validation that should halt execution on failing
// Verifies valid validation with existing DIR paths
//
// TestMustRoot 验证失败时触发 panic 的目录验证
// 测试失败时应停止执行的关键目录验证
// 验证现有目录路径的成功验证
func TestMustRoot(t *testing.T) {
	osexistpath.MustRoot(runpath.PARENT.Path())
}
