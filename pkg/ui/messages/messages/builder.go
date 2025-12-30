package messages

import (
	"strings"

	"github.com/youpele52/lazysetup/pkg/constants"
)

type MessageBuilder struct {
	lines []string
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		lines: []string{},
	}
}

func (mb *MessageBuilder) AddLine(text string) *MessageBuilder {
	mb.lines = append(mb.lines, text)
	return mb
}

func (mb *MessageBuilder) AddBlankLine() *MessageBuilder {
	mb.lines = append(mb.lines, "")
	return mb
}

func (mb *MessageBuilder) AddSeparator() *MessageBuilder {
	mb.lines = append(mb.lines, constants.ResultsSeparator)
	return mb
}

func (mb *MessageBuilder) Build() string {
	lines := make([]string, len(mb.lines))
	copy(lines, mb.lines)
	return strings.Join(lines, "\n")
}
