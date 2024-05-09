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

/*
Package szlog provides a layer on top of the standard go log package enabling
multiple levels of logging as follows.

-LevelFatal
-LevelError
-LevelWarn
-LevelInfo
-LevelDebug
-LevelTrace

Once a level is set (It defaults to LevelError) then all messages for that
level and all levels above it will be processed.  Levels below will not
be processed.

Further any individual level can be disabled independent of the current log
level set.
*/
package szlog
