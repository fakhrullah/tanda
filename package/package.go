package package_detail

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	flutter "fajarhac.com/fakhrullah/tanda/flutter"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type PackageType uint8

const (
	None    PackageType = 0
	Flutter             = 1
	Npm                 = 2
)

type PackageDetail struct {
	Name     string
	Type     PackageType
	TypeName string
	Version  string
}

func GetPackageDetail() (*PackageDetail, error) {
	// determine what file to read
	// check if pubspecyaml is exist
	isPubspecYamlExist := isFileExists("pubspec.yaml")
	isPackageJsonExist := isFileExists("package.json")

	// throw error when both file exist
	if isPackageJsonExist && isPubspecYamlExist {
		// panic("Cannot determine package type")
		return &PackageDetail{}, errors.New("Cannot determine package type")
	}

	if isPackageJsonExist {
		// packageJsonData := npm.ParsePackageJson()

		packageDetail := PackageDetail{
			Name:     "Not implement yet",
			Type:     Npm,
			TypeName: "NPM",
			Version:  "Not implement yet",
		}
		return &packageDetail, nil
	}

	if isPubspecYamlExist {
		pubspecData := flutter.ParsePubspecYaml()

		packageDetail := PackageDetail{
			Name:     pubspecData.Name,
			Type:     Flutter,
			TypeName: "Flutter",
			Version:  pubspecData.Version,
		}
		return &packageDetail, nil
	}

	// panic("Cannot determine package type")
	return &PackageDetail{}, errors.New("Cannot determine package type")
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func GetPackageTypeName(packageType PackageType) string {
	switch packageType {
	case Npm:
		return "Npm"
	case Flutter:
		return "Flutter"
	default:
		return "None"
	}
}

type T struct {
	Version yaml.Node
}

// Update the pubspec.yaml file with new updated version
// `dryRun` set to true mean, only print out the new update
// pubspec.yaml without updating the file itself.
func UpdatePackageConfigFile(newVersion string, dryRun bool) {
	pubspecYamlContent, err := ioutil.ReadFile("./pubspec.yaml")
	if err != nil {
		panic(err)
	}

	// pubspecDataNode := make(map[interface{}]interface{})
	var pubspecDataNode T
	err = yaml.Unmarshal(pubspecYamlContent, &pubspecDataNode)
	if err != nil {
		panic(err)
	}

	updatedPubspecYamlContentString, err := GetUpdatedPackageConfigFile(Flutter, pubspecYamlContent, newVersion)
	if err != nil {
		panic(err)
	}

	// Print out only or write to pubspec.yaml
	if dryRun {
		fmt.Println()
		fmt.Println("------ New pubspec.yaml should be like below -------")
		fmt.Println(updatedPubspecYamlContentString)
		fmt.Println("-------------------- End file content --------------")
		fmt.Println()
	} else {
		//
		file, err := os.Create("pubspec.yaml")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		file.WriteString(updatedPubspecYamlContentString)
	}

}

// Read package config file (pubspec.yaml or package.json)
// Then return new updated file content with new version
func GetUpdatedPackageConfigFile(
	packageType PackageType,
	originalFileContent []byte,
	newVersion string,
) (string, error) {
	var fileContent string

	if packageType == Flutter {
		var pubspecDataNode T
		err := yaml.Unmarshal(originalFileContent, &pubspecDataNode)
		if err != nil {
			return "", err
		}

		pubspecYamlContentString := string(originalFileContent)
		// Handle Windows new line
		pubspecYamlContentString = strings.ReplaceAll(pubspecYamlContentString, "\r\n", "\n")
		pubspecYamlContentStringLines := strings.Split(pubspecYamlContentString, "\n")

		// Update file with new version
		pubspecYamlContentStringLines[pubspecDataNode.Version.Line-1] = "version: " + newVersion

		updatedPubspecYamlContentString := strings.Join(pubspecYamlContentStringLines, "\n")
		log.Debug().Msg("---- START: Updated Pubspec.yaml content -----\n" + updatedPubspecYamlContentString)
		log.Debug().Msg("---- END Updated Pubspec.yaml content -----")

		fileContent = fmt.Sprint(updatedPubspecYamlContentString)
		return fileContent, nil
	}

	if packageType == Npm {
		return "", errors.New(fmt.Sprintf("Package type %v not implement yet.", GetPackageTypeName(packageType)))
	}

	return "", errors.New(fmt.Sprintf("Package type %v not found.", GetPackageTypeName(packageType)))
}
