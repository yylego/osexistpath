// This file exercises the generated Must wrapping too, so it states the constraint the generated file states
// With -tags=sure_generate the generated types vanish, and this file steps aside with them, which keeps TestGen able to run and rewrite a broken artifact
//
// 本文件也测了生成出来的 Must 包装，所以写上与生成文件相同的约束
// 带 -tags=sure_generate 时生成的类型就没了，本文件跟着一起让位，TestGen 才能照样跑起来重写破损产物

//go:build !sure_generate

package ospath_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/done"
	"github.com/yylego/osexistpath/ospath"
	"github.com/yylego/runpath"
	"github.com/yylego/runpath/runtestpath"
)

func Test_existNamespace_Path(t *testing.T) {
	for _, path := range []string{
		runpath.CurrentPath(),
		runpath.PARENT.Path(),
	} {
		t.Log(path)
		done.VBE(ospath.EXIST.IsPath(path)).TRUE()
	}
}

func Test_existNamespace_File(t *testing.T) {
	path := runtestpath.SrcPath(t)
	t.Log(path)
	done.VBE(ospath.EXIST.IsFile(path)).TRUE()
}

func Test_existNamespace_Root(t *testing.T) {
	path := runpath.PARENT.Path()
	t.Log(path)
	done.VBE(ospath.EXIST.IsRoot(path)).TRUE()
}

func Test_existNamespace_MustPath(t *testing.T) {
	ospath.EXIST.MustPath(runpath.CurrentPath())
	ospath.EXIST.MustPath(runpath.PARENT.Path())
}

func Test_existNamespace_MustFile(t *testing.T) {
	ospath.EXIST.MustFile(runtestpath.SrcPath(t))
}

func Test_existNamespace_MustRoot(t *testing.T) {
	ospath.EXIST.MustRoot(runpath.PARENT.Path())
}

// TestExistNamespaceMustPATH checks the Must wrapping gives back the path it got
// One statement thus yields a path that is sure to exist
//
// TestExistNamespaceMustPATH 检查 Must 包装会把收到的路径原样给回
// 于是一条语句就能拿到一个必然存在的路径
func TestExistNamespaceMustPATH(t *testing.T) {
	currentPath := runpath.CurrentPath()
	require.Equal(t, currentPath, ospath.EXIST.Must().PATH(currentPath))

	parentPath := runpath.PARENT.Path()
	require.Equal(t, parentPath, ospath.EXIST.Must().PATH(parentPath))
}

// TestExistNamespaceMustFILE checks the same with a path that names a file
// This test file is itself such a path, so it needs nothing set up beforehand
//
// TestExistNamespaceMustFILE 用指向文件的路径检查同一件事
// 本测试文件自身就是这么一条路径，无需事先准备任何东西
func TestExistNamespaceMustFILE(t *testing.T) {
	sourcePath := runpath.CurrentPath()
	require.Equal(t, sourcePath, ospath.EXIST.Must().FILE(sourcePath))
}

// TestExistNamespaceMustROOT checks the same with a path that names a ROOT
// TestExistNamespaceMustROOT 用指向目录的路径检查同一件事
func TestExistNamespaceMustROOT(t *testing.T) {
	parentPath := runpath.PARENT.Path()
	require.Equal(t, parentPath, ospath.EXIST.Must().ROOT(parentPath))
}
