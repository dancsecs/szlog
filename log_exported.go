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

import "io"

// LongLabels reports whether long labels (FATAL, ERROR, WARN, INFO, DEBUG,
// TRACE) are currently enabled instead of their short forms (F, E, W, I, D,
// T).
func LongLabels() bool {
	return defaultLog.LongLabels()
}

// SetLongLabels enables or disables long labels in log output. When disabled,
// short labels (F, E, W, I, D, T) are used instead. It returns the previous
// setting.
func SetLongLabels(enabled bool) bool {
	return defaultLog.SetLongLabels(enabled)
}

// Level reports the logger's current logging level.
func Level() LogLevel {
	return defaultLog.Level()
}

// SetLevel updates the logger's logging level. Valid values include
// LevelNone, LevelFatal, LevelError, LevelWarn, LevelInfo, LevelDebug,
// LevelTrace, and LevelAll.
func SetLevel(newLogLevel LogLevel) LogLevel {
	return defaultLog.SetLevel(newLogLevel)
}

// SetCustomLevels enables a custom combination of individual levels.
// LevelNone, LevelAll, and LevelCustom are ignored. Internally, this
// always results in LevelCustom being applied.
func SetCustomLevels(levels ...LogLevel) LogLevel {
	return defaultLog.SetCustomLevels(levels...)
}

// Reset restores all log settings to their default values.
func Reset() {
	defaultLog.Reset()
}

// AbsorbArgs scans the provided argument list for logging-related flags.
// It updates the log configuration (LogLevel, verbosity, quiet mode,
// LongLabels, and Language) based on the flags encountered. Recognized
// flags are removed, and the cleaned argument slice is returned.
// Multiple `-v` flags increment verbosity accordingly. If conflicting
// or invalid flags are found (e.g., combining `-v` with `--quiet`),
// an error is returned along with the original arguments.
func AbsorbArgs(argsIn []string) ([]string, error) {
	return defaultLog.AbsorbArgs(argsIn)
}

// ArgUsageInfo invokes the provided callback function for all the arguments
// szlog processes.  (Used to provide the usage information).
func ArgUsageInfo(registerArgs func(string, string)) {
	defaultLog.ArgUsageInfo(registerArgs)
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

// Close is a convenience method for safely closing any io.Closer.
// If an error occurs during Close, it is logged as a warning.
// This method is primarily intended for use in defer statements.
func Close(area string, closeable io.Closer) {
	defaultLog.Close(area, closeable)
}

// Verbose reports the logger's current verbosity level.
func Verbose() VerboseLevel {
	return defaultLog.Verbose()
}

// SetVerbose adjusts the verbosity level (-1 through 5). Level -1 silences
// all output, while higher levels progressively enable more detail.
func SetVerbose(level VerboseLevel) VerboseLevel {
	return defaultLog.SetVerbose(level)
}

// Language returns the current language setting used for localized formatting.
// An empty string indicates no localization is applied.
func Language() string {
	return defaultLog.Language()
}

// SetLanguage updates the language used for localized formatting.
// Passing an empty string ("") disables localization. It returns any
// error encountered while setting the language.
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
