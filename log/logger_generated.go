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
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprint(args...))
		} else {
			l.h.HandleLogWithSource(TraceLevel, time.Now(), fmt.Sprint(args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(TraceLevel, time.Now(), fmt.Sprint(args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(TraceLevel, time.Now(), fmt.Sprint(args...), caller(), l.fields)
		}
	}
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	if l.level < TraceLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprintf(format, args...))
		} else {
			l.h.HandleLogWithSource(TraceLevel, time.Now(), fmt.Sprintf(format, args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(TraceLevel, time.Now(), fmt.Sprintf(format, args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(TraceLevel, time.Now(), fmt.Sprintf(format, args...), caller(), l.fields)
		}
	}
}

func (l *Logger) TraceF(msg string, fields Fields) {
	if l.level < TraceLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLogWithFields(TraceLevel, time.Now(), msg, fields)
		} else {
			l.h.HandleLogWithSourceFields(TraceLevel, time.Now(), msg, caller(), fields)
		}
	} else {
		if !l.source {
			l.h.HandleLogWithContextFields(TraceLevel, time.Now(), msg, l.fields, fields)
		} else {
			l.h.HandleLogWithSourceContextFields(TraceLevel, time.Now(), msg, caller(), l.fields, fields)
		}
	}
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level >= DebugLevel
}

func (l *Logger) Debug(args ...interface{}) {
	if l.level < DebugLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprint(args...))
		} else {
			l.h.HandleLogWithSource(DebugLevel, time.Now(), fmt.Sprint(args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(DebugLevel, time.Now(), fmt.Sprint(args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(DebugLevel, time.Now(), fmt.Sprint(args...), caller(), l.fields)
		}
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level < DebugLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprintf(format, args...))
		} else {
			l.h.HandleLogWithSource(DebugLevel, time.Now(), fmt.Sprintf(format, args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(DebugLevel, time.Now(), fmt.Sprintf(format, args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(DebugLevel, time.Now(), fmt.Sprintf(format, args...), caller(), l.fields)
		}
	}
}

func (l *Logger) DebugF(msg string, fields Fields) {
	if l.level < DebugLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLogWithFields(DebugLevel, time.Now(), msg, fields)
		} else {
			l.h.HandleLogWithSourceFields(DebugLevel, time.Now(), msg, caller(), fields)
		}
	} else {
		if !l.source {
			l.h.HandleLogWithContextFields(DebugLevel, time.Now(), msg, l.fields, fields)
		} else {
			l.h.HandleLogWithSourceContextFields(DebugLevel, time.Now(), msg, caller(), l.fields, fields)
		}
	}
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level >= InfoLevel
}

func (l *Logger) Info(args ...interface{}) {
	if l.level < InfoLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprint(args...))
		} else {
			l.h.HandleLogWithSource(InfoLevel, time.Now(), fmt.Sprint(args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(InfoLevel, time.Now(), fmt.Sprint(args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(InfoLevel, time.Now(), fmt.Sprint(args...), caller(), l.fields)
		}
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level < InfoLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprintf(format, args...))
		} else {
			l.h.HandleLogWithSource(InfoLevel, time.Now(), fmt.Sprintf(format, args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(InfoLevel, time.Now(), fmt.Sprintf(format, args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(InfoLevel, time.Now(), fmt.Sprintf(format, args...), caller(), l.fields)
		}
	}
}

func (l *Logger) InfoF(msg string, fields Fields) {
	if l.level < InfoLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLogWithFields(InfoLevel, time.Now(), msg, fields)
		} else {
			l.h.HandleLogWithSourceFields(InfoLevel, time.Now(), msg, caller(), fields)
		}
	} else {
		if !l.source {
			l.h.HandleLogWithContextFields(InfoLevel, time.Now(), msg, l.fields, fields)
		} else {
			l.h.HandleLogWithSourceContextFields(InfoLevel, time.Now(), msg, caller(), l.fields, fields)
		}
	}
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level >= WarnLevel
}

func (l *Logger) Warn(args ...interface{}) {
	if l.level < WarnLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprint(args...))
		} else {
			l.h.HandleLogWithSource(WarnLevel, time.Now(), fmt.Sprint(args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(WarnLevel, time.Now(), fmt.Sprint(args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(WarnLevel, time.Now(), fmt.Sprint(args...), caller(), l.fields)
		}
	}
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level < WarnLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprintf(format, args...))
		} else {
			l.h.HandleLogWithSource(WarnLevel, time.Now(), fmt.Sprintf(format, args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(WarnLevel, time.Now(), fmt.Sprintf(format, args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(WarnLevel, time.Now(), fmt.Sprintf(format, args...), caller(), l.fields)
		}
	}
}

func (l *Logger) WarnF(msg string, fields Fields) {
	if l.level < WarnLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLogWithFields(WarnLevel, time.Now(), msg, fields)
		} else {
			l.h.HandleLogWithSourceFields(WarnLevel, time.Now(), msg, caller(), fields)
		}
	} else {
		if !l.source {
			l.h.HandleLogWithContextFields(WarnLevel, time.Now(), msg, l.fields, fields)
		} else {
			l.h.HandleLogWithSourceContextFields(WarnLevel, time.Now(), msg, caller(), l.fields, fields)
		}
	}
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level >= ErrorLevel
}

func (l *Logger) Error(args ...interface{}) {
	if l.level < ErrorLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprint(args...))
		} else {
			l.h.HandleLogWithSource(ErrorLevel, time.Now(), fmt.Sprint(args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(ErrorLevel, time.Now(), fmt.Sprint(args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(ErrorLevel, time.Now(), fmt.Sprint(args...), caller(), l.fields)
		}
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level < ErrorLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprintf(format, args...))
		} else {
			l.h.HandleLogWithSource(ErrorLevel, time.Now(), fmt.Sprintf(format, args...), caller())
		}
	} else {
		if !l.source {
			l.h.HandleLogWithFields(ErrorLevel, time.Now(), fmt.Sprintf(format, args...), l.fields)
		} else {
			l.h.HandleLogWithSourceFields(ErrorLevel, time.Now(), fmt.Sprintf(format, args...), caller(), l.fields)
		}
	}
}

func (l *Logger) ErrorF(msg string, fields Fields) {
	if l.level < ErrorLevel {
		return
	}
	if len(l.fields) == 0 {
		if !l.source {
			l.h.HandleLogWithFields(ErrorLevel, time.Now(), msg, fields)
		} else {
			l.h.HandleLogWithSourceFields(ErrorLevel, time.Now(), msg, caller(), fields)
		}
	} else {
		if !l.source {
			l.h.HandleLogWithContextFields(ErrorLevel, time.Now(), msg, l.fields, fields)
		} else {
			l.h.HandleLogWithSourceContextFields(ErrorLevel, time.Now(), msg, caller(), l.fields, fields)
		}
	}
}
