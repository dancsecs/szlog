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

// Level return the current logging level.
func Level() LogLevel {
	return defaultLog.Level()
}

// SetLevel sets the logging level.
func SetLevel(newLogLevel LogLevel) LogLevel {
	return defaultLog.SetLevel(newLogLevel)
}

// SetCustomLevels sets the logging level.
func SetCustomLevels(levels ...LogLevel) LogLevel {
	return defaultLog.SetCustomLevels(levels...)
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

// AbsorbArgs increases log level according to how many verbose flags
// encountered.
func AbsorbArgs(argsIn []string) ([]string, error) {
	return defaultLog.AbsorbArgs(argsIn)
}

// LogFatal returns true if fatal messages are being logged.
func LogFatal() bool {
	return defaultLog.LogFatal
}

// LogError returns true if error messages are being logged.
func LogError() bool {
	return defaultLog.LogError
}

// LogWarn returns true if warning messages are being logged.
func LogWarn() bool {
	return defaultLog.LogWarn
}

// LogInfo returns true if information messages are being logged.
func LogInfo() bool {
	return defaultLog.LogInfo
}

// LogDebug returns true if debug messages are being logged.
func LogDebug() bool {
	return defaultLog.LogDebug
}

// LogTrace returns true if tracing messages are being logged.
func LogTrace() bool {
	return defaultLog.LogTrace
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

// Close Invokes the default log corresponding method.
func Close(area string, closeable io.Closer) {
	defaultLog.Close(area, closeable)
}

// Verbose returns the current verbose level.
func Verbose() VerboseLevel {
	return defaultLog.Verbose()
}

// SetVerbose sets the default verbose level.
func SetVerbose(level VerboseLevel) VerboseLevel {
	return defaultLog.SetVerbose(level)
}

// Local returns the current language local string.
func Local() string {
	return defaultLog.Language()
}

// SetLanguage set the default local language formatting.
func SetLanguage(language string) error {
	return defaultLog.SetLanguage(language)
}

// S0 Invokes the default verbose corresponding method.
func S0(msg ...any) bool {
	return defaultLog.S0(msg...)
}

// S1 Invokes the default verbose corresponding method.
func S1(msg ...any) bool {
	return defaultLog.S1(msg...)
}

// S2 Invokes the default verbose corresponding method.
func S2(msg ...any) bool {
	return defaultLog.S2(msg...)
}

// S3 Invokes the default verbose corresponding method.
func S3(msg ...any) bool {
	return defaultLog.S3(msg...)
}

// S4 Invokes the default verbose corresponding method.
func S4(msg ...any) bool {
	return defaultLog.S4(msg...)
}

// S5 Invokes the default verbose corresponding method.
func S5(msg ...any) bool {
	return defaultLog.S5(msg...)
}

// Say0 Invokes the default verbose corresponding method.
func Say0(msg ...any) bool {
	return defaultLog.Say0(msg...)
}

// Say1 Invokes the default verbose corresponding method.
func Say1(msg ...any) bool {
	return defaultLog.Say1(msg...)
}

// Say2 Invokes the default verbose corresponding method.
func Say2(msg ...any) bool {
	return defaultLog.Say2(msg...)
}

// Say3 Invokes the default verbose corresponding method.
func Say3(msg ...any) bool {
	return defaultLog.Say3(msg...)
}

// Say4 Invokes the default verbose corresponding method.
func Say4(msg ...any) bool {
	return defaultLog.Say4(msg...)
}

// Say5 Invokes the default verbose corresponding method.
func Say5(msg ...any) bool {
	return defaultLog.Say5(msg...)
}

// S0f Invokes the default verbose corresponding method.
func S0f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.S0f(msgFmt, msgArgs...)
}

// S1f Invokes the default verbose corresponding method.
func S1f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.S1f(msgFmt, msgArgs...)
}

// S2f Invokes the default verbose corresponding method.
func S2f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.S2f(msgFmt, msgArgs...)
}

// S3f Invokes the default verbose corresponding method.
func S3f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.S3f(msgFmt, msgArgs...)
}

// S4f Invokes the default verbose corresponding method.
func S4f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.S4f(msgFmt, msgArgs...)
}

// S5f Invokes the default verbose corresponding method.
func S5f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.S5f(msgFmt, msgArgs...)
}

// Say0f Invokes the default verbose corresponding method.
func Say0f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Say0f(msgFmt, msgArgs...)
}

// Say1f Invokes the default verbose corresponding method.
func Say1f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Say1f(msgFmt, msgArgs...)
}

// Say2f Invokes the default verbose corresponding method.
func Say2f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Say2f(msgFmt, msgArgs...)
}

// Say3f Invokes the default verbose corresponding method.
func Say3f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Say3f(msgFmt, msgArgs...)
}

// Say4f Invokes the default verbose corresponding method.
func Say4f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Say4f(msgFmt, msgArgs...)
}

// Say5f Invokes the default verbose corresponding method.
func Say5f(msgFmt string, msgArgs ...any) bool {
	return defaultLog.Say5f(msgFmt, msgArgs...)
}
