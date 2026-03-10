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
