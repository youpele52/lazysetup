# Development Guidelines

Follow these rules strictly unless instructed otherwise.

## 1. File Size Limit
No file should exceed 250 lines. Split larger files into focused modules with clear naming (e.g., `handlers_navigation.go`, `handlers_actions.go`).

## 2. Single Responsibility
Each function should do one thing well. Use clear names, accept only necessary parameters, avoid unrelated side effects. Functions with more than 3 parameters should package arguments into a struct.

## 3. Documentation
Add doc strings to functions and structs that aren't immediately clear: complex logic, public APIs, state mutations, multiple returns, complex structs.

## 4. No Duplication
Before adding data (strings, slices, maps, structs), check if the same data already exists in source code or other files. Reuse existing constants and variables instead of duplicating them. This prevents maintenance issues when values change. Examples:
- Check `pkg/tools/tools.go` before defining `[]string{"git", "docker", ...}`
- Check `pkg/config/methods.go` before defining `[]string{"Homebrew", "APT", ...}`
- Check `pkg/constants/` before duplicating constants
- Check `pkg/models/state.go` before defining similar state-related types
