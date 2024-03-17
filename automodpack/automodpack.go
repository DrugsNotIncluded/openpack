package automodpack

import (
	"github.com/DrugsNotIncluded/openpack/cmd"
	"github.com/spf13/cobra"
)

var automodpackCmd = &cobra.Command{
	Use:     "automodpack",
	Aliases: []string{"am"},
	Short:   "Manage automodpack mods separation",
}

func init() {
	cmd.Add(automodpackCmd)
}
