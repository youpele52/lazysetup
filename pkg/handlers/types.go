package handlers

import "github.com/youpele52/lazysetup/pkg/models"

// ToolActionParams groups parameters for tool action execution functions
type ToolActionParams struct {
	State  *models.State
	Method string
	Tool   string
}
