package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Logger is a concrete type instead of interface because most logic is in handler.
// There is NO lock when calling logging methods, handlers may have locks.
// Lock is used when updating logger attributes like Level.
//
// For Printf style logging (Levelf), Logger formats string using fmt.Sprintf before passing it to handlers.
//
// 	logger.Debugf("id is %d", id)
//
// For structural logging (LevelF), Logger passes fields to handlers without any processing.
//
//	logger.DebugF("hi", log.Fields{log.Str("foo", "bar")})
//
// If you want to mix two styles, call fmt.Sprintf before calling DebugF,
//
// 	logger.DebugF(fmt.Sprintf("id is %d", id), log.Fields{log.Str("foo", "bar")})
type Logger struct {
	// mu is a Mutex instead of RWMutex because it's only for avoid concurrent write,
	// for performance reason and the natural of logging, reading stale config is not a big problem,
	// so we don't check mutex on read operation (i.e. log message) and allow race condition
	mu     sync.Mutex
	h      Handler
	level  Level
	source bool
	skip   int
	// fields contains common context, i.e. the struct is created for a specific task and it has "taskId": 0ac-123
	fields Fields

	id *Identity // use nil so we can have logger without identity
}

// Copy create a new logger with different identity, the identity is based on where Copy is called
// Normally you should call Copy inside func or method on a package/strcut logger
func (l *Logger) Copy() *Logger {
	id := NewIdentityFromCaller(1)
	c := &Logger{
		id: &id,
	}
	return newLogger(l, c)
}

// AddField add field to current logger in place, it does NOT create a copy of logger
// Use Copy if you want a copy
// It does NOT check duplication
func (l *Logger) AddField(f Field) *Logger {
	l.mu.Lock()
	// TODO: check dup or not? or may it optional
	l.fields = append(l.fields, f)
	l.mu.Unlock()
	return l
}

// AddFields add fields to current logger in place, it does NOT create a copy of logger
// Use Copy if you want a copy
// It does NOT check duplication
func (l *Logger) AddFields(fields ...Field) *Logger {
	l.mu.Lock()
	// TODO: check dup or not? or may it optional
	l.fields = append(l.fields, fields...)
	l.mu.Unlock()
	return l
}

// Flush calls Flush of its handler
func (l *Logger) Flush() {
	l.h.Flush()
}

func (l *Logger) Level() Level {
	return l.level
}

func (l *Logger) SetLevel(level Level) *Logger {
	l.mu.Lock()
	l.level = level
	l.mu.Unlock()
	return l
}

func (l *Logger) SetHandler(h Handler) *Logger {
	l.mu.Lock()
	l.h = h
	l.mu.Unlock()
	return l
}

func (l *Logger) EnableSource() *Logger {
	l.mu.Lock()
	l.source = true
	l.mu.Unlock()
	return l
}

func (l *Logger) DisableSource() *Logger {
	l.mu.Lock()
	l.source = false
	l.mu.Unlock()
	return l
}

// SetCallerSkip is used for util function to log using its caller's location instead of its own
// Without extra skip, some common util function will keep logging same line and make the real
// source hard to track.
//
// func echo(w http.ResponseWriter, r *http.Request) {
//     if r.Query().Get("word") == "" {
//          writeError(w, errors.New("word is required")
//          return
//     }
//     w.Write([]byte(r.Query().Get("word")))
// }
//
// func writeError(w http.ResponseWriter, err error) {
//     l := pkgLogger.Copy().SetCallerSkip(1)
//     l.Error(err)
//     w.Write([]byte(err.String()))
// }
func (l *Logger) SetCallerSkip(skip int) *Logger {
	l.mu.Lock()
	// ignore invalid skip, most time it should just be one
	if skip > 0 && skip < 5 {
		l.skip = skip
	}
	l.mu.Unlock()
	return l
}

// ResetCallerSkip set skip to 0, the default value
func (l *Logger) ResetCallerSkip() *Logger {
	l.mu.Lock()
	l.skip = 0
	l.mu.Unlock()
	return l
}

// Identity returns the identity set when the logger is created.
// NOTE: caller can modify the identity because all fields are public, but they should NOT do this
func (l *Logger) Identity() Identity {
	if l.id == nil {
		return UnknownIdentity
	}
	return *l.id
}

// Panic calls panic after it writes and flushes the log
func (l *Logger) Panic(args ...interface{}) {
	s := fmt.Sprint(args...)
	if len(l.fields) == 0 {
		l.h.HandleLog(PanicLevel, time.Now(), s, caller(l.skip), nil, nil)
	} else {
		l.h.HandleLog(PanicLevel, time.Now(), s, caller(l.skip), l.fields, nil)
	}
	l.h.Flush()
	panic(s)
}

// Panicf duplicates instead of calling Panic to keep source line correct
func (l *Logger) Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	if len(l.fields) == 0 {
		l.h.HandleLog(PanicLevel, time.Now(), s, caller(l.skip), nil, nil)
	} else {
		l.h.HandleLog(PanicLevel, time.Now(), s, caller(l.skip), l.fields, nil)
	}
	l.h.Flush()
	panic(s)
}

// PanicF duplicates instead of calling Panic to keep source line correct
func (l *Logger) PanicF(msg string, fields ...Field) {
	if len(l.fields) == 0 {
		l.h.HandleLog(PanicLevel, time.Now(), msg, caller(l.skip), nil, fields)
	} else {
		l.h.HandleLog(PanicLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
	l.h.Flush()
	panic(msg)
}

// Fatal calls os.Exit(1) after it writes and flushes the log
func (l *Logger) Fatal(args ...interface{}) {
	s := fmt.Sprint(args...)
	if len(l.fields) == 0 {
		l.h.HandleLog(FatalLevel, time.Now(), s, caller(l.skip), nil, nil)
	} else {
		l.h.HandleLog(FatalLevel, time.Now(), s, caller(l.skip), l.fields, nil)
	}
	l.h.Flush()
	// TODO: allow user to register hook to do cleanup before exit directly
	os.Exit(1)
}

// Fatalf duplicates instead of calling Fatal to keep source line correct
func (l *Logger) Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	if len(l.fields) == 0 {
		l.h.HandleLog(FatalLevel, time.Now(), s, caller(l.skip), nil, nil)
	} else {
		l.h.HandleLog(FatalLevel, time.Now(), s, caller(l.skip), l.fields, nil)
	}
	l.h.Flush()
	os.Exit(1)
}

// FatalF duplicates instead of calling Fatal to keep source line correct
func (l *Logger) FatalF(msg string, fields ...Field) {
	if len(l.fields) == 0 {
		l.h.HandleLog(FatalLevel, time.Now(), msg, caller(l.skip), nil, fields)
	} else {
		l.h.HandleLog(FatalLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
	l.h.Flush()
	os.Exit(1)
}

// Noop is only for test escape analysis
func (l *Logger) NoopF(msg string, fields ...Field) {
	// noop
}
