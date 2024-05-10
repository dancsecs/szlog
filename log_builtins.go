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
	"log"
)

const (
	logFatalLabel = "F:"
	logErrorLabel = "E:"
	logWarnLabel  = "W:"
	logInfoLabel  = "I:"
	logDebugLabel = "D:"
	logTraceLabel = "T:"

	logLongFatalLabel = "FATAL:"
	logLongErrorLabel = "ERROR:"
	logLongWarnLabel  = "WARN:"
	logLongInfoLabel  = "INFO:"
	logLongDebugLabel = "DEBUG:"
	logLongTraceLabel = "TRACE:"
)

func expandMsg(rawMessages ...any) string {
	messages := make([]any, len(rawMessages))

	for i, arg := range rawMessages {
		f, ok := arg.(func() Def)
		if ok {
			arg = f() // Use result of delayed parameter.
		}

		messages[i] = arg
	}

	return fmt.Sprint(messages...)
}

func expandMsgf(fmtMsg string, rawFmtArgs ...any) string {
	fmtArgs := make([]any, len(rawFmtArgs))

	for i, arg := range rawFmtArgs {
		f, ok := arg.(func() Def)
		if ok {
			arg = f() // Use result of delayed parameter.
		}

		fmtArgs[i] = arg
	}

	return fmt.Sprintf(fmtMsg, fmtArgs...)
}

func logMsg(label string, msg ...any) bool {
	log.Print(label, expandMsg(msg...))

	return true
}

func logMsgf(label, msgFmt string, msgArgs ...any) bool {
	log.Print(label, expandMsgf(msgFmt, msgArgs...))

	return true
}

func noLog(_ ...any) bool {
	return false
}

func noLogf(_ string, _ ...any) bool {
	return false
}

func logFatal(msg ...any) bool {
	return logMsg(logFatalLabel, msg...)
}

func logFatalf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logFatalLabel, msgFmt, msgArgs...)
}

func logLongFatal(msg ...any) bool {
	return logMsg(logLongFatalLabel, msg...)
}

func logLongFatalf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logLongFatalLabel, msgFmt, msgArgs...)
}

func logError(msg ...any) bool {
	return logMsg(logErrorLabel, msg...)
}

func logErrorf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logErrorLabel, msgFmt, msgArgs...)
}

func logLongError(msg ...any) bool {
	return logMsg(logLongErrorLabel, msg...)
}

func logLongErrorf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logLongErrorLabel, msgFmt, msgArgs...)
}

func logWarn(msg ...any) bool {
	return logMsg(logWarnLabel, msg...)
}

func logWarnf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logWarnLabel, msgFmt, msgArgs...)
}

func logLongWarn(msg ...any) bool {
	return logMsg(logLongWarnLabel, msg...)
}

func logLongWarnf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logLongWarnLabel, msgFmt, msgArgs...)
}

func logInfo(msg ...any) bool {
	return logMsg(logInfoLabel, msg...)
}

func logInfof(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logInfoLabel, msgFmt, msgArgs...)
}

func logLongInfo(msg ...any) bool {
	return logMsg(logLongInfoLabel, msg...)
}

func logLongInfof(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logLongInfoLabel, msgFmt, msgArgs...)
}

func logDebug(msg ...any) bool {
	return logMsg(logDebugLabel, msg...)
}

func logDebugf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logDebugLabel, msgFmt, msgArgs...)
}

func logLongDebug(msg ...any) bool {
	return logMsg(logLongDebugLabel, msg...)
}

func logLongDebugf(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logLongDebugLabel, msgFmt, msgArgs...)
}

func logTrace(msg ...any) bool {
	return logMsg(logTraceLabel, msg...)
}

func logTracef(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logTraceLabel, msgFmt, msgArgs...)
}

func logLongTrace(msg ...any) bool {
	return logMsg(logLongTraceLabel, msg...)
}

func logLongTracef(msgFmt string, msgArgs ...any) bool {
	return logMsgf(logLongTraceLabel, msgFmt, msgArgs...)
}

func validateLogLevel(area string, rawLevel LogLevel) LogLevel {
	var (
		level      LogLevel
		rangeError bool
	)

	switch {
	case rawLevel < LevelNone:
		level = LevelNone
		rangeError = true
	case rawLevel > LevelAll:
		level = LevelAll
		rangeError = true
	default:
		level = rawLevel
	}

	if rangeError {
		logWarn(
			"attempt to access out of bounds log level: ",
			int(rawLevel), // Convert to int to block LogLevel stringer.
			" from: ",
			area,
		)
	}

	return level
}
