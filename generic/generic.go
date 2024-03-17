package generic

import (
	"github.com/DrugsNotIncluded/openpack/cmd"
	"github.com/spf13/cobra"
)

var genericCmd = &cobra.Command{
	Use:     "generic",
	Aliases: []string{"gc"},
	Short:   "Manage vanilla-like mods environment",
}

func init() {
	cmd.Add(genericCmd)
}
