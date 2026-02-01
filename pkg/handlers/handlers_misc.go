package handlers

import (
	"os/exec"
	"runtime"

	"github.com/jesseduffield/gocui"
)

// OpenWebsite opens the website in the default browser
func OpenWebsite(g *gocui.Gui, v *gocui.View) error {
	url := "https://youpele.com/"

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return nil
	}

	return cmd.Start()
}
