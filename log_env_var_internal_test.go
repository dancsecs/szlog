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
	"testing"

	"github.com/dancsecs/sztest"
)

func TestSzLog_EnvVar_getEnvLevel(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	testEnvVariable := "SZLOG_TEST_LOG_LEVEL"
	defer func() {
		_ = os.Unsetenv(testEnvVariable)
	}()

	chk.NoErr(os.Unsetenv(testEnvVariable))
	chk.Int(int(getEnvLevel(testEnvVariable, 78)), 78)

	chk.NoErr(os.Setenv(testEnvVariable, "invalidLevel"))
	chk.Int(int(getEnvLevel(testEnvVariable, 62)), 62)

	chk.NoErr(os.Setenv(testEnvVariable, "NONE"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelNone))

	chk.NoErr(os.Setenv(testEnvVariable, "FATAL"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelFatal))

	chk.NoErr(os.Setenv(testEnvVariable, "ERROR"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelError))

	chk.NoErr(os.Setenv(testEnvVariable, "WARN"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelWarn))

	chk.NoErr(os.Setenv(testEnvVariable, "INFO"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelInfo))

	chk.NoErr(os.Setenv(testEnvVariable, "DEBUG"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelDebug))

	chk.NoErr(os.Setenv(testEnvVariable, "TRACE"))
	chk.Int(int(getEnvLevel(testEnvVariable, 66)), int(LevelTrace))

	chk.NoErr(os.Setenv(testEnvVariable, "ALL"))
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
	defer func() {
		_ = os.Unsetenv(testEnvVariable)
	}()

	chk.NoErr(os.Unsetenv(testEnvVariable))
	chk.False(getEnvSetting(testEnvVariable))

	chk.NoErr(os.Setenv(testEnvVariable, "invalidSetting"))
	chk.False(getEnvSetting(testEnvVariable))

	chk.NoErr(os.Setenv(testEnvVariable, "DISABLED"))
	chk.False(getEnvSetting(testEnvVariable))

	chk.NoErr(os.Setenv(testEnvVariable, "ENABLED"))
	chk.True(getEnvSetting(testEnvVariable))

	chk.Log(
		"W:szlog initialization: invalid environment override " +
			"SZLOG_TEST_LOG_LEVEL=\"invalidSetting\": " +
			"unknown log level setting error",
	)
}
