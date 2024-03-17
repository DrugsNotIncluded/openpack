package main

import (
	// Modules of openpack
	_ "github.com/DrugsNotIncluded/openpack/automodpack"
	"github.com/DrugsNotIncluded/openpack/cmd"
	_ "github.com/DrugsNotIncluded/openpack/curseforge"
	_ "github.com/DrugsNotIncluded/openpack/generic"
	_ "github.com/DrugsNotIncluded/openpack/migrate"
	_ "github.com/DrugsNotIncluded/openpack/modrinth"
	_ "github.com/DrugsNotIncluded/openpack/settings"
	_ "github.com/DrugsNotIncluded/openpack/url"
	_ "github.com/DrugsNotIncluded/openpack/utils"
)

func main() {
	cmd.Execute()
}
