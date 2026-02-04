# Agent Guidelines for lazysetup

This file provides guidelines for agentic coding agents operating in this repository.

## Build Commands

```bash
# Build the binary
CGO_ENABLED=0 go build -o lazysetup .

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

## Recent Improvements (v0.3.0)

**Static Binary Builds:**
- All builds use `CGO_ENABLED=0` for static linking
- Single binary works on all Linux distributions (Ubuntu 18.04+, CentOS 7+, Debian 9+, Alpine)
- Fixes GLIBC version compatibility issues

**Vim-Style Navigation:**
- `g` or `w` - Jump to first item
- `G` or `s` - Jump to last item
- Unified scroll state management via `PanelScrollState`

**UI Improvements:**
- Dynamic panel dimensions that adapt to content
- Fixed critical JumpToLast bug causing empty panels
- Clear status screen with `c` key
- Update application with `u` key

**Tool & Package Manager Support:**
- 27 CLI tools (added nvim, zsh, tmux, fzf, ripgrep, fd, bat, jq, node, gh, eza, zoxide, starship, python3, delta, btop, httpie, lazysql, tree, make, wget, tldr)
- 9 package managers (added Pacman, DNF, Nix)
- Command auto-generation for 8/9 package managers

## Project Structure

- `main.go` - Entry point
- `pkg/handlers/` - UI event handlers (keybindings, navigation, actions, vim-style jumps)
- `pkg/models/` - State management with thread-safe access
  - `state.go` - Core state struct
  - `state_methods.go` - UI state getters/setters
  - `state_installation.go` - Installation state methods
  - `scroll.go` - Unified scroll state for all panels
- `pkg/ui/` - UI rendering and layouts (dynamic panel dimensions, keybindings)
- `pkg/tools/` - Tool definitions (27 tools)
- `pkg/config/` - Configuration and installation methods (9 package managers)
- `pkg/commands/` - Command definitions and auto-generation
  - `utils.go` - Command generation helpers
- `pkg/executor/` - Command execution with timeout
- `pkg/constants/` - UI and application constants
- `pkg/colors/` - ANSI color schemes
- `pkg/updater/` - Update checking
- `pkg/testing/` - Shared test utilities

## Code Style Guidelines

### General Principles
- **Single Responsibility**: Each function does one thing well
- **File Size Limit**: No file should exceed 250 lines
- **No Duplication**: Reuse existing constants, avoid duplicating data
- **Static Builds**: ALWAYS use `CGO_ENABLED=0` when building (avoids GLIBC dependencies)

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
1. Check `pkg/tools/tools.go` for existing tools (currently 27 tools)
2. Check `pkg/config/methods.go` for installation methods (9 package managers)
3. Check `pkg/constants/` for existing constants
4. Check `pkg/models/state.go` for state patterns
5. Check if similar functionality exists

### Adding New Tools
When adding a new tool to the system:
1. Add tool name to `pkg/tools/tools.go` array
2. Add package name mappings to `pkg/commands/utils.go` if name differs across package managers
3. Add Curl install/update/uninstall commands (only if needed):
   - **MUST** use `/releases/latest/` for GitHub-hosted tools (never hardcode versions like `v1.2.3`)
   - **MUST** use `$(uname -m)` and `$(uname -s)` for architecture detection
   - **MUST** use wildcard patterns: `cd tool-*` not `cd tool-1.2.3`
4. Run tests: `go test -race ./pkg/commands/ -run TestCurlCommands`
5. Verify auto-generated commands work for all 8 package managers (Homebrew, APT, YUM, DNF, Pacman, Nix, Scoop, Chocolatey)

### Curl Command Requirements (CRITICAL)
Curl commands are manually maintained and **must follow these patterns**:

**GitHub-hosted tools:**
```bash
# CORRECT - Uses /latest/ for auto-updates
curl -fsSL https://github.com/user/repo/releases/latest/download/tool.tar.gz

# WRONG - Hardcoded version becomes outdated
curl -fsSL https://github.com/user/repo/releases/download/v1.2.3/tool.tar.gz
```

**Architecture detection:**
```bash
# CORRECT - Works on x86_64 and ARM
download/tool-$(uname -m)-linux-gnu.tar.gz

# WRONG - Breaks on ARM servers
download/tool-x86_64-linux-gnu.tar.gz
```

**Extraction patterns:**
```bash
# CORRECT - Version-agnostic wildcard
cd /tmp && tar -xzf tool.tar.gz && cd tool-* && make install

# WRONG - Hardcoded version in path
cd /tmp && tar -xzf tool.tar.gz && cd tool-1.2.3 && make install
```

Tests enforce these requirements: `TestCurlCommands_UseLatestVersions` and `TestCurlCommands_UseArchitectureDetection`

## Development Guidelines

Follow these rules strictly unless instructed otherwise.

### 1. File Size Limit
No file should exceed 250 lines. Split larger files into focused modules with clear naming (e.g., `handlers_navigation.go`, `handlers_actions.go`).

### 2. Single Responsibility
Each function should do one thing well. Use clear names, accept only necessary parameters, avoid unrelated side effects. Functions with more than 3 parameters should package arguments into a struct. Multiple struct arguments are acceptable when reasonable.

### 3. Documentation
Add docstrings to functions and structs that aren't immediately clear: complex logic, public APIs, state mutations, multiple returns, complex structs.

### 4. No Duplication
Before adding data (strings, slices, maps, structs), check if the same data already exists in source code or other files. Reuse existing constants and variables instead of duplicating them. This prevents maintenance issues when values change. Examples:
- Check `pkg/tools/tools.go` before defining `[]string{"git", "docker", ...}`
- Check `pkg/config/methods.go` before defining `[]string{"Homebrew", "APT", ...}`
- Check `pkg/constants/` before duplicating constants
- Check `pkg/models/state.go` before defining similar state-related types

### 5. Test Coverage by Priority
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

## Existing Documentation
- See `CLAUDE.md` for Claude Code-specific guidance (architecture, threading model, command patterns)
- See `TESTING.md` for detailed test documentation
- See `TEST_PLAN.md` for test priorities
- See `README.md` for project overview and user documentation
- See `PUBLISHING.md` for release procedures
- See `PLAN.md` for roadmap and planned features

## Dependencies
- Go 1.25.0+
- `github.com/jesseduffield/gocui` - Terminal UI framework
- `github.com/mattn/go-runewidth` - Unicode width calculation
- `github.com/nsf/termbox-go` - Terminal input handling

**Note:** `github.com/integrii/flaggy` is in go.mod but currently unused (run `go mod tidy` to remove)
