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
	"os"
	"strings"
)

// Environment variable default overrides.
const (
	EnvLogLevel      = "SZLOG_LEVEL"
	EnvLogLanguage   = "SZLOG_LANGUAGE"
	EnvVerbose       = "SZLOG_VERBOSE"
	EnvLogLongLabels = "SZLOG_LONG_LABELS"
)

//nolint:goCheckNoInits // Ok.
func init() {
	Reset()
}

func (l *Log) logEnvErrorMessage(env, envValue, errMsg string) {
	l.logWarnf(
		"szlog initialization: invalid environment override %s=%q: %s",
		env, envValue, errMsg,
	)
}

func (l *Log) setEnvVerbose() {
	var lvl VerboseLevel

	rawVerbose, ok := os.LookupEnv(EnvVerbose)

	if ok {
		switch strings.ToLower(rawVerbose) {
		case "quiet":
			lvl = -1
		case "0":
			lvl = 0
		case "1":
			lvl = 1
		case "2":
			lvl = 2
		case "3":
			lvl = 3
		case "4":
			lvl = 4
		case "5":
			lvl = 5
		default:
			l.logEnvErrorMessage(
				EnvVerbose,
				rawVerbose,
				"unknown verbose level (must be one of: "+
					"'QUIET', '0', '1', '2', '3', '4', '5'"+
					")",
			)
		}
	}

	l.SetVerbose(lvl)
}

func (l *Log) setEnvLanguage() {
	rawLanguage, ok := os.LookupEnv(EnvLogLanguage)

	if ok {
		err := l.SetLanguage(rawLanguage)
		if err != nil {
			l.logEnvErrorMessage(
				EnvLogLanguage,
				rawLanguage,
				err.Error(),
			)
		}
	} else {
		_ = l.SetLanguage("")
	}
}

func (l *Log) setEnvLevel() {
	rawEnvLevel, ok := os.LookupEnv(EnvLogLevel)
	if ok {
		err := l.parseAndSetLevel(rawEnvLevel, LevelError)
		if err != nil {
			l.logEnvErrorMessage(
				EnvLogLevel,
				rawEnvLevel,
				err.Error(),
			)
		}
	} else {
		l.SetLevel(LevelError)
	}
}

func (l *Log) setEnvLabelLength() {
	rawLabelLength, ok := os.LookupEnv(EnvLogLongLabels)

	if ok {
		switch strings.ToLower(rawLabelLength) {
		case "short":
			l.SetLongLabels(false)
		case "long":
			l.SetLongLabels(true)
		default:
			l.logEnvErrorMessage(
				EnvLogLongLabels,
				rawLabelLength,
				"unknown label length (must be 'SHORT' or 'LONG')")
		}
	} else {
		l.SetLongLabels(false)
	}
}
