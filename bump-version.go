package tanda

import (
	"fmt"
	"strings"

	tanda "fajarhac.com/fakhrullah/tanda/flutter"
	package_detail "fajarhac.com/fakhrullah/tanda/package"
	"fajarhac.com/fakhrullah/tanda/version"
	"github.com/rs/zerolog/log"
)

func BumpVersion(action string, isDryRun bool) {
	actionText := strings.Title(action)

	packageDetail, errPackageDetail := package_detail.GetPackageDetail()
	if errPackageDetail != nil {
		log.Error().Msgf("%v", errPackageDetail)
		return
	}

	ValidateVersionAction(actionText)
	ValidateGitIsClean()

	// Debug
	bumpMessage := fmt.Sprintf("%s version bump for", actionText)
	log.Info().Msgf("%s %s", bumpMessage, packageDetail.TypeName)

	// Prevent action on NPM Package
	if packageDetail.Type == package_detail.Npm {
		log.Fatal().Msg("Not implement yet for NPM package")
		panic("Not implement yet for NPM package")
	}

	// Bump version
	versionAction, err := version.StringActionToVersionAction(action)
	if err != nil {
		panic(err)
	}

	updatedVersion := version.UpdateVersion(packageDetail.Type, packageDetail.Version, versionAction)

	log.Info().Msgf("Bump the `%v` from %s to %s", packageDetail.Name, packageDetail.Version, updatedVersion)

	// Update the pubspec.yaml (package.json) file
	package_detail.UpdatePackageConfigFile(updatedVersion, isDryRun)

	if !isDryRun {
		// git commit with version message
		hash, err := AddGitCommit(updatedVersion)
		if err != nil {
			panic(err)
		}

		updatedSemverVersion := tanda.GetFlutterVersion(updatedVersion)
		// git tag with new version
		AddGitTag(updatedSemverVersion.VersionName, hash)
	}

}
