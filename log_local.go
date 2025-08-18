package szlog

import (
	"golang.org/x/text/message"
)

// SetLocal creates a printer that attempts to format data in a local manner.
// Pass an empty "" string to stop local formatting.
func (l *Log) SetLocal(language string) {
	if language == "" {
		l.printer = nil
	} else {
		l.printer = message.NewPrinter(message.MatchLanguage(language))
	}
}
