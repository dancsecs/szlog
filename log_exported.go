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

// LongLabels return true if long labels are enabled.
func LongLabels() bool {
	return defaultLog.LongLabels()
}

// SetLongLabels return true if long labels are enabled.
func SetLongLabels(enabled bool) bool {
	return defaultLog.SetLongLabels(enabled)
}

// LevelDisabled returns true if the  level is disabled.
func LevelDisabled(level LogLevel) bool {
	return defaultLog.LevelDisabled(level)
}

// DisableLevel blocks all logging at the specified level.
func DisableLevel(level LogLevel) {
	defaultLog.DisableLevel(level)
}

// Level return the current logging level.
func Level() LogLevel {
	return defaultLog.Level()
}

// SetLevel sets the logging level.
func SetLevel(newLogLevel LogLevel) LogLevel {
	return defaultLog.SetLevel(newLogLevel)
}

// IncLevel permits all logging at the specified level.
func IncLevel() LogLevel {
	return defaultLog.IncLevel()
}

// DecLevel permits all logging at the specified level.
func DecLevel() LogLevel {
	return defaultLog.DecLevel()
}

// Reset returns all log setting to default startup conditions.
func Reset() {
	defaultLog.Reset()
}

// VerboseAbsorbArgs increases log level according to how many verbose flags
// encountered.
func VerboseAbsorbArgs(argsIn []string) []string {
	return defaultLog.VerboseAbsorbArgs(argsIn)
}
