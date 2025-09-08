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
	"errors"
)

// Exported errors.
var (
	ErrUnknownLevel             = errors.New("unknown log level")
	ErrInvalidLogLevelParse     = errors.New("invalid log level string")
	ErrAmbiguousVerboseAndQuiet = errors.New("ambiguous verbose and quiet")
	ErrAmbiguousLogLevel        = errors.New("ambiguous log level")
	ErrAmbiguousLanguage        = errors.New("ambiguous language")
	ErrAmbiguousLongLabels      = errors.New("ambiguous long labels")
	ErrMissingLogLevel          = errors.New("missing log level")
	ErrMissingLanguage          = errors.New("missing language")
	ErrInvalidLanguage          = errors.New("invalid language")
)
