# lazysetup - Development Environment Setup Tool

**Goal**: TUI tool that installs dev tools (Git, Docker, Node, etc.) with AI-powered error resolution.

## Problem
- Manual tool installation is fragmented across platforms
- Error troubleshooting requires web searches
- Team onboarding is repetitive

## Core Features
- **Multi-tool installation**: Git, Docker, Node, lazygit, etc.
- **Multiple install methods**: Homebrew, curl, APT, etc.
- **AI error resolution**: OpenAI, Anthropic, OpenRouter
- **TUI interface**: Interactive progress and selection
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
lazysetup config ai.provider openai  # Set AI provider
lazysetup update            # Update all tools
lazysetup backup            # Backup configuration
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
│   ├── installer/     # Tool install logic
│   ├── ui/           # TUI components
│   ├── config/        # Configuration management
│   └── ai/           # AI integration
└── internal/         # Internal packages

3. TUI Implementation
├── InstallManager → gocui → UserInterface
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
**Phase 1**: Go project + basic installer + TUI
**Phase 2**: Multi-platform + AI integration  
**Phase 3**: Team configs + caching

## Why This Wins
- **AI troubleshooting** eliminates manual searches
- **Cross-platform** unified experience
- **"lazy" brand** recognition
- **Simple TUI** over CLI complexity

---