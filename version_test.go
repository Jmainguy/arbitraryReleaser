package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := getVersion("./examples/VERSION")
	assert.Equal(t, version, "v0.14.0", "Version should be 'v0.14.0'")
}

func TestGetChangelog(t *testing.T) {
	version := getChangelog("./examples/CHANGELOG.md")
	assert.Equal(t, version, "v0.14.0", "Version should be 'v0.14.0'")
}

func TestCompareVersion(t *testing.T) {
	versionString := getVersion("./examples/VERSION")
	assert.Equal(t, versionString, "v0.14.0", "Version should be 'v0.14.0'")
	changelogString := getChangelog("./examples/CHANGELOG.md")
	assert.Equal(t, changelogString, "v0.14.0", "Version should be 'v0.14.0'")
	result := compareVersion(versionString, changelogString)
	assert.Equal(t, result, 0, "Versions should match between CHANGELOG.md and VERSION")
}

func TestWriteChanges(t *testing.T) {
	var comments []string
	comments = append(comments, "A test")
	writeChanges(comments, "./tests/CHANGELOG.md", "./tests/VERSION", "v0.15.0")
	versionString := getVersion("./tests/VERSION")
	assert.Equal(t, versionString, "v0.15.0", "Version should be 'v0.15.0'")
	changelogString := getChangelog("./tests/CHANGELOG.md")
	assert.Equal(t, changelogString, "v0.15.0", "Version should be 'v0.15.0'")
	result := compareVersion(versionString, changelogString)
	assert.Equal(t, result, 0, "Versions should match between CHANGELOG.md and VERSION")
}

func TestBumpVersion(t *testing.T) {
	versionMajor := bumpVersion(true, true, true, "v0.14.0")
	assert.Equal(t, versionMajor, "v1.0.0", "Version should be 'v1.0.0'")
	versionMinor := bumpVersion(false, true, false, "v0.14.0")
	assert.Equal(t, versionMinor, "v0.15.0", "Version should be 'v0.15.0'")
	versionPatch := bumpVersion(false, true, true, "v0.14.0")
	assert.Equal(t, versionPatch, "v0.14.1", "Version should be 'v0.14.1'")
}
