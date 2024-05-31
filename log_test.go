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
	"errors"
	"strings"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

const (
	msgFmt    = "formatted %s %s by: %s"
	msgFmtErr = "err formatted %s %s by: %s"
)

var (
	errNil = error(nil)
	errTst = errors.New("test error (never displayed)")
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

			label + "err trace message by: szlog.TErr",
			label + "err trace message by: szlog.TraceErr",
			label + "err formatted trace message by: szlog.TErrf",
			label + "err formatted trace message by: szlog.TraceErrf",
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

			label + "err debug message by: szlog.DErr",
			label + "err debug message by: szlog.DebugErr",
			label + "err formatted debug message by: szlog.DErrf",
			label + "err formatted debug message by: szlog.DebugErrf",
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

			label + "err info message by: szlog.IErr",
			label + "err info message by: szlog.InfoErr",
			label + "err formatted info message by: szlog.IErrf",
			label + "err formatted info message by: szlog.InfoErrf",
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

			label + "err warn message by: szlog.WErr",
			label + "err warn message by: szlog.WarnErr",
			label + "err formatted warn message by: szlog.WErrf",
			label + "err formatted warn message by: szlog.WarnErrf",
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

			label + "err error message by: szlog.EErr",
			label + "err error message by: szlog.ErrorErr",
			label + "err formatted error message by: szlog.EErrf",
			label + "err formatted error message by: szlog.ErrorErrf",
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

			label + "err fatal message by: szlog.FErr",
			label + "err fatal message by: szlog.FatalErr",
			label + "err formatted fatal message by: szlog.FErrf",
			label + "err formatted fatal message by: szlog.FatalErrf",
		}
	default:
		snip = []string{
			"unknown snippet area: " + area,
		}
	}

	return strings.Join(snip, "\n")
}

// logAll issues all log forms and levels.
//
//nolint:funlen // Ok.
func logAll() {
	szlog.T("plain ", "trace ", "message by: ", "szlog.T")
	szlog.Trace("plain ", "trace ", "message by: ", "szlog.Trace")
	szlog.Tf(msgFmt, "trace", "message", "szlog.Tf")
	szlog.Tracef(msgFmt, "trace", "message", "szlog.Tracef")

	szlog.TErr(errNil, "ERR ", "TRACE ", "MESSAGE BY: ", "szlog.TErr")
	szlog.TraceErr(errNil, "ERR ", "TRACE ", "MESSAGE BY: ", "szlog.TraceErr")
	szlog.TErrf(errNil, msgFmtErr, "TRACE", "MESSAGE", "szlog.TErrf")
	szlog.TraceErrf(errNil, msgFmtErr, "TRACE", "MESSAGE", "szlog.TraceErrf")

	szlog.TErr(errTst, "err ", "trace ", "message by: ", "szlog.TErr")
	szlog.TraceErr(errTst, "err ", "trace ", "message by: ", "szlog.TraceErr")
	szlog.TErrf(errTst, msgFmtErr, "trace", "message", "szlog.TErrf")
	szlog.TraceErrf(errTst, msgFmtErr, "trace", "message", "szlog.TraceErrf")

	szlog.D("plain ", "debug ", "message by: ", "szlog.D")
	szlog.Debug("plain ", "debug ", "message by: ", "szlog.Debug")
	szlog.Df(msgFmt, "debug", "message", "szlog.Df")
	szlog.Debugf(msgFmt, "debug", "message", "szlog.Debugf")

	szlog.DErr(errNil, "ERR ", "DEBUG ", "MESSAGE BY: ", "szlog.DErr")
	szlog.DebugErr(errNil, "ERR ", "DEBUG ", "MESSAGE BY: ", "szlog.DebugErr")
	szlog.DErrf(errNil, msgFmtErr, "DEBUG", "MESSAGE", "szlog.DErrf")
	szlog.DebugErrf(errNil, msgFmtErr, "DEBUG", "MESSAGE", "szlog.DebugErrf")

	szlog.DErr(errTst, "err ", "debug ", "message by: ", "szlog.DErr")
	szlog.DebugErr(errTst, "err ", "debug ", "message by: ", "szlog.DebugErr")
	szlog.DErrf(errTst, msgFmtErr, "debug", "message", "szlog.DErrf")
	szlog.DebugErrf(errTst, msgFmtErr, "debug", "message", "szlog.DebugErrf")

	szlog.I("plain ", "info ", "message by: ", "szlog.I")
	szlog.Info("plain ", "info ", "message by: ", "szlog.Info")
	szlog.If(msgFmt, "info", "message", "szlog.If")
	szlog.Infof(msgFmt, "info", "message", "szlog.Infof")

	szlog.IErr(errNil, "ERR ", "INFO ", "MESSAGE BY: ", "szlog.IErr")
	szlog.InfoErr(errNil, "ERR ", "INFO ", "MESSAGE BY: ", "szlog.InfoErr")
	szlog.IErrf(errNil, msgFmtErr, "INFO", "MESSAGE", "szlog.IErrf")
	szlog.InfoErrf(errNil, msgFmtErr, "INFO", "MESSAGE", "szlog.InfoErrf")

	szlog.IErr(errTst, "err ", "info ", "message by: ", "szlog.IErr")
	szlog.InfoErr(errTst, "err ", "info ", "message by: ", "szlog.InfoErr")
	szlog.IErrf(errTst, msgFmtErr, "info", "message", "szlog.IErrf")
	szlog.InfoErrf(errTst, msgFmtErr, "info", "message", "szlog.InfoErrf")

	szlog.W("plain ", "warn ", "message by: ", "szlog.W")
	szlog.Warn("plain ", "warn ", "message by: ", "szlog.Warn")
	szlog.Wf(msgFmt, "warn", "message", "szlog.Wf")
	szlog.Warnf(msgFmt, "warn", "message", "szlog.Warnf")

	szlog.WErr(errNil, "ERR ", "WARN ", "MESSAGE BY: ", "szlog.WErr")
	szlog.WarnErr(errNil, "ERR ", "WARN ", "MESSAGE BY: ", "szlog.WarnErr")
	szlog.WErrf(errNil, msgFmtErr, "WARN", "MESSAGE", "szlog.WErrf")
	szlog.WarnErrf(errNil, msgFmtErr, "WARN", "MESSAGE", "szlog.WarnErrf")

	szlog.WErr(errTst, "err ", "warn ", "message by: ", "szlog.WErr")
	szlog.WarnErr(errTst, "err ", "warn ", "message by: ", "szlog.WarnErr")
	szlog.WErrf(errTst, msgFmtErr, "warn", "message", "szlog.WErrf")
	szlog.WarnErrf(errTst, msgFmtErr, "warn", "message", "szlog.WarnErrf")

	szlog.E("plain ", "error ", "message by: ", "szlog.E")
	szlog.Error("plain ", "error ", "message by: ", "szlog.Error")
	szlog.Ef(msgFmt, "error", "message", "szlog.Ef")
	szlog.Errorf(msgFmt, "error", "message", "szlog.Errorf")

	szlog.EErr(errNil, "ERR ", "ERROR ", "MESSAGE BY: ", "szlog.EErr")
	szlog.ErrorErr(errNil, "ERR ", "ERROR ", "MESSAGE BY: ", "szlog.ErrorErr")
	szlog.EErrf(errNil, msgFmtErr, "ERROR", "MESSAGE", "szlog.EErrf")
	szlog.ErrorErrf(errNil, msgFmtErr, "ERROR", "MESSAGE", "szlog.ErrorErrf")

	szlog.EErr(errTst, "err ", "error ", "message by: ", "szlog.EErr")
	szlog.ErrorErr(errTst, "err ", "error ", "message by: ", "szlog.ErrorErr")
	szlog.EErrf(errTst, msgFmtErr, "error", "message", "szlog.EErrf")
	szlog.ErrorErrf(errTst, msgFmtErr, "error", "message", "szlog.ErrorErrf")

	szlog.F("plain ", "fatal ", "message by: ", "szlog.F")
	szlog.Fatal("plain ", "fatal ", "message by: ", "szlog.Fatal")
	szlog.Ff(msgFmt, "fatal", "message", "szlog.Ff")
	szlog.Fatalf(msgFmt, "fatal", "message", "szlog.Fatalf")

	szlog.FErr(errNil, "ERR ", "FATAL ", "MESSAGE BY: ", "szlog.FErr")
	szlog.FatalErr(errNil, "ERR ", "FATAL ", "MESSAGE BY: ", "szlog.FatalErr")
	szlog.FErrf(errNil, msgFmtErr, "FATAL", "MESSAGE", "szlog.FErrf")
	szlog.FatalErrf(errNil, msgFmtErr, "FATAL", "MESSAGE", "szlog.FatalErrf")

	szlog.FErr(errTst, "err ", "fatal ", "message by: ", "szlog.FErr")
	szlog.FatalErr(errTst, "err ", "fatal ", "message by: ", "szlog.FatalErr")
	szlog.FErrf(errTst, msgFmtErr, "fatal", "message", "szlog.FErrf")
	szlog.FatalErrf(errTst, msgFmtErr, "fatal", "message", "szlog.FatalErrf")
}

func TestLog_All(t *testing.T) {
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

func TestLog_AllWithLongLabels(t *testing.T) {
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

func TestLog_LevelFatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelFatal)

	logAll()

	chk.Log(
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestLog_LevelNone(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	logAll()

	chk.Log()
}

func TestLog_LevelError(t *testing.T) {
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

func TestLog_LevelWarn(t *testing.T) {
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

func TestLog_LevelInfo(t *testing.T) {
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

func TestLog_LevelDebug(t *testing.T) {
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

func TestLog_LevelTrace(t *testing.T) {
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

func TestLog_LevelCustom_None(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelNone)

	logAll()

	chk.Log()
}

func TestLog_LevelCustom_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelTrace)

	logAll()

	chk.Log(
		expectedLogSnippet("Trace", tstShortLabels),
	)
}

func TestLog_LevelCustom_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelDebug)

	logAll()

	chk.Log(
		expectedLogSnippet("Debug", tstShortLabels),
	)
}

func TestLog_LevelCustom_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelInfo)

	logAll()

	chk.Log(
		expectedLogSnippet("Info", tstShortLabels),
	)
}

func TestLog_LevelCustom_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelWarn)

	logAll()

	chk.Log(
		expectedLogSnippet("Warn", tstShortLabels),
	)
}

func TestLog_LevelCustom_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelError)

	logAll()

	chk.Log(
		expectedLogSnippet("Error", tstShortLabels),
	)
}

func TestLog_LevelCustom_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetCustomLevels(szlog.LevelFatal)

	logAll()

	chk.Log(
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}

func TestLog_LevelCustom_All(t *testing.T) {
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

func TestLog_LevelCustom_Fatal_Trace(t *testing.T) {
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

func TestLog_DefaultLogger(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	origLog := szlog.Default()
	tstLog := szlog.New(nil)

	oldLog := szlog.SetDefault(tstLog)

	tstLog2 := szlog.SetDefault(origLog)

	chk.True(oldLog == origLog)
	chk.True(tstLog == tstLog2)
	chk.True(szlog.Default() == origLog)

	szlog.Reset()

	logAll()
	chk.Log(
		expectedLogSnippet("Error", tstShortLabels),
		expectedLogSnippet("Fatal", tstShortLabels),
	)
}
