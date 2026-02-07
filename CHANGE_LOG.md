# Change log

## v0.3.2 (6th February 2026)

**UX Enhancement - Tools Search/Filter**:
- Added real-time search/filter to Tools panel (37 tools)
- Press `/` in Tools panel to activate search mode
- Type to filter tools by name (case-insensitive)
- Live match counter shows filtered results (e.g., "Matches: 2/37")
- Press `/` again to exit search and return to full list
- Supports numbers (0-3) in search queries without switching panels
- Search query highlighted in cyan, matches in green

**Tool Optimization & Cloud-Native Expansion** (32 → 37 tools):

**Removed**:
- **htop**: Removed in favor of btop (redundant, btop provides superior features)

**Added - Cloud-Native & DevOps** (4 tools):
- **kubectl**: Kubernetes command-line tool for cluster management
- **k9s**: Terminal UI for Kubernetes with real-time cluster monitoring
- **terraform**: Infrastructure as Code tool for cloud resource management
- **helm**: Kubernetes package manager for deploying applications

**Added - Essential Utilities** (2 tools):
- **rsync**: Fast, versatile file synchronization and transfer utility (essential for servers)
- **pnpm**: Fast, disk space efficient Node.js package manager (alternative to npm/yarn)

**Package Manager Support**:
All new tools support Homebrew, APT, YUM, DNF, Pacman, Nix, Scoop, Chocolatey, and Curl fallback for maximum platform compatibility.

**Curl Commands**:
- kubectl: Uses dl.k8s.io with stable version detection and architecture mapping (x86_64→amd64, aarch64→arm64)
- k9s: Uses GitHub `/releases/latest/` pattern with platform detection
- terraform: Uses HashiCorp releases API for latest version detection
- helm: Uses official Helm v3 installation script
- pnpm: Uses official pnpm installation script
- rsync: Source build from latest release

**Documentation**:
Updated all references from 32 to 37 tools across README.md, CLAUDE.md, and AGENTS.md.

## v0.3.1 (4th February 2026)

**Tool Expansion** (29 → 32 tools):
- **JavaScript Runtime**: bun - Fast JavaScript runtime and package manager
- **Python Package Manager**: uv - Extremely fast Python package installer and resolver
- **Command Runner**: just - Modern command runner (alternative to make)

**Package Manager Support**:
- bun: Homebrew (oven-sh/bun/bun tap) and Curl
- uv: Homebrew (main repo) and Curl  
- just: Homebrew, APT (Ubuntu 24.04+), Pacman, DNF, Scoop, Nix, Chocolatey, and Curl

**Installation Methods**:
All three tools support both standard package managers and direct curl installation for maximum compatibility across platforms.

## v0.3.0 (1st February 2026)

**CRITICAL FIX - GLIBC Compatibility**: All builds now use `CGO_ENABLED=0` for static linking. Single binary works on ALL Linux distributions (Ubuntu 18.04+, CentOS 7+, Debian 9+, Alpine Linux) without GLIBC version conflicts.

**Vim-Style Navigation**:
- `g` or `w`: Jump to first item (vim-style gg)
- `G` or `s`: Jump to last item (vim-style G)
- Unified scroll state management via `PanelScrollState` with thread-safe bounds checking
- Fixed critical JumpToLast bug that caused empty panel display

**UI Enhancements**:
- Dynamic panel dimensions that adapt to content length with responsive breakpoints
- Fixed critical panel layout bugs causing empty displays and rendering issues
- Fixed double-scrolling bugs in Package Manager and Action panels
- Improved panel scrolling with automatic bounds checking and cursor visibility
- Better visual feedback for long tool lists
- UI consistency: Package managers and actions now display in lowercase (matching tools panel)

**Tool Expansion** (5 → 29 tools):
- **Editors & Shells**: nvim, zsh
- **Terminal Utilities**: tmux, fzf, starship, btop, tree
- **Modern CLI Tools**: ripgrep, fd, bat, eza, zoxide, delta
- **Development Tools**: node, gh, python3, make, jq, wget, httpie, tldr, lazysql
- **AI Assistants**: claude-code, opencode

**Package Manager Support** (6 → 9 managers):
- Added: Pacman (Arch/Manjaro), DNF (Fedora/RHEL 8+), Nix (NixOS)
- Existing: Homebrew, APT, YUM, Curl, Scoop, Chocolatey

**Command Generation Refactor**:
- Auto-generation for 8/9 package managers (only Curl commands manually maintained)
- Helper functions in `pkg/commands/utils.go` for Install/Update/Uninstall
- `PackageNameMappings` handles cross-platform package name differences
- Comprehensive test coverage for Curl commands (version detection, architecture handling)

**Architecture Improvements**:
- New `pkg/models/scroll.go` for unified scroll state management
- New `PanelHeights` struct with `calculatePanelHeights()` for responsive design
- Thread-safe scroll tracking with mutex protection and bounds validation
- Refactored navigation handlers for vim-style jumps
- Enhanced keybindings.go with scroll, clear, and update shortcuts
- Added terminal size validation with minimum requirements error handling

**Build System**:
- Static binary compilation prevents GLIBC dependency issues
- All build commands updated to include `CGO_ENABLED=0`
- CI/CD workflow updated for static builds
- Version bumped to 0.3.0

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
- **↑↓**: Navigate items
- **g** or **w**: Jump to first item (vim-style gg)
- **G** or **s**: Jump to last item (vim-style G)
- **/**: Search/filter tools (toggle on/off in Tools panel)
- **Space**: Toggle tool
- **⏎**: Confirm/Execute
- **c**: Clear status screen
- **u**: Install update (when available)
- **Esc**: Back/Cancel
- **Ctrl+C**: Quit

## Contributors

P.E.L.E. (2026)
