package ospath

import (
	"testing"

	"github.com/yylego/runpath"
	"github.com/yylego/runpath/runtestpath"
	"github.com/yylego/sure"
	"github.com/yylego/sure/sure_cls_gen"
	"github.com/yylego/syntaxgo"
	"github.com/yylego/syntaxgo/syntaxgo_ast"
)

//go:generate go test -run ^TestGen$ -tags=sure_generate
func TestGen(t *testing.T) {
	options := sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path()).
		WithNewClassNameParts("_").
		WithNamingPatternType(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE).
		MoreErrorHandlingModes(sure.MUST, sure.SOFT, sure.OMIT)

	config := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: options,
		PackageName:     syntaxgo.CurrentPackageName(),
		ImportOptions:   syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:      runtestpath.SrcPath(t),
	}

	sure_cls_gen.GenerateClasses(config, existNamespace{})
}
