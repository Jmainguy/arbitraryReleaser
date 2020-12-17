package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile2Lines(t *testing.T) {
	_, err := copy("./examples/CHANGELOG.md", "./tests/CHANGELOG.md")
	assert.Equal(t, err, nil, "Error should be nil")
	lines, err := file2lines("./tests/CHANGELOG.md")
	assert.Equal(t, len(lines), 12, "CHANGELOG.md should be 12 lines long.")
	assert.Equal(t, err, nil, "Error should be nil")
}

func TestInsertStringToFile(t *testing.T) {
	_, err := copy("./examples/CHANGELOG.md", "./tests/CHANGELOG.md")
	assert.Equal(t, err, nil, "Error should be nil")
	err = insertStringToFile("./tests/CHANGELOG.md", "Test\n", 4)
	assert.Equal(t, err, nil, "Error should be nil")
	lines, err := file2lines("./tests/CHANGELOG.md")
	assert.Equal(t, err, nil, "Error should be nil")
	assert.Equal(t, lines[4], "Test", "Line should equal 'Test'")
}
