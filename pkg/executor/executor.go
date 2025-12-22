package executor

import (
	"os/exec"
	"strings"
)

type Executor struct {
	results map[string]*ProcessResult
}

func NewExecutor() *Executor {
	return &Executor{
		results: make(map[string]*ProcessResult),
	}
}

func (e *Executor) Execute(cmd string) *ProcessResult {
	result := NewProcessResult(cmd)
	e.results[cmd] = result

	result.SetRunning()

	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		result.SetFailed("empty command")
		return result
	}

	execCmd := exec.Command(parts[0], parts[1:]...)
	output, err := execCmd.CombinedOutput()

	if err != nil {
		result.SetFailed(err.Error())
		return result
	}

	result.SetSuccess(string(output))
	return result
}

func (e *Executor) GetResult(cmd string) *ProcessResult {
	if result, ok := e.results[cmd]; ok {
		return result
	}
	return nil
}

func (e *Executor) GetAllResults() map[string]*ProcessResult {
	return e.results
}
