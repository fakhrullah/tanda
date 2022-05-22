#!/usr/bin/env bash

go build -ldflags "-X fajarhac.com/fakhrullah/tanda.BuildVersion=`git tag --sort=-version:refname | head -n 1`" cmd/main.go 

