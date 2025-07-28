package tanda

import (
	"fmt"
	"os"
	"strings"

	"fajarhac.com/fakhrullah/tanda/collection"
	"github.com/rs/zerolog/log"
)

var versionActions = [7]string{
	"major",
	"minor",
	"patch",
	// "premajor",
	"preminor",
	"prepatch",
	"prerelease",
}

// Only allow
// - version action: major, minor, patch, premajor, preminor, prepatch, prerelease
// - other command: help
func ValidateSubCommand() {

}

func ValidateVersionAction(action string) {
	isActionIsAVersionAction := collection.Includes(versionActions[:], strings.ToLower(action))

	if !isActionIsAVersionAction {
		errorMessage := fmt.Sprintf("Action %s is not available", action)
		log.Error().Msg(errorMessage)
		os.Exit(101)
	}
}

func ValidateGitIsClean() {
	isGitNotClean := !IsGitClean()
	isGitNotCleanIgnoreUntracked := !IsGitCleanIgnoreUntracked()

	if isGitNotCleanIgnoreUntracked {
		log.Error().Msg("Current git dir is not clean")
		log.Info().Msg("You MUST commit or stash all changes before update version")
		os.Exit(101)
	}

	if isGitNotClean {
		log.Warn().Msg("Current git directory is clean but has untracked files")
	}
}
