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

const (
	envLogLevel      = "SZLOG_LEVEL"
	envLogLevelFatal = "SZLOG_LEVEL_FATAL"
	envLogLevelError = "SZLOG_LEVEL_ERROR"
	envLogLevelWarn  = "SZLOG_LEVEL_WARN"
	envLogLevelInfo  = "SZLOG_LEVEL_INFO"
	envLogLevelDebug = "SZLOG_LEVEL_DEBUG"
	envLogLevelTrace = "SZLOG_LEVEL_TRACE"
	envLogLongLabels = "SZLOG_LONG_LABELS"
)

//nolint:goCheckNoInits // Ok.
func init() {
	Reset()
}

func logEnvErrorMessage(env, envValue, errMsg string) {
	logWarnf(
		"szlog initialization: invalid environment override %s=%q: %s error",
		env, envValue, errMsg,
	)
}

//nolint:cyclop // Ok.
func getEnvLevel(env string, value LogLevel) LogLevel {
	rawEnvLevel, ok := os.LookupEnv(env)
	if ok {
		switch strings.ToLower(rawEnvLevel) {
		case "none":
			value = LevelNone
		case "fatal":
			value = LevelFatal
		case "error":
			value = LevelError
		case "info":
			value = LevelInfo
		case "warn":
			value = LevelWarn
		case "debug":
			value = LevelDebug
		case "trace":
			value = LevelTrace
		case "all":
			value = LevelAll
		default:
			logEnvErrorMessage(env, rawEnvLevel, "unknown log level")
		}
	}

	return value
}

func getEnvSetting(env string) bool {
	rawEnv, ok := os.LookupEnv(env)
	enabled := false

	if ok {
		switch strings.ToLower(rawEnv) {
		case "enabled":
			enabled = true
		case "disabled":
		default:
			logEnvErrorMessage(env, rawEnv, "unknown log level setting")
		}
	}

	return enabled
}
