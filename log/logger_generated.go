// Code generated by gommon from log/logger_generated.go.tmpl DO NOT EDIT.

package log

import (
	"fmt"
	"time"
)

func (l *Logger) IsTraceEnabled() bool {
	return l.level >= TraceLevel
}

func (l *Logger) Trace(args ...interface{}) {
	if l.level < TraceLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprint(args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprint(args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	if l.level < TraceLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprintf(format, args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprintf(format, args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) TraceF(msg string, fields ...Field) {
	if l.level < TraceLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(TraceLevel, time.Now(), msg, emptyCaller, l.fields, fields)
	} else {
		l.h.HandleLog(TraceLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level >= DebugLevel
}

func (l *Logger) Debug(args ...interface{}) {
	if l.level < DebugLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprint(args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprint(args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level < DebugLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprintf(format, args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprintf(format, args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) DebugF(msg string, fields ...Field) {
	if l.level < DebugLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(DebugLevel, time.Now(), msg, emptyCaller, l.fields, fields)
	} else {
		l.h.HandleLog(DebugLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level >= InfoLevel
}

func (l *Logger) Info(args ...interface{}) {
	if l.level < InfoLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprint(args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprint(args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level < InfoLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprintf(format, args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprintf(format, args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) InfoF(msg string, fields ...Field) {
	if l.level < InfoLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(InfoLevel, time.Now(), msg, emptyCaller, l.fields, fields)
	} else {
		l.h.HandleLog(InfoLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
}

func (l *Logger) IsPrintEnabled() bool {
	return l.level >= PrintLevel
}

func (l *Logger) Print(args ...interface{}) {
	if l.level < PrintLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(PrintLevel, time.Now(), fmt.Sprint(args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(PrintLevel, time.Now(), fmt.Sprint(args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) Printf(format string, args ...interface{}) {
	if l.level < PrintLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(PrintLevel, time.Now(), fmt.Sprintf(format, args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(PrintLevel, time.Now(), fmt.Sprintf(format, args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) PrintF(msg string, fields ...Field) {
	if l.level < PrintLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(PrintLevel, time.Now(), msg, emptyCaller, l.fields, fields)
	} else {
		l.h.HandleLog(PrintLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level >= WarnLevel
}

func (l *Logger) Warn(args ...interface{}) {
	if l.level < WarnLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprint(args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprint(args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level < WarnLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprintf(format, args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprintf(format, args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) WarnF(msg string, fields ...Field) {
	if l.level < WarnLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(WarnLevel, time.Now(), msg, emptyCaller, l.fields, fields)
	} else {
		l.h.HandleLog(WarnLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level >= ErrorLevel
}

func (l *Logger) Error(args ...interface{}) {
	if l.level < ErrorLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprint(args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprint(args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level < ErrorLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprintf(format, args...), emptyCaller, l.fields, nil)
	} else {
		l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprintf(format, args...), caller(l.skip), l.fields, nil)
	}
}

func (l *Logger) ErrorF(msg string, fields ...Field) {
	if l.level < ErrorLevel {
		return
	}
	if !l.source {
		l.h.HandleLog(ErrorLevel, time.Now(), msg, emptyCaller, l.fields, fields)
	} else {
		l.h.HandleLog(ErrorLevel, time.Now(), msg, caller(l.skip), l.fields, fields)
	}
}
