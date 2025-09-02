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

import (
	"testing"

	"github.com/dancsecs/sztest"
)

func TestSzLog_EnvVar_getEnvLogLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New()

	chk.DelEnv(EnvLogLevel)
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelError))

	chk.SetEnv(EnvLogLevel, "78")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelError))

	chk.SetEnv(EnvLogLevel, "invalidLevel")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelError))

	chk.SetEnv(EnvLogLevel, "NONE")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelNone))

	chk.SetEnv(EnvLogLevel, "FATAL")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelFatal))

	chk.SetEnv(EnvLogLevel, "ERROR")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelError))

	chk.SetEnv(EnvLogLevel, "WARN")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelWarn))

	chk.SetEnv(EnvLogLevel, "INFO")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelInfo))

	chk.SetEnv(EnvLogLevel, "DEBUG")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelDebug))

	chk.SetEnv(EnvLogLevel, "TRACE")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelTrace))

	chk.SetEnv(EnvLogLevel, "ALL")
	tstLog.setEnvLevel()
	chk.Int(int(tstLog.Level()), int(LevelAll))

	chk.Log(
		"W:szlog initialization: invalid environment override "+
			"SZLOG_LEVEL=\"78\": "+
			"invalid log level string: unknown log level: '78'",
		"W:szlog initialization: invalid environment override "+
			"SZLOG_LEVEL=\"invalidLevel\": "+
			"invalid log level string: "+
			"unknown log level: 'invalidLevel'",
	)
}

func TestSzLog_EnvVar_getEnvVerbose(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New()

	chk.DelEnv(EnvVerbose)
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 0)

	chk.SetEnv(EnvVerbose, "invalidSetting")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 0)

	chk.SetEnv(EnvVerbose, "QUIET")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), -1)

	chk.SetEnv(EnvVerbose, "0")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 0)

	chk.SetEnv(EnvVerbose, "1")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 1)

	chk.SetEnv(EnvVerbose, "2")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 2)

	chk.SetEnv(EnvVerbose, "3")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 3)

	chk.SetEnv(EnvVerbose, "4")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 4)

	chk.SetEnv(EnvVerbose, "5")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 5)

	chk.SetEnv(EnvVerbose, "6")
	tstLog.setEnvVerbose()
	chk.Int(tstLog.Verbose(), 0)

	chk.Log(
		"W:szlog initialization: invalid environment override "+
			"SZLOG_VERBOSE=\"invalidSetting\": "+
			"unknown verbose level (must be one of: "+
			"'QUIET', '0', '1', '2', '3', '4', '5'"+
			")",
		"W:szlog initialization: invalid environment override "+
			"SZLOG_VERBOSE=\"6\": "+
			"unknown verbose level (must be one of: "+
			"'QUIET', '0', '1', '2', '3', '4', '5'"+
			")",
	)
}

func TestSzLog_EnvVar_getEnvLanguage(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New()

	chk.DelEnv(EnvLogLanguage)
	tstLog.setEnvLanguage()
	chk.Str(tstLog.Language(), "")

	chk.SetEnv(EnvLogLanguage, "en")
	tstLog.setEnvLanguage()
	chk.Str(tstLog.Language(), "en")

	chk.SetEnv(EnvLogLanguage, "unknown")
	tstLog.setEnvLanguage()
	chk.Str(tstLog.Language(), "")

	chk.Log(
		"W:szlog initialization: invalid environment override " +
			"SZLOG_LANGUAGE=\"unknown\": invalid language: 'unknown'",
	)
}

func TestSzLog_EnvVar_getEnvLongLabels(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	tstLog := New()

	chk.DelEnv(EnvLogLongLabels)
	tstLog.setEnvLabelLength()
	chk.False(tstLog.LongLabels())

	chk.SetEnv(EnvLogLongLabels, "SHORT")
	tstLog.setEnvLabelLength()
	chk.False(tstLog.LongLabels())

	chk.SetEnv(EnvLogLongLabels, "LONG")
	tstLog.setEnvLabelLength()
	chk.True(tstLog.LongLabels())

	chk.SetEnv(EnvLogLongLabels, "unknown")
	tstLog.setEnvLabelLength()
	chk.True(tstLog.LongLabels())

	chk.Log(
		"WARN:szlog initialization: invalid environment override " +
			"SZLOG_LONG_LABELS=\"unknown\": " +
			"unknown label length (must be 'SHORT' or 'LONG')",
	)
}
