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
	errArg   = errors.New("unexpected cleaned args")
	errLevel = errors.New("unexpected szlog level")
)

func testArgs(expectedLevel szlog.LogLevel, args []string) error {
	var (
		err   error
		cArgs []string
	)

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	cArgs = szlog.VerboseAbsorbArgs(args)

	if len(cArgs) != 5 ||
		cArgs[0] != cv1 ||
		cArgs[1] != cv2 ||
		cArgs[2] != nov ||
		cArgs[3] != cv3 ||
		cArgs[4] != cv4 {
		//
		err = fmt.Errorf("%w: %v", errArg, cArgs)
	} else if szlog.Level() != expectedLevel {
		err = fmt.Errorf("%w: %s", errLevel, szlog.Level().String())
	}

	return err
}

//nolint:funlen // Ok.
func TestLog_ArgumentAbsorption(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	chk.Int(len(szlog.VerboseAbsorbArgs(nil)), 0)

	chk.NoErr(
		testArgs(szlog.LevelNone, []string{
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelFatal, []string{
			"-v",
			cv1,
			cv2,
			nov,
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelError, []string{
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
		testArgs(szlog.LevelInfo, []string{
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
		testArgs(szlog.LevelTrace, []string{
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
		testArgs(szlog.LevelAll, []string{
			"-v",
			cv1,
			"--v",
			cv2,
			"-vv",
			nov,
			cv3,
			"--vv",
			cv4,
			"-v",
		}),
	)

	chk.Log()
}
