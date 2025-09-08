/*
   Szerszam alarm manager: szalarm.
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

//nolint:funlen // Ok.
package szlog_test

import (
	"errors"
	"fmt"
	"log"
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
	errArg        = errors.New("unexpected cleaned args")
	errLevel      = errors.New("unexpected szlog level")
	errVerbose    = errors.New("unexpected verbose level")
	errLanguage   = errors.New("unexpected language local")
	errLongLabels = errors.New("unexpected long labels")
)

//nolint:cyclop // Ok.
func testArgs(
	expectedLevel szlog.LogLevel,
	expectedVerbose szlog.VerboseLevel,
	expectedLanguage string,
	expectedLongLabels bool,
	args []string,
) error {
	var (
		err   error
		cArgs []string
	)

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	cArgs, err = szlog.AbsorbArgs(args, nil)

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
		case szlog.Language() != expectedLanguage:
			err = fmt.Errorf(
				"%w: Want: %s Got: %s", errLanguage,
				expectedLanguage,
				szlog.Language(),
			)
		case szlog.LongLabels() != expectedLongLabels:
			err = fmt.Errorf(
				"%w: Want: %t Got: %t", errLongLabels,
				expectedLongLabels,
				szlog.LongLabels(),
			)
		}
	}

	return err
}

func TestSzLog_LanguageArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()
	defer szlog.Reset()

	cleanedArgs, err := szlog.AbsorbArgs(nil, nil)
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", false, []string{
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
		testArgs(szlog.LevelError, 0, "", false, []string{
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
		testArgs(szlog.LevelError, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language", "",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "en", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--language", "en",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "fr", false, []string{
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
	defer szlog.Reset()

	cleanedArgs, err := szlog.AbsorbArgs(nil, nil)
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", false, []string{
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
		testArgs(szlog.LevelError, 0, "", false, []string{
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
		testArgs(szlog.LevelError, 0, "", false, []string{
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
		testArgs(szlog.LevelError, 0, "", false, []string{
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
		testArgs(szlog.LevelFatal, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "fatal",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelNone, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "none",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "error",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelWarn, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "warn",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelInfo, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "info",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelDebug, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "debug",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelTrace, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--log", "trace",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelAll, 0, "", false, []string{
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
	defer szlog.Reset()

	chk.Err(
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
		testArgs(szlog.LevelCustom, 0, "", false, []string{
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
	defer szlog.Reset()

	cleanedArgs, err := szlog.AbsorbArgs(nil, nil)
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, -1, "", false, []string{
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
		testArgs(szlog.LevelError, -1, "", false, []string{
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
		testArgs(szlog.LevelError, -1, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--quiet",
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "", false, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 1, "", false, []string{
			"-v",
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 2, "", false, []string{
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
		testArgs(szlog.LevelError, 4, "", false, []string{
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
		testArgs(szlog.LevelError, 6, "", false, []string{
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
		testArgs(szlog.LevelError, 7, "", false, []string{
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

func TestSzLog_LongLabelsArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()
	defer szlog.Reset()

	cleanedArgs, err := szlog.AbsorbArgs(nil, func(flag, desc string) {
		log.Printf("%s: %s", flag, desc)
	})
	chk.NoErr(err)
	chk.Int(len(cleanedArgs), 0)

	chk.Err(
		testArgs(szlog.LevelError, 0, "", true, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--long-labels",
			"--long-labels",
		}),
		chk.ErrChain(
			szlog.ErrAmbiguousLongLabels,
		),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, 0, "", true, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
			"--long-labels",
		}),
	)

	chk.Log(
		"-v[v...] | --v[v...] | --verbose: Increase the logging level for each v provided.",
		"--quiet: Sets the verbose level to -1 squashing all (non-logged) output.",
		"--log <level | (levels)>: Set the level to log (or a custom combination of levels).",
		"--language: Sets the local language used for formatting.",
		"--long-labels: Use long labels in log output.",
	)
}
