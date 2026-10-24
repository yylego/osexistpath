// Package osexistpath_test: Test suite for path existence checking
// Checks each function on three kinds of path: one that exists with the asked type, one that is missing, and one that exists with the wrong type
//
// osexistpath_test: 路径存在性检查的测试
// 每个函数都拿三种路径检查：存在且类型对的、不存在的、存在但类型不对的
package osexistpath_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/osexistpath"
	"github.com/yylego/runpath"
	"github.com/yylego/runpath/runtestpath"
	"github.com/yylego/sure"
	"github.com/yylego/sure/sure_pkg_gen"
	"github.com/yylego/syntaxgo/syntaxgo_reflect"
)

// TestGen regenerates the Soft / Must / Omit packages from osexistpath.go
// No -tags=sure_generate here: sure_pkg_gen emits standalone packages that this package does not depend on, so a broken artifact leaves this package alone
//
// TestGen 从 osexistpath.go 重新生成 Soft / Must / Omit 三个包
// 这里不带 -tags=sure_generate：sure_pkg_gen 生成的是独立新包，本包不依赖它，产物坏掉也不会连坐本包
//
//go:generate go test -run ^TestGen$
func TestGen(t *testing.T) {
	pkgPath := syntaxgo_reflect.GetPkgPathV2[osexistpath.CheckMode]()
	sure_pkg_gen.GenerateSurePackageFiles(sure_pkg_gen.NewSurePackageConfig(runpath.PARENT.Path(), sure.SOFT, pkgPath).WithNewPkgName("ossoftexist"))
	sure_pkg_gen.GenerateSurePackageFiles(sure_pkg_gen.NewSurePackageConfig(runpath.PARENT.Path(), sure.MUST, pkgPath).WithNewPkgName("osmustexist"))
	sure_pkg_gen.GenerateSurePackageFiles(sure_pkg_gen.NewSurePackageConfig(runpath.PARENT.Path(), sure.OMIT, pkgPath).WithNewPkgName("osomitexist"))
}

// TestIsPathExists checks that a file and a DIR both count as a path, and a missing one gives false with no fault
// TestIsPathExists 检查文件和目录都算路径，不存在的返回 false 且不报错
func TestIsPathExists(t *testing.T) {
	for _, mode := range []osexistpath.CheckMode{osexistpath.Noisy, osexistpath.Sweet, osexistpath.Quiet, osexistpath.Might} {
		exist, err := osexistpath.IsPathExists(runtestpath.SrcPath(t), mode)
		require.NoError(t, err)
		require.True(t, exist)

		exist, err = osexistpath.IsPathExists(runpath.PARENT.Path(), mode)
		require.NoError(t, err)
		require.True(t, exist)

		exist, err = osexistpath.IsPathExists(runpath.PARENT.Join("abc.txt"), mode)
		require.NoError(t, err)
		require.False(t, exist)
	}
}

// TestIsPathExistsStatFault checks a stat fault besides "not exist" comes back as an error in each mode
// A path beneath a file makes stat report "not a DIR", which is no "not exist"
//
// TestIsPathExistsStatFault 检查"不存在"以外的 stat 错误在每种模式下都会返回错误
// 拿文件下面的子路径去查，stat 报的是"不是目录"，而不是"不存在"
func TestIsPathExistsStatFault(t *testing.T) {
	childPath := runtestpath.SrcPath(t) + "/child"
	for _, mode := range []osexistpath.CheckMode{osexistpath.Noisy, osexistpath.Sweet, osexistpath.Quiet, osexistpath.Might} {
		exist, err := osexistpath.IsPathExists(childPath, mode)
		require.Error(t, err)
		require.False(t, exist)
	}
}

// TestIsFileExists checks a wrong type is a fault in Noisy / Sweet / Quiet, and a plain false in Might
// TestIsFileExists 检查类型不对在 Noisy / Sweet / Quiet 下算错，在 Might 下只返回 false
func TestIsFileExists(t *testing.T) {
	for _, mode := range []osexistpath.CheckMode{osexistpath.Noisy, osexistpath.Sweet, osexistpath.Quiet, osexistpath.Might} {
		exist, err := osexistpath.IsFileExists(runtestpath.SrcPath(t), mode)
		require.NoError(t, err)
		require.True(t, exist)

		exist, err = osexistpath.IsFileExists(runpath.PARENT.Join("abc.txt"), mode)
		require.NoError(t, err)
		require.False(t, exist)
	}

	for _, mode := range []osexistpath.CheckMode{osexistpath.Noisy, osexistpath.Sweet, osexistpath.Quiet} {
		exist, err := osexistpath.IsFileExists(runpath.PARENT.Path(), mode)
		require.Error(t, err)
		require.False(t, exist)
	}

	exist, err := osexistpath.IsFileExists(runpath.PARENT.Path(), osexistpath.Might)
	require.NoError(t, err)
	require.False(t, exist)
}

// TestIsRootExists checks a wrong type is a fault in Noisy / Sweet / Quiet, and a plain false in Might
// TestIsRootExists 检查类型不对在 Noisy / Sweet / Quiet 下算错，在 Might 下只返回 false
func TestIsRootExists(t *testing.T) {
	for _, mode := range []osexistpath.CheckMode{osexistpath.Noisy, osexistpath.Sweet, osexistpath.Quiet, osexistpath.Might} {
		exist, err := osexistpath.IsRootExists(runpath.PARENT.Path(), mode)
		require.NoError(t, err)
		require.True(t, exist)

		exist, err = osexistpath.IsRootExists(runpath.PARENT.Join("abc.txt"), mode)
		require.NoError(t, err)
		require.False(t, exist)
	}

	for _, mode := range []osexistpath.CheckMode{osexistpath.Noisy, osexistpath.Sweet, osexistpath.Quiet} {
		exist, err := osexistpath.IsRootExists(runtestpath.SrcPath(t), mode)
		require.Error(t, err)
		require.False(t, exist)
	}

	exist, err := osexistpath.IsRootExists(runtestpath.SrcPath(t), osexistpath.Might)
	require.NoError(t, err)
	require.False(t, exist)
}

// TestIsPathExist checks the Quiet shortcut on a present and a missing path
// TestIsPathExist 检查 Quiet 简写在存在、不存在两种路径上的结果
func TestIsPathExist(t *testing.T) {
	exist, err := osexistpath.IsPathExist(runpath.PARENT.Path())
	require.NoError(t, err)
	require.True(t, exist)

	exist, err = osexistpath.IsPathExist(runpath.PARENT.Join("abc.txt"))
	require.NoError(t, err)
	require.False(t, exist)
}

// TestIsFileExist checks the Quiet shortcut expects a file, so a DIR is a fault
// TestIsFileExist 检查 Quiet 简写期望的是文件，所以遇到目录算错
func TestIsFileExist(t *testing.T) {
	exist, err := osexistpath.IsFileExist(runtestpath.SrcPath(t))
	require.NoError(t, err)
	require.True(t, exist)

	exist, err = osexistpath.IsFileExist(runpath.PARENT.Join("abc.txt"))
	require.NoError(t, err)
	require.False(t, exist)

	_, err = osexistpath.IsFileExist(runpath.PARENT.Path())
	require.Error(t, err)
}

// TestIsRootExist checks the Quiet shortcut expects a DIR, so a file is a fault
// TestIsRootExist 检查 Quiet 简写期望的是目录，所以遇到文件算错
func TestIsRootExist(t *testing.T) {
	exist, err := osexistpath.IsRootExist(runpath.PARENT.Path())
	require.NoError(t, err)
	require.True(t, exist)

	exist, err = osexistpath.IsRootExist(runpath.PARENT.Join("abc.txt"))
	require.NoError(t, err)
	require.False(t, exist)

	_, err = osexistpath.IsRootExist(runtestpath.SrcPath(t))
	require.Error(t, err)
}

// TestIsPath checks the Might shortcut on a present and a missing path
// TestIsPath 检查 Might 简写在存在、不存在两种路径上的结果
func TestIsPath(t *testing.T) {
	exist, err := osexistpath.IsPath(runtestpath.SrcPath(t))
	require.NoError(t, err)
	require.True(t, exist)

	exist, err = osexistpath.IsPath(runpath.PARENT.Join("abc.txt"))
	require.NoError(t, err)
	require.False(t, exist)
}

// TestIsFile checks the Might shortcut answers false with no fault on a DIR
// TestIsFile 检查 Might 简写遇到目录时返回 false 且不报错
func TestIsFile(t *testing.T) {
	exist, err := osexistpath.IsFile(runtestpath.SrcPath(t))
	require.NoError(t, err)
	require.True(t, exist)

	exist, err = osexistpath.IsFile(runpath.PARENT.Path())
	require.NoError(t, err)
	require.False(t, exist)

	exist, err = osexistpath.IsFile(runpath.PARENT.Join("abc.txt"))
	require.NoError(t, err)
	require.False(t, exist)
}

// TestIsRoot checks the Might shortcut answers false with no fault on a file
// TestIsRoot 检查 Might 简写遇到文件时返回 false 且不报错
func TestIsRoot(t *testing.T) {
	exist, err := osexistpath.IsRoot(runpath.PARENT.Path())
	require.NoError(t, err)
	require.True(t, exist)

	exist, err = osexistpath.IsRoot(runtestpath.SrcPath(t))
	require.NoError(t, err)
	require.False(t, exist)

	exist, err = osexistpath.IsRoot(runpath.PARENT.Join("abc.txt"))
	require.NoError(t, err)
	require.False(t, exist)
}

// TestPATH checks PATH gives the path back when present, and os.ErrNotExist when missing
// TestPATH 检查 PATH 在路径存在时原样返回，不存在时返回 os.ErrNotExist
func TestPATH(t *testing.T) {
	path, err := osexistpath.PATH(runpath.PARENT.Path())
	require.NoError(t, err)
	require.Equal(t, runpath.PARENT.Path(), path)

	missingPath := runpath.PARENT.Join("abc.txt")
	_, err = osexistpath.PATH(missingPath)
	require.ErrorIs(t, err, os.ErrNotExist)
	require.ErrorContains(t, err, missingPath)
}

// TestFILE checks FILE gives the path back on a file, and a fault on a missing path and on a DIR
// TestFILE 检查 FILE 遇到文件时原样返回，遇到不存在的路径或目录时报错
func TestFILE(t *testing.T) {
	path, err := osexistpath.FILE(runtestpath.SrcPath(t))
	require.NoError(t, err)
	require.Equal(t, runtestpath.SrcPath(t), path)

	_, err = osexistpath.FILE(runpath.PARENT.Join("abc.txt"))
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = osexistpath.FILE(runpath.PARENT.Path())
	require.Error(t, err)
}

// TestROOT checks ROOT gives the path back on a DIR, and a fault on a missing path and on a file
// TestROOT 检查 ROOT 遇到目录时原样返回，遇到不存在的路径或文件时报错
func TestROOT(t *testing.T) {
	path, err := osexistpath.ROOT(runpath.PARENT.Path())
	require.NoError(t, err)
	require.Equal(t, runpath.PARENT.Path(), path)

	_, err = osexistpath.ROOT(runpath.PARENT.Join("abc.txt"))
	require.ErrorIs(t, err, os.ErrNotExist)

	_, err = osexistpath.ROOT(runtestpath.SrcPath(t))
	require.Error(t, err)
}

// TestMustPath checks MustPath passes on a present path and panics on a missing one
// TestMustPath 检查 MustPath 遇到存在的路径时通过，遇到不存在的就 panic
func TestMustPath(t *testing.T) {
	require.NotPanics(t, func() { osexistpath.MustPath(runpath.PARENT.Path()) })
	require.Panics(t, func() { osexistpath.MustPath(runpath.PARENT.Join("abc.txt")) })
}

// TestMustFile checks MustFile passes on a file, and panics on a missing path and on a DIR
// TestMustFile 检查 MustFile 遇到文件时通过，遇到不存在的路径或目录就 panic
func TestMustFile(t *testing.T) {
	require.NotPanics(t, func() { osexistpath.MustFile(runtestpath.SrcPath(t)) })
	require.Panics(t, func() { osexistpath.MustFile(runpath.PARENT.Join("abc.txt")) })
	require.Panics(t, func() { osexistpath.MustFile(runpath.PARENT.Path()) })
}

// TestMustRoot checks MustRoot passes on a DIR, and panics on a missing path and on a file
// TestMustRoot 检查 MustRoot 遇到目录时通过，遇到不存在的路径或文件就 panic
func TestMustRoot(t *testing.T) {
	require.NotPanics(t, func() { osexistpath.MustRoot(runpath.PARENT.Path()) })
	require.Panics(t, func() { osexistpath.MustRoot(runpath.PARENT.Join("abc.txt")) })
	require.Panics(t, func() { osexistpath.MustRoot(runtestpath.SrcPath(t)) })
}
