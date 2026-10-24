// Code generated using sure/sure_cls_gen. DO NOT EDIT.
// This file was auto generated via github.com/yylego/sure

//go:build !sure_generate

// Generated from: sure.gen_test.go:28 -> ospath.TestGen
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package ospath

import "github.com/yylego/sure"

type existNamespace_Must struct{ T *existNamespace }

func (T *existNamespace) Must() *existNamespace_Must {
	return &existNamespace_Must{T: T}
}
func (T0 *existNamespace_Must) IsPath(path string) (res bool) {
	res, err1 := T0.T.IsPath(path)
	sure.Must(err1)
	return res
}
func (T0 *existNamespace_Must) IsFile(path string) (res bool) {
	res, err1 := T0.T.IsFile(path)
	sure.Must(err1)
	return res
}
func (T0 *existNamespace_Must) IsRoot(path string) (res bool) {
	res, err1 := T0.T.IsRoot(path)
	sure.Must(err1)
	return res
}
func (T0 *existNamespace_Must) PATH(path string) (res string) {
	res, err1 := T0.T.PATH(path)
	sure.Must(err1)
	return res
}
func (T0 *existNamespace_Must) FILE(path string) (res string) {
	res, err1 := T0.T.FILE(path)
	sure.Must(err1)
	return res
}
func (T0 *existNamespace_Must) ROOT(path string) (res string) {
	res, err1 := T0.T.ROOT(path)
	sure.Must(err1)
	return res
}
func (T0 *existNamespace_Must) MustPath(path string) {
	T0.T.MustPath(path)
}
func (T0 *existNamespace_Must) MustFile(path string) {
	T0.T.MustFile(path)
}
func (T0 *existNamespace_Must) MustRoot(path string) {
	T0.T.MustRoot(path)
}

type existNamespace_Soft struct{ T *existNamespace }

func (T *existNamespace) Soft() *existNamespace_Soft {
	return &existNamespace_Soft{T: T}
}
func (T0 *existNamespace_Soft) IsPath(path string) (res bool) {
	res, err1 := T0.T.IsPath(path)
	sure.Soft(err1)
	return res
}
func (T0 *existNamespace_Soft) IsFile(path string) (res bool) {
	res, err1 := T0.T.IsFile(path)
	sure.Soft(err1)
	return res
}
func (T0 *existNamespace_Soft) IsRoot(path string) (res bool) {
	res, err1 := T0.T.IsRoot(path)
	sure.Soft(err1)
	return res
}
func (T0 *existNamespace_Soft) PATH(path string) (res string) {
	res, err1 := T0.T.PATH(path)
	sure.Soft(err1)
	return res
}
func (T0 *existNamespace_Soft) FILE(path string) (res string) {
	res, err1 := T0.T.FILE(path)
	sure.Soft(err1)
	return res
}
func (T0 *existNamespace_Soft) ROOT(path string) (res string) {
	res, err1 := T0.T.ROOT(path)
	sure.Soft(err1)
	return res
}
func (T0 *existNamespace_Soft) MustPath(path string) {
	T0.T.MustPath(path)
}
func (T0 *existNamespace_Soft) MustFile(path string) {
	T0.T.MustFile(path)
}
func (T0 *existNamespace_Soft) MustRoot(path string) {
	T0.T.MustRoot(path)
}

type existNamespace_Omit struct{ T *existNamespace }

func (T *existNamespace) Omit() *existNamespace_Omit {
	return &existNamespace_Omit{T: T}
}
func (T0 *existNamespace_Omit) IsPath(path string) (res bool) {
	res, err1 := T0.T.IsPath(path)
	sure.Omit(err1)
	return res
}
func (T0 *existNamespace_Omit) IsFile(path string) (res bool) {
	res, err1 := T0.T.IsFile(path)
	sure.Omit(err1)
	return res
}
func (T0 *existNamespace_Omit) IsRoot(path string) (res bool) {
	res, err1 := T0.T.IsRoot(path)
	sure.Omit(err1)
	return res
}
func (T0 *existNamespace_Omit) PATH(path string) (res string) {
	res, err1 := T0.T.PATH(path)
	sure.Omit(err1)
	return res
}
func (T0 *existNamespace_Omit) FILE(path string) (res string) {
	res, err1 := T0.T.FILE(path)
	sure.Omit(err1)
	return res
}
func (T0 *existNamespace_Omit) ROOT(path string) (res string) {
	res, err1 := T0.T.ROOT(path)
	sure.Omit(err1)
	return res
}
func (T0 *existNamespace_Omit) MustPath(path string) {
	T0.T.MustPath(path)
}
func (T0 *existNamespace_Omit) MustFile(path string) {
	T0.T.MustFile(path)
}
func (T0 *existNamespace_Omit) MustRoot(path string) {
	T0.T.MustRoot(path)
}
