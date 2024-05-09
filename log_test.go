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

//nolint:dupl // Ok.
package szlog_test

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

const (
	logFatalLabel = "F:"
	logErrorLabel = "E:"
	logWarnLabel  = "W:"
	logInfoLabel  = "I:"
	logDebugLabel = "D:"
	logTraceLabel = "T:"

	logFatalLongLabel = "FATAL:"
	logErrorLongLabel = "ERROR:"
	logWarnLongLabel  = "WARN:"
	logInfoLongLabel  = "INFO:"
	logDebugLongLabel = "DEBUG:"
	logTraceLongLabel = "TRACE:"
)

var (
	errArg   = errors.New("unexpected cleaned args")
	errLevel = errors.New("unexpected szlog level")
)

func calculation(numbers []float64) string {
	var sum float64

	log.Printf("calculating average")

	for _, n := range numbers {
		sum += n
	}

	avg := sum / float64(len(numbers))

	return strconv.FormatFloat(avg, 'g', -1, 64)
}

//nolint:cyclop,funlen // Ok.
func tstWrite(
	label string,
	unFmt szlog.LogFunctionUnformatted,
	fmt szlog.LogFunctionFormatted,
) []string {
	var res []string

	// Test Empty args.  Should produce empty log lines.
	if unFmt() {
		res = append(res, label) // Empty line.
	}

	if fmt("") {
		res = append(res, label) // Empty Line.
	}

	// Test Single Empty arg.  Should produce empty log lines.
	if unFmt("") {
		res = append(res, label) // Empty Line.
	}

	if fmt("%s", "") {
		res = append(res, label) // Empty Line.
	}

	// Test Single non-empty arg.  Should produce log lines.
	if unFmt("argUnFormatted") {
		res = append(res, label+"argUnFormatted")
	}

	if fmt("arg%s", "Formatted") {
		res = append(res, label+"argFormatted")
	}

	const avgStr = "Avg:3"

	tstArray := []float64{1, 2, 3, 4, 5}

	// Test Single non-delayed function.  Should produce log lines.

	res = append(res, "calculating average") // Always Called.
	if unFmt("Avg:", calculation(tstArray)) {
		res = append(res, label+avgStr)
	}

	res = append(res, "calculating average") // Always Called.
	if fmt("Avg:%s", calculation(tstArray)) {
		res = append(res, label+avgStr)
	}

	// Test Single delayed function.  Should produce log lines.
	if unFmt(
		"Avg:",
		func() szlog.Def {
			return szlog.Def(calculation(tstArray))
		},
	) {
		res = append(res, "calculating average") // Deferred Call.
		res = append(res, label+avgStr)
	}

	if fmt(
		"Avg:%s",
		func() szlog.Def {
			return szlog.Def(calculation(tstArray))
		},
	) {
		res = append(res, "calculating average") // Deferred Call.
		res = append(res, label+avgStr)
	}

	return res
}

func logAndBuildExpectedResult() []string {
	var res []string

	logger := szlog.Default()

	res = append(res, tstWrite(logTraceLabel, logger.T, logger.Tf)...)
	res = append(res, tstWrite(logTraceLabel, logger.Trace, logger.Tracef)...)

	res = append(res, tstWrite(logDebugLabel, logger.D, logger.Df)...)
	res = append(res, tstWrite(logDebugLabel, logger.Debug, logger.Debugf)...)

	res = append(res, tstWrite(logInfoLabel, logger.I, logger.If)...)
	res = append(res, tstWrite(logInfoLabel, logger.Info, logger.Infof)...)

	res = append(res, tstWrite(logWarnLabel, logger.W, logger.Wf)...)
	res = append(res, tstWrite(logWarnLabel, logger.Warn, logger.Warnf)...)

	res = append(res, tstWrite(logErrorLabel, logger.E, logger.Ef)...)
	res = append(res, tstWrite(logErrorLabel, logger.Error, logger.Errorf)...)

	res = append(res, tstWrite(logFatalLabel, logger.F, logger.Ff)...)
	res = append(res, tstWrite(logFatalLabel, logger.Fatal, logger.Fatalf)...)

	return res
}

func logAndBuildLongExpectedResult() []string {
	var res []string

	logger := szlog.Default()

	res = append(res, tstWrite(logTraceLongLabel, logger.T, logger.Tf)...)
	res = append(res, tstWrite(
		logTraceLongLabel, logger.Trace, logger.Tracef)...,
	)

	res = append(res, tstWrite(logDebugLongLabel, logger.D, logger.Df)...)
	res = append(res, tstWrite(
		logDebugLongLabel, logger.Debug, logger.Debugf)...,
	)

	res = append(res, tstWrite(logInfoLongLabel, logger.I, logger.If)...)
	res = append(res, tstWrite(logInfoLongLabel, logger.Info, logger.Infof)...)

	res = append(res, tstWrite(logWarnLongLabel, logger.W, logger.Wf)...)
	res = append(res, tstWrite(logWarnLongLabel, logger.Warn, logger.Warnf)...)

	res = append(res, tstWrite(logErrorLongLabel, logger.E, logger.Ef)...)
	res = append(res, tstWrite(
		logErrorLongLabel, logger.Error, logger.Errorf)...,
	)

	res = append(res, tstWrite(logFatalLongLabel, logger.F, logger.Ff)...)
	res = append(res, tstWrite(
		logFatalLongLabel, logger.Fatal, logger.Fatalf)...,
	)

	return res
}

func TestCfg_LogError_Default(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	origLog := szlog.Default()
	tstLog := szlog.New(nil)

	oldLog := szlog.SetDefault(tstLog)

	tstLog2 := szlog.SetDefault(origLog)

	chk.True(oldLog == origLog)
	chk.True(tstLog == tstLog2)
	chk.True(szlog.Default() == origLog)

	szlog.Reset()

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 72, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_None(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 24, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_Fatal(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelFatal)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 48, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_Error(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelError)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 72, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_Warn(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelWarn)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 96, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_Info(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelInfo)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 120, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_Debug(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelDebug)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 144, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogError_Trace(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelTrace)

	expLog := logAndBuildExpectedResult()

	chk.Int(len(expLog), 168, "Unexpected number of lines")
	chk.Log(expLog...)
}

func TestCfg_LogLabelLength(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.Reset()
	szlog.SetLevel(szlog.LevelAll)

	expLog := logAndBuildExpectedResult()

	chk.False(szlog.LongLabels())

	chk.False(szlog.SetLongLabels(true))
	chk.True(szlog.LongLabels())

	expLog = append(expLog, logAndBuildLongExpectedResult()...)

	chk.True(szlog.SetLongLabels(false))

	expLog = append(expLog, logAndBuildExpectedResult()...)

	chk.Log(expLog...)
}

const (
	cv1 = "-v1"
	cv2 = "--v2"
	cv3 = "-vvv1x"
	cv4 = "--vvv2x"
)

func testArgs(expectedLevel szlog.LogLevel, args []string) error {
	var (
		err   error
		cArgs []string
	)

	szlog.Reset()
	szlog.SetLevel(szlog.LevelNone)

	cArgs = szlog.VerboseAbsorbArgs(args)

	if len(cArgs) != 4 ||
		cArgs[0] != cv1 ||
		cArgs[1] != cv2 ||
		cArgs[2] != cv3 ||
		cArgs[3] != cv4 {
		//
		err = fmt.Errorf("%w: %v", errArg, cArgs)
	} else if szlog.Level() != expectedLevel {
		err = fmt.Errorf("%w: %d", errLevel, szlog.Level())
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
			cv3,
			cv4,
		}),
	)

	chk.NoErr(
		testArgs(szlog.LevelFatal, []string{
			"-v",
			cv1,
			cv2,
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
			cv3,
			"--vv",
			cv4,
			"-v",
		}),
	)

	chk.Log()
}
