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

	tstLog := New(nil)

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

	chk.Str(tstLog.expandMsg(), "")
	chk.Str(tstLog.expandMsg("abc"), "abc")
	chk.Str(tstLog.expandMsg("abc", "def"), "abcdef")
	chk.Str(
		tstLog.expandMsg(nonDeferredFunc, "cd"),
		printNon(nonDeferredFunc)+"cd",
	)
	chk.Str(tstLog.expandMsg(deferredFunc, "cd"), "abcd")
}

func TestSzLog_Builtin_ExpandMsgf(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tstLog := New(nil)

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

	chk.Str(tstLog.expandMsgf(""), "")
	chk.Str(tstLog.expandMsgf("%v", "abc"), "abc")
	chk.Str(tstLog.expandMsgf("%v%v", "ab", "cd"), "abcd")
	chk.Str(
		tstLog.expandMsgf("%v%v", nonDeferredFunc, "cd"),
		printNon(nonDeferredFunc)+"cd",
	)
	chk.Str(tstLog.expandMsgf("%v%v", deferredFunc, "cd"), "abcd")
}

func TestSzLog_Builtin_LogMsg(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logMsg("U:", "ab", "cd"))
	chk.True(tstLog.logMsgf("F:", "%v%v", "ab", "cd"))

	chk.Log(
		"U:abcd",
		"F:abcd",
	)
}

func TestSzLog_Builtin_NoLog(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)
	chk.False(tstLog.noLog("U:", "ab", "cd"))
	chk.False(tstLog.noLogf("F:", "%v%v", "ab", "cd"))

	chk.False(tstLog.noLogErr(errNil, "will not be logged"))
	chk.False(tstLog.noLogErrf(errNil, "will not be %s", "logged"))

	chk.Log()
}

func TestSzLog_Builtin_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logFatal("ab", "cd"))
	chk.True(tstLog.logFatalf("fmt:%v%v", "ab", "cd"))

	chk.True(tstLog.logLongFatal("ab", "cd"))
	chk.True(tstLog.logLongFatalf("fmt:%v%v", "ab", "cd"))

	chk.False(tstLog.logFatalErr(errNil, "will not be logged"))
	chk.False(tstLog.logFatalErrf(errNil, "will not be %s", "logged"))

	chk.False(tstLog.logLongFatalErr(errNil, "will not be logged"))
	chk.False(tstLog.logLongFatalErrf(errNil, "will not be %s", "logged"))

	chk.True(tstLog.logFatalErr(errTst, "error ", "not ", "nil"))
	chk.True(
		tstLog.logFatalErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"),
	)

	chk.True(tstLog.logLongFatalErr(errTst, "error ", "not ", "nil"))
	chk.True(
		tstLog.logLongFatalErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"),
	)

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

func TestSzLog_Builtin_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logError("ab", "cd"))
	chk.True(tstLog.logErrorf("fmt:%v%v", "ab", "cd"))

	chk.True(tstLog.logLongError("ab", "cd"))
	chk.True(tstLog.logLongErrorf("fmt:%v%v", "ab", "cd"))

	chk.False(tstLog.logErrorErr(errNil, "will not be logged"))
	chk.False(tstLog.logErrorErrf(errNil, "will not be %s", "logged"))

	chk.False(tstLog.logLongErrorErr(errNil, "will not be logged"))
	chk.False(tstLog.logLongErrorErrf(errNil, "will not be %s", "logged"))

	chk.True(tstLog.logErrorErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logErrorErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(tstLog.logLongErrorErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logLongErrorErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

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

func TestSzLog_Builtin_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logWarn("ab", "cd"))
	chk.True(tstLog.logWarnf("fmt:%v%v", "ab", "cd"))

	chk.True(tstLog.logLongWarn("ab", "cd"))
	chk.True(tstLog.logLongWarnf("fmt:%v%v", "ab", "cd"))

	chk.False(tstLog.logWarnErr(errNil, "will not be logged"))
	chk.False(tstLog.logWarnErrf(errNil, "will not be %s", "logged"))

	chk.False(tstLog.logLongWarnErr(errNil, "will not be logged"))
	chk.False(tstLog.logLongWarnErrf(errNil, "will not be %s", "logged"))

	chk.True(tstLog.logWarnErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logWarnErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(tstLog.logLongWarnErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logLongWarnErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

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

func TestSzLog_Builtin_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logInfo("ab", "cd"))
	chk.True(tstLog.logInfof("fmt:%v%v", "ab", "cd"))

	chk.True(tstLog.logLongInfo("ab", "cd"))
	chk.True(tstLog.logLongInfof("fmt:%v%v", "ab", "cd"))

	chk.False(tstLog.logInfoErr(errNil, "will not be logged"))
	chk.False(tstLog.logInfoErrf(errNil, "will not be %s", "logged"))

	chk.False(tstLog.logLongInfoErr(errNil, "will not be logged"))
	chk.False(tstLog.logLongInfoErrf(errNil, "will not be %s", "logged"))

	chk.True(tstLog.logInfoErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logInfoErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(tstLog.logLongInfoErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logLongInfoErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

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

func TestSzLog_Builtin_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logDebug("ab", "cd"))
	chk.True(tstLog.logDebugf("fmt:%v%v", "ab", "cd"))

	chk.True(tstLog.logLongDebug("ab", "cd"))
	chk.True(tstLog.logLongDebugf("fmt:%v%v", "ab", "cd"))

	chk.False(tstLog.logDebugErr(errNil, "will not be logged"))
	chk.False(tstLog.logDebugErrf(errNil, "will not be %s", "logged"))

	chk.False(tstLog.logLongDebugErr(errNil, "will not be logged"))
	chk.False(tstLog.logLongDebugErrf(errNil, "will not be %s", "logged"))

	chk.True(tstLog.logDebugErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logDebugErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(tstLog.logLongDebugErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logLongDebugErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

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

func TestSzLog_Builtin_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New(nil)

	chk.True(tstLog.logTrace("ab", "cd"))
	chk.True(tstLog.logTracef("fmt:%v%v", "ab", "cd"))

	chk.True(tstLog.logLongTrace("ab", "cd"))
	chk.True(tstLog.logLongTracef("fmt:%v%v", "ab", "cd"))

	chk.False(tstLog.logTraceErr(errNil, "will not be logged"))
	chk.False(tstLog.logTraceErrf(errNil, "will not be %s", "logged"))

	chk.False(tstLog.logLongTraceErr(errNil, "will not be logged"))
	chk.False(tstLog.logLongTraceErrf(errNil, "will not be %s", "logged"))

	chk.True(tstLog.logTraceErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logTraceErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

	chk.True(tstLog.logLongTraceErr(errTst, "error ", "not ", "nil"))
	chk.True(tstLog.logLongTraceErrf(errTst, "fmt:%s %s %s", "error", "not", "nil"))

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

func TestSzLog_Builtin_ValidateLogLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	const tstArea = "testingArea"

	tstLog := New(nil)

	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelNone-1)), int(LevelNone))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelNone)), int(LevelNone))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelFatal)), int(LevelFatal))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelError)), int(LevelError))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelWarn)), int(LevelWarn))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelInfo)), int(LevelInfo))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelDebug)), int(LevelDebug))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelTrace)), int(LevelTrace))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelAll)), int(LevelAll))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelCustom)), int(LevelCustom))
	chk.Int(int(tstLog.validateLogLevel(tstArea, LevelAll+1)), int(LevelAll))

	chk.Log(
		"W:attempt to access out of bounds log level: -1 from: testingArea",
		"W:attempt to access out of bounds log level: 8 from: testingArea",
	)
}
