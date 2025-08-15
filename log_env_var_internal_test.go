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
	"testing"

	"github.com/dancsecs/sztest"
)

func TestSzLog_EnvVar_getEnvLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	testEnvVariable := "SZLOG_TEST_LOG_LEVEL"

	chk.DelEnv(testEnvVariable)
	chk.Int(int(getEnvLevel(testEnvVariable, 78)), 78)

	chk.SetEnv(testEnvVariable, "invalidLevel")
	chk.Int(int(getEnvLevel(testEnvVariable, 62)), 62)

	chk.SetEnv(testEnvVariable, "NONE")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelNone))

	chk.SetEnv(testEnvVariable, "FATAL")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelFatal))

	chk.SetEnv(testEnvVariable, "ERROR")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelError))

	chk.SetEnv(testEnvVariable, "WARN")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelWarn))

	chk.SetEnv(testEnvVariable, "INFO")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelInfo))

	chk.SetEnv(testEnvVariable, "DEBUG")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelDebug))

	chk.SetEnv(testEnvVariable, "TRACE")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelTrace))

	chk.SetEnv(testEnvVariable, "ALL")
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelAll))

	chk.Log(
		"W:szlog initialization: invalid environment override " +
			"SZLOG_TEST_LOG_LEVEL=\"invalidLevel\": " +
			"unknown log level error",
	)
}

func TestSzLog_EnvVar_getEnvSetting(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	testEnvVariable := "SZLOG_TEST_LOG_LEVEL"

	chk.DelEnv(testEnvVariable)
	chk.False(getEnvSetting(testEnvVariable))

	chk.SetEnv(testEnvVariable, "invalidSetting")
	chk.False(getEnvSetting(testEnvVariable))

	chk.SetEnv(testEnvVariable, "DISABLED")
	chk.False(getEnvSetting(testEnvVariable))

	chk.SetEnv(testEnvVariable, "ENABLED")
	chk.True(getEnvSetting(testEnvVariable))

	chk.Log(
		"W:szlog initialization: invalid environment override " +
			"SZLOG_TEST_LOG_LEVEL=\"invalidSetting\": " +
			"unknown log level setting error",
	)
}
