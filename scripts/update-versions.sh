#!/usr/bin/env bash
set -e # exit if a command fails
set -u # unbound variable is an error

HELP=$(cat <<EOF
Usage: update-version.sh [-s] <docs path> <old version> <new version>

    <docs path>      The path of the docs folder that need update (i.e. docs-ai/en)
    <old version>    The version to replace throughout the
                     (i.e. {{site.ALLUXIO_VERSION_STRING}} or
                     {{site.ALLUXIO_OPERATOR_VERSION_STRING}} or
                     existing versions 2.2.0-SNAPSHOT)
    <new version>    The new version in the codebase (i.e. 2.2.0-RC1)

EOF
)

function update_string() {
    find "${1}" -name '*.md' | xargs -t -n 1 perl -pi -e "s/\Q${2}\E/${3}/g"
}


function main() {
    if [[ "$#" -lt 3 ]]; then
        echo "Arguments '<docs path>', '<old version>' and '<new version>' must be provided."
        echo "${HELP}"
        exit 1
    fi
    local _path="${1}"
    local _old="${2}"
    local _new="${3}"

    update_string "$_path" "$_old" "$_new"

    exit 0
}

main "$@"
