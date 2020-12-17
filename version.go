package main

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"os"
	"strings"
)

func getVersion(path string) string {
	lines, err := file2lines(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var line string
	if len(lines) > 0 {
		line = lines[0]
	}
	return line
}

func getChangelog(path string) string {
	lines, err := file2lines(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var l string
	if len(lines) > 4 {
		l = lines[4]
	}
	line := strings.Split(l, "* ")
	if len(line) == 0 {
		fmt.Printf("Error, line didnt split. line: %v\n", line)
		os.Exit(1)
	}
	return line[1]
}

func compareVersion(version, changelog string) int {
	version1, err := semver.NewVersion(version)
	if err != nil {
		panic(err)
	}
	version2, err := semver.NewVersion(changelog)
	if err != nil {
		panic(err)
	}

	result := version1.Compare(version2)

	return result
}

func writeChanges(comments arrayFlags, changelogPath, versionPath, version string) {
	// Add comments to CHANGELOG.md
	for _, c := range comments {
		formattedComment := "  * " + c + "\n"
		err := insertStringToFile(changelogPath, formattedComment, 4)
		if err != nil {
			panic(err)
		}
	}
	// Add version to CHANGELOG.md
	formattedVersion := "* " + version + "\n"
	err := insertStringToFile(changelogPath, formattedVersion, 4)
	if err != nil {
		panic(err)
	}
	// Add version to VERSION
	f, err := os.Create(versionPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(version)
	if err != nil {
		panic(err)
	}
}

func bumpVersion(major, minor, patch bool, version string) string {
	var sv semver.Version
	semVersion, err := semver.NewVersion(version)
	if err != nil {
		panic(err)
	}
	if major {
		sv = semVersion.IncMajor()
	} else if patch {
		sv = semVersion.IncPatch()
	} else {
		sv = semVersion.IncMinor()
	}
	versionString := sv.Original()
	return versionString
}
