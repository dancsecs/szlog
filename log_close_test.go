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

package szlog_test

import (
	"errors"
	"os"
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

type tstCloseable struct {
	err error
}

func (t tstCloseable) Close() error {
	return t.err
}

var errCloseError = errors.New("close error")

func TestSzLog_Close(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.SetLevel(szlog.LevelAll)

	noErrClose := tstCloseable{
		err: nil,
	}

	errClose := tstCloseable{
		err: errCloseError,
	}

	szlog.Close("no error", noErrClose)

	szlog.Close("with error", errClose)

	chk.Log(
		"W:Closing: with error caused: close error",
	)
}

func TestSzLog_AlreadyClosed(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	szlog.SetLevel(szlog.LevelAll)

	file := chk.CreateTmpFile(nil)

	fHandle, err := os.Open(file) //nolint:gosec // Ok to test.
	chk.NoErr(err)

	chk.NoErr(fHandle.Close())

	chk.Err(
		fHandle.Close(),
		chk.ErrChain(
			"close "+file,
			os.ErrClosed,
		),
	)

	// do it again to make sure we still get the same error.
	chk.Err(
		fHandle.Close(),
		chk.ErrChain(
			"close "+file,
			os.ErrClosed,
		),
	)

	// Already closed error should be ignored.
	szlog.Close("should not be displayed", fHandle)

	chk.Log()
}
