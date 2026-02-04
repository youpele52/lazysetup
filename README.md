# lazysetup

A modern, interactive terminal UI for installing development tools across multiple platforms. Select your package manager, choose which tools to install, and watch the installation progress in real-time.

## Features

- **Multi-Panel Interface**: Intuitive side-by-side layout for installation method selection, tool selection, and status monitoring
- **Dynamic Panel Sizing**: Responsive panel dimensions that adapt to content
- **Multiple Package Managers**: Support for Homebrew, APT, YUM, DNF, Pacman, Nix, Scoop, Chocolatey, and Curl
- **Tool Selection**: Multi-select interface for choosing from 32 development tools
- **Vim-Style Navigation**: Jump to first (g/w), last (G/s), and scroll through lists efficiently
- **Real-Time Status**: Live installation status with spinner animation and output display
- **Parallel Installation**: Install multiple tools simultaneously for faster setup
- **Retry Logic**: Automatic retry mechanism with random delays for failed installations
- **Installation Duration**: Track how long each tool takes to install
- **Colored Output**: Green for successful installations, red for failures
- **Error Handling**: Clear error messages and validation at each step
- **Static Binary**: Single binary works across all Linux distributions (no GLIBC version issues)

## Installation

### Option 1: Curl (Recommended) - Safest & Easiest

The safest way to install lazysetup with automatic checksum verification:

```bash
# Latest stable release (recommended)
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/install.sh | bash

# Specific version
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/install.sh | bash -s v0.1.0
```

**Features:**
- ✅ Auto-detects OS and architecture
- ✅ Fetches latest release automatically
- ✅ Downloads pre-built binary
- ✅ Verifies checksum for security
- ✅ Installs to `/usr/local/bin`
- ✅ Requests sudo only if needed

**Verify installation:**
```bash
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/verify.sh | bash
```

### Option 2: Go Install

If you have Go 1.16+ installed:

```bash
# Latest version
go install github.com/youpele52/lazysetup@latest

# Specific version
go install github.com/youpele52/lazysetup@v0.1.0
```

Binary will be installed to `$GOPATH/bin/lazysetup` (usually `~/go/bin/lazysetup`)

### Option 3: Build from Source

```bash
git clone https://github.com/youpele52/lazysetup.git
cd lazysetup
CGO_ENABLED=0 go build -o lazysetup
./lazysetup
```

### Uninstall

**If installed via Curl:**

Interactive uninstall (prompts for confirmation):
```bash
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/uninstall.sh -o /tmp/uninstall.sh && bash /tmp/uninstall.sh
```

Or manually:
```bash
# Find installation location
which lazysetup

# Remove it
sudo rm /usr/local/bin/lazysetup
```

**If installed via Go:**
```bash
# Remove from GOPATH
rm ~/go/bin/lazysetup
```

**If built from source:**
```bash
# Remove the binary you built
rm ./lazysetup
```

### Requirements

- **Curl method**: curl, sha256sum/shasum (for verification)
- **Go method**: Go 1.16 or higher
- **Source method**: Go 1.16 or higher
- **Runtime**: A supported package manager (Homebrew, APT, YUM, Curl, Scoop, or Chocolatey)

## Usage

### Starting the Application

```bash
./lazysetup
```

### Navigation

| Key | Action |
|-----|--------|
| `Tab` / `Shift+Tab` | Cycle through panels (left/right) |
| `0`, `1`, `2`, `3` | Jump to specific panels (Status, Package Manager, Action, Tools) |
| `↑` `↓` | Navigate within active panel |
| `g` or `w` | Jump to first item (vim-style) |
| `G` or `s` | Jump to last item (vim-style) |
| `Space` | Toggle tool selection |
| `Enter` | Confirm selection or proceed to next panel |
| `c` | Clear status screen and reset state |
| `u` | Update application (when update available) |
| `Esc` (double-tap) | Cancel and return to main menu |
| `Ctrl+C` | Quit application |

### Workflow

1. **Panel 1 (Package Manager)**: Select your package manager (Homebrew, APT, YUM, DNF, Pacman, Nix, Scoop, Chocolatey, Curl)
2. **Panel 2 (Action)**: Choose action - Check, Install, Update, or Uninstall
3. **Panel 3 (Tools)**: Select which tools to install/update/uninstall (32 tools available)
4. **Panel 0 (Status)**: Watch real-time progress with spinner animation and results

## Supported Tools (32 Total)

### Version Control & Development
- **git**: Version control system
- **gh**: GitHub CLI for pull requests and issues
- **lazygit**: Terminal UI for git
- **delta**: Syntax-highlighting pager for git diffs

### Containers
- **docker**: Container platform
- **lazydocker**: Terminal UI for docker

### Modern CLI Replacements
- **ripgrep** (rg): Fast grep alternative with .gitignore support
- **fd**: User-friendly find alternative
- **bat**: Cat with syntax highlighting
- **eza**: Better ls with git integration and colors
- **zoxide**: Smarter cd that learns your habits
- **fzf**: Fuzzy finder for files and command history

### Editors & Shells
- **nvim**: Modern Vim-based editor with LSP support
- **zsh**: Superior shell with better completion

### Terminal Utilities
- **tmux**: Terminal multiplexer for persistent sessions
- **starship**: Beautiful, fast cross-shell prompt
- **htop**: Interactive process viewer
- **btop**: Modern htop with more features
- **tree**: Directory structure visualizer

### Development Tools
- **node**: JavaScript runtime
- **python3**: Python interpreter and pip
- **bun**: Fast JavaScript/TypeScript runtime and package manager
- **uv**: Ultra-fast Python package manager and environment manager
- **just**: Modern command runner (alternative to make)
- **make**: Build automation tool
- **jq**: JSON processor for APIs
- **wget**: File downloader
- **httpie**: User-friendly HTTP client
- **tldr**: Simplified man pages
- **lazysql**: Terminal UI for databases

## Supported Package Managers (9 Total)

- **Homebrew**: macOS and Linux
- **APT**: Debian/Ubuntu
- **YUM**: RHEL/CentOS (older versions)
- **DNF**: Fedora/RHEL 8+
- **Pacman**: Arch/Manjaro
- **Nix**: NixOS and cross-platform
- **Scoop**: Windows
- **Chocolatey**: Windows
- **Curl**: Universal fallback (downloads and installs from GitHub releases)

## Architecture

```
lazysetup/
├── main.go                 # Application entry point
├── pkg/
│   ├── commands/          # Installation command definitions (install, update, uninstall)
│   ├── config/            # Configuration (install methods, etc.)
│   ├── constants/         # UI constants and messages
│   ├── colors/            # Color scheme definitions
│   ├── executor/          # Command execution with timeout and cancellation
│   ├── handlers/          # Event handlers and keybindings
│   │   ├── handlers_navigation.go    # Panel navigation and cursor movement
│   │   ├── handlers_actions.go       # Action selection and execution
│   │   ├── handlers_execution.go     # Tool execution and command runners
│   │   └── handlers_legacy.go        # Legacy single-page installation handler
│   ├── models/            # State management
│   │   ├── state.go                  # Application state struct
│   │   ├── state_methods.go          # UI state getters/setters
│   │   └── state_installation.go     # Installation-related state methods
│   ├── tools/             # Tool definitions
│   └── ui/                # UI layout and rendering
│       ├── layout_multipanel.go      # Main 4-panel layout orchestration
│       ├── layout_panels.go          # Individual panel rendering functions
│       ├── messages.go               # Message builder for consistent formatting
│       └── keybindings.go            # Centralized keybinding setup
└── go.mod                 # Go module definition
```

## Key Components

### UI Package
- `layout_multipanel.go`: 4-panel layout management (Status, Package Manager, Action, Tools)
- `layout_panels.go`: Individual panel rendering (Package Manager, Action, Tools)
- `messages.go`: Dynamic message builder with action-specific success/failure messages
- `keybindings.go`: Centralized keybinding setup

### Handlers Package
- `handlers_navigation.go`: Panel navigation and cursor movement
- `handlers_actions.go`: Action selection and execution initialization
- `handlers_execution.go`: Concurrent tool execution and command runners
- `handlers_legacy.go`: Legacy single-page installation handler

### Models Package
- `state.go`: Application state struct with ActionType enum
- `state_methods.go`: UI state getters/setters with mutex protection
- `state_installation.go`: Installation-related state methods

### Commands Package
- Installation, update, and uninstall command definitions for each package manager
- Support for Homebrew, APT, YUM, Curl, Scoop, and Chocolatey

## Terminal Theme Recommendation

For optimal visual experience, use a terminal with:
- **Background**: Dark Navy/Charcoal (#2d2d44)
- **Foreground**: White/Light Gray
- **Highlight**: Bright Magenta
- **Accent**: Bright Cyan

## Development

### Building

```bash
CGO_ENABLED=0 go build -o lazysetup
```

### Running Tests

```bash
go test ./...
```

For detailed testing documentation including test structure, race detection, and test priorities, see [docs/TESTING.md](docs/TESTING.md).

### Code Structure

The project follows a modular architecture with clear separation of concerns:
- **UI Layer**: Handles all terminal rendering and layout
- **Handler Layer**: Manages user input and business logic
- **Model Layer**: Maintains application state
- **Command Layer**: Executes system commands for installation

## License

Copyright 2026 Youpele Michael

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests.

## Troubleshooting

### Installation fails with "command not found"
- Ensure the package manager is installed on your system
- The application checks for package manager availability before installation

### No tools selected error
- Select at least one tool before pressing Enter to start installation
- Use Space to toggle tool selection

### Installation hangs
- Check your internet connection
- Some installations may take time; the spinner indicates progress

### Double-tap Esc not working
- Ensure you press Esc twice within 500ms to cancel and return to main menu
- Single Esc press marks the time; second press within 500ms triggers abort

## Changelog

For a detailed history of changes and version updates, see [docs/CHANGE_LOG.md](docs/CHANGE_LOG.md).

## Roadmap

For planned features, enhancements, and development roadmap, see [docs/PLAN.md](docs/PLAN.md).
