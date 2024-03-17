package automodpack

import (
	"fmt"
	"os"

	"github.com/DrugsNotIncluded/openpack/core"
	"github.com/spf13/cobra"
)

// constructs automodpack config where all server-side mods excluded
// (check if any of them are mismarked as server-side dependencies of client/both side mods and exclude them)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the current modpack into a automodpack format",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// Load pack
		pack, err := core.LoadPack()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Load index
		index, err := pack.LoadIndex()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Do a refresh to ensure files are up to date
		err = index.Refresh()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = index.Write()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = pack.UpdateIndexHash()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = pack.Write()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	automodpackCmd.AddCommand(exportCmd)
}
