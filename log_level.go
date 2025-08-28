/*
   Szerszam logging library: szlog.
   Copyright (C) 2024  Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package szlog

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// LogLevel represents the current minium level of message to log.
type LogLevel int

// Available logging levels.
const (
	LevelNone LogLevel = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
	LevelCustom = LogLevel(math.MaxInt)
)

const (
	enabledFatal = 1 << iota
	enabledError
	enabledWarn
	enabledInfo
	enabledDebug
	enabledTrace

	enableLevelNone  = 0
	enableLevelFatal = enableLevelNone | enabledFatal
	enableLevelError = enableLevelFatal | enabledError
	enableLevelWarn  = enableLevelError | enabledWarn
	enableLevelInfo  = enableLevelWarn | enabledInfo
	enableLevelDebug = enableLevelInfo | enabledDebug
	enableLevelTrace = enableLevelDebug | enabledTrace
	enableLevelAll   = enableLevelTrace
)

// String implements the Stringer interface.
//
//nolint:cyclop // Ok.
func (ll LogLevel) String() string {
	switch ll {
	case LevelNone:
		return "NONE"
	case LevelFatal:
		return "FATAL"
	case LevelError:
		return "ERROR"
	case LevelWarn:
		return "WARN"
	case LevelInfo:
		return "INFO"
	case LevelDebug:
		return "DEBUG"
	case LevelTrace:
		return "TRACE"
	case LevelAll:
		return "ALL"
	case LevelCustom:
		return "CUSTOM"
	default:
		return "UNKNOWN:" + strconv.FormatInt(int64(ll), 10)
	}
}

// LogLevelFromString converts a string to a LogLevel.
//
//nolint:cyclop   // Ok.
func LogLevelFromString(raw string) (LogLevel, error) {
	switch strings.ToUpper(raw) {
	case "NONE":
		return LevelNone, nil
	case "FATAL":
		return LevelFatal, nil
	case "ERROR":
		return LevelError, nil
	case "WARN":
		return LevelWarn, nil
	case "INFO":
		return LevelInfo, nil
	case "DEBUG":
		return LevelDebug, nil
	case "TRACE":
		return LevelTrace, nil
	case "ALL":
		return LevelAll, nil
	case "CUSTOM":
		return LevelCustom, nil
	default:
		return 0, fmt.Errorf("%w: '%s'", ErrUnknownLevel, raw)
	}
}

//nolint:cyclop //Ok.
func (l *Log) parseAndSetLevel(raw string, defLevel LogLevel) error {
	var (
		lvl LogLevel
		err error
	)

	clnRaw := strings.Trim(raw, "()")

	if clnRaw != raw { //nolint:nestif // Ok.
		customLevels := make([]LogLevel, 0)

		for _, lvlStr := range strings.Split(clnRaw, "|") {
			lvl, err = LogLevelFromString(strings.TrimSpace(lvlStr))
			if err == nil {
				if lvl != LevelCustom {
					customLevels = append(customLevels, lvl)
				} else {
					err = fmt.Errorf("%w: '%s'", ErrUnknownLevel, raw)
				}
			} else {
				break
			}
		}

		if err == nil && len(customLevels) > 0 {
			l.SetCustomLevels(customLevels...)
		}

		if err == nil {
			return nil
		}
	}

	lvl, err = LogLevelFromString(raw)
	if err == nil && lvl == LevelCustom {
		err = fmt.Errorf("%w: '%s'", ErrUnknownLevel, raw)
	}

	if err == nil {
		l.SetLevel(lvl)

		return nil
	}

	l.SetLevel(defLevel)

	return fmt.Errorf("%w: %w", ErrInvalidLogLevelParse, err)
}

// SetLevel sets the logging level.
//
//nolint:funlen //Ok.
func (l *Log) SetLevel(newLogLevel LogLevel) LogLevel {
	oldLogLevel := l.level
	l.level = l.validateLogLevel("SetLevel", newLogLevel)

	enable := 0

	switch l.level {
	case LevelNone:
		enable = enableLevelNone
	case LevelFatal:
		enable = enableLevelFatal
	case LevelError:
		enable = enableLevelError
	case LevelWarn:
		enable = enableLevelWarn
	case LevelInfo:
		enable = enableLevelInfo
	case LevelDebug:
		enable = enableLevelDebug
	case LevelTrace:
		enable = enableLevelTrace
	case LevelAll:
		enable = enableLevelAll
	case LevelCustom:
		enable = l.customLevelsEnabled
	}

	l.LogFatal = enable&enabledFatal > 0
	l.F = l.selectLog(l.LogFatal, l.logFatal)
	l.Fatal = l.F
	l.Ff = l.selectLogf(l.LogFatal, l.logFatalf)
	l.Fatalf = l.Ff

	l.LogError = enable&enabledError > 0
	l.E = l.selectLog(l.LogError, l.logError)
	l.Error = l.E
	l.Ef = l.selectLogf(l.LogError, l.logErrorf)
	l.Errorf = l.Ef

	l.LogWarn = enable&enabledWarn > 0
	l.W = l.selectLog(l.LogWarn, l.logWarn)
	l.Warn = l.W
	l.Wf = l.selectLogf(l.LogWarn, l.logWarnf)
	l.Warnf = l.Wf

	l.LogInfo = enable&enabledInfo > 0
	l.I = l.selectLog(l.LogInfo, l.logInfo)
	l.Info = l.I
	l.If = l.selectLogf(l.LogInfo, l.logInfof)
	l.Infof = l.If

	l.LogDebug = enable&enabledDebug > 0
	l.D = l.selectLog(l.LogDebug, l.logDebug)
	l.Debug = l.D
	l.Df = l.selectLogf(l.LogDebug, l.logDebugf)
	l.Debugf = l.Df

	l.LogTrace = enable&enabledTrace > 0
	l.T = l.selectLog(l.LogTrace, l.logTrace)
	l.Trace = l.T
	l.Tf = l.selectLogf(l.LogTrace, l.logTracef)
	l.Tracef = l.Tf

	return oldLogLevel
}

// IncLevel permits all logging at the specified level.
func (l *Log) IncLevel() LogLevel {
	lastLevel := l.level

	if lastLevel != LevelCustom {
		l.SetLevel(l.validateLogLevel("IncLevel", l.level+1))
	}

	return lastLevel
}

// DecLevel permits all logging at the specified level.
func (l *Log) DecLevel() LogLevel {
	lastLevel := l.level

	if lastLevel != LevelCustom {
		l.SetLevel(l.validateLogLevel("DecLevel", l.level-1))
	}

	return lastLevel
}

// Level return the current logging level.
func (l *Log) Level() LogLevel {
	return l.level
}

func (l *Log) selectLog(
	enabled bool,
	shortLog LogFunc,
) LogFunc {
	if enabled {
		return shortLog
	}

	return l.noLog
}

func (l *Log) selectLogf(
	enabled bool,
	shortLogf LogFuncf,
) LogFuncf {
	if enabled {
		return shortLogf
	}

	return l.noLogf
}

// SetCustomLevels permits the selective enabling of individual levels.
//
//nolint:cyclop // Ok.
func (l *Log) SetCustomLevels(levels ...LogLevel) LogLevel {
	enabledLevels := enableLevelNone

	for _, level := range levels {
		switch level {
		case LevelNone:
			enabledLevels = enableLevelNone
		case LevelFatal:
			enabledLevels |= enabledFatal
		case LevelError:
			enabledLevels |= enabledError
		case LevelWarn:
			enabledLevels |= enabledWarn
		case LevelInfo:
			enabledLevels |= enabledInfo
		case LevelDebug:
			enabledLevels |= enabledDebug
		case LevelTrace:
			enabledLevels |= enabledTrace
		case LevelAll:
			enabledLevels = enableLevelAll
		case LevelCustom:
		}
	}

	l.customLevelsEnabled = enabledLevels

	return l.SetLevel(LevelCustom)
}
