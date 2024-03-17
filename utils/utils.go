package utils

import (
	"github.com/DrugsNotIncluded/openpack/cmd"
	"github.com/spf13/cobra"
)

// utilsCmd represents the base command when called without any subcommands
var utilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "Utilities for managing openpack itself",
}

func init() {
	cmd.Add(utilsCmd)
}
