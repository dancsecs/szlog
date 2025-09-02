/*
   Szerszam alarm manager: szalarm.
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

package szlog_test

import (
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

func TestSzLog_Exported_SetLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

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

	szlog.E("E message")
	szlog.Error("Error message")
	szlog.Ef("fmt: %s", "Ef message")
	szlog.Errorf("fmt: %s", "Errorf message")

	szlog.W("W message")
	szlog.Warn("Warn message")
	szlog.Wf("fmt: %s", "Wf message")
	szlog.Warnf("fmt: %s", "Warnf message")

	szlog.I("I message")
	szlog.Info("Info message")
	szlog.If("fmt: %s", "If message")
	szlog.Infof("fmt: %s", "Infof message")

	szlog.D("D message")
	szlog.Debug("Debug message")
	szlog.Df("fmt: %s", "Df message")
	szlog.Debugf("fmt: %s", "Debugf message")

	szlog.T("T message")
	szlog.Trace("Trace message")
	szlog.Tf("fmt: %s", "Tf message")
	szlog.Tracef("fmt: %s", "Tracef message")

	chk.Log(
		"F:F message",
		"F:Fatal message",
		"F:fmt: Ff message",
		"F:fmt: Fatalf message",

		"E:E message",
		"E:Error message",
		"E:fmt: Ef message",
		"E:fmt: Errorf message",

		"W:W message",
		"W:Warn message",
		"W:fmt: Wf message",
		"W:fmt: Warnf message",

		"I:I message",
		"I:Info message",
		"I:fmt: If message",
		"I:fmt: Infof message",

		"D:D message",
		"D:Debug message",
		"D:fmt: Df message",
		"D:fmt: Debugf message",

		"T:T message",
		"T:Trace message",
		"T:fmt: Tf message",
		"T:fmt: Tracef message",
	)
}
