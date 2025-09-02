/*
   Szerszam logging library: szlog.
   Copyright (C) 2024-2025  Leslie Dancsecs

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
	"testing"

	"github.com/dancsecs/sztest"
)

func TestSzLog_Builtin_ExpandMsg(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tLog := New()

	nonDeferredFunc := func() string {
		return "ab"
	}

	deferredFunc := func() Defer {
		return "ab"
	}

	// Needed to get past compiler blocking displaying the address of a func.
	printNon := func(a any) string {
		return fmt.Sprint(a)
	}

	chk.Str(tLog.expandMsg(), "")
	chk.Str(tLog.expandMsg("abc"), "abc")
	chk.Str(tLog.expandMsg("abc", "def"), "abcdef")
	chk.Str(
		tLog.expandMsg(nonDeferredFunc, "cd"),
		printNon(nonDeferredFunc)+"cd",
	)
	chk.Str(tLog.expandMsg(deferredFunc, "cd"), "abcd")

	// Local formatting

	chk.Str(tLog.expandMsg(1234), "1234")
	chk.NoErr(tLog.SetLanguage("en"))
	chk.Str(tLog.expandMsg(1234), "1,234")
	chk.NoErr(tLog.SetLanguage(""))
	chk.Str(tLog.expandMsg(1234), "1234")
}

func TestSzLog_Builtin_ExpandMsgf(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tLog := New()

	nonDeferredFunc := func() string {
		return "ab"
	}

	deferredFunc := func() Defer {
		return "ab"
	}

	// Needed to get past compiler blocking displaying the address of a func.
	printNon := func(a any) string {
		return fmt.Sprint(a)
	}

	chk.Str(tLog.expandMsgf(""), "")
	chk.Str(tLog.expandMsgf("%v", "abc"), "abc")
	chk.Str(tLog.expandMsgf("%v%v", "ab", "cd"), "abcd")
	chk.Str(
		tLog.expandMsgf("%v%v", nonDeferredFunc, "cd"),
		printNon(nonDeferredFunc)+"cd",
	)
	chk.Str(tLog.expandMsgf("%v%v", deferredFunc, "cd"), "abcd")

	// Local formatting

	chk.Str(tLog.expandMsgf("%d", 1234), "1234")
	chk.NoErr(tLog.SetLanguage("en"))
	chk.Str(tLog.expandMsgf("%d", 1234), "1,234")
	chk.NoErr(tLog.SetLanguage(""))
	chk.Str(tLog.expandMsgf("%d", 1234), "1234")
}

func TestSzLog_Builtin_LogMsg(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logMsg(labelDebug, "ab", "cd"))
	chk.True(tLog.logMsgf(labelFatal, "%v%v", "ab", "cd"))

	chk.Log(
		"D:abcd",
		"F:abcd",
	)
}

func TestSzLog_Builtin_NoLog(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.False(tLog.noLog("U:", "ab", "cd"))
	chk.False(tLog.noLogf("F:", "%v%v", "ab", "cd"))

	chk.Log()
}

func TestSzLog_Builtin_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logFatal("ab", "cd"))
	chk.True(tLog.logFatalf("fmt:%v%v", "ab", "cd"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logFatal("ab", "cd"))
	chk.True(tLog.logFatalf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"F:abcd",
		"F:fmt:abcd",
		"FATAL:abcd",
		"FATAL:fmt:abcd",
	)
}

func TestSzLog_Builtin_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logError("ab", "cd"))
	chk.True(tLog.logErrorf("fmt:%v%v", "ab", "cd"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logError("ab", "cd"))
	chk.True(tLog.logErrorf("fmt:%v%v", "ab", "cd"))
	chk.Log(
		"E:abcd",
		"E:fmt:abcd",
		"ERROR:abcd",
		"ERROR:fmt:abcd",
	)
}

func TestSzLog_Builtin_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logWarn("ab", "cd"))
	chk.True(tLog.logWarnf("fmt:%v%v", "ab", "cd"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logWarn("ab", "cd"))
	chk.True(tLog.logWarnf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"W:abcd",
		"W:fmt:abcd",
		"WARN:abcd",
		"WARN:fmt:abcd",
	)
}

func TestSzLog_Builtin_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logInfo("ab", "cd"))
	chk.True(tLog.logInfof("fmt:%v%v", "ab", "cd"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logInfo("ab", "cd"))
	chk.True(tLog.logInfof("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"I:abcd",
		"I:fmt:abcd",
		"INFO:abcd",
		"INFO:fmt:abcd",
	)
}

func TestSzLog_Builtin_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logDebug("ab", "cd"))
	chk.True(tLog.logDebugf("fmt:%v%v", "ab", "cd"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logDebug("ab", "cd"))
	chk.True(tLog.logDebugf("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"D:abcd",
		"D:fmt:abcd",
		"DEBUG:abcd",
		"DEBUG:fmt:abcd",
	)
}

func TestSzLog_Builtin_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New()

	chk.True(tLog.logTrace("ab", "cd"))
	chk.True(tLog.logTracef("fmt:%v%v", "ab", "cd"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logTrace("ab", "cd"))
	chk.True(tLog.logTracef("fmt:%v%v", "ab", "cd"))

	chk.Log(
		"T:abcd",
		"T:fmt:abcd",
		"TRACE:abcd",
		"TRACE:fmt:abcd",
	)
}

func TestSzLog_Builtin_ValidateLogLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	const tstArea = "testingArea"

	tLog := New()

	chk.Int(int(tLog.validateLogLevel(tstArea, LevelNone-1)), int(LevelNone))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelNone)), int(LevelNone))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelFatal)), int(LevelFatal))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelError)), int(LevelError))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelWarn)), int(LevelWarn))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelInfo)), int(LevelInfo))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelDebug)), int(LevelDebug))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelTrace)), int(LevelTrace))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelAll)), int(LevelAll))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelCustom)), int(LevelCustom))
	chk.Int(int(tLog.validateLogLevel(tstArea, LevelAll+1)), int(LevelAll))

	chk.Log(
		"W:attempt to access out of bounds log level: -1 from: testingArea",
		"W:attempt to access out of bounds log level: 8 from: testingArea",
	)
}
