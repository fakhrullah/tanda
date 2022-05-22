package tanda

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type FlutterVersion struct {
	VersionName string
	VersionCode uint
}

func throwIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

type PubSpecData struct {
	Name    string
	Version string
}

func (fv *FlutterVersion) String() string {
	return fmt.Sprintf("%v+%v", fv.VersionName, fv.VersionCode)
}

func ParsePubspecYaml() PubSpecData {
	var versionString string
	var packageName string

	// read pubspec.yaml
	pubspecYamlContent, err := ioutil.ReadFile("./pubspec.yaml")
	throwIfErr(err)

	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(pubspecYamlContent, &data)
	throwIfErr(err2)

	for key, val := range data {
		if key == "version" {
			versionString = fmt.Sprintf("%v", val)
		}
		if key == "name" {
			packageName = fmt.Sprintf("%v", val)
		}
	}

	return PubSpecData{
		Name:    packageName,
		Version: versionString,
	}
}

func GetFlutterVersion(versionString string) FlutterVersion {
	// Split version to - versionName (semver) & versionCode (buildNumber)
	splittedVersion := strings.Split(versionString, "+")
	versionCodeString := splittedVersion[1]

	versionCode, err := strconv.ParseInt(versionCodeString, 0, 16)
	throwIfErr(err)

	flutterVersion := FlutterVersion{
		VersionName: splittedVersion[0],
		VersionCode: uint(versionCode),
	}

	return flutterVersion
}
