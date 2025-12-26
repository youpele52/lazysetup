package handlers

import (
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/tools"
)

// MultiPanelSelectMethod selects installation method in the Package Manager panel
func MultiPanelSelectMethod(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelPackageManager {
			state.SelectedMethod = state.InstallMethods[state.SelectedIndex]
			state.CheckStatus, state.Error = checkInstallation(state.SelectedMethod)

			if state.Error == "" {
				state.Tools = tools.Tools
				state.SelectedTools = make(map[string]bool)
				state.ToolsIndex = 0
				state.SetActivePanel(models.PanelAction)
			}
		}
		return nil
	}
}

// MultiPanelSelectAction selects action in the Action panel
func MultiPanelSelectAction(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelAction {
			state.SelectedAction = models.ActionType(state.ActionIndex)
			state.SetActivePanel(models.PanelTools)
		}
		return nil
	}
}

// MultiPanelExecuteAction executes the selected action on chosen tools
func MultiPanelExecuteAction(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() != models.PageMultiPanel || state.GetActivePanel() != models.PanelTools {
			return nil
		}

		selectedTools := state.GetSelectedTools()
		selectedCount := 0
		for _, selected := range selectedTools {
			if selected {
				selectedCount++
			}
		}

		if selectedCount == 0 {
			state.Error = "No tools selected"
			return nil
		}

		// Check action doesn't need sudo confirmation
		if state.SelectedAction == models.ActionCheck {
			return executeCheckAction(state)
		}

		// Validate tool-method compatibility
		method := state.GetSelectedMethod()
		if method == "Curl" {
			// Check if htop is selected - htop cannot be installed via Curl
			if selectedTools["htop"] {
				state.Error = constants.ErrorHtopCurlNotSupported
				return nil
			}
		}

		// Only show sudo popup for package managers that need it (APT, YUM)
		// Curl and Homebrew don't require sudo
		needsSudo := method == "APT" || method == "YUM"

		if needsSudo {
			// Show sudo confirmation popup
			state.SetPendingAction(state.SelectedAction)
			state.SetShowSudoConfirm(true)
			return nil
		}

		// For Homebrew, Scoop, Chocolatey - execute directly without sudo
		switch state.SelectedAction {
		case models.ActionInstall:
			return executeInstallAction(state)
		case models.ActionUpdate:
			return executeUpdateAction(state)
		case models.ActionUninstall:
			return executeUninstallAction(state)
		}

		return nil
	}
}

// executeInstallAction initializes and runs installation for selected tools
func executeInstallAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())
	state.ActionCompletionTime = 0
	state.LastRenderedResultCount = 0

	go runToolAction(state, constants.ToolActionInstall)
	return nil
}

// executeUpdateAction initializes and runs update for selected tools
func executeUpdateAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())
	state.ActionCompletionTime = 0
	state.LastRenderedResultCount = 0

	go runToolAction(state, constants.ToolActionUpdate)
	return nil
}

// executeUninstallAction initializes and runs uninstall for selected tools
func executeUninstallAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())
	state.ActionCompletionTime = 0
	state.LastRenderedResultCount = 0

	go runToolAction(state, constants.ToolActionUninstall)
	return nil
}

func executeCheckAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())
	state.ActionCompletionTime = 0
	state.LastRenderedResultCount = 0

	go runToolAction(state, constants.ToolActionCheck)
	return nil
}

// ConfirmSudoPopup handles Enter key on sudo confirmation popup
func ConfirmSudoPopup(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if !state.GetShowSudoConfirm() {
			return nil
		}

		// Save the password from input buffer
		password := state.GetPasswordInput()
		if password == "" {
			// Don't proceed without a password
			return nil
		}
		state.SetSudoPassword(password)
		state.SetPasswordInput("") // Clear input buffer

		// Hide popup
		state.SetShowSudoConfirm(false)

		// Execute the pending action
		pendingAction := state.GetPendingAction()
		switch pendingAction {
		case models.ActionInstall:
			return executeInstallAction(state)
		case models.ActionUpdate:
			return executeUpdateAction(state)
		case models.ActionUninstall:
			return executeUninstallAction(state)
		}

		return nil
	}
}

// CancelSudoPopup handles Esc key on sudo confirmation popup
func CancelSudoPopup(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if !state.GetShowSudoConfirm() {
			return nil
		}

		// Hide popup and clear pending action
		state.SetShowSudoConfirm(false)
		state.SetPendingAction(models.ActionCheck)

		return nil
	}
}
