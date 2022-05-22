package version

import (
	"errors"
	"io/ioutil"

	"fajarhac.com/fakhrullah/tanda/collection"
	flutter "fajarhac.com/fakhrullah/tanda/flutter"
	package_detail "fajarhac.com/fakhrullah/tanda/package"
	"github.com/blang/semver"
	"github.com/rs/zerolog/log"
)

type VersionAction string

const (
	Major      VersionAction = "major"
	Minor                    = "minor"
	Patch                    = "patch"
	Premajor                 = "premajor"
	Preminor                 = "preminor"
	Prepatch                 = "prepatch"
	Prerelease               = "prerelease"
	NoAction                 = ""
)

func throwIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

type T struct {
	name    string
	version string
}

func hasPrerelease(semverVersion semver.Version) bool {
	if len(semverVersion.Pre) > 0 {
		return true
	}
	return false
}

func major(currentVersion semver.Version) semver.Version {
	currentMajor := currentVersion.Major
	newMajor := currentMajor + 1

	newVersion, err := semver.Make("0.0.0")
	if err != nil {
		panic(err)
	}

	if hasPrerelease(currentVersion) && currentVersion.Minor == 0 && currentVersion.Patch == 0 {
		// When update major from prerelease premajor
		newVersion.Major = currentMajor
	} else {
		newVersion.Major = newMajor
	}

	return newVersion
}

func premajor(currentVersion semver.Version) semver.Version {
	newVersion := major(currentVersion)
	updatedPremajorVersion := newVersion
	pre := semver.PRVersion{
		VersionNum: 0,
		IsNum:      true,
	}
	updatedPremajorVersion.Pre = append(updatedPremajorVersion.Pre, pre)

	return updatedPremajorVersion
}

func minor(currentVersion semver.Version) semver.Version {
	currentMinor := currentVersion.Minor
	newMinor := currentMinor + 1

	newVersion, err := semver.Make("0.0.0")
	if err != nil {
		panic(err)
	}
	newVersion.Major = currentVersion.Major

	// Update from prerelease
	if len(currentVersion.Pre) > 0 && currentVersion.Patch == 0 {
		// When is prerelease from preminor
		newVersion.Minor = currentMinor
	} else {
		newVersion.Minor = newMinor
	}

	return newVersion
}

func preminor(currentVersion semver.Version) semver.Version {
	newVersion := minor(currentVersion)
	updatedPreminorVersion := newVersion
	pre := semver.PRVersion{
		VersionNum: 0,
		IsNum:      true,
	}
	updatedPreminorVersion.Pre = append(updatedPreminorVersion.Pre, pre)

	return updatedPreminorVersion
}

func patch(currentVersion semver.Version) semver.Version {
	currentPatch := currentVersion.Patch
	newPatch := currentPatch + 1

	newVersion, err := semver.Make("0.0.0")
	if err != nil {
		panic(err)
	}
	newVersion.Major = currentVersion.Major
	newVersion.Minor = currentVersion.Minor

	// Update from prerelease
	if len(currentVersion.Pre) == 0 {
		newVersion.Patch = newPatch
	} else {
		newVersion.Patch = currentPatch
	}
	return newVersion
}

func prepatch(currentVersion semver.Version) semver.Version {
	newVersion := patch(currentVersion)
	updatedVersion := newVersion
	pre := semver.PRVersion{
		VersionNum: 0,
		IsNum:      true,
	}

	updatedVersion.Pre = append(updatedVersion.Pre, pre)

	return updatedVersion
}

func prerelease(currentVersion semver.Version) semver.Version {
	updatedPRVersion := currentVersion
	currentPre := updatedPRVersion.Pre
	updatedPre := currentPre
	for i, v := range currentPre {
		if v.IsNum {
			updatedPre[i].VersionNum = v.VersionNum + 1
		}
	}

	updatedPRVersion.Pre = append(updatedPre)

	log.Info().Msgf("%v", updatedPre[0])

	return updatedPRVersion
}

func createNewSemVerVersion(currentVersion semver.Version, partToUpdate VersionAction) semver.Version {
	var newVersion semver.Version

	switch partToUpdate {
	case Major:
		newVersion = major(currentVersion)
	case Minor:
		newVersion = minor(currentVersion)
	case Patch:
		newVersion = patch(currentVersion)
	case Premajor:
		newVersion = premajor(currentVersion)
	case Preminor:
		newVersion = preminor(currentVersion)
	case Prepatch:
		newVersion = prepatch(currentVersion)
	case Prerelease:
		newVersion = prerelease(currentVersion)
	default:
		newVersion = currentVersion
	}

	return newVersion
}

// UpdatePackage do 2 things
// 1. Update file that indicate version for the package. `pubspec.yaml` for flutter,
// `package.json` for npm package.
// 2. Create git tag and commit the changes to git.
func UpdatePackage() {
	// Handle
}

func UpdateVersion(
	packageType package_detail.PackageType,
	currentVersion string,
	versionAction VersionAction,
) string {

	var currentSemverVersion string
	var newVersionString string

	if packageType == package_detail.Flutter {
		flutterVersion := flutter.GetFlutterVersion(currentVersion)
		currentSemverVersion = flutterVersion.VersionName

		updateVersionName := UpdateSemverVersion(currentSemverVersion, versionAction)
		updateVersionCode := flutterVersion.VersionCode + 1

		updatedFlutterVersion := flutter.FlutterVersion{
			VersionName: updateVersionName,
			VersionCode: updateVersionCode,
		}

		newVersionString = updatedFlutterVersion.String()
	}

	if packageType == package_detail.Npm {
		log.Error().Msg("Not implement yet for NPM")
		currentSemverVersion = currentVersion
		updateVersion := UpdateSemverVersion(currentSemverVersion, versionAction)
		newVersionString = updateVersion

	}

	return newVersionString
}

func UpdateSemverVersion(currentSemverVersion string, versionAction VersionAction) string {
	parsedCurrentVersion, err := semver.Make(currentSemverVersion)
	newVersion := createNewSemVerVersion(parsedCurrentVersion, versionAction)

	newVersionString := newVersion.String()
	if err != nil {
		log.Fatal()
	}
	return newVersionString
}

func GetPackageTypeFromCurrentDir() (package_detail.PackageType, error) {
	filenames, err := listDirFilenames()
	if err != nil {
		return package_detail.None, err
	}
	return GetPackageType(filenames)
}

func GetPackageType(filenames []string) (package_detail.PackageType, error) {

	if isNpmPackage(filenames) {
		return package_detail.Npm, nil
	}
	if isFlutterPackage(filenames) {
		return package_detail.Flutter, nil
	}

	return package_detail.None, errors.New("Cannot determine package type.")
}

func listDirFilenames() ([]string, error) {
	var filenames []string
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal().Msgf("%s", err)
		return filenames, err
	}

	filenames = collection.Map(files)

	return filenames, nil
}

func isNpmPackage(filenames []string) bool {
	// Check if current directory contain `package.json` file
	return collection.Includes(filenames, "package.json")
}

func isFlutterPackage(filenames []string) bool {
	// Check if current directory contain `pubspec.yaml`
	return collection.Includes(filenames, "pubspec.yaml")
}

func StringActionToVersionAction(action string) (VersionAction, error) {
	switch action {
	case "major":
		return Major, nil
	case "minor":
		return Minor, nil
	case "patch":
		return Patch, nil
	case "prepatch":
		return Prepatch, nil
	case "preminor":
		return Preminor, nil
	case "premajor":
		return Premajor, nil
	case "prerelease":
		return Prerelease, nil
	default:
		return NoAction, errors.New("Action no support")
	}
}
