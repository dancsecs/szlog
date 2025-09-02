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
	"io"
)

// Close is a convenience method for safely closing any io.Closer.
// If an error occurs during Close, it is logged as a warning.
// This method is primarily intended for use in defer statements.
func (l *Log) Close(area string, closeable io.Closer) {
	err := closeable.Close()
	if err != nil && l.LogWarn {
		l.Warn(area, " caused: ", err)
	}
}
