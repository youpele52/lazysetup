# Testing Guide

## Running Tests

```bash
# All tests
go test ./...

# Specific package
go test ./pkg/models/

# With race detection (required for P0/P1)
go test -race ./...

# Verbose output
go test -v ./...

# Coverage
go test -cover ./...
```

## Test Priorities

| Priority | Focus |
|----------|--------|
| P0 | Critical: Race conditions, concurrent access |
| P1 | High: Core features, common paths |
| P2 | Medium: Standard features, edge cases |
| P3 | Low: UI details, constants |

## Implemented Tests

### pkg/models (`state_test.go`)

| Test | Priority | Description |
|------|----------|-------------|
| `TestNewState` | P2 | Validates initial state |
| `TestState_ConcurrentAccess` | P0 | Multiple goroutines read/write safely |
| `TestState_PasswordMethods_ThreadSafety` | P0 | Concurrent password access |
| `TestState_PasswordInputMethods_ThreadSafety` | P0 | Concurrent input operations |
| `TestState_SelectedTools_ConcurrentOperations` | P0 | Concurrent tool selection |
| `TestState_InstallOutputMethods_ThreadSafety` | P0 | Concurrent output operations |
| `TestState_InstallResults_ConcurrentAccess` | P0 | Concurrent results collection |
| `TestState_InstallationDoneFlag_ThreadSafety` | P0 | Concurrent done flag |
| `TestState_PagePanelMethods_ThreadSafety` | P0 | Concurrent page/panel ops |
| `TestState_ToolStartTime_ConcurrentAccess` | P0 | Concurrent start time ops |
| `TestState_CancelContext_ThreadSafety` | P0 | Concurrent context ops |
| `TestState_AbortFlag_ThreadSafety` | P0 | Concurrent abort flag |
| `TestState_Reset` | P2 | Reset clears all fields |

**Run:** `go test -race ./pkg/models/`

### pkg/handlers (`handlers_execution_test.go`)

| Test | Priority | Description |
|------|----------|-------------|
| `TestRunToolAction_SingleToolExecution` | P1 | Executes action on single tool |
| `TestRunToolAction_ConcurrentExecution` | P0 | Multiple tools execute concurrently |
| `TestRunToolAction_AbortFlagStopsExecution` | P1 | Abort flag stops execution |
| `TestRunToolAction_ResultsCollectedCorrectly` | P1 | Results collected correctly |
| `TestRunToolAction_SpinnerFrameIncrements` | P2 | Spinner increments during execution |
| `TestCheckToolWithOutput_TimeoutHandling` | P1 | Check handles timeout |
| `TestCheckToolWithOutput_CancelledHandling` | P1 | Check handles cancellation |
| `TestUpdateToolWithOutput_SuccessCase` | P1 | Update tool success |
| `TestUpdateToolWithOutput_FailureWithError` | P1 | Update tool handles failure |
| `TestUninstallToolWithOutput_SuccessCase` | P1 | Uninstall tool success |

**Run:** `go test -race ./pkg/handlers/ -run "RunToolAction|CheckTool|UpdateTool|UninstallTool"`

### pkg/handlers (`installation_handlers_test.go`)

| Test | Priority | Description |
|------|----------|-------------|
| `TestMultiPanelStartInstallation_NoToolsError` | P2 | Error when no tools selected |
| `TestMultiPanelStartInstallation_StateInitialization` | P1 | Initializes state before install |
| `TestMultiPanelStartInstallation_ParallelGoroutines` | P0 | Launches parallel goroutines safely |
| `TestMultiPanelStartInstallation_FullIntegrationFlow` | P0 | Full installation flow with tools |

**Run:** `go test -race ./pkg/handlers/ -run "MultiPanelStartInstallation"`

## Test Utilities

Shared helpers in `pkg/testing/testutil.go`:

```go
// Create fresh state
state := testing.MockState()

// Assertions
testing.AssertEqual(t, want, got)
testing.AssertTrue(t, value)
testing.AssertNil(t, value)

// Wait with timeout
testing.WaitForCondition(t, func() bool {
    return state.GetInstallationDone()
}, 10*time.Second, "installation did not complete")
```

## Quick Commands

```bash
# All P0 tests (critical)
go test -race ./pkg/models/ ./pkg/handlers/ -run "Concurrent|Parallel|FullIntegration"

# All models tests
go test -race ./pkg/models/ -v

# All handlers tests
go test -race ./pkg/handlers/ -v

# Coverage report
go test -cover ./...
```

## Notes

- All concurrent tests MUST use `-race` flag
- Tests with gocui require TTY (may fail in some CI environments)
- P0 tests should complete in < 2 minutes

For detailed test plan, see [TEST_PLAN.md](TEST_PLAN.md).
