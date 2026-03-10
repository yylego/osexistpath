// Package ospath: Namespace-style path existence checking API
// Provides EXIST namespace with methods to check path, file, and DIR existence
// Uses consistent 4-letter naming (PATH, FILE, ROOT) to keep code neat
//
// ospath: 命名空间风格的路径存在性检查 API
// 提供 EXIST 命名空间，包含检查路径、文件和目录存在性的方法
// 使用一致的 4 字母命名 (PATH, FILE, ROOT) 保持代码整洁
package ospath

import (
	"github.com/yylego/osexistpath"
)

type existNamespace struct{}

// EXIST is the namespace instance to access path checking methods
// EXIST 是访问路径检查方法的命名空间实例
var EXIST = &existNamespace{}

// IsPath checks if something exists at the specified path
// IsPath 检查这个路径下是否有内容
func (*existNamespace) IsPath(path string) (bool, error) {
	return osexistpath.IsPath(path)
}

// IsFile checks if a file exists at the specified path
// IsFile 检查文件是否存在
func (*existNamespace) IsFile(path string) (bool, error) {
	return osexistpath.IsFile(path)
}

// IsRoot checks if a DIR exists at the specified path
// Uses "root" to represent DIR, keeping function names at 4 letters
// IsRoot 检查目录是否存在
// 使用 "root" 表示目录，保持函数名都是 4 个字母
func (*existNamespace) IsRoot(path string) (bool, error) {
	return osexistpath.IsRoot(path)
}

// PATH validates and returns the path if it exists
// PATH 验证路径存在性，存在则返回路径，否则返回错误
func (T *existNamespace) PATH(path string) (string, error) {
	return osexistpath.PATH(path)
}

// FILE validates and returns the path if the file exists
// FILE 验证文件存在性，存在则返回路径，否则返回错误
func (T *existNamespace) FILE(path string) (string, error) {
	return osexistpath.FILE(path)
}

// ROOT validates and returns the path if the DIR exists
// ROOT 验证目录存在性，存在则返回路径，否则返回错误
func (T *existNamespace) ROOT(path string) (string, error) {
	return osexistpath.ROOT(path)
}

// MustPath validates path existence, panics if not found
// MustPath 验证路径存在性，不存在则 panic
func (T *existNamespace) MustPath(path string) {
	osexistpath.MustPath(path)
}

// MustFile validates file existence, panics if not found
// MustFile 验证文件存在性，不存在则 panic
func (T *existNamespace) MustFile(path string) {
	osexistpath.MustFile(path)
}

// MustRoot validates DIR existence, panics if not found
// MustRoot 验证目录存在性，不存在则 panic
func (T *existNamespace) MustRoot(path string) {
	osexistpath.MustRoot(path)
}
