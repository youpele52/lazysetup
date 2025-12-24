# lazysetup

A modern, interactive terminal UI for installing development tools across multiple platforms. Select your package manager, choose which tools to install, and watch the installation progress in real-time.

## Features

- **Multi-Panel Interface**: Intuitive side-by-side layout for installation method selection, tool selection, and status monitoring
- **Multiple Package Managers**: Support for Homebrew, APT, YUM, Curl, Scoop, and Chocolatey
- **Tool Selection**: Multi-select interface for choosing which tools to install (git, docker, lazygit, lazydocker)
- **Real-Time Status**: Live installation status with spinner animation and output display
- **Parallel Installation**: Install multiple tools simultaneously for faster setup
- **Retry Logic**: Automatic retry mechanism with random delays for failed installations
- **Installation Duration**: Track how long each tool takes to install
- **Colored Output**: Green for successful installations, red for failures
- **Error Handling**: Clear error messages and validation at each step

## Installation

### Quick Install (Curl) - Recommended

The safest and easiest way to install lazysetup:

```bash
# Latest stable release
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/install.sh | bash

# Or specific version
curl -fsSL https://github.com/youpele52/lazysetup/releases/download/v0.0.1/install.sh | bash
```

**What it does:**
- ✅ Detects your OS and architecture automatically
- ✅ Downloads the latest pre-built binary
- ✅ Verifies checksum for security
- ✅ Installs to `/usr/local/bin`
- ✅ Requests sudo only if needed

**Verify installation:**
```bash
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/verify.sh | bash
```

### Uninstall

To remove lazysetup:

```bash
# Interactive uninstall (prompts for confirmation)
curl -fsSL https://github.com/youpele52/lazysetup/releases/latest/download/uninstall.sh | bash
```

Or manually:
```bash
# Find where it's installed
which lazysetup

# Remove it
sudo rm /usr/local/bin/lazysetup
```

### From Source

```bash
git clone https://github.com/youpele52/lazysetup.git
cd lazysetup
go build -o lazysetup
./lazysetup
```

### Requirements

- Go 1.16 or higher (for building from source)
- A supported package manager (Homebrew, APT, YUM, Curl, Scoop, or Chocolatey)

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
| `Space` | Toggle tool selection |
| `Enter` | Confirm selection or proceed to next panel |
| `Esc` (double-tap) | Cancel and return to main menu |
| `Ctrl+C` | Quit application |

### Workflow

1. **Panel 1 (Package Manager)**: Select your package manager (Homebrew, APT, YUM, Curl, Scoop, Chocolatey)
2. **Panel 2 (Action)**: Choose action - Install, Update, or Uninstall
3. **Panel 3 (Tools)**: Select which tools to install/update/uninstall (git, docker, lazygit, lazydocker)
4. **Panel 0 (Status)**: Watch real-time progress with spinner animation and results

## Supported Tools

- **git**: Version control system
- **docker**: Container platform
- **lazygit**: Terminal UI for git
- **lazydocker**: Terminal UI for docker

## Supported Package Managers

- **Homebrew**: macOS and Linux
- **APT**: Debian/Ubuntu
- **YUM**: RedHat/CentOS
- **Curl**: Universal (downloads and installs)
- **Scoop**: Windows
- **Chocolatey**: Windows

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
go build -o lazysetup
```

### Running Tests

```bash
go test ./...
```

### Code Structure

The project follows a modular architecture with clear separation of concerns:
- **UI Layer**: Handles all terminal rendering and layout
- **Handler Layer**: Manages user input and business logic
- **Model Layer**: Maintains application state
- **Command Layer**: Executes system commands for installation

## License

Copyright 2025 Youpele Michael

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

## Roadmap

For planned features, enhancements, and development roadmap, see [PLAN.md](PLAN.md).
