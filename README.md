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

### From Source

```bash
git clone https://github.com/youpele52/lazysetup.git
cd lazysetup
go build -o lazysetup
./lazysetup
```

### Requirements

- Go 1.16 or higher
- A supported package manager (Homebrew, APT, YUM, Curl, Scoop, or Chocolatey)

## Usage

### Starting the Application

```bash
./lazysetup
```

### Navigation

| Key | Action |
|-----|--------|
| `Tab` | Cycle through panels |
| `0`, `1`, `2` | Jump to specific panels (Installation, Tools, Progress) |
| `↑` `↓` | Navigate within active panel |
| `Space` | Toggle tool selection |
| `Enter` | Confirm selection or start installation |
| `Esc` | Go back |
| `Ctrl+C` | Quit |

### Workflow

1. **Panel 0 (Installation)**: Select your package manager (Homebrew is default)
2. **Panel 1 (Tools)**: Select which tools to install (git, docker, lazygit, lazydocker)
3. **Panel 2 (Status)**: Watch real-time installation status and results

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
│   ├── commands/          # Installation command definitions
│   ├── config/            # Configuration (install methods, etc.)
│   ├── constants/         # UI constants and messages
│   ├── colors/            # Color scheme definitions
│   ├── errors/            # Error handling
│   ├── handlers/          # Event handlers and keybindings
│   ├── models/            # State management
│   ├── tools/             # Tool definitions
│   └── ui/                # UI layout and rendering
└── go.mod                 # Go module definition
```

## Key Components

### UI Package
- `layout.go`: Multi-panel layout management
- `messages.go`: Message builder for consistent formatting
- `keybindings.go`: Centralized keybinding setup

### Handlers Package
- `keybindings.go`: Event handlers for user input and installation logic

### Models Package
- `state.go`: Application state management

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

## Future Enhancements

- [ ] Configuration file support
- [ ] Installation history
- [ ] Custom tool definitions
- [ ] Progress persistence
- [ ] Rollback functionality
