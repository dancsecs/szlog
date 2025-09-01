/*
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

func TestSzLog_Language(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstOut := szlog.New()

	chk.Str(tstOut.Language(), "")

	chk.NoErr(tstOut.SetLanguage("en"))

	chk.Str(tstOut.Language(), "en")

	chk.Err(
		tstOut.SetLanguage("unknown-language"),
		chk.ErrChain(
			szlog.ErrInvalidLanguage,
			"'unknown-language'",
		),
	)

	chk.NoErr(tstOut.SetLanguage(""))

	chk.Str(tstOut.Language(), "")

	chk.Log()
}
