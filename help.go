package manager

import (
	"fmt"
)

type HelpCmd struct {
	Label string
	Cmd string
}

func (h HelpCmd) Render() string {
	return fmt.Sprintf(
		"%s %s",
		GuideLabelStyle.Render(fmt.Sprintf("%s:", h.Label)),
		GuideCmdStyle.Render(h.Cmd),
	)
}