#!/bin/bash

BASE_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly BASE_DIR

parse_platform () {
    local platform

    platform=${1:-$(go version | awk '{ print $4 }')}

    os=$(echo "${platform}" | awk -F '/' '{ print $1 }')
    arch=$(echo "${platform}" | awk -F '/' '{ print $2 }')

#    echo "$os" "$arch"

    export CGO_ENABLED=0
    export GOOS=$os
    export GOARCH=$arch
}

gen_bin_file () {
    echo "${APP_NAME}_${GOOS}_${GOARCH}"
}


do_build () {
    local output_file
    output_file=$(realpath "$BASE_DIR/../build")/$(gen_bin_file)
    local cmd_line=(go build --ldflags="-s -w" "$@" -o "$output_file" "$MODULE_PATH")
    echo "[ OS: $os  Architecture: $arch ]"
    echo "${cmd_line[@]}"
    "${cmd_line[@]}"
}


# $1 应用名称
# $2 模块路径
main () {
    APP_NAME=${1:-app}
    MODULE_PATH=${2:-'./cmd'}

    source "$BASE_DIR/supported_platforms.sh"

    shift 2
    for platform in "${SUPPORTED_PLATFORMS_LIST[@]}"; do
        parse_platform "$platform"
        do_build "$@"
    done
}

main "$@"
