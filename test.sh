#!/usr/bin/env bash
set -e

PKGS=$(go list ./... | grep -v /examples | grep -v /vendor)

go test $PKGS -cover
go vet $PKGS
