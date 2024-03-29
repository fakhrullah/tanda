package main

import (
	"flag"
	"fmt"
	"os"

	"fajarhac.com/fakhrullah/tanda"
	"fajarhac.com/fakhrullah/tanda/collection"
	package_detail "fajarhac.com/fakhrullah/tanda/package"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func isVersionActionAllow(action string) bool {
	allowActions := [7]string{
		"major",
		"minor",
		"patch",
		"premajor",
		"preminor",
		"prepatch",
		"prerelease",
	}

	return collection.Includes(allowActions[:], action)
}

func isSubcommandIsVersionAction(subcommand string) bool {
	return isVersionActionAllow(subcommand)
}

func main() {
	args := os.Args

	var dryRun bool

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	tanda.SetupLoggerFormat()

	tanda.LogSoftwareDetail()
	tanda.LogDebugArguments(args)

	if len(args) == 1 {
		packageDetail, err := package_detail.GetPackageDetail()
		if err != nil {

			// When using Snap package
			// tanda cannot access file not inside HOME dir
			isSnapAndNotHomeDir := tanda.IsSnapAndNotHomeDir()
			log.Debug().Msgf("is snap and not home dir %v", isSnapAndNotHomeDir)

			if isSnapAndNotHomeDir {
				log.Info().Msg("`tanda` is installed with snap (snapcraft), because of snap security, `tanda` cannot read file not inside HOME dir.")
				log.Info().Msg("You can install using other packager NOT snap (snapcraft).")
			}

			log.Fatal().Msgf("%v", err)

			return
		}

		fmt.Println()
		log.Info().Msgf("NAME:     %v", packageDetail.Name)
		log.Info().Msgf("VERSION:  %v", packageDetail.Version)
		log.Info().Msgf("TYPE:     %v\n", packageDetail.TypeName)
	}

	if len(args) >= 2 {
		versionCommand := flag.NewFlagSet("version", flag.ExitOnError)
		// sub command - [<newversion> | major | minor | patch | premajor | preminor | prepatch | prerelease | from-git]

		versionCommand.BoolVar(&dryRun, "dry-run", false, "Print output without changes the file")

		subcommand := os.Args[1]

		// Bump version
		if isSubcommandIsVersionAction(subcommand) {
			versionCommand.Parse(os.Args[2:])
			tanda.BumpVersion(subcommand, dryRun)
			return
		}

		log.Error().Msgf("Command %v is not found", subcommand)
	}
}
