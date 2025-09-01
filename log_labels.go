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

package szlog

type labelIdx int

const (
	labelFatal labelIdx = iota
	labelError
	labelWarn
	labelInfo
	labelDebug
	labelTrace
)

//nolint:goCheckNoGlobals // Package level local lookups.
var (
	shortLabel = []string{
		"F:",
		"E:",
		"W:",
		"I:",
		"D:",
		"T:",
	}
	longLabel = []string{
		"FATAL:",
		"ERROR:",
		"WARN:",
		"INFO:",
		"DEBUG:",
		"TRACE:",
	}
)

// LongLabels reports whether long labels (FATAL, ERROR, WARN, INFO, DEBUG,
// TRACE) are currently enabled instead of their short forms (F, E, W, I, D,
// T).
func (l *Log) LongLabels() bool {
	return l.longLabels
}

// SetLongLabels enables or disables long labels in log output. When disabled,
// short labels (F, E, W, I, D, T) are used instead. It returns the previous
// setting.
func (l *Log) SetLongLabels(enable bool) bool {
	orig := l.longLabels
	l.longLabels = enable

	return orig
}
