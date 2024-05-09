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

import "strings"

// LogFunctionUnformatted defines the signature of an unformatted log function.
type LogFunctionUnformatted func(msg ...any) bool

// LogFunctionFormatted defines the signature of a formatted log function.
type LogFunctionFormatted func(msgFmt string, msgArgs ...any) bool

// LogLevel represents the current minium level of message to log.
type LogLevel int

// Def represents the return value of a function provided to a log
// function that will only be executed if the message is logged.
type Def string

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
)

// Usage suitable strings for verbose argument absorption.
const (
	VerboseFlag     = "-v[v...], --v[v...]"
	VerboseFlagDesc = "Increase the logging level for each v provided."
)

// Log represents a szlog logging object.
type Log struct {
	envOverride []string
	level       LogLevel
	longLabels  bool

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

	F, Fatal   LogFunctionUnformatted
	Ff, Fatalf LogFunctionFormatted
	E, Error   LogFunctionUnformatted
	Ef, Errorf LogFunctionFormatted
	W, Warn    LogFunctionUnformatted
	Wf, Warnf  LogFunctionFormatted
	I, Info    LogFunctionUnformatted
	If, Infof  LogFunctionFormatted
	D, Debug   LogFunctionUnformatted
	Df, Debugf LogFunctionFormatted
	T, Trace   LogFunctionUnformatted
	Tf, Tracef LogFunctionFormatted
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

// LongLabels enables/disables the use of longer labels in log output.
func (l *Log) LongLabels() bool {
	return l.longLabels
}

// SetLongLabels enables/disables the use of longer labels in log output.
func (l *Log) SetLongLabels(enable bool) bool {
	orig := l.longLabels
	l.longLabels = enable
	l.SetLevel(l.level)

	return orig
}

// LevelDisabled return true if the  level is disabled.
//
//nolint:cyclop // Ok.
func (l *Log) LevelDisabled(level LogLevel) bool {
	result := true

	switch validateLogLevel("LevelDisabled", level) {
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
	}

	return result
}

// DisableLevel blocks all logging at the specified level.
func (l *Log) DisableLevel(level LogLevel) {
	switch validateLogLevel("DisableLevel", level) {
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
	}

	l.SetLevel(l.level)
}

// Level return the current logging level.
func (l *Log) Level() LogLevel {
	return l.level
}

func selectLog(
	disabled bool,
	useLong bool,
	shortLog, longLog LogFunctionUnformatted,
) LogFunctionUnformatted {
	if !disabled {
		if useLong {
			return longLog
		}

		return shortLog
	}

	return noLog
}

func selectLogf(
	disabled bool,
	useLong bool,
	shortLogf, longLogf LogFunctionFormatted,
) LogFunctionFormatted {
	if !disabled {
		if useLong {
			return longLogf
		}

		return shortLogf
	}

	return noLogf
}

// SetLevel sets the logging level.
func (l *Log) SetLevel(newLogLevel LogLevel) LogLevel {
	oldLogLevel := l.level
	l.level = validateLogLevel("SetLevel", newLogLevel)

	disable := l.level < LevelFatal || l.disableLevelFatal
	l.F = selectLog(disable, l.longLabels, logFatal, logLongFatal)
	l.Fatal = l.F
	l.Ff = selectLogf(disable, l.longLabels, logFatalf, logLongFatalf)
	l.Fatalf = l.Ff

	disable = l.level < LevelError || l.disableLevelError
	l.E = selectLog(disable, l.longLabels, logError, logLongError)
	l.Error = l.E
	l.Ef = selectLogf(disable, l.longLabels, logErrorf, logLongErrorf)
	l.Errorf = l.Ef

	disable = l.level < LevelWarn || l.disableLevelWarn
	l.W = selectLog(disable, l.longLabels, logWarn, logLongWarn)
	l.Warn = l.W
	l.Wf = selectLogf(disable, l.longLabels, logWarnf, logLongWarnf)
	l.Warnf = l.Wf

	disable = l.level < LevelInfo || l.disableLevelInfo
	l.I = selectLog(disable, l.longLabels, logInfo, logLongInfo)
	l.Info = l.I
	l.If = selectLogf(disable, l.longLabels, logInfof, logLongInfof)
	l.Infof = l.If

	disable = l.level < LevelDebug || l.disableLevelDebug
	l.D = selectLog(disable, l.longLabels, logDebug, logLongDebug)
	l.Debug = l.D
	l.Df = selectLogf(disable, l.longLabels, logDebugf, logLongDebugf)
	l.Debugf = l.Df

	disable = l.level < LevelTrace || l.disableLevelTrace
	l.T = selectLog(disable, l.longLabels, logTrace, logLongTrace)
	l.Trace = l.T
	l.Tf = selectLogf(disable, l.longLabels, logTracef, logLongTracef)
	l.Tracef = l.Tf

	return oldLogLevel
}

// IncLevel permits all logging at the specified level.
func (l *Log) IncLevel() LogLevel {
	lastLevel := l.level

	l.SetLevel(validateLogLevel("IncLevel", l.level+1))

	return lastLevel
}

// DecLevel permits all logging at the specified level.
func (l *Log) DecLevel() LogLevel {
	lastLevel := l.level

	l.SetLevel(validateLogLevel("DecLevel", l.level-1))

	return lastLevel
}

// Reset returns all log setting to default startup conditions.
func (l *Log) Reset() {
	l.disableLevelFatal = getEnvSetting(envLogLevelFatal)
	l.disableLevelError = getEnvSetting(envLogLevelError)
	l.disableLevelWarn = getEnvSetting(envLogLevelWarn)
	l.disableLevelInfo = getEnvSetting(envLogLevelInfo)
	l.disableLevelDebug = getEnvSetting(envLogLevelDebug)
	l.disableLevelTrace = getEnvSetting(envLogLevelTrace)

	l.SetLevel(getEnvLevel(envLogLevel, LevelError))
}

// VerboseAbsorbArgs scans an argument list for verbose flags increasing
// the log level for each verbose flag encountered.  These flags are removed
// and a cleaned up arg list is returned.  Verbose flags can be a single (or
// multiple letter 'v's with the corresponding number of log level increments
// made.
func (l *Log) VerboseAbsorbArgs(argsIn []string) []string {
	argsOut := make([]string, 0, len(argsIn))

	for _, rArg := range argsIn {
		if strings.HasPrefix(rArg, "-v") || strings.HasPrefix(rArg, "--v") {
			keepArg := false
			cleanArg := strings.TrimLeft(rArg, "-")

			for i, mi := 0, len(cleanArg); i < mi && !keepArg; i++ {
				keepArg = cleanArg[i] != 'v'
			}

			if keepArg {
				argsOut = append(argsOut, rArg)
			} else {
				for range len(cleanArg) {
					l.IncLevel()
				}
			}
		}
	}

	return argsOut
}
