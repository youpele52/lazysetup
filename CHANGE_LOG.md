# Change log

## v0.1.4 (30th December 2025)

**Build System**: Fixed critical package conflict caused by duplicate `messages.go` files that prevented compilation on all platforms.

**Version Management**: Updated version constant to 0.1.4 and cleaned up incorrectly versioned release tags for accurate auto-update checks.

**Check Action Enhancement**: Version information now displays for successful tool checks with proper formatting and color coding.

**Code Quality**: Consolidated hardcoded action verb strings into reusable constants, eliminating duplication across the codebase.

**Test Coverage**: Added comprehensive test documentation with priority labels (P0-P3) across all test files. Added 59 new tests covering configuration validation, navigation, and edge cases.

**Development Guidelines**: Enhanced AGENT.md with rules for code duplication prevention and test coverage requirements by priority level. Expanded TEST_PLAN.md with concrete examples and decision criteria for each priority level.

**Bug Fixes**: Fixed check action output display, improved default action verb handling, standardized version tag naming convention.

## v0.1.2 (27th December 2025)

**Auto-Update System**: Automatically checks GitHub for updates on startup. Yellow notification banner displays for 10 seconds. Press 'u' to install.

**Tool Management**: 
- htop Installation: Now fully supported via Homebrew and APT package managers
- Version Checking: Users can check installed versions of all tools (git, docker, lazygit, lazydocker, htop)
- Installation Methods: Supports Homebrew, APT, YUM, Curl, Scoop, and Chocolatey

**Enhanced Status Panel**: Error messages display in red, update notifications in yellow. Auto-hide after timeout.

**UI/UX Improvements**: 
- Press 'c' to clear status panel
- Press 'u' to install updates
- Rolling credits display (latest at top, oldest at bottom)
- Improved keybinding hints in status bar

**Bug Fixes**: Fixed sudo password escaping, removed unnecessary sudo for Curl, fixed state reset for sequential actions, improved rolling credits display.

## Installation

```bash
brew install lazysetup  # macOS
apt-get install lazysetup  # Linux (APT)
yum install lazysetup  # Linux (YUM)
```

## Key Bindings

- **Tab/0-3**: Switch panels
- **↑↓**: Navigate
- **Space**: Toggle tool
- **⏎**: Confirm
- **C**: Clear status
- **U**: Install update
- **Esc**: Back
- **Ctrl+C**: Quit

## Contributors

P.E.L.E. (2025)
