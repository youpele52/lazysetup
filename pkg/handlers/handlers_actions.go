package handlers

import (
	"time"

	"github.com/jesseduffield/gocui"
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

	go runToolAction(state, "install")
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

	go runToolAction(state, "update")
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

	go runToolAction(state, "uninstall")
	return nil
}
