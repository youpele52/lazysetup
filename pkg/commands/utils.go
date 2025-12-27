package commands

type LifecycleCommandsType map[string]map[string]string
type CheckCommandsType map[string]string

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	m := make(map[K]V)
	for _, src := range maps {
		for k, v := range src {
			m[k] = v
		}
	}
	return m
}

type GetLifecycleCommandType struct {
	LifecycleCommandsType
	method, tool string
}

func GetLifecycleCommand(params GetLifecycleCommandType) string {
	if methodCmds, ok := params.LifecycleCommandsType[params.method]; ok {
		if cmd, ok := methodCmds[params.tool]; ok {
			return cmd
		}
	}
	return ""
}

type GetCheckCommandType struct {
	CheckCommandsType
	method string
}

func GetCheckCommandBase(params GetCheckCommandType) string {
	if cmd, ok := params.CheckCommandsType[params.method]; ok {
		return cmd
	}
	return ""
}
