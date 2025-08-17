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

	return fmt.Sprintf(fmtMsg, fmtArgs...)
}

func (l *Log) logMsg(label string, msg ...any) bool {
	log.Print(label, l.expandMsg(msg...))

	return true
}

func (l *Log) logMsgf(label, msgFmt string, msgArgs ...any) bool {
	log.Print(label, l.expandMsgf(msgFmt, msgArgs...))

	return true
}

func (l *Log) noLog(_ ...any) bool {
	return false
}

func (l *Log) noLogErr(_ error, _ ...any) bool {
	return false
}

func (l *Log) noLogf(_ string, _ ...any) bool {
	return false
}

func (l *Log) noLogErrf(_ error, _ string, _ ...any) bool {
	return false
}

func (l *Log) logFatal(msg ...any) bool {
	return l.logMsg(logFatalLabel, msg...)
}

func (l *Log) logFatalErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logFatalLabel, msg...)
}

func (l *Log) logFatalf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logFatalLabel, msgFmt, msgArgs...)
}

func (l *Log) logFatalErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logFatalLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongFatal(msg ...any) bool {
	return l.logMsg(logLongFatalLabel, msg...)
}

func (l *Log) logLongFatalErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logLongFatalLabel, msg...)
}

func (l *Log) logLongFatalf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logLongFatalLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongFatalErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logLongFatalLabel, msgFmt, msgArgs...)
}

func (l *Log) logError(msg ...any) bool {
	return l.logMsg(logErrorLabel, msg...)
}

func (l *Log) logErrorErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logErrorLabel, msg...)
}

func (l *Log) logErrorf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logErrorLabel, msgFmt, msgArgs...)
}

func (l *Log) logErrorErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logErrorLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongError(msg ...any) bool {
	return l.logMsg(logLongErrorLabel, msg...)
}

func (l *Log) logLongErrorErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logLongErrorLabel, msg...)
}

func (l *Log) logLongErrorf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logLongErrorLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongErrorErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logLongErrorLabel, msgFmt, msgArgs...)
}

func (l *Log) logWarn(msg ...any) bool {
	return l.logMsg(logWarnLabel, msg...)
}

func (l *Log) logWarnErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logWarnLabel, msg...)
}

func (l *Log) logWarnf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logWarnLabel, msgFmt, msgArgs...)
}

func (l *Log) logWarnErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logWarnLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongWarn(msg ...any) bool {
	return l.logMsg(logLongWarnLabel, msg...)
}

func (l *Log) logLongWarnErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logLongWarnLabel, msg...)
}

func (l *Log) logLongWarnf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logLongWarnLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongWarnErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logLongWarnLabel, msgFmt, msgArgs...)
}

func (l *Log) logInfo(msg ...any) bool {
	return l.logMsg(logInfoLabel, msg...)
}

func (l *Log) logInfoErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logInfoLabel, msg...)
}

func (l *Log) logInfof(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logInfoLabel, msgFmt, msgArgs...)
}

func (l *Log) logInfoErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logInfoLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongInfo(msg ...any) bool {
	return l.logMsg(logLongInfoLabel, msg...)
}

func (l *Log) logLongInfoErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logLongInfoLabel, msg...)
}

func (l *Log) logLongInfof(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logLongInfoLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongInfoErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logLongInfoLabel, msgFmt, msgArgs...)
}

func (l *Log) logDebug(msg ...any) bool {
	return l.logMsg(logDebugLabel, msg...)
}

func (l *Log) logDebugErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logDebugLabel, msg...)
}

func (l *Log) logDebugf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logDebugLabel, msgFmt, msgArgs...)
}

func (l *Log) logDebugErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logDebugLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongDebug(msg ...any) bool {
	return l.logMsg(logLongDebugLabel, msg...)
}

func (l *Log) logLongDebugErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logLongDebugLabel, msg...)
}

func (l *Log) logLongDebugf(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logLongDebugLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongDebugErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logLongDebugLabel, msgFmt, msgArgs...)
}

func (l *Log) logTrace(msg ...any) bool {
	return l.logMsg(logTraceLabel, msg...)
}

func (l *Log) logTraceErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logTraceLabel, msg...)
}

func (l *Log) logTracef(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logTraceLabel, msgFmt, msgArgs...)
}

func (l *Log) logTraceErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logTraceLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongTrace(msg ...any) bool {
	return l.logMsg(logLongTraceLabel, msg...)
}

func (l *Log) logLongTraceErr(err error, msg ...any) bool {
	return err != nil && l.logMsg(logLongTraceLabel, msg...)
}

func (l *Log) logLongTracef(msgFmt string, msgArgs ...any) bool {
	return l.logMsgf(logLongTraceLabel, msgFmt, msgArgs...)
}

func (l *Log) logLongTraceErrf(err error, msgFmt string, msgArgs ...any) bool {
	return err != nil && l.logMsgf(logLongTraceLabel, msgFmt, msgArgs...)
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
