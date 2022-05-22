package package_detail_test

import (
	"io/ioutil"
	"testing"

	package_detail "fajarhac.com/fakhrullah/tanda/package"
)

func TestGetUpdatedPackageConfigFile_shouldUpdateFlutterVersion(t *testing.T) {
	original, err := ioutil.ReadFile("./use_in_test-pubspec.yaml")
	if err != nil {
		t.Fatal(err)
	}

	expected, err := ioutil.ReadFile("./use_in_test-pubspec.yaml.updated")
	if err != nil {
		t.Fatal(err)
	}

	updatedVersion, err := package_detail.GetUpdatedPackageConfigFile(package_detail.Flutter, original, "1.13.0+15")
	if err != nil {
		t.Fatal(err)
	}

	if updatedVersion != string(expected) {
		t.Fatalf("Expect pubspec.yaml content to update version to %v. But got %v", "1.13.0+15", updatedVersion)
	}
}
