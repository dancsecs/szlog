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

/*
Package szlog is a flexible output and logging library for Go. It builds on
the standard log while adding:

  - six standard logging levels for categorizing messages by severity (Fatal,
    Error, Warn, Info, Debug, Trace with Fatalf, Errorf, Warnf, Infof, Debugf,
    Tracef for formatting) layered on top of the standard log package

  - six independent verbosity levels (Say0–Say5 with Say0f–Say5f for
    formatting)

  - lazy evaluation with `func() DEFER` so expensive values are only
    computed when needed

  - efficient function swapping so disabled functions are minimal

  - local formatting through the Go text package by setting a locale string
    ("en", "en-CA", "fr", etc.)

## Overview

A lightweight logging and verbosity framework for Go that extends the standard
library with structured levels, verbosity control, and deferred evaluation.
Logging levels target structured diagnostic messages, while verbosity levels
are designed for user-facing or progress output

Each output function is a variable member that may be swapped at runtime. When
disabled, a function is replaced with a no-op and incurs minimal runtime cost.
When enabled, the function points to an optimized writer.  Each function
mirrors standard fmt and log package function Print and Printf.  Two versions
of each function is supplied with long and short identifiers.

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

The builtin default logger may be accessed/replaced with:

```go

// Default returns the package's current default logger.
func Default() *Log

// SetDefault replaces the package's default logger with the provided one.
// It returns the previous default logger.
func SetDefault(newDefaultLog *Log) *Log

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
plain or formatted messages in addition to normal program output.

The verbose output functions provide fine-grained control over program
messaging through six levels: Say0…Say5 and their formatted counterparts
Say0f…Say5f. Level 0 represents the default program output, active when no
verbosity flags are given or when explicitly set with SetVerbose(0).
Increasing levels (1–5) are enabled by supplying additional verbosity flags or
calling SetVerbose(1..5), progressively revealing more detailed output. All
verbose functions respect the global setting: output is suppressed entirely
when --quiet is specified or SetVerbose(-1) is used, ensuring even default
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
	|           | S5f, Say5f     | Formatted maximum verbosity             |
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

that are processed with the program functions:

```go

// AbsorbArgs scans the provided argument list for enabled logging-related
// flags and updates the log configuration accordingly. Only arguments
// specified in the enable set are recognized; all others are ignored.
// Recognized flags are removed, and the cleaned argument slice is returned.
// Multiple `-v` flags increment verbosity, while invalid or conflicting
// combinations (e.g., `-v` with `--quiet`) return an error along with
// the original arguments. If no enable set is provided, EnableAll is used.
func (l *Log) AbsorbArgs(argsIn []string) ([]string, error)
func AbsorbArgs(argsIn []string) ([]string, error)

```

while usage information may be gathered with the program functions:

```go

// ArgUsageInfo reports usage information for all enabled arguments by
// invoking the provided callback for each one. Only arguments permitted
// in the enable set are included, allowing applications to present
// accurate help/usage output tailored to their configuration.
func (l *Log) ArgUsageInfo(registerArgs func(string, string))
func ArgUsageInfo(registerArgs func(string, string))

```

### Selective Argument Enabling

The `AbsorbArgs` function accepts an optional set of `EnableArg` constants
to restrict which flags are recognized. This allows applications to
individually enable or disable handling of built-in arguments.
If no constants are provided, `EnableAll` is assumed.

	+-----------------------|----------------------|-------------------------+
	| Flag / Option         | `EnableArg` constant | Description             |
	|-----------------------|----------------------|-------------------------|
	| `-v`, `--verbose`     | `EnableVerbose`      | Increase verbosity      |
	|                       |                      | (multiple `-v` allowed) |
	| `--quiet`             | `EnableQuiet`        | Suppress output         |
	| `--log <level>`       | `EnableLogLevel`     | Set log level (all,     |
	|                       |                      | trace, debug, info,     |
	|                       |                      | warn error, fatal,      |
	|                       |                      | none)                   |
	| `--language <locale>` | `EnableLanguage`     | Set message language    |
	| `--long-labels`       | `EnableLongLabels`   | Use extended labels in  |
	|                       |                      | log output              |
	| *all of the above*    | `EnableAll`          | Default (recognize all  |
	|                       |                      | arguments)              |
	+-----------------------|----------------------|-------------------------+

Example usage:

```go

// Only absorb verbosity and quiet flags, ignore all others.
args, err := szlog.AbsorbArgs(os.Args, EnableVerbose, EnableQuiet)

```

### When to Disable Arguments

By default, `szlog` absorbs all supported flags.
In some applications you may wish to disable certain arguments:

  - **Custom flag parsing**: If your program already defines `--quiet` or
    `--log`, you can prevent `szlog` from intercepting them by omitting
    `EnableQuiet` or `EnableLogLevel`.
  - **Minimal CLI surface**: Small utilities may only need `-v` for verbosity
    and want to skip everything else.
  - **Conflict avoidance**: If another library introduces overlapping flags,
    selectively enabling arguments avoids collisions.
  - **Explicit control**: Developers who prefer handling configuration
    programmatically can turn off all arguments and manage `szlog` options
    directly.

This flexibility ensures that `szlog` integrates smoothly with a variety of
command-line setups without forcing a fixed argument model.

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
	| Verbose     | Verbose()                            |
	|             | SetVerbose(level VerboseLevel)       |
	+-------------+--------------------------------------+

implemented as follows:

```go

// Language returns the current language setting used for localized formatting.
// An empty string indicates no localization is applied.
func (l *Log) Language() string
func Language() string

// SetLanguage updates the language used for localized formatting.
// Passing an empty string ("") disables localization. It returns any
// error encountered while setting the language.
func (l *Log) SetLanguage(langStr string) error
func SetLanguage(language string) error

// LongLabels reports whether long labels (FATAL, ERROR, WARN, INFO, DEBUG,
// TRACE) are currently enabled instead of their short forms (F, E, W, I, D,
// T).
func (l *Log) LongLabels() bool
func LongLabels() bool

// SetLongLabels enables or disables long labels in log output. When disabled,
// short labels (F, E, W, I, D, T) are used instead. It returns the previous
// setting.
func (l *Log) SetLongLabels(enable bool) bool
func SetLongLabels(enabled bool) bool

// Level reports the logger's current logging level.
func (l *Log) Level() LogLevel
func Level() LogLevel

// SetLevel updates the logger's logging level. Valid values include
// LevelNone, LevelFatal, LevelError, LevelWarn, LevelInfo, LevelDebug,
// LevelTrace, and LevelAll.
func (l *Log) SetLevel(newLogLevel LogLevel) LogLevel
func SetLevel(newLogLevel LogLevel) LogLevel

// SetCustomLevels enables a custom combination of individual levels.
// LevelNone, LevelAll, and LevelCustom are ignored. Internally, this
// always results in LevelCustom being applied.
func (l *Log) SetCustomLevels(levels ...LogLevel) LogLevel
func SetCustomLevels(levels ...LogLevel) LogLevel

// Verbose reports the logger's current verbosity level.
func (l *Log) Verbose() VerboseLevel
func Verbose() VerboseLevel

// SetVerbose adjusts the verbosity level (-1 through 5). Level -1 silences
// all output, while higher levels progressively enable more detail.
func (l *Log) SetVerbose(newLevel VerboseLevel) VerboseLevel
func SetVerbose(level VerboseLevel) VerboseLevel

```

## Deferred Evaluation

Output function arguments can be wrapped in a func() DEFER, where DEFER is a
string type. Deferred functions are only invoked if the target output function
is enabled. This avoids the cost of constructing expensive strings or reports
that would otherwise be discarded.

This pattern makes it safe to pass in expensive computations without worrying
about wasted work if the message is suppressed.

```go

	szlog.Info("Report:\n", func() szlog.DEFER {
	  return generateReport()
	})

```

If Info logging is disabled, generateReport is never executed.

## Localization

Szlog can bind to the Go text package for message localization. SetLanguage
accepts a locale string such as "en" or "fr" to select the appropriate
translation.

## Convenience Function(s)

```go

// Close is a convenience method for safely closing any io.Closer.
// If an error occurs during Close, it is logged as a warning.
// This method is primarily intended for use in defer statements.
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

	import "github.com/dancsecs/szlog"

	func main() {
	  var configFileName string

	  args,err:= szlog.AbsorbArgs(os.args,nil) // Set verbose and log level.

	  if len(args) != 2 {  // Ignore extra args.
	    configFileName = "default.cfg
	  } else {
	    configFileName = args[1]
	  }

	  szlog.Info("Application starting\n")

	  szlog.Say0f("Processing configuration: %s\n",configFileName)

	  szlog.Info("Report:\n", func() szlog.DEFER {
	    return longRunningReport()
	  })
	}

```

NOTE: If Info logging is disabled, longRunningReport() is never executed.

---

## Dedication

This project is dedicated to Reem.
Your brilliance, courage, and quiet strength continue to inspire me.
Every line is written in gratitude for the light and hope you brought into my
life.

---

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package szlog
