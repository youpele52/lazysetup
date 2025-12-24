# lazysetup - Development Environment Setup Tool

**Goal**: TUI tool that installs dev tools (Git, Docker, Node, etc.) with AI-powered error resolution.

## Problem
- Manual tool installation is fragmented across platforms
- Error troubleshooting requires web searches
- Team onboarding is repetitive

## Core Features
- **Multi-tool management**: Install, update, and delete tools (Git, Docker, Node, lazygit, etc.)
- **Multiple install methods**: Homebrew, curl, APT, etc.
- **AI error resolution**: OpenAI, Anthropic, OpenRouter
- **TUI interface**: Interactive progress and selection with multi-panel layout
- **Team config sharing**: YAML-based setup files

## Installation Methods
- **Homebrew**: macOS/Linux package manager
- **Curl**: Direct download scripts
- **APT/YUM**: Linux distribution package managers  
- **Scoop/Chocolatey**: Windows package managers

## Usage
```bash
lazysetup init              # Interactive setup
lazysetup install git docker # Install specific tools
lazysetup update            # Update installed tools
lazysetup delete            # Delete/uninstall tools
lazysetup config ai.provider openai  # Set AI provider
lazysetup backup            # Backup configuration
```

## Tool Management Operations

### Install Panel
- Select installation method (Homebrew, APT, Curl, etc.)
- Choose tools to install from available list
- Monitor installation progress with spinner animation
- View installation results and error messages

### Update Panel
- Display currently installed tools with versions
- Select tools to update
- Check for available updates
- Monitor update progress
- Show update results and changelog

### Delete Panel
- Display installed tools with version info
- Select tools to uninstall
- Confirm deletion with safety checks
- Monitor uninstall progress
- Show deletion results and cleanup status

### Multi-Panel Layout
```
┌─────────────────────────────────────────────────┐
│ [1]-Installation  [2]-Update  [3]-Delete        │
├──────────────┬──────────────┬──────────────────┤
│ Methods      │ Tools        │ Progress/Results │
│              │              │                  │
│ ○ Homebrew   │ ☑ git        │ Installing...    │
│ ○ APT        │ ☑ docker     │ ✓ git (5s)       │
│ ○ Curl       │ ☐ node       │ ✗ docker (err)   │
│              │ ☐ lazygit    │                  │
└──────────────┴──────────────┴──────────────────┘
```

## AI Integration
**User Setup Flow:**
```
First Run → Select AI Provider → Enter API Key → Test Connection → Ready
```

**Error Resolution Flow:**
```
Install Fails → Classify Error → Check Cache → AI Provider → Suggest Solutions → User Decision → Apply Fix
```

**User Selection Interface:**
```
┌─ AI Provider Setup ────────────────────┐
│                                      │
│ Choose AI provider for troubleshooting:  │
│                                      │
│ ○ OpenAI (GPT-3.5-turbo)           │
│   Cost: ~$0.002 per resolution        │
│   Fastest response time                │
│                                      │
│ ○ Anthropic (Claude Haiku)           │
│   Cost: ~$0.00025 per 1K tokens     │
│   Better reasoning                    │
│                                      │
│ ○ OpenRouter (Mixtral)              │
│   Cost: ~$0.0005 per 1K tokens      │
│   Budget-friendly                    │
│                                      │
│ Provider: [OpenAI ▼]                 │
│ API Key: [_________________]           │
│                                      │
│ [Test Connection] [Save] [Skip]       │
└──────────────────────────────────────────┘
```

**Providers & Models:**
- OpenAI: GPT-3.5-turbo (~$0.002/resolution)
- Anthropic: Claude Haiku (~$0.00025/1K tokens)  
- OpenRouter: Mixtral (~$0.0005/1K tokens)

**Configuration:**
```yaml
ai:
  enabled: true
  provider: "openai"  # openai, anthropic, openrouter
  api_key: "sk-..."   # encrypted local storage
  model: "gpt-3.5-turbo"
  cache_solutions: true
  fallback_providers: ["anthropic", "openrouter"]  # try others if primary fails
```

## Implementation Details

### Tool Management Commands
```go
// Install command
installToolWithRetry(state, method, tool) → (status, error, output)

// Update command (new)
updateToolWithRetry(state, method, tool) → (status, error, output)
getInstalledToolVersion(tool) → (version, error)
checkToolUpdates(tool) → (hasUpdate, newVersion, error)

// Delete command (new)
deleteToolWithRetry(state, tool) → (status, error, output)
verifyToolInstalled(tool) → (installed, version, error)
```

### UI Pages/Panels
```
Current:
- PageMenu: Installation method selection
- PageSelection: Tool selection
- PageTools: Tool list display
- PageInstalling: Installation progress
- PageResults: Installation results
- PageMultiPanel: 3-panel layout (Installation/Tools/Progress)

New:
- PageUpdate: Update management (similar to Install)
- PageDelete: Delete/uninstall management
- PageUpdateProgress: Update progress tracking
- PageDeleteProgress: Delete progress tracking
```

### State Extensions
```go
// Add to State struct:
CurrentOperation string  // "install", "update", "delete"
InstalledTools map[string]string  // tool → version
AvailableUpdates map[string]string  // tool → newVersion
```

## Development Workflow

```
1. Project Setup
├── go mod init github.com/user/lazysetup
├── go get github.com/jroimartin/gocui
└── go get github.com/sashabaranov/go-openai

2. Core Structure
├── cmd/
│   └── root.go        # CLI commands
├── pkg/
│   ├── installer/     # Tool install/update/delete logic
│   ├── ui/           # TUI components & pages
│   ├── config/        # Configuration management
│   └── ai/           # AI integration
└── internal/         # Internal packages

3. TUI Implementation
├── InstallManager → gocui → UserInterface
├── UpdateManager → gocui → UserInterface (new)
├── DeleteManager → gocui → UserInterface (new)
├── ErrorHandler → AI → SuggestionDisplay
└── ConfigLoader → UserPreferences

4. AI Integration
├── ErrorClassifier → PatternMatching
├── AIProvider → OpenAI/Anthropic/OpenRouter
├── SolutionCache → LocalStorage
└── SolutionApplier → AutoFix

5. Testing
├── Unit Tests → Installer Logic
├── Integration Tests → AI Providers
└── E2E Tests → Full Workflow
```

## Roadmap
**Phase 1**: Go project + basic installer + TUI ✓ (Current)
- [x] Multi-tool installation support
- [x] Multiple install methods
- [x] TUI with multi-panel layout
- [ ] Update tool functionality
- [ ] Delete/uninstall tool functionality

**Phase 2**: Update & Delete Operations
- [ ] Implement update command handlers
- [ ] Implement delete command handlers
- [ ] Add version tracking for installed tools
- [ ] Create Update and Delete UI pages
- [ ] Add update progress tracking
- [ ] Add delete confirmation dialogs

**Phase 3**: Multi-platform + AI integration  
- [ ] Cross-platform compatibility testing
- [ ] AI error resolution integration
- [ ] Solution caching

**Phase 4**: Team configs + caching
- [ ] Team configuration sharing
- [ ] Solution caching system
- [ ] Advanced error handling

## Why This Wins
- **AI troubleshooting** eliminates manual searches
- **Cross-platform** unified experience
- **"lazy" brand** recognition
- **Simple TUI** over CLI complexity

---