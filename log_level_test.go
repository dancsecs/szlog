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

package szlog_test

import (
	"strings"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

const (
	msgFmt = "formatted %s %s by: %s"
)

type tstLabelType bool

const (
	tstLongLabels  = tstLabelType(true)
	tstShortLabels = tstLabelType(false)
)

//nolint:funlen,cyclop // Ok.
func expectedLogSnippet(area string, longLabels tstLabelType) string {
	var snip []string

	switch area {
	case "Trace":
		var label string

		if longLabels {
			label = logTraceLongLabel
		} else {
			label = logTraceLabel
		}

		snip = []string{
			label + "plain trace message by: szlog.T",
			label + "plain trace message by: szlog.Trace",
			label + "formatted trace message by: szlog.Tf",
			label + "formatted trace message by: szlog.Tracef",
		}
	case "Debug":
		var label string

		if longLabels {
			label = logDebugLongLabel
		} else {
			label = logDebugLabel
		}

		snip = []string{
			label + "plain debug message by: szlog.D",
			label + "plain debug message by: szlog.Debug",
			label + "formatted debug message by: szlog.Df",
			label + "formatted debug message by: szlog.Debugf",
		}
	case "Info":
		var label string

		if longLabels {
			label = logInfoLongLabel
		} else {
			label = logInfoLabel
		}

		snip = []string{
			label + "plain info message by: szlog.I",
			label + "plain info message by: szlog.Info",
			label + "formatted info message by: szlog.If",
			label + "formatted info message by: szlog.Infof",
		}
	case "Warn":
		var label string

		if longLabels {
			label = logWarnLongLabel
		} else {
			label = logWarnLabel
		}

		snip = []string{
			label + "plain warn message by: szlog.W",
			label + "plain warn message by: szlog.Warn",
			label + "formatted warn message by: szlog.Wf",
			label + "formatted warn message by: szlog.Warnf",
		}
	case "Error":
		var label string

		if longLabels {
			label = logErrorLongLabel
		} else {
			label = logErrorLabel
		}

		snip = []string{
			label + "plain error message by: szlog.E",
			label + "plain error message by: szlog.Error",
			label + "formatted error message by: szlog.Ef",
			label + "formatted error message by: szlog.Errorf",
		}
	case "Fatal":
		var label string

		if longLabels {
			label = logFatalLongLabel
		} else {
			label = logFatalLabel
		}

		snip = []string{
			label + "plain fatal message by: szlog.F",
			label + "plain fatal message by: szlog.Fatal",
			label + "formatted fatal message by: szlog.Ff",
			label + "formatted fatal message by: szlog.Fatalf",
		}
	default:
		snip = []string{
			"unknown snippet area: " + area,
		}
	}

	return strings.Join(snip, "\n")
}

// logAll issues all log forms and levels.
func logAll() {
	szlog.T("plain ", "trace ", "message by: ", "szlog.T")
	szlog.Trace("plain ", "trace ", "message by: ", "szlog.Trace")
	szlog.Tf(msgFmt, "trace", "message", "szlog.Tf")
	szlog.Tracef(msgFmt, "trace", "message", "szlog.Tracef")

	szlog.D("plain ", "debug ", "message by: ", "szlog.D")
	szlog.Debug("plain ", "debug ", "message by: ", "szlog.Debug")
	szlog.Df(msgFmt, "debug", "message", "szlog.Df")
	szlog.Debugf(msgFmt, "debug", "message", "szlog.Debugf")

	szlog.I("plain ", "info ", "message by: ", "szlog.I")
	szlog.Info("plain ", "info ", "message by: ", "szlog.Info")
	szlog.If(msgFmt, "info", "message", "szlog.If")
	szlog.Infof(msgFmt, "info", "message", "szlog.Infof")

	szlog.W("plain ", "warn ", "message by: ", "szlog.W")
	szlog.Warn("plain ", "warn ", "message by: ", "szlog.Warn")
	szlog.Wf(msgFmt, "warn", "message", "szlog.Wf")
	szlog.Warnf(msgFmt, "warn", "message", "szlog.Warnf")

	szlog.E("plain ", "error ", "message by: ", "szlog.E")
	szlog.Error("plain ", "error ", "message by: ", "szlog.Error")
	szlog.Ef(msgFmt, "error", "message", "szlog.Ef")
	szlog.Errorf(msgFmt, "error", "message", "szlog.Errorf")

	szlog.F("plain ", "fatal ", "message by: ", "szlog.F")
	szlog.Fatal("plain ", "fatal ", "message by: ", "szlog.Fatal")
	szlog.Ff(msgFmt, "fatal", "message", "szlog.Ff")
	szlog.Fatalf(msgFmt, "fatal", "message", "szlog.Fatalf")
}

func TestSzLog_LogLevel_String(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Str((szlog.LevelNone - 1).String(), "UNKNOWN:-1")
	chk.Str((szlog.LevelNone).String(), "NONE")
	chk.Str((szlog.LevelFatal).String(), "FATAL")
	chk.Str((szlog.LevelError).String(), "ERROR")
	chk.Str((szlog.LevelWarn).String(), "WARN")
	chk.Str((szlog.LevelInfo).String(), "INFO")
	chk.Str((szlog.LevelDebug).String(), "DEBUG")
	chk.Str((szlog.LevelTrace).String(), "TRACE")
	chk.Str((szlog.LevelAll).String(), "ALL")
	chk.Str((szlog.LevelCustom).String(), "CUSTOM")
	chk.Str((szlog.LevelAll + 1).String(), "UNKNOWN:8")
}

func TestSzLog_All(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstShortLabels),
		expectedLogSnippet("Debug", tstShortLabels),
		expectedLogSnippet("Info", tstShortLabels),
		expectedLogSnippet("Warn", tstShortLabels),
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_AllWithLongLabels(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)
	chk.False(szlog.LongLabels())
	szlog.SetLongLabels(true)
	chk.True(szlog.LongLabels())

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstLongLabels),
		expectedLogSnippet("Debug", tstLongLabels),
		expectedLogSnippet("Info", tstLongLabels),
		expectedLogSnippet("Warn", tstLongLabels),
		expectedLogSnippet("Error", tstLongLabels),
		expectedLogSnippet("Fatal", tstLongLabels),
	)
}

func TestSzLog_LevelFatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelFatal)

	logAll()

	chk.Log(
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelNone(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	logAll()

	chk.Log()
}

func TestSzLog_LevelError(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelError)

	logAll()

	chk.Log(
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelWarn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelWarn)

	logAll()

	chk.Log(
		expectedLogSnippet("Warn", tstShortLabels),
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelInfo(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelInfo)

	logAll()

	chk.Log(
		expectedLogSnippet("Info", tstShortLabels),
		expectedLogSnippet("Warn", tstShortLabels),
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelDebug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelDebug)

	logAll()

	chk.Log(
		expectedLogSnippet("Debug", tstShortLabels),
		expectedLogSnippet("Info", tstShortLabels),
		expectedLogSnippet("Warn", tstShortLabels),
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelTrace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelTrace)

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstShortLabels),
		expectedLogSnippet("Debug", tstShortLabels),
		expectedLogSnippet("Info", tstShortLabels),
		expectedLogSnippet("Warn", tstShortLabels),
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_None(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelNone)

	logAll()

	chk.Log()
}

func TestSzLog_LevelCustom_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelTrace)

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelDebug)

	logAll()

	chk.Log(
		expectedLogSnippet("Debug", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelInfo)

	logAll()

	chk.Log(
		expectedLogSnippet("Info", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelWarn)

	logAll()

	chk.Log(
		expectedLogSnippet("Warn", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelError)

	logAll()

	chk.Log(
		expectedLogSnippet("Error", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelFatal)

	logAll()

	chk.Log(
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_All(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelAll)

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstShortLabels),
		expectedLogSnippet("Debug", tstShortLabels),
		expectedLogSnippet("Info", tstShortLabels),
		expectedLogSnippet("Warn", tstShortLabels),
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestSzLog_LevelCustom_Fatal_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelFatal, szlog.LevelTrace)

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}
