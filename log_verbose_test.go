/*
   Szerszam alarm manager: szalarm.
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

//nolint:funlen // Ok.
package szlog_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

const (
	cv1 = "-v1"
	cv2 = "--v2"
	nov = "notVerbose"
	cv3 = "-vvv1x"
	cv4 = "--vvv2x"
)

var (
	errArg      = errors.New("unexpected cleaned args")
	errLevel    = errors.New("unexpected szlog level")
	errVerbose  = errors.New("unexpected verbose level")
	errLanguage = errors.New("unexpected language local")
)

//nolint:cyclop // Ok.
func testArgs(
	expectedLevel szlog.LogLevel,
	expectedVerbose szlog.VerboseLevel,
	expectedLanguage string,
	args []string,
) error {
	var (
		err   error
		cArgs []string
	)

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	cArgs, err = szlog.AbsorbArgs(args)

	if err == nil {
		switch {
		case len(cArgs) != 5 ||
			cArgs[0] != cv1 ||
			cArgs[1] != cv2 ||
			cArgs[2] != nov ||
			cArgs[3] != cv3 ||
			cArgs[4] != cv4:
			//
			err = fmt.Errorf("%w: %v", errArg, cArgs)
		case szlog.Level() != expectedLevel:
			err = fmt.Errorf(
				"%w: Want: %s Got: %s", errLevel,
				expectedLevel.String(),
				szlog.Level().String(),
			)
		case szlog.Verbose() != expectedVerbose:
			err = fmt.Errorf(
				"%w: Want: %d Got: %d", errVerbose,
				expectedVerbose,
				szlog.Verbose(),
			)
		case szlog.Local() != expectedLanguage:
			err = fmt.Errorf(
				"%w: Want: %s Got: %s", errLanguage,
				expectedLanguage,
				szlog.Local(),
			)
		}
	}

	return err
}

func TestSzLog_LanguageArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	cleanedArgs, err := szlog.AbsorbArgs(nil)
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language",
		}),
		szlog.ErrMissingLanguage.Error(),
	)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language", "en",
			"--language", "fr",
		}),
		szlog.ErrAmbiguousLanguage.Error(),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language", "",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "en", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language", "en",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "fr", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language", "fr",
		}),
	)

	chk.Log()
}

func TestSzLog_LogLevelArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	cleanedArgs, err := szlog.AbsorbArgs(nil)
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log",
		}),
		szlog.ErrMissingLogLevel.Error(),
	)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "unknown",
		}),
		chk.ErrChain(
			szlog.ErrInvalidLogLevelParse,
			szlog.ErrUnknownLevel,
			"'unknown'",
		),
	)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "custom",
		}),
		chk.ErrChain(
			szlog.ErrInvalidLogLevelParse,
			szlog.ErrUnknownLevel,
			"'custom'",
		),
	)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "error",
			"--log", "unknown",
		}),
		chk.ErrChain(
			szlog.ErrAmbiguousLogLevel,
		),
	)

	chk.NoErr(
		testArgs(szlog.LevelFatal, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "fatal",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelNone, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "none",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "error",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelWarn, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "warn",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelInfo, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "info",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelDebug, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "debug",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelTrace, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "trace",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelAll, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "all",
		}),
	)

	chk.Log()
}

func TestSzLog_CustomLogLevelArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.Err(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "()",
		}),
		chk.ErrChain(
			szlog.ErrInvalidLogLevelParse,
			szlog.ErrUnknownLevel,
			"'()'",
		),
	)

	chk.Err(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(custom)",
		}),
		chk.ErrChain(
			szlog.ErrInvalidLogLevelParse,
			szlog.ErrUnknownLevel,
			"'(custom)'",
		),
	)

	chk.Err(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(trace|bad)",
		}),
		chk.ErrChain(
			szlog.ErrInvalidLogLevelParse,
			szlog.ErrUnknownLevel,
			"'(trace|bad)'",
		),
	)

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(trace)",
		}),
	)
	chk.False(szlog.LogFatal())
	chk.False(szlog.LogError())
	chk.False(szlog.LogWarn())
	chk.False(szlog.LogInfo())
	chk.False(szlog.LogDebug())
	chk.True(szlog.LogTrace())

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(debug)",
		}),
	)
	chk.False(szlog.LogFatal())
	chk.False(szlog.LogError())
	chk.False(szlog.LogWarn())
	chk.False(szlog.LogInfo())
	chk.True(szlog.LogDebug())
	chk.False(szlog.LogTrace())

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(Info)",
		}),
	)
	chk.False(szlog.LogFatal())
	chk.False(szlog.LogError())
	chk.False(szlog.LogWarn())
	chk.True(szlog.LogInfo())
	chk.False(szlog.LogDebug())
	chk.False(szlog.LogTrace())

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(WARN)",
		}),
	)
	chk.False(szlog.LogFatal())
	chk.False(szlog.LogError())
	chk.True(szlog.LogWarn())
	chk.False(szlog.LogInfo())
	chk.False(szlog.LogDebug())
	chk.False(szlog.LogTrace())

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(ERROR)",
		}),
	)
	chk.False(szlog.LogFatal())
	chk.True(szlog.LogError())
	chk.False(szlog.LogWarn())
	chk.False(szlog.LogInfo())
	chk.False(szlog.LogDebug())
	chk.False(szlog.LogTrace())

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(FaTal)",
		}),
	)
	chk.True(szlog.LogFatal())
	chk.False(szlog.LogError())
	chk.False(szlog.LogWarn())
	chk.False(szlog.LogInfo())
	chk.False(szlog.LogDebug())
	chk.False(szlog.LogTrace())

	chk.NoErr(
		testArgs(szlog.LevelCustom, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "(FaTal|trace)",
		}),
	)
	chk.True(szlog.LogFatal())
	chk.False(szlog.LogError())
	chk.False(szlog.LogWarn())
	chk.False(szlog.LogInfo())
	chk.False(szlog.LogDebug())
	chk.True(szlog.LogTrace())

	chk.Log()
}

func TestSzLog_VerboseArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	cleanedArgs, err := szlog.AbsorbArgs(nil)
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, -1, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--quiet",
			"-v",
		}),
		chk.ErrChain(
			szlog.ErrAmbiguousVerboseAndQuiet,
		),
	)

	chk.Err(
		testArgs(szlog.LevelError, -1, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"-v",
			"--quiet",
		}),
		chk.ErrChain(
			szlog.ErrAmbiguousVerboseAndQuiet,
		),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, -1, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--quiet",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "", []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 1, "", []string{
			"-v",
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 2, "", []string{
			"-v",
			cv1,
			"--v",
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 4, "", []string{
			"-v",
			cv1,
			"--v",
			cv2,
			"-vv",
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 6, "", []string{
			"-v",
			cv1,
			"--v",
			cv2,
			"-vv",
			nov,
			cv3,
			"--vv",
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 7, "", []string{
			"-v",
			cv1,
			"--v",
			cv2,
			"-vv",
			nov,
			cv3,
			"--vv",
			cv4,
			"--verbose",
		}),
	)

	chk.Log()
}

func TestSzLog_Verbose_Quiet(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	tLog := szlog.New()

	tLog.SetVerbose(-1)

	tLog.S0("Short Hello: 0\n")
	tLog.S1("Short Hello: 1\n")
	tLog.S2("Short Hello: 2\n")
	tLog.S3("Short Hello: 3\n")
	tLog.S4("Short Hello: 4\n")
	tLog.S5("Short Hello: 5\n")

	tLog.Say0("Long Hello: 0\n")
	tLog.Say1("Long Hello: 1\n")
	tLog.Say2("Long Hello: 2\n")
	tLog.Say3("Long Hello: 3\n")
	tLog.Say4("Long Hello: 4\n")
	tLog.Say5("Long Hello: 5\n")

	tLog.S0f("Short Formatted Hello: %d\n", 0)
	tLog.S1f("Short Formatted Hello: %d\n", 1)
	tLog.S2f("Short Formatted Hello: %d\n", 2)
	tLog.S3f("Short Formatted Hello: %d\n", 3)
	tLog.S4f("Short Formatted Hello: %d\n", 4)
	tLog.S5f("Short Formatted Hello: %d\n", 5)

	tLog.Say0f("Long Formatted Hello: %d\n", 0)
	tLog.Say1f("Long Formatted Hello: %d\n", 1)
	tLog.Say2f("Long Formatted Hello: %d\n", 2)
	tLog.Say3f("Long Formatted Hello: %d\n", 3)
	tLog.Say4f("Long Formatted Hello: %d\n", 4)
	tLog.Say5f("Long Formatted Hello: %d\n", 5)

	chk.Log()
	chk.Stdout()
}

func TestSzLog_Verbose_Default(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	tLog := szlog.New()

	tLog.SetVerbose(0)

	tLog.S0("Short Hello: 0\n")
	tLog.S1("Short Hello: 1\n")
	tLog.S2("Short Hello: 2\n")
	tLog.S3("Short Hello: 3\n")
	tLog.S4("Short Hello: 4\n")
	tLog.S5("Short Hello: 5\n")

	tLog.Say0("Long Hello: 0\n")
	tLog.Say1("Long Hello: 1\n")
	tLog.Say2("Long Hello: 2\n")
	tLog.Say3("Long Hello: 3\n")
	tLog.Say4("Long Hello: 4\n")
	tLog.Say5("Long Hello: 5\n")

	tLog.S0f("Short Formatted Hello: %d\n", 0)
	tLog.S1f("Short Formatted Hello: %d\n", 1)
	tLog.S2f("Short Formatted Hello: %d\n", 2)
	tLog.S3f("Short Formatted Hello: %d\n", 3)
	tLog.S4f("Short Formatted Hello: %d\n", 4)
	tLog.S5f("Short Formatted Hello: %d\n", 5)

	tLog.Say0f("Long Formatted Hello: %d\n", 0)
	tLog.Say1f("Long Formatted Hello: %d\n", 1)
	tLog.Say2f("Long Formatted Hello: %d\n", 2)
	tLog.Say3f("Long Formatted Hello: %d\n", 3)
	tLog.Say4f("Long Formatted Hello: %d\n", 4)
	tLog.Say5f("Long Formatted Hello: %d\n", 5)

	tLog.S0("A Short Local Test: ", 1234, "\n")
	tLog.S0f("A Short Local Formatted Test: %d\n", 2345)
	tLog.Say0("A Long Local Test: ", 3456, "\n")
	tLog.Say0f("A Long Local Formatted Test: %d\n", 4567)

	chk.Log()
	chk.Stdout(
		"Short Hello: 0",
		"Long Hello: 0",
		"Short Formatted Hello: 0",
		"Long Formatted Hello: 0",
		"A Short Local Test: 1234",
		"A Short Local Formatted Test: 2345",
		"A Long Local Test: 3456",
		"A Long Local Formatted Test: 4567",
	)
}

func TestSzLog_Verbose_V1(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetVerbose(1)
	chk.NoErr(szlog.SetLanguage("en"))

	szlog.S0("Short Hello: 0\n")
	szlog.S1("Short Hello: 1\n")
	szlog.S2("Short Hello: 2\n")
	szlog.S3("Short Hello: 3\n")
	szlog.S4("Short Hello: 4\n")
	szlog.S5("Short Hello: 5\n")

	szlog.Say0("Long Hello: 0\n")
	szlog.Say1("Long Hello: 1\n")
	szlog.Say2("Long Hello: 2\n")
	szlog.Say3("Long Hello: 3\n")
	szlog.Say4("Long Hello: 4\n")
	szlog.Say5("Long Hello: 5\n")

	szlog.S0f("Short Formatted Hello: %d\n", 0)
	szlog.S1f("Short Formatted Hello: %d\n", 1)
	szlog.S2f("Short Formatted Hello: %d\n", 2)
	szlog.S3f("Short Formatted Hello: %d\n", 3)
	szlog.S4f("Short Formatted Hello: %d\n", 4)
	szlog.S5f("Short Formatted Hello: %d\n", 5)

	szlog.Say0f("Long Formatted Hello: %d\n", 0)
	szlog.Say1f("Long Formatted Hello: %d\n", 1)
	szlog.Say2f("Long Formatted Hello: %d\n", 2)
	szlog.Say3f("Long Formatted Hello: %d\n", 3)
	szlog.Say4f("Long Formatted Hello: %d\n", 4)
	szlog.Say5f("Long Formatted Hello: %d\n", 5)

	szlog.S0("A Short Local Test: ", 1234, "\n")
	szlog.S0f("A Short Local Formatted Test: %d\n", 2345)
	szlog.Say0("A Long Local Test: ", 3456, "\n")
	szlog.Say0f("A Long Local Formatted Test: %d\n", 4567)

	chk.Log()
	chk.Stdout(
		"Short Hello: 0",
		"Short Hello: 1",
		"Long Hello: 0",
		"Long Hello: 1",
		"Short Formatted Hello: 0",
		"Short Formatted Hello: 1",
		"Long Formatted Hello: 0",
		"Long Formatted Hello: 1",
		"A Short Local Test: 1,234",
		"A Short Local Formatted Test: 2,345",
		"A Long Local Test: 3,456",
		"A Long Local Formatted Test: 4,567",
	)
}

func TestSzLog_Verbose_V2(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetVerbose(2)
	chk.NoErr(szlog.SetLanguage("en"))

	szlog.S0("Short Hello: 0\n")
	szlog.S1("Short Hello: 1\n")
	szlog.S2("Short Hello: 2\n")
	szlog.S3("Short Hello: 3\n")
	szlog.S4("Short Hello: 4\n")
	szlog.S5("Short Hello: 5\n")

	szlog.Say0("Long Hello: 0\n")
	szlog.Say1("Long Hello: 1\n")
	szlog.Say2("Long Hello: 2\n")
	szlog.Say3("Long Hello: 3\n")
	szlog.Say4("Long Hello: 4\n")
	szlog.Say5("Long Hello: 5\n")

	szlog.S0f("Short Formatted Hello: %d\n", 0)
	szlog.S1f("Short Formatted Hello: %d\n", 1)
	szlog.S2f("Short Formatted Hello: %d\n", 2)
	szlog.S3f("Short Formatted Hello: %d\n", 3)
	szlog.S4f("Short Formatted Hello: %d\n", 4)
	szlog.S5f("Short Formatted Hello: %d\n", 5)

	szlog.Say0f("Long Formatted Hello: %d\n", 0)
	szlog.Say1f("Long Formatted Hello: %d\n", 1)
	szlog.Say2f("Long Formatted Hello: %d\n", 2)
	szlog.Say3f("Long Formatted Hello: %d\n", 3)
	szlog.Say4f("Long Formatted Hello: %d\n", 4)
	szlog.Say5f("Long Formatted Hello: %d\n", 5)

	szlog.S0("A Short Local Test: ", 1234, "\n")
	szlog.S0f("A Short Local Formatted Test: %d\n", 2345)
	szlog.Say0("A Long Local Test: ", 3456, "\n")
	szlog.Say0f("A Long Local Formatted Test: %d\n", 4567)

	chk.Log()
	chk.Stdout(
		"Short Hello: 0",
		"Short Hello: 1",
		"Short Hello: 2",
		"Long Hello: 0",
		"Long Hello: 1",
		"Long Hello: 2",
		"Short Formatted Hello: 0",
		"Short Formatted Hello: 1",
		"Short Formatted Hello: 2",
		"Long Formatted Hello: 0",
		"Long Formatted Hello: 1",
		"Long Formatted Hello: 2",
		"A Short Local Test: 1,234",
		"A Short Local Formatted Test: 2,345",
		"A Long Local Test: 3,456",
		"A Long Local Formatted Test: 4,567",
	)
}

//nolint:funlen // Ok
func TestSzLog_Verbose_V3(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetVerbose(3)
	chk.NoErr(szlog.SetLanguage("en"))

	szlog.S0("Short Hello: 0\n")
	szlog.S1("Short Hello: 1\n")
	szlog.S2("Short Hello: 2\n")
	szlog.S3("Short Hello: 3\n")
	szlog.S4("Short Hello: 4\n")
	szlog.S5("Short Hello: 5\n")

	szlog.Say0("Long Hello: 0\n")
	szlog.Say1("Long Hello: 1\n")
	szlog.Say2("Long Hello: 2\n")
	szlog.Say3("Long Hello: 3\n")
	szlog.Say4("Long Hello: 4\n")
	szlog.Say5("Long Hello: 5\n")

	szlog.S0f("Short Formatted Hello: %d\n", 0)
	szlog.S1f("Short Formatted Hello: %d\n", 1)
	szlog.S2f("Short Formatted Hello: %d\n", 2)
	szlog.S3f("Short Formatted Hello: %d\n", 3)
	szlog.S4f("Short Formatted Hello: %d\n", 4)
	szlog.S5f("Short Formatted Hello: %d\n", 5)

	szlog.Say0f("Long Formatted Hello: %d\n", 0)
	szlog.Say1f("Long Formatted Hello: %d\n", 1)
	szlog.Say2f("Long Formatted Hello: %d\n", 2)
	szlog.Say3f("Long Formatted Hello: %d\n", 3)
	szlog.Say4f("Long Formatted Hello: %d\n", 4)
	szlog.Say5f("Long Formatted Hello: %d\n", 5)

	szlog.S0("A Short Local Test: ", 1234, "\n")
	szlog.S0f("A Short Local Formatted Test: %d\n", 2345)
	szlog.Say0("A Long Local Test: ", 3456, "\n")
	szlog.Say0f("A Long Local Formatted Test: %d\n", 4567)

	chk.Log()
	chk.Stdout(
		"Short Hello: 0",
		"Short Hello: 1",
		"Short Hello: 2",
		"Short Hello: 3",
		"Long Hello: 0",
		"Long Hello: 1",
		"Long Hello: 2",
		"Long Hello: 3",
		"Short Formatted Hello: 0",
		"Short Formatted Hello: 1",
		"Short Formatted Hello: 2",
		"Short Formatted Hello: 3",
		"Long Formatted Hello: 0",
		"Long Formatted Hello: 1",
		"Long Formatted Hello: 2",
		"Long Formatted Hello: 3",
		"A Short Local Test: 1,234",
		"A Short Local Formatted Test: 2,345",
		"A Long Local Test: 3,456",
		"A Long Local Formatted Test: 4,567",
	)
}

//nolint:funlen // Ok
func TestSzLog_Verbose_V4(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetVerbose(4)
	chk.NoErr(szlog.SetLanguage("en"))

	szlog.S0("Short Hello: 0\n")
	szlog.S1("Short Hello: 1\n")
	szlog.S2("Short Hello: 2\n")
	szlog.S3("Short Hello: 3\n")
	szlog.S4("Short Hello: 4\n")
	szlog.S5("Short Hello: 5\n")

	szlog.Say0("Long Hello: 0\n")
	szlog.Say1("Long Hello: 1\n")
	szlog.Say2("Long Hello: 2\n")
	szlog.Say3("Long Hello: 3\n")
	szlog.Say4("Long Hello: 4\n")
	szlog.Say5("Long Hello: 5\n")

	szlog.S0f("Short Formatted Hello: %d\n", 0)
	szlog.S1f("Short Formatted Hello: %d\n", 1)
	szlog.S2f("Short Formatted Hello: %d\n", 2)
	szlog.S3f("Short Formatted Hello: %d\n", 3)
	szlog.S4f("Short Formatted Hello: %d\n", 4)
	szlog.S5f("Short Formatted Hello: %d\n", 5)

	szlog.Say0f("Long Formatted Hello: %d\n", 0)
	szlog.Say1f("Long Formatted Hello: %d\n", 1)
	szlog.Say2f("Long Formatted Hello: %d\n", 2)
	szlog.Say3f("Long Formatted Hello: %d\n", 3)
	szlog.Say4f("Long Formatted Hello: %d\n", 4)
	szlog.Say5f("Long Formatted Hello: %d\n", 5)

	szlog.S0("A Short Local Test: ", 1234, "\n")
	szlog.S0f("A Short Local Formatted Test: %d\n", 2345)
	szlog.Say0("A Long Local Test: ", 3456, "\n")
	szlog.Say0f("A Long Local Formatted Test: %d\n", 4567)

	chk.Log()
	chk.Stdout(
		"Short Hello: 0",
		"Short Hello: 1",
		"Short Hello: 2",
		"Short Hello: 3",
		"Short Hello: 4",
		"Long Hello: 0",
		"Long Hello: 1",
		"Long Hello: 2",
		"Long Hello: 3",
		"Long Hello: 4",
		"Short Formatted Hello: 0",
		"Short Formatted Hello: 1",
		"Short Formatted Hello: 2",
		"Short Formatted Hello: 3",
		"Short Formatted Hello: 4",
		"Long Formatted Hello: 0",
		"Long Formatted Hello: 1",
		"Long Formatted Hello: 2",
		"Long Formatted Hello: 3",
		"Long Formatted Hello: 4",
		"A Short Local Test: 1,234",
		"A Short Local Formatted Test: 2,345",
		"A Long Local Test: 3,456",
		"A Long Local Formatted Test: 4,567",
	)
}

//nolint:funlen // Ok
func TestSzLog_Verbose_V5(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetVerbose(5)
	chk.NoErr(szlog.SetLanguage("en"))

	szlog.S0("Short Hello: 0\n")
	szlog.S1("Short Hello: 1\n")
	szlog.S2("Short Hello: 2\n")
	szlog.S3("Short Hello: 3\n")
	szlog.S4("Short Hello: 4\n")
	szlog.S5("Short Hello: 5\n")

	szlog.Say0("Long Hello: 0\n")
	szlog.Say1("Long Hello: 1\n")
	szlog.Say2("Long Hello: 2\n")
	szlog.Say3("Long Hello: 3\n")
	szlog.Say4("Long Hello: 4\n")
	szlog.Say5("Long Hello: 5\n")

	szlog.S0f("Short Formatted Hello: %d\n", 0)
	szlog.S1f("Short Formatted Hello: %d\n", 1)
	szlog.S2f("Short Formatted Hello: %d\n", 2)
	szlog.S3f("Short Formatted Hello: %d\n", 3)
	szlog.S4f("Short Formatted Hello: %d\n", 4)
	szlog.S5f("Short Formatted Hello: %d\n", 5)

	szlog.Say0f("Long Formatted Hello: %d\n", 0)
	szlog.Say1f("Long Formatted Hello: %d\n", 1)
	szlog.Say2f("Long Formatted Hello: %d\n", 2)
	szlog.Say3f("Long Formatted Hello: %d\n", 3)
	szlog.Say4f("Long Formatted Hello: %d\n", 4)
	szlog.Say5f("Long Formatted Hello: %d\n", 5)

	szlog.S0("A Short Local Test: ", 1234, "\n")
	szlog.S0f("A Short Local Formatted Test: %d\n", 2345)
	szlog.Say0("A Long Local Test: ", 3456, "\n")
	szlog.Say0f("A Long Local Formatted Test: %d\n", 4567)

	chk.Log()
	chk.Stdout(
		"Short Hello: 0",
		"Short Hello: 1",
		"Short Hello: 2",
		"Short Hello: 3",
		"Short Hello: 4",
		"Short Hello: 5",
		"Long Hello: 0",
		"Long Hello: 1",
		"Long Hello: 2",
		"Long Hello: 3",
		"Long Hello: 4",
		"Long Hello: 5",
		"Short Formatted Hello: 0",
		"Short Formatted Hello: 1",
		"Short Formatted Hello: 2",
		"Short Formatted Hello: 3",
		"Short Formatted Hello: 4",
		"Short Formatted Hello: 5",
		"Long Formatted Hello: 0",
		"Long Formatted Hello: 1",
		"Long Formatted Hello: 2",
		"Long Formatted Hello: 3",
		"Long Formatted Hello: 4",
		"Long Formatted Hello: 5",
		"A Short Local Test: 1,234",
		"A Short Local Formatted Test: 2,345",
		"A Long Local Test: 3,456",
		"A Long Local Formatted Test: 4,567",
	)
}
