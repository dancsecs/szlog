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

//nolint:funlen // Ok.
func TestLog_All(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)

	logAll()

	chk.Log(
		"T:plain trace message by: szlog.T",
		"T:plain trace message by: szlog.Trace",
		"T:formatted trace message by: szlog.Tf",
		"T:formatted trace message by: szlog.Tracef",

		"T:err trace message by: szlog.TErr",
		"T:err trace message by: szlog.TraceErr",
		"T:err formatted trace message by: szlog.TErrf",
		"T:err formatted trace message by: szlog.TraceErrf",

		"D:plain debug message by: szlog.D",
		"D:plain debug message by: szlog.Debug",
		"D:formatted debug message by: szlog.Df",
		"D:formatted debug message by: szlog.Debugf",

		"D:err debug message by: szlog.DErr",
		"D:err debug message by: szlog.DebugErr",
		"D:err formatted debug message by: szlog.DErrf",
		"D:err formatted debug message by: szlog.DebugErrf",

		"I:plain info message by: szlog.I",
		"I:plain info message by: szlog.Info",
		"I:formatted info message by: szlog.If",
		"I:formatted info message by: szlog.Infof",

		"I:err info message by: szlog.IErr",
		"I:err info message by: szlog.InfoErr",
		"I:err formatted info message by: szlog.IErrf",
		"I:err formatted info message by: szlog.InfoErrf",

		"W:plain warn message by: szlog.W",
		"W:plain warn message by: szlog.Warn",
		"W:formatted warn message by: szlog.Wf",
		"W:formatted warn message by: szlog.Warnf",

		"W:err warn message by: szlog.WErr",
		"W:err warn message by: szlog.WarnErr",
		"W:err formatted warn message by: szlog.WErrf",
		"W:err formatted warn message by: szlog.WarnErrf",

		"E:plain error message by: szlog.E",
		"E:plain error message by: szlog.Error",
		"E:formatted error message by: szlog.Ef",
		"E:formatted error message by: szlog.Errorf",

		"E:err error message by: szlog.EErr",
		"E:err error message by: szlog.ErrorErr",
		"E:err formatted error message by: szlog.EErrf",
		"E:err formatted error message by: szlog.ErrorErrf",

		"F:plain fatal message by: szlog.F",
		"F:plain fatal message by: szlog.Fatal",
		"F:formatted fatal message by: szlog.Ff",
		"F:formatted fatal message by: szlog.Fatalf",

		"F:err fatal message by: szlog.FErr",
		"F:err fatal message by: szlog.FatalErr",
		"F:err formatted fatal message by: szlog.FErrf",
		"F:err formatted fatal message by: szlog.FatalErrf",
	)
}

//nolint:funlen // Ok.
func TestLog_AllWithLongLabels(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)
	szlog.SetLongLabels(true)

	logAll()

	chk.Log(
		"TRACE:plain trace message by: szlog.T",
		"TRACE:plain trace message by: szlog.Trace",
		"TRACE:formatted trace message by: szlog.Tf",
		"TRACE:formatted trace message by: szlog.Tracef",

		"TRACE:err trace message by: szlog.TErr",
		"TRACE:err trace message by: szlog.TraceErr",
		"TRACE:err formatted trace message by: szlog.TErrf",
		"TRACE:err formatted trace message by: szlog.TraceErrf",

		"DEBUG:plain debug message by: szlog.D",
		"DEBUG:plain debug message by: szlog.Debug",
		"DEBUG:formatted debug message by: szlog.Df",
		"DEBUG:formatted debug message by: szlog.Debugf",

		"DEBUG:err debug message by: szlog.DErr",
		"DEBUG:err debug message by: szlog.DebugErr",
		"DEBUG:err formatted debug message by: szlog.DErrf",
		"DEBUG:err formatted debug message by: szlog.DebugErrf",

		"INFO:plain info message by: szlog.I",
		"INFO:plain info message by: szlog.Info",
		"INFO:formatted info message by: szlog.If",
		"INFO:formatted info message by: szlog.Infof",

		"INFO:err info message by: szlog.IErr",
		"INFO:err info message by: szlog.InfoErr",
		"INFO:err formatted info message by: szlog.IErrf",
		"INFO:err formatted info message by: szlog.InfoErrf",

		"WARN:plain warn message by: szlog.W",
		"WARN:plain warn message by: szlog.Warn",
		"WARN:formatted warn message by: szlog.Wf",
		"WARN:formatted warn message by: szlog.Warnf",

		"WARN:err warn message by: szlog.WErr",
		"WARN:err warn message by: szlog.WarnErr",
		"WARN:err formatted warn message by: szlog.WErrf",
		"WARN:err formatted warn message by: szlog.WarnErrf",

		"ERROR:plain error message by: szlog.E",
		"ERROR:plain error message by: szlog.Error",
		"ERROR:formatted error message by: szlog.Ef",
		"ERROR:formatted error message by: szlog.Errorf",

		"ERROR:err error message by: szlog.EErr",
		"ERROR:err error message by: szlog.ErrorErr",
		"ERROR:err formatted error message by: szlog.EErrf",
		"ERROR:err formatted error message by: szlog.ErrorErrf",

		"FATAL:plain fatal message by: szlog.F",
		"FATAL:plain fatal message by: szlog.Fatal",
		"FATAL:formatted fatal message by: szlog.Ff",
		"FATAL:formatted fatal message by: szlog.Fatalf",

		"FATAL:err fatal message by: szlog.FErr",
		"FATAL:err fatal message by: szlog.FatalErr",
		"FATAL:err formatted fatal message by: szlog.FErrf",
		"FATAL:err formatted fatal message by: szlog.FatalErrf",
	)
}
