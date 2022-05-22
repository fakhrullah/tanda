package version

import (
	"testing"

	package_detail "fajarhac.com/fakhrullah/tanda/package"
)

func TestGetPackage_shouldReturnFlutter(t *testing.T) {
	list := []string{"miaw", "hoit", "pubspec.yaml", "meh"}

	isFlutter, err := GetPackageType(list)

	if isFlutter != package_detail.Flutter || err != nil {
		t.Fatalf("Expect flutter but err found: %s", err)
	}
}

func TestGetPackage_shouldReturnNpm(t *testing.T) {
	list := []string{"miaw", "hoit", "meh", "package.json", "src"}

	isNpm, err := GetPackageType(list)

	if isNpm != package_detail.Npm {
		t.Fatalf("Expect Npm but err found: %s", err)
	}
}

func TestGetPackage_shouldError(t *testing.T) {
	nonFlutterNorNpmDir := []string{"miaw", "hoit", "meh"}
	_, err := GetPackageType(nonFlutterNorNpmDir)

	if err == nil {
		t.Fatalf("Expect error thrown but no error. Error: %s", err)
	}
}

func TestUpdateVersion_shouldUpdateMajor(t *testing.T) {
	currentVersion := "32.576.45+23"

	updateMajor := UpdateVersion(package_detail.Flutter, currentVersion, Major)

	if updateMajor != "33.0.0+24" {
		t.Fatalf("Expect new version: %s but got %s", updateMajor, "33.0.0+24")
	}

}

func TestUpdateVersion_shouldUpdatePremajor(t *testing.T) {
	currentVersion := "6.32.14+64"
	updateMajor := UpdateVersion(package_detail.Flutter, currentVersion, Premajor)
	if updateMajor != "7.0.0-0+65" {
		t.Fatalf("Expect new version: %s but got %s", "7.0.0-0+65", updateMajor)
	}
}

func TestUpdateVersion_shouldUpdateMajorFromPrerelase(t *testing.T) {
	currentVersion := "1.42.2-3+7"
	updateMajorFromPrerelease := UpdateVersion(package_detail.Flutter, currentVersion, Major)
	if updateMajorFromPrerelease != "2.0.0+8" {
		t.Fatalf("Expect new version: %s but got %s", "2.0.0+8", updateMajorFromPrerelease)
	}

	// Update major from prerelease premajor
	currentVersion2 := "2.0.0-3+7"
	updatePremajor2 := UpdateVersion(package_detail.Flutter, currentVersion2, Major)
	if updatePremajor2 != "2.0.0+8" {
		t.Fatalf("Expect new version: %s but got %s", "2.0.0+8", updatePremajor2)
	}
}

func TestUpdateVersion_shouldUpdateMinor(t *testing.T) {
	currentVersion := "8.35.980-beta"

	updateMinor := UpdateVersion(package_detail.Npm, currentVersion, Minor)

	if updateMinor != "8.36.0" {
		t.Fatalf("Expect new version: %s but got %s", updateMinor, "8.36.0")
	}
}

func TestUpdateVersion_shouldUpdatePreMinor(t *testing.T) {
	currentVersion := "2.12.43+54"

	updatedPreMinor := UpdateVersion(package_detail.Flutter, currentVersion, Preminor)

	if updatedPreMinor != "2.13.0-0+55" {
		t.Fatalf("Expect new version: %s but got %s", "2.13.0-0+55", updatedPreMinor)
	}
}

func TestUpdateVersion_shouldUpdateMinorFromPreRelease(t *testing.T) {
	// Test when update minor from prerelase on patch
	currentVersion := "12.879.78-7+699"
	updateMinor := UpdateVersion(package_detail.Flutter, currentVersion, Minor)
	if updateMinor != "12.880.0+700" {
		t.Fatalf("Expect new version: %s but got %s", "12.880.0+700", updateMinor)
	}

	// Test when update minor from preminor
	currentVersion2 := "2.14.0-4+49"
	updateMinor2 := UpdateVersion(package_detail.Flutter, currentVersion2, Minor)
	if updateMinor2 != "2.14.0+50" {
		t.Fatalf("Expect new version: %s but got %s", "2.14.0+50", updateMinor2)
	}
}

func TestUpdateVersion_shouldUpdatePatch(t *testing.T) {
	currentVersion := "0.0.1-alpha.preview+123.github"

	updatePatch := UpdateVersion(package_detail.Npm, currentVersion, Patch)

	if updatePatch != "0.0.2" {
		t.Fatalf("Expect new version: %s but got %s", updatePatch, "0.0.2")
	}
}

func TestUpdateVersion_shouldUpdatePatchFromPreRelease(t *testing.T) {
	currentVersion := "2.1.4-14+32"

	updatePatch := UpdateVersion(package_detail.Flutter, currentVersion, Patch)

	if updatePatch != "2.1.4+33" {
		t.Fatalf("Expect new version: %s but got %s", "2.1.4+33", updatePatch)
	}
}

func TestUpdateVersion_shouldUpdatePrepatch(t *testing.T) {
	currentVersion := "8.35.14+234"
	updatedPrepatch := UpdateVersion(package_detail.Flutter, currentVersion, Prepatch)
	if updatedPrepatch != "8.35.15-0+235" {
		t.Fatalf("Expect new Flutter pacakge version %s but got %s", "8.35.15-0+235", updatedPrepatch)
	}

	// currentVersionNpmAlreadyPrepatched := "2.12.76-5"
	// updatedPrepatchNpm := UpdateVersion(package_detail.Npm, currentVersionNpmAlreadyPrepatched, Prepatch)
	// if updatedPrepatchNpm != "2.12.77-0" {
	// 	t.Fatalf("Expect new Npm pacakge version %s but got %s", updatedPrepatchNpm, "2.12.76-6")
	// }
}

func TestUpdateVersion_shouldUpdatePrereleaseNumber(t *testing.T) {
	currentVersion := "2.3.12-5+23"
	updatedPrerelease := UpdateVersion(package_detail.Flutter, currentVersion, Prerelease)
	if updatedPrerelease != "2.3.12-6+24" {
		t.Fatalf("Expect new Flutter pacakge version %s but got %s", "2.3.12-6+24", updatedPrerelease)
	}

	currentVersion2 := "4.3.12-beta.12+54"
	updatedPrerelease2 := UpdateVersion(package_detail.Flutter, currentVersion2, Prerelease)
	if updatedPrerelease2 != "4.3.12-beta.13+55" {
		t.Fatalf("Expect new Flutter pacakge version %s but got %s", "4.3.12-beta.13+55", updatedPrerelease2)
	}

	currentVersion3 := "0.1.1-0+11"
	updatedPrerelease3 := UpdateVersion(package_detail.Flutter, currentVersion3, Prerelease)
	if updatedPrerelease3 != "0.1.1-1+12" {
		t.Fatalf("Expect new Flutter pacakge version %s but got %s", "0.1.1-1+12", updatedPrerelease3)
	}
}

func TestUpdateSemverVersion_shouldUpdateSemver(t *testing.T) {
	currentVersion := "14.4.65"

	updatedMajor := UpdateSemverVersion(currentVersion, Major)

	if updatedMajor != "15.0.0" {
		t.Fatalf("Expect new semver version: %s but got %s", updatedMajor, "15.0.0")
	}
}
