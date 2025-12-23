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
	HighlightFg = gocui.ColorBlack

	// Accent colors
	AccentPrimary = gocui.ColorMagenta
	AccentText    = gocui.ColorCyan

	// Status colors
	SuccessColor = gocui.ColorGreen
	FailureColor = gocui.ColorRed
)

// ANSI color codes for inline text coloring
const (
	ANSIGreen = "\033[32m"
	ANSIRed   = "\033[31m"
	ANSIReset = "\033[0m"
)
