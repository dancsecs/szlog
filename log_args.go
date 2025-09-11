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
	"math"
	"strings"
)

// Usage suitable strings for verbose argument absorption.
const (
	VerboseFlag     = "[-v | --verbose ...]"
	VerboseFlagDesc = "Increase the verbose level for each v provided."

	QuietFlag     = "[--quiet]"
	QuietFlagDesc = "Sets the verbose level to -1 squashing all " +
		"(non-logged) output."

	LogLevelFlag     = "[--log <level | (levels)>]"
	LogLevelFlagDesc = "Set the level to log (or a custom combination of " +
		"levels).  Valid levels are: " +
		"None, FATAL, ERROR, WARN, INFO, DEBUG,TRACE, ALL."

	LanguageFlag     = "[--language <lang>]"
	LanguageFlagDesc = "Sets the local language used for formatting."

	LongLabelFlag     = "[--long-labels]"
	LongLabelFlagDesc = "Use long labels in log output."
)

// EnableArg specifies which command-line arguments szlog should recognize
// and process. These values can be combined and passed to AbsorbArgs to
// restrict handling to a subset of arguments. If no values are provided,
// EnableAll is assumed by default.
type EnableArg int

// EnableArg flags define which arguments are recognized by AbsorbArgs and
// ArgUsageInfo. Multiple values can be provided (or combined with a
// bitwise OR.)
//
//	EnableVerbose    - enable verbosity flags (-v, --verbose)
//	EnableQuiet      - enable quiet flag (--quiet)
//	EnableLogLevel   - enable log level flag (--log <level>)
//	EnableLanguage   - enable language/locale flag (--language <locale>)
//	EnableLongLabels - enable long-labels flag (--long-labels)
//	EnableAll        - enable all supported argument flags
const (
	EnableVerbose EnableArg = 1 << iota
	EnableQuiet
	EnableLogLevel
	EnableLanguage
	EnableLongLabels

	EnableAll EnableArg = math.MaxInt
)

// AbsorbArgs scans the provided argument list for enabled logging-related
// flags and updates the log configuration accordingly. Only arguments
// specified in the enable set are recognized; all others are ignored.
// Recognized flags are removed, and the cleaned argument slice is returned.
// Multiple `-v` flags increment verbosity, while invalid or conflicting
// combinations (e.g., `-v` with `--quiet`) return an error along with
// the original arguments. If no enable set is provided, EnableAll is used.
//
//nolint:gocognit,cyclop,gocyclo,funlen // complexity is intentional.
func (l *Log) AbsorbArgs(
	argsIn []string, enable ...EnableArg,
) ([]string, error) {
	err := error(nil)
	captureLogLevel := false
	captureLanguage := false
	logLevelSet := false
	languageSet := false
	longLabelsSet := false
	vCount := VerboseLevel(0)
	argsOut := make([]string, 0, len(argsIn))

	if len(enable) == 0 {
		l.argsEnabled = EnableAll
	} else {
		l.argsEnabled = 0
		for _, enArg := range enable {
			l.argsEnabled |= enArg
		}
	}

	for _, rArg := range argsIn {
		if err != nil {
			break
		}

		if rArg == "--log" && l.argsEnabled&EnableLogLevel > 0 {
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

		if rArg == "--language" && l.argsEnabled&EnableLanguage > 0 {
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

		if rArg == "--quiet" && l.argsEnabled&EnableQuiet > 0 {
			if vCount > 0 { // Already processed a --quiet argument.
				err = ErrAmbiguousVerboseAndQuiet
			}

			vCount = -1

			continue
		}

		if rArg == "--verbose" && l.argsEnabled&EnableVerbose > 0 {
			rArg = "-v"
		}

		if strings.HasPrefix(rArg, "-v") && l.argsEnabled&EnableVerbose > 0 {
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

		if rArg == "--long-labels" && l.argsEnabled&EnableLongLabels > 0 {
			if longLabelsSet {
				err = ErrAmbiguousLongLabels
			} else {
				longLabelsSet = true
				l.SetLongLabels(true)
			}

			continue
		}

		argsOut = append(argsOut, rArg)
	}

	if err == nil && captureLogLevel {
		err = ErrMissingLogLevel
	}

	if err == nil && captureLanguage {
		err = ErrMissingLanguage
	}

	if l.argsEnabled&EnableVerbose > 0 {
		// Final reset level to adjust verbose settings.
		l.SetVerbose(vCount)
	}

	if !logLevelSet && l.argsEnabled&EnableLogLevel > 0 {
		l.SetLevel(LevelError)
	}

	if err == nil {
		return argsOut, nil
	}

	return argsIn, err
}

// ArgUsageInfo reports usage information for all enabled arguments by
// invoking the provided callback for each one. Only arguments permitted
// in the enable set are included, allowing applications to present
// accurate help/usage output tailored to their configuration.
func (l *Log) ArgUsageInfo(registerArg func(string, string)) {
	if l.argsEnabled&EnableVerbose > 0 {
		registerArg(VerboseFlag, VerboseFlagDesc)
	}

	if l.argsEnabled&EnableQuiet > 0 {
		registerArg(QuietFlag, QuietFlagDesc)
	}

	if l.argsEnabled&EnableLogLevel > 0 {
		registerArg(LogLevelFlag, LogLevelFlagDesc)
	}

	if l.argsEnabled&EnableLanguage > 0 {
		registerArg(LanguageFlag, LanguageFlagDesc)
	}

	if l.argsEnabled&EnableLongLabels > 0 {
		registerArg(LongLabelFlag, LongLabelFlagDesc)
	}
}
