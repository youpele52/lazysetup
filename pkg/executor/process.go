package executor

type ProcessState string

const (
	Pending ProcessState = "pending"
	Running ProcessState = "running"
	Success ProcessState = "success"
	Failed  ProcessState = "failed"
)

type ProcessResult struct {
	Command string
	State   ProcessState
	Output  string
	Error   string
}

func NewProcessResult(cmd string) *ProcessResult {
	return &ProcessResult{
		Command: cmd,
		State:   Pending,
		Output:  "",
		Error:   "",
	}
}

func (p *ProcessResult) SetRunning() {
	p.State = Running
}

func (p *ProcessResult) SetSuccess(output string) {
	p.State = Success
	p.Output = output
}

func (p *ProcessResult) SetFailed(errMsg string) {
	p.State = Failed
	p.Error = errMsg
}
