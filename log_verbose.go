/*
   Szerszam logging library: szlog.
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

package szlog

import (
	"fmt"
	"os"
)

// VerboseLevel indicates how much non-essential output is allowed.
// Levels range from -1 (silent) to 5 (most verbose).
type VerboseLevel = int

// Verbose reports the logger's current verbosity level.
func (l *Log) Verbose() VerboseLevel {
	return l.verboseLevel
}

// SetVerbose adjusts the verbosity level (-1 through 5). Level -1 silences
// all output, while higher levels progressively enable more detail.
func (l *Log) SetVerbose(newLevel VerboseLevel) VerboseLevel {
	origLevel := l.verboseLevel

	l.S0 = l.selectLog(newLevel > -1, l.vPrint)
	l.Say0 = l.S0
	l.S0f = l.selectLogf(newLevel > -1, l.vPrintf)
	l.Say0f = l.S0f

	l.S1 = l.selectLog(newLevel > 0, l.vPrint)
	l.Say1 = l.S1
	l.S1f = l.selectLogf(newLevel > 0, l.vPrintf)
	l.Say1f = l.S1f

	l.S2 = l.selectLog(newLevel > 1, l.vPrint)
	l.Say2 = l.S2
	l.S2f = l.selectLogf(newLevel > 1, l.vPrintf)
	l.Say2f = l.S2f

	l.S3 = l.selectLog(newLevel > 2, l.vPrint) //nolint:mnd // Ok.
	l.Say3 = l.S3
	l.S3f = l.selectLogf(newLevel > 2, l.vPrintf) //nolint:mnd // Ok.
	l.Say3f = l.S3f

	l.S4 = l.selectLog(newLevel > 3, l.vPrint) //nolint:mnd // Ok.
	l.Say4 = l.S4
	l.S4f = l.selectLogf(newLevel > 3, l.vPrintf) //nolint:mnd // Ok.
	l.Say4f = l.S4f

	l.S5 = l.selectLog(newLevel > 4, l.vPrint) //nolint:mnd // Ok.
	l.Say5 = l.S5
	l.S5f = l.selectLogf(newLevel > 4, l.vPrintf) //nolint:mnd // Ok.
	l.Say5f = l.S5f

	l.verboseLevel = newLevel

	return origLevel
}

func (l *Log) vPrint(msg ...any) bool {
	if l.printer != nil {
		_, _ = l.printer.Fprint(os.Stdout, msg...)
	} else {
		_, _ = fmt.Fprint(os.Stdout, msg...)
	}

	return true
}

func (l *Log) vPrintf(msgFmt string, msgArgs ...any) bool {
	if l.printer != nil {
		_, _ = l.printer.Fprintf(os.Stdout, msgFmt, msgArgs...)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, msgFmt, msgArgs...)
	}

	return true
}
