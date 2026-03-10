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
