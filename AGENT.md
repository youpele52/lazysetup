# Development Guidelines

Follow these rules strictly unless instructed otherwise.

## 1. File Size Limit
No file should exceed 200 lines. Split larger files into focused modules with clear naming (e.g., `handlers_navigation.go`, `handlers_actions.go`).

## 2. Single Responsibility
Each function should do one thing well. Use clear names, accept only necessary parameters, avoid unrelated side effects.

## 3. Documentation
Add doc strings to functions and structs that aren't immediately clear: complex logic, public APIs, state mutations, multiple returns, complex structs.

## 4. No Duplication in Tests
Before adding test data (strings, slices, maps, structs), check if the same data already exists in source code or other test files. Reuse existing constants and variables instead of recreating them. This prevents maintenance issues when source values change. Examples:
- Check `pkg/tools/tools.go` before defining `[]string{"git", "docker", ...}`
- Check `pkg/config/methods.go` before defining `[]string{"Homebrew", "APT", ...}`
- Check `pkg/constants/` before recreating constants in tests
