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

//nolint:dupl // Ok.
package szlog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dancsecs/sztest"
)

var (
	errNil = error(nil)
	errTst = errors.New("this error will not be displayed")
)

func TestSzLog_Builtin_ExpandMsg(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tLog := New(nil)

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
	tLog.SetLocal("en")
	chk.Str(tLog.expandMsg(1234), "1,234")
	tLog.SetLocal("")
	chk.Str(tLog.expandMsg(1234), "1234")
}

func TestSzLog_Builtin_ExpandMsgf(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tLog := New(nil)

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
	tLog.SetLocal("en")
	chk.Str(tLog.expandMsgf("%d", 1234), "1,234")
	tLog.SetLocal("")
	chk.Str(tLog.expandMsgf("%d", 1234), "1234")
}

func TestSzLog_Builtin_LogMsg(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

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

	tLog := New(nil)

	chk.False(tLog.noLog("U:", "ab", "cd"))
	chk.False(tLog.noLogf("F:", "%v%v", "ab", "cd"))

	chk.False(tLog.noLogErr(errNil, "will not be logged"))
	chk.False(tLog.noLogErrf(errNil, "will not be %s", "logged"))

	chk.Log()
}

func TestSzLog_Builtin_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

	chk.True(tLog.logFatal("ab", "cd"))
	chk.True(tLog.logFatalf("fmt:%v%v", "ab", "cd"))

	chk.False(tLog.logFatalErr(errNil, "will not be logged"))
	chk.False(tLog.logFatalErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logFatalErr(errTst, "error ", "not ", "nil"))
	chk.True(
		tLog.logFatalErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"),
	)

	tLog.SetLongLabels(true)

	chk.True(tLog.logFatal("ab", "cd"))
	chk.True(tLog.logFatalf("fmt:%v%v", "ab", "cd"))
	chk.False(tLog.logFatalErr(errNil, "will not be logged"))
	chk.False(tLog.logFatalErrf(errNil, "will not be %s", "logged"))
	chk.True(tLog.logFatalErr(errTst, "error ", "not ", "nil"))
	chk.True(
		tLog.logFatalErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"),
	)

	chk.Log(
		"F:abcd",
		"F:fmt:abcd",
		"F:error not nil",
		"F:fmt:error not nil",
		"FATAL:abcd",
		"FATAL:fmt:abcd",
		"FATAL:error not nil",
		"FATAL:fmt:error not nil",
	)
}

func TestSzLog_Builtin_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

	chk.True(tLog.logError("ab", "cd"))
	chk.True(tLog.logErrorf("fmt:%v%v", "ab", "cd"))

	chk.False(tLog.logErrorErr(errNil, "will not be logged"))
	chk.False(tLog.logErrorErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logErrorErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logErrorErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logError("ab", "cd"))
	chk.True(tLog.logErrorf("fmt:%v%v", "ab", "cd"))
	chk.False(tLog.logErrorErr(errNil, "will not be logged"))
	chk.False(tLog.logErrorErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logErrorErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logErrorErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"E:abcd",
		"E:fmt:abcd",
		"E:error not nil",
		"E:fmt:error not nil",
		"ERROR:abcd",
		"ERROR:fmt:abcd",
		"ERROR:error not nil",
		"ERROR:fmt:error not nil",
	)
}

func TestSzLog_Builtin_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

	chk.True(tLog.logWarn("ab", "cd"))
	chk.True(tLog.logWarnf("fmt:%v%v", "ab", "cd"))

	chk.False(tLog.logWarnErr(errNil, "will not be logged"))
	chk.False(tLog.logWarnErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logWarnErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logWarnErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	tLog.SetLongLabels(true)

	chk.False(tLog.logWarnErr(errNil, "will not be logged"))
	chk.False(tLog.logWarnErrf(errNil, "will not be %s", "logged"))
	chk.True(tLog.logWarn("ab", "cd"))
	chk.True(tLog.logWarnf("fmt:%v%v", "ab", "cd"))

	chk.True(tLog.logWarnErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logWarnErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"W:abcd",
		"W:fmt:abcd",
		"W:error not nil",
		"W:fmt:error not nil",
		"WARN:abcd",
		"WARN:fmt:abcd",
		"WARN:error not nil",
		"WARN:fmt:error not nil",
	)
}

func TestSzLog_Builtin_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

	chk.True(tLog.logInfo("ab", "cd"))
	chk.True(tLog.logInfof("fmt:%v%v", "ab", "cd"))

	chk.False(tLog.logInfoErr(errNil, "will not be logged"))
	chk.False(tLog.logInfoErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logInfoErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logInfoErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	tLog.SetLongLabels(true)

	chk.False(tLog.logInfoErr(errNil, "will not be logged"))
	chk.False(tLog.logInfoErrf(errNil, "will not be %s", "logged"))
	chk.True(tLog.logInfo("ab", "cd"))
	chk.True(tLog.logInfof("fmt:%v%v", "ab", "cd"))
	chk.True(tLog.logInfoErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logInfoErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"I:abcd",
		"I:fmt:abcd",
		"I:error not nil",
		"I:fmt:error not nil",
		"INFO:abcd",
		"INFO:fmt:abcd",
		"INFO:error not nil",
		"INFO:fmt:error not nil",
	)
}

func TestSzLog_Builtin_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

	chk.True(tLog.logDebug("ab", "cd"))
	chk.True(tLog.logDebugf("fmt:%v%v", "ab", "cd"))

	chk.False(tLog.logDebugErr(errNil, "will not be logged"))
	chk.False(tLog.logDebugErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logDebugErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logDebugErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logDebug("ab", "cd"))
	chk.True(tLog.logDebugf("fmt:%v%v", "ab", "cd"))
	chk.False(tLog.logDebugErr(errNil, "will not be logged"))
	chk.False(tLog.logDebugErrf(errNil, "will not be %s", "logged"))
	chk.True(tLog.logDebugErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logDebugErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"D:abcd",
		"D:fmt:abcd",
		"D:error not nil",
		"D:fmt:error not nil",
		"DEBUG:abcd",
		"DEBUG:fmt:abcd",
		"DEBUG:error not nil",
		"DEBUG:fmt:error not nil",
	)
}

func TestSzLog_Builtin_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tLog := New(nil)

	chk.True(tLog.logTrace("ab", "cd"))
	chk.True(tLog.logTracef("fmt:%v%v", "ab", "cd"))

	chk.False(tLog.logTraceErr(errNil, "will not be logged"))
	chk.False(tLog.logTraceErrf(errNil, "will not be %s", "logged"))

	chk.True(tLog.logTraceErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logTraceErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	tLog.SetLongLabels(true)

	chk.True(tLog.logTrace("ab", "cd"))
	chk.True(tLog.logTracef("fmt:%v%v", "ab", "cd"))
	chk.False(tLog.logTraceErr(errNil, "will not be logged"))
	chk.False(tLog.logTraceErrf(errNil, "will not be %s", "logged"))
	chk.True(tLog.logTraceErr(errTst, "error ", "not ", "nil"))
	chk.True(tLog.logTraceErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"T:abcd",
		"T:fmt:abcd",
		"T:error not nil",
		"T:fmt:error not nil",
		"TRACE:abcd",
		"TRACE:fmt:abcd",
		"TRACE:error not nil",
		"TRACE:fmt:error not nil",
	)
}

func TestSzLog_Builtin_ValidateLogLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	const tstArea = "testingArea"

	tLog := New(nil)

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
