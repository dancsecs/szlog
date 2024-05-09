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
	"testing"

	"github.com/dancsecs/sztest"
)

func TestLOgBuiltin_ExpandMsg(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Str(expandMsg(), "")
	chk.Str(expandMsg("abc"), "abc")
	chk.Str(expandMsg("abc", "def"), "abcdef")
	chk.Str(expandMsg(func() Def { return "abc" }, "def"), "abcdef")
}

func TestLogBuiltin_ExpandMsgf(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Str(expandMsgf(""), "")
	chk.Str(expandMsgf("%v", "abc"), "abc")
	chk.Str(expandMsgf("%v%v", "ab", "cd"), "abcd")
	chk.Str(expandMsgf("%v%v", func() Def { return "ab" }, "cd"), "abcd")
}

func TestLogBuiltin_LogMsg(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logMsg("U:", "ab", "cd"))
	chk.True(logMsgf("F:", "%v%v", "ab", "cd"))

	chk.Log(
		"U:abcd",
		"F:abcd",
	)
}

func TestLogBuiltin_NoLog(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.False(noLog("U:", "ab", "cd"))
	chk.False(noLogf("F:", "%v%v", "ab", "cd"))

	chk.Log()
}

func TestLogBuiltin_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logFatal("ab", "cd"))
	chk.True(logFatalf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongFatal("ab", "cd"))
	chk.True(logLongFatalf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"F:abcd",
		"F:fmt:abcd",
		"FATAL:abcd",
		"FATAL:fmt:abcd",
	)
}

func TestLogBuiltin_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logError("ab", "cd"))
	chk.True(logErrorf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongError("ab", "cd"))
	chk.True(logLongErrorf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"E:abcd",
		"E:fmt:abcd",
		"ERROR:abcd",
		"ERROR:fmt:abcd",
	)
}

func TestLogBuiltin_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logWarn("ab", "cd"))
	chk.True(logWarnf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongWarn("ab", "cd"))
	chk.True(logLongWarnf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"W:abcd",
		"W:fmt:abcd",
		"WARN:abcd",
		"WARN:fmt:abcd",
	)
}

func TestLogBuiltin_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logInfo("ab", "cd"))
	chk.True(logInfof("fmt:%v%v", "ab", "cd"))

	chk.True(logLongInfo("ab", "cd"))
	chk.True(logLongInfof("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"I:abcd",
		"I:fmt:abcd",
		"INFO:abcd",
		"INFO:fmt:abcd",
	)
}

func TestLogBuiltin_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logDebug("ab", "cd"))
	chk.True(logDebugf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongDebug("ab", "cd"))
	chk.True(logLongDebugf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"D:abcd",
		"D:fmt:abcd",
		"DEBUG:abcd",
		"DEBUG:fmt:abcd",
	)
}

func TestLogBuiltin_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logTrace("ab", "cd"))
	chk.True(logTracef("fmt:%v%v", "ab", "cd"))

	chk.True(logLongTrace("ab", "cd"))
	chk.True(logLongTracef("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"T:abcd",
		"T:fmt:abcd",
		"TRACE:abcd",
		"TRACE:fmt:abcd",
	)
}

func TestLogBuiltin_ValidateLogLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	const tstArea = "testingArea"

	chk.Int(int(validateLogLevel(tstArea, LevelNone-1)), int(LevelNone))
	chk.Int(int(validateLogLevel(tstArea, LevelNone)), int(LevelNone))
	chk.Int(int(validateLogLevel(tstArea, LevelFatal)), int(LevelFatal))
	chk.Int(int(validateLogLevel(tstArea, LevelError)), int(LevelError))
	chk.Int(int(validateLogLevel(tstArea, LevelWarn)), int(LevelWarn))
	chk.Int(int(validateLogLevel(tstArea, LevelInfo)), int(LevelInfo))
	chk.Int(int(validateLogLevel(tstArea, LevelDebug)), int(LevelDebug))
	chk.Int(int(validateLogLevel(tstArea, LevelTrace)), int(LevelTrace))
	chk.Int(int(validateLogLevel(tstArea, LevelAll)), int(LevelAll))
	chk.Int(int(validateLogLevel(tstArea, LevelAll+1)), int(LevelAll))

	chk.Log(
		"W:attempt to access out of bounds log level: -1 from: testingArea",
		"W:attempt to access out of bounds log level: 8 from: testingArea",
	)
}
