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
	"math"
	"strconv"
)

// LogLevel represents the current minium level of message to log.
type LogLevel int

// Available logging levels.
const (
	LevelNone LogLevel = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
	LevelCustom = LogLevel(math.MaxInt)
)

const (
	enabledFatal = 1 << iota
	enabledError
	enabledWarn
	enabledInfo
	enabledDebug
	enabledTrace

	enableLevelNone  = 0
	enableLevelFatal = enableLevelNone | enabledFatal
	enableLevelError = enableLevelFatal | enabledError
	enableLevelWarn  = enableLevelError | enabledWarn
	enableLevelInfo  = enableLevelWarn | enabledInfo
	enableLevelDebug = enableLevelInfo | enabledDebug
	enableLevelTrace = enableLevelDebug | enabledTrace
	enableLevelAll   = enableLevelTrace
)

// String implements the Stringer interface.
//
//nolint:cyclop // Ok.
func (ll LogLevel) String() string {
	switch ll {
	case LevelNone:
		return "NONE"
	case LevelFatal:
		return "FATAL"
	case LevelError:
		return "ERROR"
	case LevelWarn:
		return "WARN"
	case LevelInfo:
		return "INFO"
	case LevelDebug:
		return "DEBUG"
	case LevelTrace:
		return "TRACE"
	case LevelAll:
		return "ALL"
	case LevelCustom:
		return "CUSTOM"
	default:
		return "UNKNOWN:" + strconv.FormatInt(int64(ll), 10)
	}
}
