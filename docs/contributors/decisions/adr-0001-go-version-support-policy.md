# **ADR-0001** Go Version Support Policy

**Author**: Pierre Fouilloux

![Accepted](https://img.shields.io/badge/status-accepted-darkgreen) ![Date](https://img.shields.io/badge/Date-02_Mar_2025-lightblue)

## Context and Problem Statement

The project needs a clear policy for Go version support across different components. Specifically, we need to determine which Go versions should be supported by packages (`pkg/`) versus applications (`cmd/`). This decision impacts backward compatibility, maintenance overhead, and the ability to leverage new language features.

## Considered Options

* Option 1: Support only the latest Go version for all components
* Option 2: Support the earliest supported Go version for all components
* Option 3: Support the earliest supported Go version for packages (`pkg/`) and only the latest Go version for applications (`cmd/`)
* Option 4: Custom support matrix for each component based on specific needs

## Decision Outcome

Chosen option: "Option 3: Support the earliest supported Go version for packages (`pkg/`) and only the latest Go version for applications (`cmd/`)", because it provides an optimal balance between backward compatibility for libraries and the ability to leverage new language features in applications.

For example, if the current Go version is 1.24, packages would declare 1.23 in their go.mod files, while applications would declare 1.24.

### Consequences

* Good, because packages (`pkg/`) maintain backward compatibility with projects using the previous Go version
* Good, because applications (`cmd/`) can leverage the latest language features and optimizations
* Good, because it establishes a clear, consistent policy across the codebase
* Good, because it simplifies dependency management with a predictable versioning scheme
* Bad, because packages cannot immediately use the newest Go features
* Bad, because it requires tracking and updating the Go version declarations when new Go versions are released
