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
	"strings"
)

// Usage suitable strings for verbose argument absorption.
const (
	VerboseFlag     = "-v[v...] | --v[v...] | --verbose"
	VerboseFlagDesc = "Increase the logging level for each v provided."

	QuietFlag     = "--quiet"
	QuietFlagDesc = "Sets the verbose level to -1 squashing all " +
		"(non-logged) output."

	LogLevelFlag     = "--log <level | (levels)>"
	LogLevelFlagDesc = "Set the level to log (or a custom combination of " +
		"levels)."

	LanguageFlag     = "--language"
	LanguageFlagDesc = "Sets the local language used for formatting."

	LongLabelFlag     = "--long-labels"
	LongLabelFlagDesc = "Use long labels in log output."
)

// AbsorbArgs scans the provided argument list for logging-related flags.
// It updates the log configuration (LogLevel, verbosity, quiet mode,
// LongLabels, and Language) based on the flags encountered. Recognized
// flags are removed, and the cleaned argument slice is returned.
// Multiple `-v` flags increment verbosity accordingly. If conflicting
// or invalid flags are found (e.g., combining `-v` with `--quiet`),
// an error is returned along with the original arguments.  Optionally
// a function that registers argument flags and their description can
// be provided (usually for usage information.)
//
//nolint:gocognit,cyclop,funlen // OK.
func (l *Log) AbsorbArgs(
	argsIn []string, registerArgs func(string, string),
) ([]string, error) {
	err := error(nil)
	captureLogLevel := false
	captureLanguage := false
	logLevelSet := false
	languageSet := false
	longLabelsSet := false
	vCount := VerboseLevel(0)
	argsOut := make([]string, 0, len(argsIn))

	if registerArgs != nil {
		registerArgs(VerboseFlag, VerboseFlagDesc)
		registerArgs(QuietFlag, QuietFlagDesc)
		registerArgs(LogLevelFlag, LogLevelFlagDesc)
		registerArgs(LanguageFlag, LanguageFlagDesc)
		registerArgs(LongLabelFlag, LongLabelFlagDesc)
	}

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

		if rArg == "--long-labels" {
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
