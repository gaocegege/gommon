package log

import (
	"fmt"
	"testing"

	"github.com/dyweb/gommon/util/testutil"
	"github.com/stretchr/testify/assert"
)

var lg = NewPackageLogger()

func foo() *Logger {
	funcLog := NewFunctionLogger(lg)
	return funcLog
}

func fooUseCopy() *Logger {
	return lg.Copy()
}

type Foo struct {
	log *Logger
}

func (f *Foo) GetLogger() *Logger {
	return f.log
}

func (f *Foo) SetLogger(logger *Logger) {
	f.log = logger
}

func (f *Foo) LoggerIdentity(justCallMe func() Identity) Identity {
	return justCallMe()
}

func (f *Foo) method() *Logger {
	mlog := NewMethodLogger(f.log)
	return mlog
}

func (f *Foo) methodUseCopy() *Logger {
	return f.log.Copy()
}

func (f *Foo) methodOrphanCopy() *Logger {
	return lg.Copy()
}

var dummyFoo = &Foo{} // used for get struct logger identity

func TestNewPackageLogger(t *testing.T) {
	id := lg.id
	assert.Equal(t, PackageLogger, id.Type)
	assert.Equal(t, "pkg", id.Type.String())
	assert.Equal(t, "init", id.Function)
	assert.Equal(t, testutil.GOPATH()+"/src/github.com/dyweb/gommon/log/logger_identity_test.go:11",
		fmt.Sprintf("%s:%d", id.File, id.Line))
}

func TestNewFunctionLogger(t *testing.T) {
	assert.Equal(t, FunctionLogger, foo().id.Type)
	assert.Equal(t, FunctionLogger, fooUseCopy().id.Type)
}

func TestNewStructLogger(t *testing.T) {
	slog := NewStructLogger(lg, dummyFoo)
	id := slog.id
	assert.Equal(t, StructLogger, id.Type)
	assert.Equal(t, "struct", id.Type.String())
	assert.Equal(t, "Foo", id.Struct)
	assert.Equal(t, MagicStructLoggerFunctionName, id.Function)
	assert.Equal(t, testutil.GOPATH()+"/src/github.com/dyweb/gommon/log/logger_identity_test.go:35",
		fmt.Sprintf("%s:%d", id.File, id.Line))
}

func TestNewMethodLogger(t *testing.T) {
	slog := NewStructLogger(lg, dummyFoo)
	dummyFoo.log = slog
	mlog := dummyFoo.method()
	id := mlog.id
	assert.Equal(t, MethodLogger, id.Type)
	assert.Equal(t, "method", id.Type.String())
	assert.Equal(t, "Foo", id.Struct)
	assert.Equal(t, "method", id.Function)
	assert.Equal(t, testutil.GOPATH()+"/src/github.com/dyweb/gommon/log/logger_identity_test.go:39",
		fmt.Sprintf("%s:%d", id.File, id.Line))

	assert.Equal(t, MethodLogger, dummyFoo.methodUseCopy().id.Type)

	orphanLog := dummyFoo.methodOrphanCopy()
	assert.Equal(t, MethodLogger, orphanLog.id.Type)
	assert.Equal(t, "Foo", orphanLog.id.Struct)
	assert.Equal(t, "methodOrphanCopy", orphanLog.id.Function)
}
