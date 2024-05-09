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
func TestSzLogExported_LevelDisabled(t *testing.T) {
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

func TestSzLogExported_SetLevel(t *testing.T) {
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

func TestSzLogExported_IncDecLevel(t *testing.T) {
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

func TestSzLogExported_LogRedirects(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)

	szlog.F("f error")
	szlog.Fatal("fatal error")
	szlog.Ff("fmt: %s", "ff error")
	szlog.Fatalf("fmt: %s", "fatalf error")

	szlog.E("e error")
	szlog.Error("error error") //nolint:dupword // Ok.
	szlog.Ef("fmt: %s", "ef error")
	szlog.Errorf("fmt: %s", "errorf error")

	szlog.W("w error")
	szlog.Warn("warn error")
	szlog.Wf("fmt: %s", "wf error")
	szlog.Warnf("fmt: %s", "warnf error")

	szlog.I("i error")
	szlog.Info("info error")
	szlog.If("fmt: %s", "if error")
	szlog.Infof("fmt: %s", "infof error")

	szlog.D("d error")
	szlog.Debug("debug error")
	szlog.Df("fmt: %s", "df error")
	szlog.Debugf("fmt: %s", "debugf error")

	szlog.T("t error")
	szlog.Trace("trace error")
	szlog.Tf("fmt: %s", "tf error")
	szlog.Tracef("fmt: %s", "tracef error")

	chk.Log(
		"F:f error",
		"F:fatal error",
		"F:fmt: ff error",
		"F:fmt: fatalf error",

		"E:e error",
		"E:error error",
		"E:fmt: ef error",
		"E:fmt: errorf error",

		"W:w error",
		"W:warn error",
		"W:fmt: wf error",
		"W:fmt: warnf error",

		"I:i error",
		"I:info error",
		"I:fmt: if error",
		"I:fmt: infof error",

		"D:d error",
		"D:debug error",
		"D:fmt: df error",
		"D:fmt: debugf error",

		"T:t error",
		"T:trace error",
		"T:fmt: tf error",
		"T:fmt: tracef error",
	)
}
