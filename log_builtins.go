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

func (l *Log) expandMsg(rawMessages ...any) string {
	messages := make([]any, len(rawMessages))

	for i, arg := range rawMessages {
		f, ok := arg.(func() Defer)
		if ok {
			arg = f() // Use result of delayed parameter.
		}

		messages[i] = arg
	}

	if l.printer != nil {
		return l.printer.Sprint(messages...)
	}

	return fmt.Sprint(messages...)
}

func (l *Log) expandMsgf(fmtMsg string, rawFmtArgs ...any) string {
	fmtArgs := make([]any, len(rawFmtArgs))

	for i, arg := range rawFmtArgs {
		f, ok := arg.(func() Defer)
		if ok {
			arg = f() // Use result of delayed parameter.
		}

		fmtArgs[i] = arg
	}

	if l.printer != nil {
		return l.printer.Sprintf(fmtMsg, fmtArgs...)
	}

	return fmt.Sprintf(fmtMsg, fmtArgs...)
}

func (l *Log) logMsg(idx labelIdx, msg ...any) bool {
	var msgLabel string

	if l.longLabels {
		msgLabel = longLabel[idx]
	} else {
		msgLabel = shortLabel[idx]
	}

	log.Print(msgLabel, l.expandMsg(msg...))

	return true
}

func (l *Log) logMsgf(idx labelIdx, msgFmt string, msgArgs ...any) bool {
	var msgLabel string

	if l.longLabels {
		msgLabel = longLabel[idx]
	} else {
		msgLabel = shortLabel[idx]
	}

	log.Print(msgLabel, l.expandMsgf(msgFmt, msgArgs...))

	return true
}

func (l *Log) noLog(_ ...any) bool {
	return false
}

func (l *Log) noLogf(_ string, _ ...any) bool {
	return false
}

func (l *Log) logFatal(msg ...any) bool {
	return l.logMsg(labelFatal, msg...)
}

func (l *Log) logFatalf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(labelFatal, msgFmt, msgArgs...)
}

func (l *Log) logError(msg ...any) bool {
	return l.logMsg(labelError, msg...)
}

func (l *Log) logErrorf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(labelError, msgFmt, msgArgs...)
}

func (l *Log) logWarn(msg ...any) bool {
	return l.logMsg(labelWarn, msg...)
}

func (l *Log) logWarnf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(labelWarn, msgFmt, msgArgs...)
}

func (l *Log) logInfo(msg ...any) bool {
	return l.logMsg(labelInfo, msg...)
}

func (l *Log) logInfof(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(labelInfo, msgFmt, msgArgs...)
}

func (l *Log) logDebug(msg ...any) bool {
	return l.logMsg(labelDebug, msg...)
}

func (l *Log) logDebugf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(labelDebug, msgFmt, msgArgs...)
}

func (l *Log) logTrace(msg ...any) bool {
	return l.logMsg(labelTrace, msg...)
}

func (l *Log) logTracef(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(labelTrace, msgFmt, msgArgs...)
}

func (l *Log) validateLogLevel(area string, rawLevel LogLevel) LogLevel {
	var (
		level      LogLevel
		rangeError bool
	)

	switch {
	case rawLevel == LevelCustom:
		level = LevelCustom
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
		l.logWarn(
			"attempt to access out of bounds log level: ",
			int(rawLevel), // Convert to int to block LogLevel stringer.
			" from: ",
			area,
		)
	}

	return level
}
