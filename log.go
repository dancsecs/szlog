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

// LogFunc defines the signature of an unformatted log function.
type LogFunc func(msg ...any) bool

// LogFuncf defines the signature of a formatted log function.
type LogFuncf func(msgFmt string, msgArgs ...any) bool

// LogErrFunc defines the signature of an unformatted error logging function.
type LogErrFunc func(err error, msg ...any) bool

// LogErrFuncf defines the signature of a formatted error logging function.
type LogErrFuncf func(err error, msgFmt string, msgArgs ...any) bool

// Log represents a szlog logging object.
type Log struct {
	envOverride           []string
	level                 LogLevel
	customLevelsPermitted int
	longLabels            bool

	disableLevelFatal bool
	disableLevelError bool
	disableLevelWarn  bool
	disableLevelInfo  bool
	disableLevelDebug bool
	disableLevelTrace bool

	LogFatal bool
	LogError bool
	LogWarn  bool
	LogInfo  bool
	LogDebug bool
	LogTrace bool

	F, Fatal         LogFunc
	FErr, FatalErr   LogErrFunc
	Ff, Fatalf       LogFuncf
	FErrf, FatalErrf LogErrFuncf
	E, Error         LogFunc
	EErr, ErrorErr   LogErrFunc
	Ef, Errorf       LogFuncf
	EErrf, ErrorErrf LogErrFuncf
	W, Warn          LogFunc
	WErr, WarnErr    LogErrFunc
	Wf, Warnf        LogFuncf
	WErrf, WarnErrf  LogErrFuncf
	I, Info          LogFunc
	IErr, InfoErr    LogErrFunc
	If, Infof        LogFuncf
	IErrf, InfoErrf  LogErrFuncf
	D, Debug         LogFunc
	DErr, DebugErr   LogErrFunc
	Df, Debugf       LogFuncf
	DErrf, DebugErrf LogErrFuncf
	T, Trace         LogFunc
	TErr, TraceErr   LogErrFunc
	Tf, Tracef       LogFuncf
	TErrf, TraceErrf LogErrFuncf
}

//nolint:goCheckNoGlobals // Default package logger.
var defaultLog = New(nil)

// New creates a new log object.
func New(envOverrides []string) *Log {
	l := new(Log)
	l.envOverride = envOverrides
	l.Reset()

	return l
}

// Default Returns the package's default logger.
func Default() *Log {
	return defaultLog
}

// SetDefault sets the package's default logger.
func SetDefault(newDefaultLog *Log) *Log {
	origLog := defaultLog
	defaultLog = newDefaultLog

	return origLog
}

// LevelDisabled return true if the  level is disabled.
//
//nolint:cyclop // Ok.
func (l *Log) LevelDisabled(level LogLevel) bool {
	result := true

	switch l.validateLogLevel("LevelDisabled", level) {
	case LevelNone:
		result = !(l.disableLevelFatal ||
			l.disableLevelError ||
			l.disableLevelWarn ||
			l.disableLevelInfo ||
			l.disableLevelDebug ||
			l.disableLevelTrace)
	case LevelFatal:
		result = l.disableLevelFatal
	case LevelError:
		result = l.disableLevelError
	case LevelWarn:
		result = l.disableLevelWarn
	case LevelInfo:
		result = l.disableLevelInfo
	case LevelDebug:
		result = l.disableLevelDebug
	case LevelTrace:
		result = l.disableLevelTrace
	case LevelAll:
		result = l.disableLevelFatal &&
			l.disableLevelError &&
			l.disableLevelWarn &&
			l.disableLevelInfo &&
			l.disableLevelDebug &&
			l.disableLevelTrace
	case LevelCustom:
	}

	return result
}

// DisableLevel blocks all logging at the specified level.
func (l *Log) DisableLevel(level LogLevel) {
	switch l.validateLogLevel("DisableLevel", level) {
	case LevelNone:
		l.disableLevelFatal = false
		l.disableLevelError = false
		l.disableLevelWarn = false
		l.disableLevelInfo = false
		l.disableLevelDebug = false
		l.disableLevelTrace = false
	case LevelFatal:
		l.disableLevelFatal = true
	case LevelError:
		l.disableLevelError = true
	case LevelWarn:
		l.disableLevelWarn = true
	case LevelInfo:
		l.disableLevelInfo = true
	case LevelDebug:
		l.disableLevelDebug = true
	case LevelTrace:
		l.disableLevelTrace = true
	case LevelAll:
		l.disableLevelFatal = true
		l.disableLevelError = true
		l.disableLevelWarn = true
		l.disableLevelInfo = true
		l.disableLevelDebug = true
		l.disableLevelTrace = true
	case LevelCustom:
	}

	l.SetLevel(l.level)
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

func (l *Log) selectLogErr(
	enabled bool,
	shortLog LogErrFunc,
) LogErrFunc {
	if enabled {
		return shortLog
	}

	return l.noLogErr
}

func (l *Log) selectLogErrf(
	enabled bool,
	shortLogf LogErrFuncf,
) LogErrFuncf {
	if enabled {
		return shortLogf
	}

	return l.noLogErrf
}

// SetCustomLevels permits the selective enabling of individual levels.
//
//nolint:cyclop // Ok.
func (l *Log) SetCustomLevels(levels ...LogLevel) LogLevel {
	permittedLevels := enableLevelNone

	for _, level := range levels {
		switch level {
		case LevelNone:
			permittedLevels = enableLevelNone
		case LevelFatal:
			permittedLevels |= enabledFatal
		case LevelError:
			permittedLevels |= enabledError
		case LevelWarn:
			permittedLevels |= enabledWarn
		case LevelInfo:
			permittedLevels |= enabledInfo
		case LevelDebug:
			permittedLevels |= enabledDebug
		case LevelTrace:
			permittedLevels |= enabledTrace
		case LevelAll:
			permittedLevels = enableLevelAll
		case LevelCustom:
		}
	}

	l.customLevelsPermitted = permittedLevels

	return l.SetLevel(LevelCustom)
}

// SetLevel sets the logging level.
//
//nolint:funlen,cyclop // OK.
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
		enable = l.customLevelsPermitted
	}

	l.LogFatal = enable&enabledFatal > 0 && !l.disableLevelFatal
	l.F = l.selectLog(l.LogFatal, l.logFatal)
	l.Fatal = l.F
	l.Ff = l.selectLogf(l.LogFatal, l.logFatalf)
	l.Fatalf = l.Ff
	l.FErr = l.selectLogErr(l.LogFatal, l.logFatalErr)
	l.FatalErr = l.FErr
	l.FErrf = l.selectLogErrf(l.LogFatal, l.logFatalErrf)
	l.FatalErrf = l.FErrf

	l.LogError = enable&enabledError > 0 && !l.disableLevelError
	l.E = l.selectLog(l.LogError, l.logError)
	l.Error = l.E
	l.Ef = l.selectLogf(l.LogError, l.logErrorf)
	l.Errorf = l.Ef
	l.EErr = l.selectLogErr(l.LogError, l.logErrorErr)
	l.ErrorErr = l.EErr
	l.EErrf = l.selectLogErrf(l.LogError, l.logErrorErrf)
	l.ErrorErrf = l.EErrf

	l.LogWarn = enable&enabledWarn > 0 && !l.disableLevelWarn
	l.W = l.selectLog(l.LogWarn, l.logWarn)
	l.Warn = l.W
	l.Wf = l.selectLogf(l.LogWarn, l.logWarnf)
	l.Warnf = l.Wf
	l.WErr = l.selectLogErr(l.LogWarn, l.logWarnErr)
	l.WarnErr = l.WErr
	l.WErrf = l.selectLogErrf(l.LogWarn, l.logWarnErrf)
	l.WarnErrf = l.WErrf

	l.LogInfo = enable&enabledInfo > 0 && !l.disableLevelInfo
	l.I = l.selectLog(l.LogInfo, l.logInfo)
	l.Info = l.I
	l.If = l.selectLogf(l.LogInfo, l.logInfof)
	l.Infof = l.If
	l.IErr = l.selectLogErr(l.LogInfo, l.logInfoErr)
	l.InfoErr = l.IErr
	l.IErrf = l.selectLogErrf(l.LogInfo, l.logInfoErrf)
	l.InfoErrf = l.IErrf

	l.LogDebug = enable&enabledDebug > 0 && !l.disableLevelDebug
	l.D = l.selectLog(l.LogDebug, l.logDebug)
	l.Debug = l.D
	l.Df = l.selectLogf(l.LogDebug, l.logDebugf)
	l.Debugf = l.Df
	l.DErr = l.selectLogErr(l.LogDebug, l.logDebugErr)
	l.DebugErr = l.DErr
	l.DErrf = l.selectLogErrf(l.LogDebug, l.logDebugErrf)
	l.DebugErrf = l.DErrf

	l.LogTrace = enable&enabledTrace > 0 && !l.disableLevelTrace
	l.T = l.selectLog(l.LogTrace, l.logTrace)
	l.Trace = l.T
	l.Tf = l.selectLogf(l.LogTrace, l.logTracef)
	l.Tracef = l.Tf
	l.TErr = l.selectLogErr(l.LogTrace, l.logTraceErr)
	l.TraceErr = l.TErr
	l.TErrf = l.selectLogErrf(l.LogTrace, l.logTraceErrf)
	l.TraceErrf = l.TErrf

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

// Reset returns all log setting to default startup conditions.
func (l *Log) Reset() {
	l.longLabels = l.getEnvSetting(envLogLongLabels)

	l.disableLevelFatal = l.getEnvSetting(envLogLevelFatal)
	l.disableLevelError = l.getEnvSetting(envLogLevelError)
	l.disableLevelWarn = l.getEnvSetting(envLogLevelWarn)
	l.disableLevelInfo = l.getEnvSetting(envLogLevelInfo)
	l.disableLevelDebug = l.getEnvSetting(envLogLevelDebug)
	l.disableLevelTrace = l.getEnvSetting(envLogLevelTrace)

	l.customLevelsPermitted = enableLevelError // Later: load from env

	l.SetLevel(l.getEnvLevel(envLogLevel, LevelError))
}
