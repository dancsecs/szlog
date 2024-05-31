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

const (
	logFatalLabel = "F:"
	logErrorLabel = "E:"
	logWarnLabel  = "W:"
	logInfoLabel  = "I:"
	logDebugLabel = "D:"
	logTraceLabel = "T:"

	logLongFatalLabel = "FATAL:"
	logLongErrorLabel = "ERROR:"
	logLongWarnLabel  = "WARN:"
	logLongInfoLabel  = "INFO:"
	logLongDebugLabel = "DEBUG:"
	logLongTraceLabel = "TRACE:"
)

// LongLabels enables/disables the use of longer labels in log output.
func (l *Log) LongLabels() bool {
	return l.longLabels
}

// SetLongLabels enables/disables the use of longer labels in log output.
func (l *Log) SetLongLabels(enable bool) bool {
	orig := l.longLabels
	l.longLabels = enable
	l.SetLevel(l.level)

	return orig
}
