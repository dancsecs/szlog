package szlog

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Language return the current local language string.
func (l *Log) Language() string {
	return l.language
}

// SetLanguage creates a printer that attempts to format data in a local
// manner. Pass an empty "" string to stop local formatting.
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
