# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

lazysetup is an interactive TUI for installing development tools across multiple platforms. It abstracts away package manager differences (Homebrew, APT, YUM, DNF, Pacman, Nix, Scoop, Chocolatey, Curl) to provide a unified interface for installing 37+ CLI tools.

**Key Use Case:** Fresh Linux server setup where developers repeatedly install the same tools on VPS instances (Hetzner, DigitalOcean, etc.).

## Build & Test Commands

```bash
# Build
CGO_ENABLED=0 go build -o lazysetup

# Run all tests
go test ./...

# Run tests with race detection (REQUIRED for concurrent code)
go test -race ./...

# Run specific package tests
go test -race ./pkg/commands/
go test -race ./pkg/models/
go test -race ./pkg/handlers/

# Run specific test
go test -race ./pkg/commands/ -run TestCurlCommands_UseLatestVersions

# Coverage
go test -cover ./...

# Clean up unused dependencies
go mod tidy

# Build for release (see .github/workflows for multi-platform builds)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lazysetup-linux-amd64
```

## Architecture Overview

### Threading Model (Critical)

The application has **two execution contexts** that share state:

1. **Main UI Thread:** gocui event loop handling user input and rendering
2. **Installation Goroutines:** Parallel tool installations launched by handlers

**State Management:**
- All state in `models.State` is protected by `sync.RWMutex`
- ALWAYS use getter/setter methods (e.g., `state.GetCurrentTool()`, `state.SetCurrentTool()`)
- NEVER access struct fields directly (causes data races)
- All concurrent code MUST pass `-race` tests

### Package Responsibilities

**pkg/models/state.go**
- Central state struct with mutex-protected access
- Thread-safe getters/setters for all fields
- Four state files: `state.go` (struct), `state_methods.go` (accessors), `state_installation.go` (installation state), `scroll.go` (scroll state management)
- `PanelScrollState` provides unified scroll tracking for all panels with vim-style navigation support

**pkg/commands/**
- Command generation abstraction layer
- `PackageNameMappings` handles cross-platform package name differences (e.g., `nvim` vs `neovim`)
- `install.go`, `update.go`, `uninstall.go`: Lifecycle command maps
- `utils.go`: `GenerateInstallCommand()`, `GenerateUpdateCommand()`, `GenerateUninstallCommand()` auto-generate commands for 8/9 package managers
- Only Curl commands are manually maintained (GitHub releases, official downloads)

**pkg/handlers/**
- `handlers_navigation.go`: Tab/arrow key navigation, panel switching, vim-style jumps (JumpToFirst, JumpToLast)
- `handlers_actions.go`: Action (Install/Update/Uninstall) selection
- `handlers_execution.go`: Parallel tool execution with goroutines
- `installation_handlers.go`: Installation orchestration
- Each handler modifies state and triggers UI updates via gocui's `g.Update()`

**pkg/ui/**
- `layout_multipanel.go`: Main 4-panel layout orchestrator with dynamic panel dimensions
- `layout_panels.go`: Individual panel renderers (Package Manager, Action, Tools) with scroll support
- `keybindings.go`: Centralized keybinding registration (navigation, vim-style jumps, clear, update)
- Each layout function returns `error` for gocui compatibility

**pkg/executor/**
- Command execution with timeout and context cancellation
- Handles abort signals during installation

### Command Generation Pattern

When adding a new tool:

1. Add to `pkg/tools/tools.go` array
2. Add package name mapping in `pkg/commands/utils.go` if name differs across package managers
3. Add Curl install/update/uninstall commands in respective files (only if needed)
4. Auto-generation handles Homebrew, APT, YUM, DNF, Pacman, Nix, Scoop, Chocolatey

**Curl Command Requirements:**
- MUST use `/releases/latest/` for GitHub-hosted tools (not hardcoded versions)
- MUST use `$(uname -m)` and `$(uname -s)` for architecture detection
- MUST use wildcard patterns for extraction: `cd tool-*` not `cd tool-1.2.3`
- See `TestCurlCommands_UseLatestVersions` for enforcement

### UI Flow

```
PageMultiPanel (main interface)
  ├─ Panel 0: Status (installation progress)
  ├─ Panel 1: Package Manager (Homebrew, APT, etc.)
  ├─ Panel 2: Action (Install, Update, Uninstall)
  └─ Panel 3: Tools (37 tools, multi-select)

User workflow:
1. Select package manager (Panel 1)
2. Select action (Panel 2)
3. Select tools (Panel 3, Space to toggle)
4. Press Enter → handlers_execution.go launches goroutines
5. Status panel (Panel 0) shows real-time progress

Navigation:
- Tab/Shift+Tab: Cycle panels
- 0/1/2/3: Jump to specific panels
- ↑/↓: Navigate items
- g/w: Jump to first item (vim-style)
- G/s: Jump to last item (vim-style)
- c: Clear status screen
- u: Update application (when available)
```

### Scroll State Management

All panels use unified `PanelScrollState` (pkg/models/scroll.go):
- Thread-safe cursor tracking with mutex protection
- Automatic bounds checking prevents empty panel bugs
- Supports dynamic item lists (installation results grow over time)
- Vim-style navigation (gg, G) integrated with scroll state

### State Lifecycle During Installation

```go
// handlers_actions.go: User presses Enter
MultiPanelStartInstallation()
  └─ Validates tools selected
  └─ Sets state.SetInstallationStarted(true)
  └─ Launches goroutines in handlers_execution.go

// handlers_execution.go: Parallel execution
RunToolAction() // One goroutine per tool
  ├─ state.SetCurrentTool(tool)
  ├─ Executes command via executor
  ├─ state.AddInstallResult(result)
  └─ Checks state.GetAbortInstallation() flag

// UI thread: Continuous refresh
main.go: refreshUI() ticker
  └─ Calls g.Update() every 100ms
  └─ layout_panels.go reads state and renders
```

## Adding New Tools

1. Add tool name to `pkg/tools/tools.go`
2. Add package name mapping to `pkg/commands/utils.go` (if differs)
3. Add Curl commands to `install.go`, `update.go`, `uninstall.go`:
   - Use `/releases/latest/` for GitHub tools
   - Use `$(uname -m)` for architecture
   - Use wildcards: `cd tool-*` not hardcoded versions
4. Add uninstall cleanup paths to `uninstall.go` (buildCurlUninstallMap)
5. Run tests: `go test -race ./pkg/commands/`

## Adding New Package Managers

1. Add to `pkg/config/methods.go` InstallMethods array
2. Add check command to `pkg/commands/checks.go`
3. Add mapping to `PackageNameMappings` if needed
4. Update `GenerateInstallCommand()`, `GenerateUpdateCommand()`, `GenerateUninstallCommand()` in `utils.go`
5. Commands auto-generate for all tools

## Test Priorities

**P0 (Critical - Race Detection Required):**
- Concurrent state access
- Parallel goroutine execution
- All tests with "Concurrent", "Parallel", "ThreadSafety" in name

**P1 (High):**
- Command generation
- Installation flow
- Abort handling

**P2 (Medium):**
- UI rendering
- State initialization

Run P0 tests: `go test -race ./pkg/models/ ./pkg/handlers/ -run "Concurrent|Parallel|ThreadSafety"`

## Common Pitfalls

1. **Direct State Access:** Never do `state.CurrentTool = "git"`. Always use `state.SetCurrentTool("git")`
2. **Hardcoded Versions:** Curl commands must use `/latest/` or latest-stable URLs, not `v1.2.3`
3. **Missing Race Tests:** All concurrent code must pass `go test -race`
4. **Package Name Assumptions:** Check `PackageNameMappings` - tools have different names per package manager
5. **Curl Architecture:** Don't hardcode `x86_64`, use `$(uname -m)`

## Version Management

Version is in `pkg/version/version.go`. Update before releases.

Current: Check `git describe --tags --abbrev=0` and sync with version.go constant.

## Key Design Decisions

**Why TUI over CLI?** Interactive selection is better than maintaining config files for infrequent server setup tasks (1-2x per month).

**Why 9 Package Managers?** Cross-platform abstraction is the core value prop. Users on different distros/OSes use the same tool.

**Why Curl Fallback?** Many tools aren't in default repos (especially modern CLI tools like `eza`, `zoxide`, `delta`). Curl provides universal installation.

**Why Auto-Generation?** Maintaining 37 tools × 9 package managers = 333 commands is unsustainable. Auto-generate 8/9, manually maintain Curl only.

**Why Thread-Safe State?** UI thread reads state for rendering while goroutines write installation results. Mutex prevents races.

**Why Static Builds (CGO_ENABLED=0)?** Single static binary works on ALL Linux distributions (Ubuntu 18.04+, CentOS 7+, Debian 9+, Alpine) without GLIBC version conflicts. All dependencies (gocui, termbox-go, go-runewidth) are pure Go with no CGO requirements.

## Code Style

- Follow existing patterns in each package
- Tests MUST have priority comments (P0/P1/P2) and descriptions
- Curl commands MUST be tested by `TestCurlCommands_UseLatestVersions`
- All mutex-protected methods follow Get/Set prefix convention
- Comments explain "why" not "what" (code is self-documenting)
