#!/usr/bin/env bash
set -e

# get all go files that aren't under vendor directory
go_files=$(find . -path '*/vendor/*' -prune -o -name '*.go' -type f -print)

gofmt="gofmt -s -d"

if [[ "$1" == "fix" ]]; then
	gofmt="${gofmt} -w"
fi

exit_code=0
for go_file in ${go_files}; do
	cmd="$gofmt $go_file"
	echo "==> $cmd"
	$cmd
	test -z $(gofmt -l ${go_file})
	[[ $? -ne 0 ]] && exit_code=1
done
exit ${exit_code}
