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
	"strings"
)

// Usage suitable strings for verbose argument absorption.
const (
	VerboseFlag     = "-v[v...], --v[v...]"
	VerboseFlagDesc = "Increase the logging level for each v provided."
)

// VerboseAbsorbArgs scans an argument list for verbose flags increasing
// the log level for each verbose flag encountered.  These flags are removed
// and a cleaned up arg list is returned.  Verbose flags can be a single (or
// multiple letter 'v's with the corresponding number of log level increments
// made.
func (l *Log) VerboseAbsorbArgs(argsIn []string) []string {
	argsOut := make([]string, 0, len(argsIn))

	for _, rArg := range argsIn {
		keepArg := true

		if strings.HasPrefix(rArg, "-v") || strings.HasPrefix(rArg, "--v") {
			keepArg = false
			cleanArg := strings.TrimLeft(rArg, "-")

			for i, mi := 0, len(cleanArg); i < mi && !keepArg; i++ {
				keepArg = cleanArg[i] != 'v'
			}

			if !keepArg {
				for range len(cleanArg) {
					l.IncLevel()
				}
			}
		}

		if keepArg {
			argsOut = append(argsOut, rArg)
		}
	}

	return argsOut
}
