package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var comments arrayFlags
	flag.Var(&comments, "c", "Comment to explain whats included in version")
	majorPtr := flag.Bool("major", false, "Bump major version")
	minorPtr := flag.Bool("minor", true, "Bump minor version")
	patchPtr := flag.Bool("patch", false, "Bump patch version")
	changelogPtr := flag.String("changelog", "./CHANGELOG.md", "Path to CHANGELOG.md")
	versionPtr := flag.String("version", "./VERSION", "Path to VERSION")
	flag.Parse()
	if len(comments) == 0 {
		fmt.Println("Please pass a comment with -c for the release")
		os.Exit(1)
	}
	versionString := getVersion(*versionPtr)
	changelogString := getChangelog(*changelogPtr)
	result := compareVersion(versionString, changelogString)
	if result == 0 {
		version := bumpVersion(*majorPtr, *minorPtr, *patchPtr, versionString)
		fmt.Printf("Bumping %s and %s to version %s\n", *versionPtr, *changelogPtr, version)
		writeChanges(comments, *changelogPtr, *versionPtr, version)
	} else {
		msg := "Version in " + *changelogPtr + " is " + changelogString + " which does not match " + versionString + " which is in " + *versionPtr
		fmt.Println(msg)
	}
}
