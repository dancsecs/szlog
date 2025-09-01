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
