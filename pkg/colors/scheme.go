package colors

import "github.com/jesseduffield/gocui"

const (
	// Background colors
	BgDark = gocui.ColorDefault

	// Text colors
	TextPrimary   = gocui.ColorWhite
	TextSecondary = gocui.ColorCyan

	// Highlight colors
	HighlightBg = gocui.ColorMagenta
	HighlightFg = gocui.ColorGreen

	// Accent colors
	AccentPrimary = gocui.ColorMagenta
	AccentText    = gocui.ColorCyan

	// Status colors
	SuccessColor = gocui.ColorGreen
	FailureColor = gocui.ColorRed

	// Active panel border color
	ActiveBorderColor = gocui.ColorGreen
)

// ANSI color codes for inline text coloring
const (
	ANSIGreen   = "\033[32m"
	ANSIRed     = "\033[31m"
	ANSIYellow  = "\033[33m"
	ANSIMagenta = "\033[35m"
	ANSICyan    = "\033[36m"
	ANSIWhite   = "\033[37m"
	ANSIReset   = "\033[0m"
)
