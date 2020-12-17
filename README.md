# arbitraryReleaser
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/arbitraryReleaser)](https://goreportcard.com/badge/github.com/Jmainguy/arbitraryReleaser)
[![Release](https://img.shields.io/github/release/Jmainguy/arbitraryReleaser.svg?style=flat-square)](https://github.com/Jmainguy/arbitraryReleaser/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/arbitraryReleaser/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/arbitraryReleaser?branch=main)

This command is designed to to bump the version in CHANGELOG.md and RELEASE

## Usage
Usage of arbitraryReleaser:
  -c value
    	Comment to explain whats included in version
  -changelog string
    	Path to CHANGELOG.md (default "./CHANGELOG.md")
  -major
    	Bump major version
  -minor
    	Bump minor version (default true)
  -patch
    	Bump patch version
  -version string
    	Path to VERSION (default "./VERSION")


```/bin/bash
./arbitraryReleaser -c "Fixed tests" -c "Added support for github actions" -c "Increased synergy"
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/arbitraryReleaser/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```
