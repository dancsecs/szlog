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
	"bytes"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

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

func TestSzLog_SetStdout(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	szlog.Reset()

	szlog.Say0("This should go to os.Stdout before change\n")

	var testOutput bytes.Buffer

	szlog.SetStdout(&testOutput)

	szlog.Say0("This should go to bytes.Buffer")

	chk.ByteSlice(
		testOutput.Bytes(),
		[]byte("This should go to bytes.Buffer"),
	)

	szlog.SetStdout(nil)

	szlog.Say0("This should go to os.Stdout after change\n")

	chk.Log()
	chk.Stdout(
		"This should go to os.Stdout before change",
		"This should go to os.Stdout after change",
	)
}
