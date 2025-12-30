# Development Guidelines

Follow these rules strictly unless instructed otherwise.

## 1. File Size Limit
No file should exceed 250 lines. Split larger files into focused modules with clear naming (e.g., `handlers_navigation.go`, `handlers_actions.go`).

## 2. Single Responsibility
Each function should do one thing well. Use clear names, accept only necessary parameters, avoid unrelated side effects. Functions with more than 3 parameters should package arguments into a struct. Multiple struct arguments are acceptable when reasonable.

## 3. Documentation
Add doc strings to functions and structs that aren't immediately clear: complex logic, public APIs, state mutations, multiple returns, complex structs.

## 4. No Duplication
Before adding data (strings, slices, maps, structs), check if the same data already exists in source code or other files. Reuse existing constants and variables instead of duplicating them. This prevents maintenance issues when values change. Examples:
- Check `pkg/tools/tools.go` before defining `[]string{"git", "docker", ...}`
- Check `pkg/config/methods.go` before defining `[]string{"Homebrew", "APT", ...}`
- Check `pkg/constants/` before duplicating constants
- Check `pkg/models/state.go` before defining similar state-related types

## 5. Test Coverage by Priority
Every new function or feature must be graded against TEST_PLAN.md to determine its priority level (P0, P1, P2, P3):

**P0 & P1 Tests (Critical & High Priority):**
- Tests MUST be added immediately after function/feature creation
- Do not proceed without tests for P0/P1 items
- These tests are mandatory and non-negotiable

**P2 & P3 Tests (Medium & Low Priority):**
- Ask the user whether tests should be added immediately or deferred to later
- Provide context: "This feature is P2/P3. Should I add tests now or defer them?"
- Respect user's preference on timing

**Process:**
1. Identify the function/feature being added
2. Check TEST_PLAN.md for its priority classification
3. If P0/P1: Add tests immediately after implementation
4. If P2/P3: Ask user about test timing preference
5. Document the priority in test doc strings
