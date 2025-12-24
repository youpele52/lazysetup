package handlers

import (
	"sync"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// StartInstallation initiates parallel installation of selected tools
// It launches goroutines for each tool, collects results in a channel,
// and updates the UI in real-time with spinner animation
// NOTE: This function is for old single-page layout, use MultiPanelStartInstallation for multi-panel mode
func StartInstallation(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		// Initialize installation state using thread-safe methods
		state.ClearInstallResults()
		state.ClearToolStartTimes()
		state.ClearInstallOutput()
		state.SetCurrentPage(models.PageInstalling)
		state.SetInstallingIndex(0)
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())
		state.SetCurrentTool("")

		go func() {
			spinnerTicker := time.NewTicker(100 * time.Millisecond)
			defer spinnerTicker.Stop()

			spinnerDone := make(chan bool)

			go func() {
				for {
					select {
					case <-spinnerDone:
						return
					case <-spinnerTicker.C:
						if !state.GetInstallationDone() {
							state.IncrementSpinnerFrame()
						}
					}
				}
			}()

			var wg sync.WaitGroup
			var mu sync.Mutex
			resultsChan := make(chan models.InstallResult, len(state.Tools))

			for _, tool := range state.Tools {
				if state.SelectedTools[tool] {
					wg.Add(1)
					go func(toolName string) {
						defer wg.Done()
						startTime := time.Now().Unix()
						state.SetToolStartTime(toolName, startTime)
						status, errMsg, output := installToolWithRetry(state, state.SelectedMethod, toolName)

						mu.Lock()
						state.AppendInstallOutput("Tool: " + toolName + "\n" + output + "\n")
						mu.Unlock()

						duration := time.Now().Unix() - state.GetToolStartTime(toolName)
						result := models.InstallResult{
							Tool:     toolName,
							Success:  status == constants.StatusSuccess,
							Error:    errMsg,
							Duration: duration,
							Retries:  0,
						}
						resultsChan <- result
					}(tool)
				}
			}

			go func() {
				wg.Wait()
				close(resultsChan)
			}()

			for result := range resultsChan {
				state.AddInstallResult(result)
				state.IncrementInstallingIndex()
			}

			state.SetInstallationDone(true)
			spinnerDone <- true
			time.Sleep(1 * time.Second)
			state.CurrentPage = models.PageResults
		}()

		return nil
	}
}

// MultiPanelStartInstallation initiates parallel installation in multi-panel mode
// Validates tool selection, launches goroutines, collects results, and handles abort requests
func MultiPanelStartInstallation(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelTools {
			// Check if at least one tool is selected
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

			// Initialize installation state using thread-safe methods
			state.ClearInstallResults()
			state.ClearToolStartTimes()
			state.ClearInstallOutput()
			state.SetInstallingIndex(0)
			state.SetInstallationDone(false)
			state.SetInstallStartTime(time.Now().Unix())

			go func() {
				spinnerTicker := time.NewTicker(100 * time.Millisecond)
				defer spinnerTicker.Stop()

				spinnerDone := make(chan bool)

				go func() {
					for {
						select {
						case <-spinnerDone:
							return
						case <-spinnerTicker.C:
							if !state.GetInstallationDone() {
								state.IncrementSpinnerFrame()
							}
						}
					}
				}()

				var wg sync.WaitGroup
				var mu sync.Mutex
				resultsChan := make(chan models.InstallResult, len(state.Tools))

				for _, tool := range state.Tools {
					if state.SelectedTools[tool] {
						// Check if abort was requested
						if state.GetAbortInstallation() {
							break
						}

						wg.Add(1)
						go func(toolName string) {
							defer wg.Done()

							// Check abort flag before starting installation
							if state.GetAbortInstallation() {
								return
							}

							startTime := time.Now().Unix()
							state.SetToolStartTime(toolName, startTime)
							status, errMsg, output := installToolWithRetry(state, state.SelectedMethod, toolName)

							mu.Lock()
							state.AppendInstallOutput("Tool: " + toolName + "\n" + output + "\n")
							mu.Unlock()

							duration := time.Now().Unix() - state.GetToolStartTime(toolName)
							result := models.InstallResult{
								Tool:     toolName,
								Success:  status == constants.StatusSuccess,
								Error:    errMsg,
								Duration: duration,
								Retries:  0,
							}
							resultsChan <- result
						}(tool)
					}
				}

				go func() {
					wg.Wait()
					close(resultsChan)
				}()

				for result := range resultsChan {
					state.AddInstallResult(result)
					state.IncrementInstallingIndex()
				}

				state.SetInstallationDone(true)
				spinnerDone <- true
				time.Sleep(1 * time.Second)
			}()
		}
		return nil
	}
}
