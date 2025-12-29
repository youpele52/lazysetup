# Test Plan for lazysetup

## Priority Legend
- **P0**: Critical - Core functionality, security, or data safety at risk
- **P1**: High - Major features, common user paths, or complex concurrency
- **P2**: Medium - Standard features, edge cases, or validation
- **P3**: Low - Nice-to-have, UI details, or simple constants

---

## pkg/executor

### File: command.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Execute() - Successful command execution | P1 | Core execution path for all installations |
| Unit Test - ExecuteWithTimeout() - Timeout behavior | P1 | Prevents hanging installations indefinitely |
| Unit Test - ExecuteWithTimeout() - Cancellation behavior | P1 | User abort functionality depends on this |
| Unit Test - ExecuteWithTimeout() - Exit code handling | P1 | Incorrect exit codes cause false success/failure reporting |
| Unit Test - ExecuteWithSudo() - Valid password | P1 | APT/YUM installations require sudo to work |
| Unit Test - ExecuteWithSudo() - Invalid password | P1 | Security - prevent credentials from being exposed |
| Integration Test - Shell operators (&&, \>, \|) | P2 | Ensures complex installation commands work correctly |

---

## pkg/models

### File: state.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - NewState() - Initial state validation | P2 | Prevents nil pointer panics on startup |
| Unit Test - Reset() - State reset clears all fields | P2 | Prevents state corruption between operations |
| Unit Test - Concurrent Access - Multiple goroutines | P0 | Race conditions can cause crashes or data corruption |
| Unit Test - Password Methods - Thread-safety | P1 | Password security must be maintained under concurrent access |
| Unit Test - PasswordInput Methods - Thread-safety | P1 | Input buffer corruption causes incorrect passwords |

### File: state_installation.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - InstallOutput Methods - Thread-safety | P1 | Output display corruption breaks user feedback |
| Unit Test - InstallResults Methods - Concurrent access | P1 | Results tracking corruption causes incorrect status |
| Unit Test - InstallationDone Flags - Thread-safety | P1 | Premature completion flag blocks UI updates |
| Unit Test - InstallingIndex - Concurrent calls | P2 | Progress display corruption affects user experience |

### File: state_methods.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Page/Panel Methods - Thread-safety | P1 | UI navigation breaks with incorrect state |
| Unit Test - SelectedTools Map - Concurrent operations | P1 | Tool selection corruption causes wrong installations |
| Unit Test - SpinnerFrame - Multiple goroutines | P2 | Spinner animation breaks if frame index corrupts |

---

## pkg/commands

### File: install.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - GetInstallCommand() - Valid returns command | P1 | Incorrect commands cause all installations to fail |
| Unit Test - GetInstallCommand() - Invalid returns empty | P1 | Empty commands cause confusing errors |
| Integration Test - Command correctness (6 managers × 5 tools) | P1 | Every combination must work for users |
| Integration Test - Shell syntax validation | P2 | Syntax errors cause installation failures |

### File: update.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - GetUpdateCommand() - Valid returns command | P1 | Update functionality is critical for security |
| Unit Test - GetUpdateCommand() - Invalid returns empty | P1 | Prevents silent update failures |
| Integration Test - Update command correctness | P1 | Users rely on updates for bug fixes |

### File: uninstall.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - GetUninstallCommand() - Valid returns command | P1 | Users need to cleanly remove tools |
| Unit Test - GetUninstallCommand() - Invalid returns empty | P1 | Prevents partial uninstalls |
| Integration Test - Uninstall command correctness | P1 | Incomplete uninstalls leave system dirty |

### File: checks.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - GetCheckCommand() - Returns correct command | P1 | Package manager detection is essential |
| Unit Test - GetToolCheckCommand() - Returns correct command | P1 | Tool version checking is a core feature |
| Unit Test - MergeMaps() - Merges correctly | P2 | Data structure validation |
| Unit Test - MergeMaps() - Handles empty maps | P2 | Edge case handling |
| Integration Test - Commands actually check availability | P2 | Validates real-world behavior |

### File: utils.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - GetLifecycleCommand() - Valid returns command | P1 | Used by all install/update/uninstall operations |
| Unit Test - GetLifecycleCommand() - Invalid returns empty | P1 | Error handling depends on this |
| Unit Test - GetCheckCommandBase() - Valid returns value | P2 | Helper function validation |
| Unit Test - GetCheckCommandBase() - Invalid returns empty | P2 | Edge case handling |

---

## pkg/handlers

### File: keybindings.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - installToolWithRetry() - Success on first attempt | P1 | Happy path for installations |
| Unit Test - installToolWithRetry() - Success after retries | P1 | Retry logic is critical for network reliability |
| Unit Test - installToolWithRetry() - Failure after max retries | P1 | Prevents infinite loops, reports failure |
| Unit Test - installToolWithRetry() - Exponential backoff timing | P2 | Prevents server overload from rapid retries |
| Unit Test - installToolWithOutput() - Success with sudo | P1 | APT/YUM installations won't work without this |
| Unit Test - installToolWithOutput() - Success without sudo | P1 | Homebrew/Scoop/Chocolatey depend on this |
| Unit Test - installToolWithOutput() - Timeout handling | P1 | Prevents installations from hanging forever |
| Unit Test - installToolWithOutput() - Cancellation handling | P1 | User abort functionality requires this |
| Unit Test - checkInstallation() - Installed detected | P1 | Core flow depends on correct detection |
| Unit Test - checkInstallation() - Not installed detected | P1 | Error messages depend on correct detection |
| Unit Test - GoBack() - Single escape records timestamp | P2 | Prevents accidental double-trigger |
| Unit Test - GoBack() - Double escape triggers abort | P1 | Critical for user control |
| Unit Test - GoBack() - Spaced escapes don't trigger | P2 | Prevents unintended aborts |

### File: handlers_actions.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - MultiPanelSelectMethod() - Updates state | P2 | Navigation correctness |
| Unit Test - MultiPanelSelectAction() - Updates state | P2 | Action selection flow |
| Unit Test - MultiPanelExecuteAction() - No tools error | P2 | Prevents confusing empty execution |
| Unit Test - MultiPanelExecuteAction() - Check without sudo | P1 | Check action should work without passwords |
| Unit Test - MultiPanelExecuteAction() - APT/YUM shows sudo popup | P1 | Security - sudo required for these managers |
| Unit Test - MultiPanelExecuteAction() - htop+Curl error | P1 | Prevents installation attempt that will fail |
| Unit Test - ConfirmSudoPopup() - Password saved and action executed | P1 | Complete sudo flow must work |
| Unit Test - ConfirmSudoPopup() - Empty password doesn't proceed | P1 | Security - prevents blank passwords |
| Unit Test - CancelSudoPopup() - Cancels pending action | P2 | User can change their mind |

### File: handlers_execution.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - runToolAction() - Single tool execution | P1 | Basic installation flow |
| Unit Test - runToolAction() - Multiple tools concurrent | P0 | Race condition risk with parallel execution |
| Unit Test - runToolAction() - Abort flag stops execution | P1 | User control during parallel execution |
| Unit Test - runToolAction() - Results collected correctly | P1 | Results display depends on this |
| Unit Test - runToolAction() - Spinner frame increments | P2 | UI feedback for progress |
| Unit Test - checkToolWithOutput() - Success returns version | P1 | Version display feature |
| Unit Test - checkToolWithOutput() - Timeout handling | P1 | Prevents hanging checks |
| Unit Test - checkToolWithOutput() - Cancelled handling | P1 | User abort during check |
| Unit Test - updateToolWithOutput() - Success case | P1 | Core update feature |
| Unit Test - updateToolWithOutput() - Failure with error | P1 | Error reporting for updates |
| Unit Test - uninstallToolWithOutput() - Success case | P1 | Core uninstall feature |
| Integration Test - Parallel execution with 4 tools | P0 | Main user scenario must work |

### File: installation_handlers.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - MultiPanelStartInstallation() - No tools error | P2 | Validation before execution |
| Unit Test - MultiPanelStartInstallation() - State initialization | P1 | Prevents state corruption on installation |
| Unit Test - MultiPanelStartInstallation() - Parallel goroutines | P0 | Complex concurrency requires testing |
| Integration Test - Full installation flow with multiple tools | P0 | Main user journey |

### File: handlers_navigation.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Navigation methods - Panel switching | P2 | UI usability depends on this |
| Unit Test - Cursor movement - Index bounds | P2 | Prevents out-of-bounds panics |

### File: handlers_update.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Update handlers - All scenarios | P1 | Update feature is critical for security |

---

## pkg/updater

### File: updater.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - CheckForUpdates() - Successful check | P1 | Users must know about updates |
| Unit Test - CheckForUpdates() - Network error handling | P1 | Graceful failure on network issues |
| Unit Test - CheckForUpdates() - HTTP error handling | P1 | Handle GitHub API rate limits/errors |
| Unit Test - CheckForUpdates() - JSON parsing error | P1 | Invalid API response shouldn't crash app |
| Unit Test - findDownloadURL() - Correct asset found | P1 | Update download depends on this |
| Unit Test - findDownloadURL() - No match returns empty | P2 | Edge case handling |
| Unit Test - isNewerVersion() - Newer returns true | P1 | Update notification logic |
| Unit Test - isNewerVersion() - Same/older returns false | P1 | Prevents unnecessary update prompts |
| Unit Test - isNewerVersion() - Edge cases (v prefix, length) | P2 | Version string parsing robustness |
| Unit Test - DownloadAndInstall() - Successful | P1 | Update installation is critical feature |
| Unit Test - DownloadAndInstall() - Download error | P1 | Graceful failure on download issues |
| Unit Test - DownloadAndInstall() - Cross-device copy fallback | P1 | Filesystem compatibility |
| Unit Test - copyFile() - File copied correctly | P2 | Helper function validation |
| Unit Test - RestartApplication() - Command prepared | P2 | Restart mechanism validation |

---

## pkg/ui

### File: layout.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - getSpinner() - Frame cycling | P3 | Visual feedback correctness |
| Unit Test - Layout() - Correct layout function | P2 | Page navigation correctness |

### File: layout_multipanel.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - layoutMultiPanel() - All 4 panels rendered | P2 | UI completeness |
| Unit Test - layoutMultiPanel() - Correct panel dimensions | P3 | Visual layout correctness |

### File: layout_panels.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Each panel render function - Content displayed | P3 | UI content correctness |
| Unit Test - Active panel highlighting - Only one panel | P3 | Visual feedback correctness |

### File: keybindings.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - SetupKeybindings() - All keys bound | P2 | Navigation depends on this |

### File: messages.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Message builders - Correct formatting | P3 | User message clarity |

---

## pkg/tools

### File: tools.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Tools slice - Contains expected tools | P3 | Configuration validation |

---

## pkg/config

### File: methods.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - InstallMethods - Contains 6 managers | P3 | Configuration validation |
| Unit Test - Actions - Contains 4 actions | P3 | Configuration validation |

---

## pkg/colors

### File: scheme.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Color codes - Valid ANSI escape sequences | P3 | Terminal display correctness |

---

## pkg/constants

### File: ui.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - Constants - All defined and non-empty | P3 | Configuration validation |

---

## main

### File: main.go

| Test | Priority | Why |
|------|----------|-----|
| Unit Test - refreshUI() - Ticker interval | P2 | UI responsiveness |
| Unit Test - checkForUpdates() - Update info set correctly | P1 | Startup update check |
| Unit Test - checkForUpdates() - No update on error | P2 | Error handling |
| Integration Test - main() - App runs without panicking | P1 | Application stability |

---

## Critical Integration Tests

| Test | Priority | Why |
|------|----------|-----|
| Full Installation Flow - Complete install cycle | P0 | Main user journey |
| Sudo Flow - Password input and execution | P0 | APT/YUM installations require this |
| Abort Flow - Double escape aborts installation | P1 | User control during operations |
| Update Check Flow - Update notification on startup | P1 | Security and feature updates |
| Concurrent Installation - 4 tools in parallel | P0 | Race condition risk with parallel execution |
| Tool Compatibility - htop+Curl validation | P1 | Prevents failed installations |
| Multi-Panel Navigation - Panel switching works | P2 | Core user interaction |

---

## End-to-End Tests

| Test | Priority | Why |
|------|----------|-----|
| Complete Workflow - Install and verify | P0 | Validates entire application |
| Error Recovery - Retry logic works | P1 | Network reliability feature |
| Update Installation - Download and restart | P1 | Critical for application updates |

---

## Summary Statistics

- **P0 Tests (Critical)**: 13
- **P1 Tests (High)**: 52
- **P2 Tests (Medium)**: 30
- **P3 Tests (Low)**: 10
- **Total**: 105 tests

### Recommended Implementation Order

1. **Phase 1 (Week 1-2)**: All P0 tests - Critical functionality and race conditions
2. **Phase 2 (Week 3-4)**: All P1 tests - Core features and common paths
3. **Phase 3 (Week 5-6)**: P2 tests - Standard features and edge cases
4. **Phase 4 (Week 7)**: P3 tests - UI details and constants

---

## Testing Framework Recommendations

- **Unit Tests**: Standard Go `testing` package
- **Race Detection**: `go test -race` flag mandatory for P0/P1 concurrent tests
- **Mocking**: Use `testify/mock` for HTTP requests and command execution
- **Integration Tests**: Use `testcontainers` for package manager simulation (optional)
- **Coverage**: Target 80%+ code coverage for P0/P1 packages

---

## Notes

- All concurrent tests MUST be run with `-race` flag
- Tests involving actual package managers should be skipped in CI or use mocks
- HTTP tests should use mock servers to avoid external dependencies
- File operation tests should use temporary directories
- Consider test execution time: P0/P1 tests should complete in < 2 minutes
