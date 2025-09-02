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
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Language returns the current language setting used for localized formatting.
// An empty string indicates no localization is applied.
func (l *Log) Language() string {
	return l.language
}

// SetLanguage updates the language used for localized formatting.
// Passing an empty string ("") disables localization. It returns any
// error encountered while setting the language.
func (l *Log) SetLanguage(langStr string) error {
	var (
		languageTag language.Tag
		err         error
	)

	l.printer = nil
	l.language = ""

	if langStr != "" {
		languageTag, err = language.Parse(langStr)
		if err == nil {
			l.printer = message.NewPrinter(languageTag)
			l.language = langStr
		}
	}

	if err == nil {
		return nil
	}

	return fmt.Errorf("%w: '%s'", ErrInvalidLanguage, langStr)
}
