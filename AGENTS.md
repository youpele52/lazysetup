# Agent Guidelines for lazysetup

This file provides guidelines for agentic coding agents operating in this repository.

## Build Commands

```bash
# Build the binary
go build -o lazysetup .

# Build for all platforms
./build.sh  # or see PUBLISHING.md for cross-compilation

# Run locally
go run main.go
```

## Test Commands

```bash
# All tests
go test ./...

# Specific package
go test ./pkg/models/
go test ./pkg/handlers/

# With race detection (REQUIRED for P0/P1 tests)
go test -race ./...

# Verbose output
go test -v ./...

# Coverage report
go test -cover ./...

# Run single test
go test -race ./pkg/models/ -run TestState_ConcurrentAccess

# Run test category
go test -race ./pkg/handlers/ -run "RunToolAction|CheckTool|UpdateTool|UninstallTool"

# Quick P0 tests (critical concurrent tests)
go test -race ./pkg/models/ ./pkg/handlers/ -run "Concurrent|Parallel|FullIntegration"
```

## Linting & Formatting

```bash
# Format code
gofmt -w .

# Vet code
go vet ./...

# Check for issues
go vet -shadow ./...

```

## Project Structure

- `main.go` - Entry point
- `pkg/handlers/` - UI event handlers (keybindings, navigation, actions)
- `pkg/models/` - State management with thread-safe access
- `pkg/ui/` - UI rendering and layouts
- `pkg/tools/` - Tool definitions
- `pkg/config/` - Configuration and installation methods
- `pkg/commands/` - Command definitions
- `pkg/executor/` - Command execution
- `pkg/constants/` - UI and application constants
- `pkg/colors/` - ANSI color schemes
- `pkg/updater/` - Update checking
- `pkg/testing/` - Shared test utilities

## Code Style Guidelines

### General Principles
- **Single Responsibility**: Each function does one thing well
- **File Size Limit**: No file should exceed 250 lines
- **No Duplication**: Reuse existing constants, avoid duplicating data

### Imports
- Standard library first, then third-party, then internal packages
- Group imports with clear separation:
```go
import (
    "context"
    "fmt"
    "sync"
    "time"

    "github.com/jesseduffield/gocui"
    "github.com/youpele52/lazysetup/pkg/models"
    "github.com/youpele52/lazysetup/pkg/ui"
)
```

### Naming Conventions
- **Variables**: camelCase for locals, PascalCase for exported
- **Constants**: PascalCase for exported, camelCase for unexported
- **Structs**: PascalCase with clear names (e.g., `InstallResult`, `State`)
- **Interfaces**: PascalCase, often with "er" suffix for capabilities (e.g., `Reader`)
- **Files**: snake_case for test files (`handlers_execution_test.go`)

### Thread Safety
- State structs must use `sync.RWMutex` for concurrent access
- Provide getter/setter methods with proper locking
- Access shared state only through thread-safe methods
- Always use `-race` flag when testing concurrent code

### Error Handling
- Return errors to callers rather than logging directly
- Use `log.Panicln` only for truly unrecoverable errors (e.g., GUI initialization failure)
- Context cancellations should be handled gracefully
- Check `err != nil` immediately after calls that can fail

### Documentation
- Add docstrings to:
  - Exported functions and types
  - Complex logic or state mutations
  - Functions with multiple return values
  - Public APIs and stateful methods
- Comment implementation details for non-obvious code

### Function Parameters
- Functions with >3 parameters should use a struct
- Use pointers for large structs to avoid copying
- Accept interfaces for dependencies (e.g., `io.Reader`)

### Testing Requirements

All new code must be tested according to `TEST_PLAN.md` priority levels:

**P0 (Critical) - Tests REQUIRED immediately:**
- Race conditions, concurrent access
- Core features, common paths
- Must pass before merge

**P1 (High) - Tests REQUIRED immediately:**
- Important features, standard operations

**P2 (Medium) - Ask user:**
- Edge cases, standard features
- Ask: "Should I add tests now or defer?"

**P3 (Low) - Ask user:**
- UI details, constants
- Ask: "Should I add tests now or defer?"

Test utilities are available in `pkg/testing/testutil.go`:
```go
state := testing.MockState()
testing.AssertEqual(t, want, got)
testing.AssertTrue(t, value)
testing.WaitForCondition(t, func() bool { ... }, timeout, message)
```

### Before Adding New Code
1. Check `pkg/tools/tools.go` for existing tools
2. Check `pkg/config/methods.go` for installation methods
3. Check `pkg/constants/` for existing constants
4. Check `pkg/models/state.go` for state patterns
5. Check if similar functionality exists

## Existing Documentation
- See `AGENT.md` for additional development guidelines
- See `TESTING.md` for detailed test documentation
- See `TEST_PLAN.md` for test priorities
- See `README.md` for project overview
- See `PUBLISHING.md` for release procedures

## Dependencies
- Go 1.25.0+
- `github.com/jesseduffield/gocui` - Terminal UI framework
- `github.com/integrii/flaggy` - CLI argument parsing
- `github.com/mattn/go-runewidth` - Unicode width calculation
- `github.com/nsf/termbox-go` - Terminal input handling
