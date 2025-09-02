/*
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

// Test versions of labels.
const (
	logFatalLabel = "F:"
	logErrorLabel = "E:"
	logWarnLabel  = "W:"
	logInfoLabel  = "I:"
	logDebugLabel = "D:"
	logTraceLabel = "T:"

	logFatalLongLabel = "FATAL:"
	logErrorLongLabel = "ERROR:"
	logWarnLongLabel  = "WARN:"
	logInfoLongLabel  = "INFO:"
	logDebugLongLabel = "DEBUG:"
	logTraceLongLabel = "TRACE:"
)

func TestSzLog_Labels(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstOut := szlog.New()

	chk.False(tstOut.LongLabels())

	chk.False(tstOut.SetLongLabels(true))

	chk.True(tstOut.LongLabels())

	chk.True(tstOut.SetLongLabels(false))

	chk.False(tstOut.LongLabels())

	chk.Log()
}
