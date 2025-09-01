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

import (
	"io"

	"golang.org/x/text/message"
)

// LogFunc defines the signature of an unformatted log function.
type LogFunc func(msg ...any) bool

// LogFuncf defines the signature of a formatted log function.
type LogFuncf func(msgFmt string, msgArgs ...any) bool

// Log represents a szlog logging object.
type Log struct {
	level               LogLevel
	verboseLevel        VerboseLevel
	language            string
	customLevelsEnabled int
	longLabels          bool
	printer             *message.Printer
	stdout              io.Writer

	LogFatal bool
	LogError bool
	LogWarn  bool
	LogInfo  bool
	LogDebug bool
	LogTrace bool

	// Logging functions
	F, Fatal   LogFunc
	Ff, Fatalf LogFuncf
	E, Error   LogFunc
	Ef, Errorf LogFuncf
	W, Warn    LogFunc
	Wf, Warnf  LogFuncf
	I, Info    LogFunc
	If, Infof  LogFuncf
	D, Debug   LogFunc
	Df, Debugf LogFuncf
	T, Trace   LogFunc
	Tf, Tracef LogFuncf

	// Verbose functions
	S0, Say0   LogFunc
	S0f, Say0f LogFuncf
	S1, Say1   LogFunc
	S1f, Say1f LogFuncf
	S2, Say2   LogFunc
	S2f, Say2f LogFuncf
	S3, Say3   LogFunc
	S3f, Say3f LogFuncf
	S4, Say4   LogFunc
	S4f, Say4f LogFuncf
	S5, Say5   LogFunc
	S5f, Say5f LogFuncf
}

// New creates a new log object.
func New() *Log {
	l := new(Log)
	l.Reset()

	return l
}

// Reset returns all log setting to default startup conditions.
func (l *Log) Reset() {
	l.setEnvLabelLength()
	l.setEnvLevel()
	l.setEnvLanguage()
	l.setEnvVerbose()
	l.SetStdout(nil)
}
