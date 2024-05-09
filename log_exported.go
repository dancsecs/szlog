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

// F Invokes the default log corresponding method.
func F(msg ...any) bool {
	return defaultLog.F(msg...)
}

// Fatal Invokes the default log corresponding method.
func Fatal(msg ...any) bool {
	return defaultLog.Fatal(msg...)
}

// Ff Invokes the default log corresponding method.
func Ff(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Ff(msgFmt, msgArgs...)
}

// Fatalf Invokes the default log corresponding method.
func Fatalf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Fatalf(msgFmt, msgArgs...)
}

// E Invokes the default log corresponding method.
func E(msg ...any) bool {
	return defaultLog.E(msg...)
}

// Error Invokes the default log corresponding method.
func Error(msg ...any) bool {
	return defaultLog.Error(msg...)
}

// Ef Invokes the default log corresponding method.
func Ef(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Ef(msgFmt, msgArgs...)
}

// Errorf Invokes the default log corresponding method.
func Errorf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Errorf(msgFmt, msgArgs...)
}

// W Invokes the default log corresponding method.
func W(msg ...any) bool {
	return defaultLog.W(msg...)
}

// Warn Invokes the default log corresponding method.
func Warn(msg ...any) bool {
	return defaultLog.Warn(msg...)
}

// Wf Invokes the default log corresponding method.
func Wf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Wf(msgFmt, msgArgs...)
}

// Warnf Invokes the default log corresponding method.
func Warnf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Warnf(msgFmt, msgArgs...)
}

// I Invokes the default log corresponding method.
func I(msg ...any) bool {
	return defaultLog.I(msg...)
}

// Info Invokes the default log corresponding method.
func Info(msg ...any) bool {
	return defaultLog.Info(msg...)
}

// If Invokes the default log corresponding method.
func If(msgFmt string, msgArgs ...any) bool {
	return defaultLog.If(msgFmt, msgArgs...)
}

// Infof Invokes the default log corresponding method.
func Infof(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Infof(msgFmt, msgArgs...)
}

// D Invokes the default log corresponding method.
func D(msg ...any) bool {
	return defaultLog.D(msg...)
}

// Debug Invokes the default log corresponding method.
func Debug(msg ...any) bool {
	return defaultLog.Debug(msg...)
}

// Df Invokes the default log corresponding method.
func Df(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Df(msgFmt, msgArgs...)
}

// Debugf Invokes the default log corresponding method.
func Debugf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Debugf(msgFmt, msgArgs...)
}

// T Invokes the default log corresponding method.
func T(msg ...any) bool {
	return defaultLog.T(msg...)
}

// Trace Invokes the default log corresponding method.
func Trace(msg ...any) bool {
	return defaultLog.Trace(msg...)
}

// Tf Invokes the default log corresponding method.
func Tf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Tf(msgFmt, msgArgs...)
}

// Tracef Invokes the default log corresponding method.
func Tracef(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Tracef(msgFmt, msgArgs...)
}
