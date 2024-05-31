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

import "io"

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

// FErr Invokes the default log corresponding method.
func FErr(err error, msg ...any) bool {
	return defaultLog.FErr(err, msg...)
}

// Fatal Invokes the default log corresponding method.
func Fatal(msg ...any) bool {
	return defaultLog.Fatal(msg...)
}

// FatalErr Invokes the default log corresponding method.
func FatalErr(err error, msg ...any) bool {
	return defaultLog.FatalErr(err, msg...)
}

// Ff Invokes the default log corresponding method.
func Ff(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Ff(msgFmt, msgArgs...)
}

// FErrf Invokes the default log corresponding method.
func FErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.FErrf(err, msgFmt, msgArgs...)
}

// Fatalf Invokes the default log corresponding method.
func Fatalf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Fatalf(msgFmt, msgArgs...)
}

// FatalErrf Invokes the default log corresponding method.
func FatalErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.FatalErrf(err, msgFmt, msgArgs...)
}

// E Invokes the default log corresponding method.
func E(msg ...any) bool {
	return defaultLog.E(msg...)
}

// EErr Invokes the default log corresponding method.
func EErr(err error, msg ...any) bool {
	return defaultLog.EErr(err, msg...)
}

// Error Invokes the default log corresponding method.
func Error(msg ...any) bool {
	return defaultLog.Error(msg...)
}

// ErrorErr Invokes the default log corresponding method.
func ErrorErr(err error, msg ...any) bool {
	return defaultLog.ErrorErr(err, msg...)
}

// Ef Invokes the default log corresponding method.
func Ef(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Ef(msgFmt, msgArgs...)
}

// EErrf Invokes the default log corresponding method.
func EErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.EErrf(err, msgFmt, msgArgs...)
}

// Errorf Invokes the default log corresponding method.
func Errorf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Errorf(msgFmt, msgArgs...)
}

// ErrorErrf Invokes the default log corresponding method.
func ErrorErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.ErrorErrf(err, msgFmt, msgArgs...)
}

// W Invokes the default log corresponding method.
func W(msg ...any) bool {
	return defaultLog.W(msg...)
}

// WErr Invokes the default log corresponding method.
func WErr(err error, msg ...any) bool {
	return defaultLog.WErr(err, msg...)
}

// Warn Invokes the default log corresponding method.
func Warn(msg ...any) bool {
	return defaultLog.Warn(msg...)
}

// WarnErr Invokes the default log corresponding method.
func WarnErr(err error, msg ...any) bool {
	return defaultLog.WarnErr(err, msg...)
}

// Wf Invokes the default log corresponding method.
func Wf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Wf(msgFmt, msgArgs...)
}

// WErrf Invokes the default log corresponding method.
func WErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.WErrf(err, msgFmt, msgArgs...)
}

// Warnf Invokes the default log corresponding method.
func Warnf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Warnf(msgFmt, msgArgs...)
}

// WarnErrf Invokes the default log corresponding method.
func WarnErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.WarnErrf(err, msgFmt, msgArgs...)
}

// I Invokes the default log corresponding method.
func I(msg ...any) bool {
	return defaultLog.I(msg...)
}

// IErr Invokes the default log corresponding method.
func IErr(err error, msg ...any) bool {
	return defaultLog.IErr(err, msg...)
}

// Info Invokes the default log corresponding method.
func Info(msg ...any) bool {
	return defaultLog.Info(msg...)
}

// InfoErr Invokes the default log corresponding method.
func InfoErr(err error, msg ...any) bool {
	return defaultLog.InfoErr(err, msg...)
}

// If Invokes the default log corresponding method.
func If(msgFmt string, msgArgs ...any) bool {
	return defaultLog.If(msgFmt, msgArgs...)
}

// IErrf Invokes the default log corresponding method.
func IErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.IErrf(err, msgFmt, msgArgs...)
}

// Infof Invokes the default log corresponding method.
func Infof(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Infof(msgFmt, msgArgs...)
}

// InfoErrf Invokes the default log corresponding method.
func InfoErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.InfoErrf(err, msgFmt, msgArgs...)
}

// D Invokes the default log corresponding method.
func D(msg ...any) bool {
	return defaultLog.D(msg...)
}

// DErr Invokes the default log corresponding method.
func DErr(err error, msg ...any) bool {
	return defaultLog.DErr(err, msg...)
}

// Debug Invokes the default log corresponding method.
func Debug(msg ...any) bool {
	return defaultLog.Debug(msg...)
}

// DebugErr Invokes the default log corresponding method.
func DebugErr(err error, msg ...any) bool {
	return defaultLog.DebugErr(err, msg...)
}

// Df Invokes the default log corresponding method.
func Df(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Df(msgFmt, msgArgs...)
}

// DErrf Invokes the default log corresponding method.
func DErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.DErrf(err, msgFmt, msgArgs...)
}

// Debugf Invokes the default log corresponding method.
func Debugf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Debugf(msgFmt, msgArgs...)
}

// DebugErrf Invokes the default log corresponding method.
func DebugErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.DebugErrf(err, msgFmt, msgArgs...)
}

// T Invokes the default log corresponding method.
func T(msg ...any) bool {
	return defaultLog.T(msg...)
}

// TErr Invokes the default log corresponding method.
func TErr(err error, msg ...any) bool {
	return defaultLog.TErr(err, msg...)
}

// Trace Invokes the default log corresponding method.
func Trace(msg ...any) bool {
	return defaultLog.Trace(msg...)
}

// TraceErr Invokes the default log corresponding method.
func TraceErr(err error, msg ...any) bool {
	return defaultLog.TraceErr(err, msg...)
}

// Tf Invokes the default log corresponding method.
func Tf(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Tf(msgFmt, msgArgs...)
}

// TErrf Invokes the default log corresponding method.
func TErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.TErrf(err, msgFmt, msgArgs...)
}

// Tracef Invokes the default log corresponding method.
func Tracef(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Tracef(msgFmt, msgArgs...)
}

// TraceErrf Invokes the default log corresponding method.
func TraceErrf(err error, msgFmt string, msgArgs ...any) bool {
	return defaultLog.TraceErrf(err, msgFmt, msgArgs...)
}

// Close Invokes the default log corresponding method.
func Close(area string, closeable io.Closer) {
	defaultLog.Close(area, closeable)
}
