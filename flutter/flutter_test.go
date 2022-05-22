package tanda

import "testing"

func TestGetFlutterVersion(t *testing.T) {
	flutterVersion := GetFlutterVersion("3.41.97+47")

	versionName := flutterVersion.VersionName
	versionCode := flutterVersion.VersionCode

	if versionName != "3.41.97" {
		t.Fatalf("Expect versionName %s but got %s", "3.41.97", versionName)
	}

	if versionCode != 47 {
		t.Fatalf("Expect versionCode be %v but got %v", 47, versionCode)
	}
}
