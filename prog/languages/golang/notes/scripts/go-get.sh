#!/bin/bash

set -o errexit

[ "$1" == '-h' ] && {
cat << \eof
options:
	-c: also executes `go clean`
eof
exit
}

[ "$1" == '-c' ] && {
	echo -e '\e[3;36m> go clean -modcache\e[m'
	go clean -modcache
}

echo -e '\e[3;36m> go list -u -m all\e[m'
go list -u -m all

echo -e '\e[3;36m> go get -u -t ./...\e[m'
go get -u -t ./...

echo -e '\e[3;36m> go get -u -t all\e[m'
go get -u -t all

echo -e '\e[3;36m> go mod tidy\e[m'
go mod tidy
