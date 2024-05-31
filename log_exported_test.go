/*
   Szerszam alarm manager: szalarm.
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
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

//nolint:funlen // Ok.
func TestSzLog_Exported_LevelDisabled(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.DisableLevel(szlog.LevelAll)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.True(szlog.LevelDisabled(szlog.LevelFatal))
	chk.True(szlog.LevelDisabled(szlog.LevelError))
	chk.True(szlog.LevelDisabled(szlog.LevelWarn))
	chk.True(szlog.LevelDisabled(szlog.LevelInfo))
	chk.True(szlog.LevelDisabled(szlog.LevelDebug))
	chk.True(szlog.LevelDisabled(szlog.LevelTrace))
	chk.True(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	chk.True(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelAll)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.True(szlog.LevelDisabled(szlog.LevelFatal))
	chk.True(szlog.LevelDisabled(szlog.LevelError))
	chk.True(szlog.LevelDisabled(szlog.LevelWarn))
	chk.True(szlog.LevelDisabled(szlog.LevelInfo))
	chk.True(szlog.LevelDisabled(szlog.LevelDebug))
	chk.True(szlog.LevelDisabled(szlog.LevelTrace))
	chk.True(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelFatal)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.True(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelError)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.True(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelWarn)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.True(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelInfo)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.True(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelDebug)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.True(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelTrace)
	chk.False(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.True(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelNone)
	chk.True(szlog.LevelDisabled(szlog.LevelNone))
	chk.False(szlog.LevelDisabled(szlog.LevelFatal))
	chk.False(szlog.LevelDisabled(szlog.LevelError))
	chk.False(szlog.LevelDisabled(szlog.LevelWarn))
	chk.False(szlog.LevelDisabled(szlog.LevelInfo))
	chk.False(szlog.LevelDisabled(szlog.LevelDebug))
	chk.False(szlog.LevelDisabled(szlog.LevelTrace))
	chk.False(szlog.LevelDisabled(szlog.LevelAll))

	szlog.DisableLevel(szlog.LevelAll + 1)
	szlog.DisableLevel(szlog.LevelNone)
	szlog.DisableLevel(szlog.LevelNone - 1)

	chk.True(szlog.LevelDisabled(szlog.LevelNone - 1))
	chk.False(szlog.LevelDisabled(szlog.LevelAll + 1))

	chk.Log(
		"W:attempt to access out of bounds log level: 8 from: DisableLevel",
		"W:attempt to access out of bounds log level: -1 from: DisableLevel",
		"W:attempt to access out of bounds log level: -1 from: LevelDisabled",
		"W:attempt to access out of bounds log level: 8 from: LevelDisabled",
	)
}

func TestSzLog_Exported_SetLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.DisableLevel(szlog.LevelNone)

	lastLevel := szlog.Level()
	chk.Int(int(szlog.SetLevel(szlog.LevelNone)), int(lastLevel))
	chk.Int(int(szlog.Level()), int(szlog.LevelNone))

	chk.Int(int(szlog.SetLevel(szlog.LevelFatal)), int(szlog.LevelNone))
	chk.Int(int(szlog.Level()), int(szlog.LevelFatal))

	chk.Int(int(szlog.SetLevel(szlog.LevelError)), int(szlog.LevelFatal))
	chk.Int(int(szlog.Level()), int(szlog.LevelError))

	chk.Int(int(szlog.SetLevel(szlog.LevelWarn)), int(szlog.LevelError))
	chk.Int(int(szlog.Level()), int(szlog.LevelWarn))

	chk.Int(int(szlog.SetLevel(szlog.LevelInfo)), int(szlog.LevelWarn))
	chk.Int(int(szlog.Level()), int(szlog.LevelInfo))

	chk.Int(int(szlog.SetLevel(szlog.LevelDebug)), int(szlog.LevelInfo))
	chk.Int(int(szlog.Level()), int(szlog.LevelDebug))

	chk.Int(int(szlog.SetLevel(szlog.LevelTrace)), int(szlog.LevelDebug))
	chk.Int(int(szlog.Level()), int(szlog.LevelTrace))

	chk.Int(int(szlog.SetLevel(szlog.LevelAll)), int(szlog.LevelTrace))
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	chk.Int(int(szlog.SetLevel(szlog.LevelAll+1)), int(szlog.LevelAll))
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	chk.Int(int(szlog.SetLevel(szlog.LevelNone-1)), int(szlog.LevelAll))
	chk.Int(int(szlog.Level()), int(szlog.LevelNone))

	chk.Log(
		"W:attempt to access out of bounds log level: 8 from: SetLevel",
		"W:attempt to access out of bounds log level: -1 from: SetLevel",
	)
}

func TestSzLog_Exported_IncDecLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.SetLevel(szlog.LevelNone)

	chk.Int(int(szlog.Level()), int(szlog.LevelNone))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelNone))
	chk.Int(int(szlog.Level()), int(szlog.LevelFatal))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelFatal))
	chk.Int(int(szlog.Level()), int(szlog.LevelError))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelError))
	chk.Int(int(szlog.Level()), int(szlog.LevelWarn))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelWarn))
	chk.Int(int(szlog.Level()), int(szlog.LevelInfo))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelInfo))
	chk.Int(int(szlog.Level()), int(szlog.LevelDebug))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelDebug))
	chk.Int(int(szlog.Level()), int(szlog.LevelTrace))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelTrace))
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	chk.Int(int(szlog.IncLevel()), int(szlog.LevelAll))
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelAll))
	chk.Int(int(szlog.Level()), int(szlog.LevelTrace))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelTrace))
	chk.Int(int(szlog.Level()), int(szlog.LevelDebug))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelDebug))
	chk.Int(int(szlog.Level()), int(szlog.LevelInfo))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelInfo))
	chk.Int(int(szlog.Level()), int(szlog.LevelWarn))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelWarn))
	chk.Int(int(szlog.Level()), int(szlog.LevelError))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelError))
	chk.Int(int(szlog.Level()), int(szlog.LevelFatal))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelFatal))
	chk.Int(int(szlog.Level()), int(szlog.LevelNone))

	chk.Int(int(szlog.DecLevel()), int(szlog.LevelNone))
	chk.Int(int(szlog.Level()), int(szlog.LevelNone))

	chk.Log(
		"W:attempt to access out of bounds log level: 8 from: IncLevel",
		"W:attempt to access out of bounds log level: -1 from: DecLevel",
	)
}

//nolint:funlen // Ok.
func TestSzLog_Exported_LogRedirects(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)

	szlog.F("F message")
	szlog.Fatal("Fatal message")
	szlog.Ff("fmt: %s", "Ff message")
	szlog.Fatalf("fmt: %s", "Fatalf message")
	//
	szlog.FErr(errNil, "FErr message NOT DISPLAYED")
	szlog.FatalErr(errNil, "FatalErr message NOT DISPLAYED")
	szlog.FErrf(errNil, "fmt: %s", "FErrf message NOT DISPLAYED")
	szlog.FatalErrf(errNil, "fmt: %s", "FatalErrf message NOT DISPLAYED")
	//
	szlog.FErr(errTst, "FErr message displayed")
	szlog.FatalErr(errTst, "FatalErr message displayed")
	szlog.FErrf(errTst, "fmt: %s", "FErrf message displayed")
	szlog.FatalErrf(errTst, "fmt: %s", "FatalErrf message displayed")

	szlog.E("E message")
	szlog.Error("Error message")
	szlog.Ef("fmt: %s", "Ef message")
	szlog.Errorf("fmt: %s", "Errorf message")
	//
	szlog.EErr(errNil, "EErr message NOT DISPLAYED")
	szlog.ErrorErr(errNil, "ErrorErr message NOT DISPLAYED")
	szlog.EErrf(errNil, "fmt: %s", "EErr message NOT DISPLAYED")
	szlog.ErrorErrf(errNil, "fmt: %s", "ErrorErrf message NOT DISPLAYED")
	//
	szlog.EErr(errTst, "EErr message displayed")
	szlog.ErrorErr(errTst, "ErrorErr message displayed")
	szlog.EErrf(errTst, "fmt: %s", "EErrf message displayed")
	szlog.ErrorErrf(errTst, "fmt: %s", "ErrorErrf message displayed")

	szlog.W("W message")
	szlog.Warn("Warn message")
	szlog.Wf("fmt: %s", "Wf message")
	szlog.Warnf("fmt: %s", "Warnf message")
	//
	szlog.WErr(errNil, "WErr message NOT DISPLAYED")
	szlog.WarnErr(errNil, "WarnErr message NOT DISPLAYED")
	szlog.WErrf(errNil, "fmt: %s", "WErrf message NOT DISPLAYED")
	szlog.WarnErrf(errNil, "fmt: %s", "WarnErrf message NOT DISPLAYED")
	//
	szlog.WErr(errTst, "WErr message displayed")
	szlog.WarnErr(errTst, "WarnErr message displayed")
	szlog.WErrf(errTst, "fmt: %s", "WErrf message displayed")
	szlog.WarnErrf(errTst, "fmt: %s", "WarnErrf message displayed")

	szlog.I("I message")
	szlog.Info("Info message")
	szlog.If("fmt: %s", "If message")
	szlog.Infof("fmt: %s", "Infof message")
	//
	szlog.IErr(errNil, "IErr message NOT DISPLAYED")
	szlog.InfoErr(errNil, "InfoErr message NOT DISPLAYED")
	szlog.IErrf(errNil, "fmt: %s", "IErrf message NOT DISPLAYED")
	szlog.InfoErrf(errNil, "fmt: %s", "InfoErrf message NOT DISPLAYED")
	//
	szlog.IErr(errTst, "IErr message displayed")
	szlog.InfoErr(errTst, "InfoErr message displayed")
	szlog.IErrf(errTst, "fmt: %s", "IErrf message displayed")
	szlog.InfoErrf(errTst, "fmt: %s", "InfoErrf message displayed")

	szlog.D("D message")
	szlog.Debug("Debug message")
	szlog.Df("fmt: %s", "Df message")
	szlog.Debugf("fmt: %s", "Debugf message")
	//
	szlog.DErr(errNil, "DErr message NOT DISPLAYED")
	szlog.DebugErr(errNil, "DebugErr message NOT DISPLAYED")
	szlog.DErrf(errNil, "fmt: %s", "DErrf message NOT DISPLAYED")
	szlog.DebugErrf(errNil, "fmt: %s", "DebugErrf message NOT DISPLAYED")
	//
	szlog.DErr(errTst, "DErr message displayed")
	szlog.DebugErr(errTst, "DebugErr message displayed")
	szlog.DErrf(errTst, "fmt: %s", "DErrf message displayed")
	szlog.DebugErrf(errTst, "fmt: %s", "DebugErrf message displayed")

	szlog.T("T message")
	szlog.Trace("Trace message")
	szlog.Tf("fmt: %s", "Tf message")
	szlog.Tracef("fmt: %s", "Tracef message")
	//
	szlog.TErr(errNil, "TErr message NOT DISPLAYED")
	szlog.TraceErr(errNil, "TraceErr message NOT DISPLAYED")
	szlog.TErrf(errNil, "fmt: %s", "TErrf message NOT DISPLAYED")
	szlog.TraceErrf(errNil, "fmt: %s", "TraceErrf message NOT DISPLAYED")
	//
	szlog.TErr(errTst, "TErr message displayed")
	szlog.TraceErr(errTst, "TraceErr message displayed")
	szlog.TErrf(errTst, "fmt: %s", "TErrf message displayed")
	szlog.TraceErrf(errTst, "fmt: %s", "TraceErrf message displayed")

	chk.Log(
		"F:F message",
		"F:Fatal message",
		"F:fmt: Ff message",
		"F:fmt: Fatalf message",

		"F:FErr message displayed",
		"F:FatalErr message displayed",
		"F:fmt: FErrf message displayed",
		"F:fmt: FatalErrf message displayed",

		"E:E message",
		"E:Error message",
		"E:fmt: Ef message",
		"E:fmt: Errorf message",

		"E:EErr message displayed",
		"E:ErrorErr message displayed",
		"E:fmt: EErrf message displayed",
		"E:fmt: ErrorErrf message displayed",

		"W:W message",
		"W:Warn message",
		"W:fmt: Wf message",
		"W:fmt: Warnf message",

		"W:WErr message displayed",
		"W:WarnErr message displayed",
		"W:fmt: WErrf message displayed",
		"W:fmt: WarnErrf message displayed",

		"I:I message",
		"I:Info message",
		"I:fmt: If message",
		"I:fmt: Infof message",

		"I:IErr message displayed",
		"I:InfoErr message displayed",
		"I:fmt: IErrf message displayed",
		"I:fmt: InfoErrf message displayed",

		"D:D message",
		"D:Debug message",
		"D:fmt: Df message",
		"D:fmt: Debugf message",

		"D:DErr message displayed",
		"D:DebugErr message displayed",
		"D:fmt: DErrf message displayed",
		"D:fmt: DebugErrf message displayed",

		"T:T message",
		"T:Trace message",
		"T:fmt: Tf message",
		"T:fmt: Tracef message",

		"T:TErr message displayed",
		"T:TraceErr message displayed",
		"T:fmt: TErrf message displayed",
		"T:fmt: TraceErrf message displayed",
	)
}
