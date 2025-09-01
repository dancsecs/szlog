<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

<!---
   Szerszam alarm manager: szalarm.
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
-->

# Package szlog

<!--- gotomd::Bgn::doc::./package -->
```go
package szlog
```

Package szlog is a flexible output and logging library for Go. It builds on
the standard log while adding:

  - six standard logging levels for categorizing messages by severity (Fatal,
    Error, Warn, Info, Debug, Trace with Fatalf, Errorf, Warnf, Infof, Debugf,
    Tracef for formatting) layered on top of the standard log package

  - six independent verbosity levels (Say0–Say4 with Say0f–Say4f for
    formatting) with settable output defaulting to ```os.Stdout```

  - lazy evaluation with ```func() DEFER```` so expensive values are only
    computed when needed

  - efficient function swapping so disabled functions are minimal

  - local formatting through the Go text package by setting a locale string
    ("en", "en-CA", "fr", etc.)

## Overview

Szout provides two separate independent families of output functions.  Each
function is a variable member that may be swapped at runtime. When disabled,
a function is replaced with a no-op and incurs minimal runtime cost. When
enabled, the function points to an optimized writer.  Each function mirrors
standard fmt and log package function Print and Printf.  Two versions of each
function is supplied with long and short identifiers.

It further offers two ways to output messages. For full control, create a *Log
instance with New(), giving you an independent logger with its own output,
levels, and settings. For quick and convenient logging, use the package-level
functions (Info, Error, Warn, etc.), which operate through a default internal
logger automatically created for you. Both provide the same API; choose
instance-based loggers when you need separate configurations, or rely on the
package-level functions for simple, global logging.

An instance can be created as follows:
```go

// New returns a new Log instance with default settings.
func New() *Log

```

While any instance or the default inner log can be reset to initial conditions
with the Reset method as follows:

```go

// Reset restores all log settings to their default values.
func (l *Log) Reset()
func Reset()

```

### Logging Levels

Logging functions categorize messages by severity writing to the standard go
supplied log package. These functions report plain or formatted diagnostic
and status messages:

    +-----------+----------------+-----------------------------------------+
    | Level     | Functions      | Description                             |
    +-----------+----------------+-----------------------------------------+
    | Fatal     | F, Fatal       | Plain Critical Errors, program exits    |
    |           | Ff, Fatalf     | Formatted Critical Errors, program exits|
    | Error     | E, Error       | Plain errors that allow continuation    |
    |           | Ef, Errorf     | Formatted errors that allow continuation|
    | Warn      | W, Warn        | Plain warnings needing attention        |
    |           | Wf, Warnf      | Formatted warnings needing attention    |
    | Info      | I, Info        | Plain progress and status info          |
    |           | If, Infof      | Formatted progress and status info      |
    | Debug     | D, Debug       | Plain developer-focused diagnostics     |
    |           | Df, Debugf     | Formatted developer-focused diagnostics |
    | Trace     | T, Trace       | Plain detailed program flow             |
    |           | Tf, Tracef     | Formatted detailed program flow         |
    +-----------+----------------+-----------------------------------------+

### Verbose Output Levels

Verbosity output is completely separate from logging providing 5 levels of
plain or formatted messages in addition to normal program output written to
a io.Writer defaulting to os.Stdout.

The verbose output functions provide fine-grained control over program
messaging through six levels: Say0…Say5 and their formatted counterparts
Say0f…Say5f. Level 0 represents the default program output, active when no
verbosity flags are given or when explicitly set with SetVerbosity(0).
Increasing levels (1–5) are enabled by supplying additional verbosity flags or
calling SetVerbosity(1..5), progressively revealing more detailed output. All
verbose functions respect the global setting: output is suppressed entirely
when --quiet is specified or SetVerbosity(-1) is used, ensuring even default
level 0 messages are silenced.

    +-----------+----------------+-----------------------------------------+
    | Level     | Functions      | Typical Use                             |
    +-----------+----------------+-----------------------------------------+
    | 0         | S0, Say0       | Plain normal program messages           |
    |           | S0f, Say0f     | Formatted normal program messages       |
    | 1         | S1, Say1       | Plain additional messages               |
    |           | S1f, Say1f     | Formatted additional messages           |
    | 2         | S2, Say2       | Plain moderate detailed messages        |
    |           | S2f, Say2f     | Formatted moderate detailed messages    |
    | 3         | S3, Say3       | Plain detailed messages                 |
    |           | S3f, Say3f     | Formatted detailed messages             |
    | 4         | S4, Say4       | Plain highly detailed messages          |
    |           | S4f, Say4f     | Formatted highly detailed messages      |
    | 5         | S5, Say5       | Plain maximum verbosity                 |
    |           | S4f, Say4f     | Formatted maximum verbosity             |
    +-----------+----------------+-----------------------------------------+

## Settings

### Defaults

Settings have defaults which can be overridden by an environment variables.

    +-------------+-------------+-------------------------------+
    | Area        | Default     | Environment Variable Override |
    |-------------|-------------|-------------------------------|
    | LongLabels  | false       | SZLOG_LONG_LABELS             |
    | LogLevel    | LevelError  | SZLOG_LEVEL                   |
    | Language    | ""          | SZLOG_LANGUAGE                |
    | Verbose     | 0           | SZLOG_VERBOSE                 |
    +-------------+-------------+-------------------------------+

### Command Line Arguments

Settings may also be specified in command line arguments

    +-------------+-----------------------------------------------+
    | Area        | Argument                                      |
    |-------------|-----------------------------------------------|
    | LongLabels  | --long-labels                                 |
    | LogLevel    | --log <level>                                 |
    | Language    | --language <local>                            |
    | Verbose     | -v[v...] | --v[v...] | --verbose | --quiet    |
    +-------------+-----------------------------------------------+

that are processed with the program function:

```go

// AbsorbArgs scans the provided argument list for logging-related flags.
// It updates the log configuration (LogLevel, verbosity, quiet mode,
// LongLabels, and Language) based on the flags encountered. Recognized
// flags are removed, and the cleaned argument slice is returned.
// Multiple `-v` flags increment verbosity accordingly. If conflicting
// or invalid flags are found (e.g., combining `-v` with `--quiet`),
// an error is returned along with the original arguments.
func (l *Log) AbsorbArgs(argsIn []string) ([]string, error)
func AbsorbArgs(argsIn []string) ([]string, error)

```

### Program Functions

Finally settings man be queried and changed programmatically using library
functions.

    +-------------+--------------------------------------+
    | Area        | Function(s)                          |
    |-------------|--------------------------------------|
    | Language    | Language()                           |
    |             | SetLanguage(language string)         |
    | LongLabels  | LongLabels()                         |
    |             | SetLongLabels(enable bool)           |
    | LogLevel    | Level()                              |
    |             | SetLevel(newLogLevel LogLevel)       |
    |             | SetCustomLevels(levels ...LogLevel)  |
    |             | IncLevel()                           |
    |             | DecLevel()                           |
    | Verbose     | Verbose()                            |
    |             | SetVerbose(level VerboseLevel)       |
    |             | SetStdout(newWriter io.Writer)       |
    +-------------+--------------------------------------+

implemented as follows:

```go

// Language return the current local language string.
func (l *Log) Language() string
func Language() string

// SetLanguage creates a printer that attempts to format data in a local
// manner. Pass an empty "" string to stop local formatting.
func (l *Log) SetLanguage(langStr string) error
func SetLanguage(language string) error

// LongLabels returns true if long labels are currently enabled..
func (l *Log) LongLabels() bool
func LongLabels() bool

// SetLongLabels enables/disables the use of longer labels in log output.
func (l *Log) SetLongLabels(enable bool) bool
func SetLongLabels(enabled bool) bool

// Level return the current logging level.
func (l *Log) Level() LogLevel
func Level() LogLevel

// SetLevel sets the logging level.
func (l *Log) SetLevel(newLogLevel LogLevel) LogLevel
func SetLevel(newLogLevel LogLevel) LogLevel

// SetCustomLevels permits the selective enabling of individual levels.
func (l *Log) SetCustomLevels(levels ...LogLevel) LogLevel
func SetCustomLevels(levels ...LogLevel) LogLevel

// IncLevel permits all logging at the specified level.
func (l *Log) IncLevel() LogLevel
func IncLevel() LogLevel

// DecLevel permits all logging at the specified level.
func (l *Log) DecLevel() LogLevel
func DecLevel() LogLevel

// Verbose returns the current verbose level.
func (l *Log) Verbose() VerboseLevel
func Verbose() VerboseLevel

// SetVerbose set the level of output to permit.
func (l *Log) SetVerbose(newLevel VerboseLevel) VerboseLevel
func SetVerbose(level VerboseLevel) VerboseLevel

// SetStdout changes the io.Writer used by verbose output functions.  A nil
// writer will result in the default os.Stdout being used.  To cut off all
// verbose output see the --quiet argument of SetVerbose(-1).
func (l *Log) SetStdout(newWriter io.Writer) {
func SetStdout(newWriter io.Writer) {

```

## Deferred Evaluation

Output function arguments can be wrapped in a func() DEFER, where DEFER is a
string type. Deferred functions are only invoked if the target output function
is enabled. This avoids the cost of constructing expensive strings or reports
that would otherwise be discarded.

```go

    szout.Info("Report:\n", func() szout.DEFER {
      return generateReport()
    })

```

If Info logging is disabled, generateReport is never executed.

## Output Control

Szout maintains its own output pointer for verbosity, defaulting to os.Stdout
and can be redirected to any io.Writer using SetStdout. Logging output is
controlled through the built in log package and may be redirected using that
package..

## Localization

Szout can bind to the Go text package for message localization. SetLanguage
accepts a locale string such as "en" or "fr" to select the appropriate
translation.

## Convenience Function(s)

```go

// Close provides a convenience function to close anything implementing
// io.Closer and log any error returned as a warning.  Mainly to be used
// in defer functions.
func (l *Log) Close(area string, closeable io.Closer)
func Close(area string, closeable io.Closer)

```
## Use Cases

  - Use logging functions for structured severity-based output.
  - Use verbosity functions for progress and detailed tracing independent of
    logging configuration.
  - Combine deferred evaluation with either system for maximum performance.

### Quick Start

```go

    import "github.com/dancsecs/szout"

    func main() {
      szout.Info("Application starting\n")

      szout.Say1("Loading configuration\n")

      szout.Info("Report:\n", func() szout.DEFER {
        return longRunningReport()
      })
    }

```

NOTE: If Info logging is disabled, longRunningReport() is never executed.

## Dedication

This project is dedicated to Reem.
Your brilliance, courage, and quiet strength continue to inspire me.
Every line is written in gratitude for the light and hope you brought into my
life.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
<!--- gotomd::End::doc::./package -->
