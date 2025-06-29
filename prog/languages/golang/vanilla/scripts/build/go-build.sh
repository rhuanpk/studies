#!/bin/bash

[ "$1" = '-h' ] && {
	cat <<- eof
		go-build args:
		- 1: main.go path
		- 2: output path
		- 3: output name
	eof
	exit
}

declare -A oses=(
	['windows']='exe'
	['darwin']='app'
	['linux']='bin'
)
archs=(
	'amd64'
	'arm64'
)

main="${1:-.}"
path="${2:-.}"
name="${3:-app}"

for arch in "${archs[@]}"; do
	for os in "${!oses[@]}"; do
		ext="${oses[$os]}"
		app="${name}_$os-$arch.$ext"
		GOOS="$os" GOARCH="$arch" \
		go build \
			-ldflags '-s -w' \
			-o "${path%/}/$app" \
			"$main"
		echo "$app OK"
	done
done
