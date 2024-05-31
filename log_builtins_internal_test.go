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

func TestLOgBuiltin_ExpandMsg(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

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

	chk.Str(expandMsg(), "")
	chk.Str(expandMsg("abc"), "abc")
	chk.Str(expandMsg("abc", "def"), "abcdef")
	chk.Str(
		expandMsg(nonDeferredFunc, "cd"),
		printNon(nonDeferredFunc)+"cd",
	)
	chk.Str(expandMsg(deferredFunc, "cd"), "abcd")
}

func TestLogBuiltin_ExpandMsgf(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

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

	chk.Str(expandMsgf(""), "")
	chk.Str(expandMsgf("%v", "abc"), "abc")
	chk.Str(expandMsgf("%v%v", "ab", "cd"), "abcd")
	chk.Str(
		expandMsgf("%v%v", nonDeferredFunc, "cd"),
		printNon(nonDeferredFunc)+"cd",
	)
	chk.Str(expandMsgf("%v%v", deferredFunc, "cd"), "abcd")
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

	chk.False(noLogErr(errNil, "will not be logged"))
	chk.False(noLogErrf(errNil, "will not be %s", "logged"))

	chk.Log()
}

func TestLogBuiltin_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logFatal("ab", "cd"))
	chk.True(logFatalf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongFatal("ab", "cd"))
	chk.True(logLongFatalf("fmt:%v%v", "ab", "cd"))

	chk.False(logFatalErr(errNil, "will not be logged"))
	chk.False(logFatalErrf(errNil, "will not be %s", "logged"))

	chk.False(logLongFatalErr(errNil, "will not be logged"))
	chk.False(logLongFatalErrf(errNil, "will not be %s", "logged"))

	chk.True(logFatalErr(errTst, "error ", "not ", "nil"))
	chk.True(logFatalErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(logLongFatalErr(errTst, "error ", "not ", "nil"))
	chk.True(logLongFatalErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"F:abcd",
		"F:fmt:abcd",
		"FATAL:abcd",
		"FATAL:fmt:abcd",
		"F:error not nil",
		"F:fmt:error not nil",
		"FATAL:error not nil",
		"FATAL:fmt:error not nil",
	)
}

func TestLogBuiltin_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logError("ab", "cd"))
	chk.True(logErrorf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongError("ab", "cd"))
	chk.True(logLongErrorf("fmt:%v%v", "ab", "cd"))

	chk.False(logErrorErr(errNil, "will not be logged"))
	chk.False(logErrorErrf(errNil, "will not be %s", "logged"))

	chk.False(logLongErrorErr(errNil, "will not be logged"))
	chk.False(logLongErrorErrf(errNil, "will not be %s", "logged"))

	chk.True(logErrorErr(errTst, "error ", "not ", "nil"))
	chk.True(logErrorErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(logLongErrorErr(errTst, "error ", "not ", "nil"))
	chk.True(logLongErrorErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"E:abcd",
		"E:fmt:abcd",
		"ERROR:abcd",
		"ERROR:fmt:abcd",
		"E:error not nil",
		"E:fmt:error not nil",
		"ERROR:error not nil",
		"ERROR:fmt:error not nil",
	)
}

func TestLogBuiltin_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logWarn("ab", "cd"))
	chk.True(logWarnf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongWarn("ab", "cd"))
	chk.True(logLongWarnf("fmt:%v%v", "ab", "cd"))

	chk.False(logWarnErr(errNil, "will not be logged"))
	chk.False(logWarnErrf(errNil, "will not be %s", "logged"))

	chk.False(logLongWarnErr(errNil, "will not be logged"))
	chk.False(logLongWarnErrf(errNil, "will not be %s", "logged"))

	chk.True(logWarnErr(errTst, "error ", "not ", "nil"))
	chk.True(logWarnErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(logLongWarnErr(errTst, "error ", "not ", "nil"))
	chk.True(logLongWarnErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"W:abcd",
		"W:fmt:abcd",
		"WARN:abcd",
		"WARN:fmt:abcd",
		"W:error not nil",
		"W:fmt:error not nil",
		"WARN:error not nil",
		"WARN:fmt:error not nil",
	)
}

func TestLogBuiltin_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logInfo("ab", "cd"))
	chk.True(logInfof("fmt:%v%v", "ab", "cd"))

	chk.True(logLongInfo("ab", "cd"))
	chk.True(logLongInfof("fmt:%v%v", "ab", "cd"))

	chk.False(logInfoErr(errNil, "will not be logged"))
	chk.False(logInfoErrf(errNil, "will not be %s", "logged"))

	chk.False(logLongInfoErr(errNil, "will not be logged"))
	chk.False(logLongInfoErrf(errNil, "will not be %s", "logged"))

	chk.True(logInfoErr(errTst, "error ", "not ", "nil"))
	chk.True(logInfoErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(logLongInfoErr(errTst, "error ", "not ", "nil"))
	chk.True(logLongInfoErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"I:abcd",
		"I:fmt:abcd",
		"INFO:abcd",
		"INFO:fmt:abcd",
		"I:error not nil",
		"I:fmt:error not nil",
		"INFO:error not nil",
		"INFO:fmt:error not nil",
	)
}

func TestLogBuiltin_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logDebug("ab", "cd"))
	chk.True(logDebugf("fmt:%v%v", "ab", "cd"))

	chk.True(logLongDebug("ab", "cd"))
	chk.True(logLongDebugf("fmt:%v%v", "ab", "cd"))

	chk.False(logDebugErr(errNil, "will not be logged"))
	chk.False(logDebugErrf(errNil, "will not be %s", "logged"))

	chk.False(logLongDebugErr(errNil, "will not be logged"))
	chk.False(logLongDebugErrf(errNil, "will not be %s", "logged"))

	chk.True(logDebugErr(errTst, "error ", "not ", "nil"))
	chk.True(logDebugErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(logLongDebugErr(errTst, "error ", "not ", "nil"))
	chk.True(logLongDebugErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"D:abcd",
		"D:fmt:abcd",
		"DEBUG:abcd",
		"DEBUG:fmt:abcd",
		"D:error not nil",
		"D:fmt:error not nil",
		"DEBUG:error not nil",
		"DEBUG:fmt:error not nil",
	)
}

func TestLogBuiltin_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.True(logTrace("ab", "cd"))
	chk.True(logTracef("fmt:%v%v", "ab", "cd"))

	chk.True(logLongTrace("ab", "cd"))
	chk.True(logLongTracef("fmt:%v%v", "ab", "cd"))

	chk.False(logTraceErr(errNil, "will not be logged"))
	chk.False(logTraceErrf(errNil, "will not be %s", "logged"))

	chk.False(logLongTraceErr(errNil, "will not be logged"))
	chk.False(logLongTraceErrf(errNil, "will not be %s", "logged"))

	chk.True(logTraceErr(errTst, "error ", "not ", "nil"))
	chk.True(logTraceErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(logLongTraceErr(errTst, "error ", "not ", "nil"))
	chk.True(logLongTraceErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.Log(
		"T:abcd",
		"T:fmt:abcd",
		"TRACE:abcd",
		"TRACE:fmt:abcd",
		"T:error not nil",
		"T:fmt:error not nil",
		"TRACE:error not nil",
		"TRACE:fmt:error not nil",
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
	chk.Int(int(validateLogLevel(tstArea, LevelCustom)), int(LevelCustom))
	chk.Int(int(validateLogLevel(tstArea, LevelAll+1)), int(LevelAll))

	chk.Log(
		"W:attempt to access out of bounds log level: -1 from: testingArea",
		"W:attempt to access out of bounds log level: 8 from: testingArea",
	)
}
