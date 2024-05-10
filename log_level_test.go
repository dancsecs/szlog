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

//nolint:dupl // Ok.
package szlog_test

import (
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

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
	chk.Str((szlog.LevelAll + 1).String(), "UNKNOWN:8")
}
