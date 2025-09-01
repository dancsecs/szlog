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
	VerboseFlag     = "-v[v...] | --v[v...] | --verbose"
	VerboseFlagDesc = "Increase the logging level for each v provided."

	LogLevelFlag     = "--log <level | (levels)>"
	LogLevelFlagDesc = "Set the level to log (or a custom combination of " +
		"levels)"

	QuietFlag     = "--quiet"
	QuietFlagDesc = "Sets the verbose level to -1 squashing all " +
		"(non-logged) output"

	LanguageFlag     = "--language"
	LanguageFlagDesc = "Sets the local language used for formatting."
)

// AbsorbArgs scans an argument list for verbose flags increasing
// the log level for each verbose flag encountered.  These flags are removed
// and a cleaned up arg list is returned.  Verbose flags can be a single (or
// multiple letter 'v's with the corresponding number of log level increments
// made.  If an error is encountered it is returned with the original
// unchanged argument slice.
//
//nolint:gocognit,cyclop,funlen // OK.
func (l *Log) AbsorbArgs(argsIn []string) ([]string, error) {
	err := error(nil)
	captureLogLevel := false
	captureLanguage := false
	logLevelSet := false
	languageSet := false
	vCount := VerboseLevel(0)
	argsOut := make([]string, 0, len(argsIn))

	for _, rArg := range argsIn {
		if err != nil {
			break
		}

		if rArg == "--log" {
			if logLevelSet {
				err = ErrAmbiguousLogLevel
			} else {
				logLevelSet = true
				captureLogLevel = true
			}

			continue
		}

		if captureLogLevel {
			captureLogLevel = false
			err = l.parseAndSetLevel(rArg, LevelError)

			continue
		}

		if rArg == "--language" {
			captureLanguage = true

			continue
		}

		if captureLanguage {
			if languageSet {
				err = ErrAmbiguousLanguage
			} else {
				languageSet = true
				captureLanguage = false

				err = l.SetLanguage(rArg)
			}

			continue
		}

		if rArg == "--quiet" {
			if vCount > 0 { // Already processed a --quiet argument.
				err = ErrAmbiguousVerboseAndQuiet
			}

			vCount = -1

			continue
		}

		if strings.HasPrefix(rArg, "-v") || strings.HasPrefix(rArg, "--v") {
			if rArg == "--verbose" {
				rArg = "-v"
			}

			isVerbose := true
			cleanArg := strings.TrimLeft(rArg, "-")

			for i, mi := 0, len(cleanArg); i < mi && isVerbose; i++ {
				isVerbose = cleanArg[i] == 'v'
			}

			if isVerbose {
				if vCount == -1 { // Already processed a --quiet argument.
					err = ErrAmbiguousVerboseAndQuiet
				}

				vCount += len(cleanArg)

				continue
			}
		}

		argsOut = append(argsOut, rArg)
	}

	if err == nil && captureLogLevel {
		err = ErrMissingLogLevel
	}

	if err == nil && captureLanguage {
		err = ErrMissingLanguage
	}

	// Final reset level to adjust verbose settings.
	l.SetVerbose(vCount)

	if !logLevelSet {
		l.SetLevel(LevelError)
	}

	if err == nil {
		return argsOut, nil
	}

	return argsIn, err
}
