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

  - five independent verbosity levels (Say0–Say4 with Say0f–Say4f for
    formatting) with settable output defaulting to ```os.Stdout```

  - lazy evaluation with ```func() DEFER```` so expensive values are only
    computed when needed

  - efficient function swapping so disabled functions are minimal

  - local formatting through the Go text package by setting a locale string
    ("en", "fr", etc.)

## Overview

Szout offers two independent families of output functions:

  - Logging functions: Fatal, Error, Warn, Info, Debug, Trace
  - Verbosity functions: Say0–Say4 and Say0f–Say4f

Logging and verbosity are separate systems. Logging functions categorize
messages by severity, while verbosity functions emit step-by-step or
progress-style output at increasing levels of detail.

Each function is a variable member that may be swapped at runtime. When
disabled, a function is replaced with a no-op and incurs minimal runtime cost.
When enabled, the function points to an optimized writer.

## Deferred Evaluation

Arguments can be wrapped in a func() DEFER, where DEFER is a string type.
Deferred functions are only invoked if the target output function is enabled.
This avoids the cost of constructing expensive strings or reports that would
otherwise be discarded.

    szout.Info("Report:\n", func() szout.DEFER { return generateReport()
    })

If Info logging is disabled, generateReport is never executed.

## Output Control

Szout maintains its own output pointer for verbosity, defaulting to os.Stdout
and can be redirected to any io.Writer using SetStdout. Logging output is
controlled through the built in log package.

## Localization

Szout can bind to the Go text package for message localization. SetLanguage
accepts a locale string such as "en" or "fr" to select the appropriate
translation.

## Use Cases

  - Use logging functions for structured severity-based output.
  - Use verbosity functions for progress and detailed tracing independent of
    logging configuration.
  - Combine deferred evaluation with either system for maximum performance.

## Quick Start

     import "github.com/dancsecs/szout"

    func main() { szout.Info("Application starting\n")

        szout.Say1("Loading configuration\n")

        szout.Info("Report:\n", func() szout.DEFER {
            return longRunningReport()
        })
    }

NOTE: If Info logging is disabled, longRunningReport() is never executed.
<!--- gotomd::End::doc::./package -->
