# Development Guidelines


Follow these rules strictly unless instructed otherwise.

## 1. File Size Limit
No file should exceed 200 lines. Split larger files into focused modules with clear naming (e.g., `handlers_navigation.go`, `handlers_actions.go`).

## 2. Single Responsibility
Each function should do one thing well. Use clear names, accept only necessary parameters, avoid unrelated side effects.

## 3. Documentation
Add doc strings to functions and structs that aren't immediately clear: complex logic, public APIs, state mutations, multiple returns, complex structs.
