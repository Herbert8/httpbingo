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
    # 输出目录
    local output_dir=$BASE_DIR/../build
    # 由于 macOS 版的 realpath 必须针对存在的路径操作，所以先创建，再取绝对路径
    mkdir -p "$output_dir" && output_dir=$(realpath "$output_dir")

    # 输出文件
    local output_file
    output_file=$output_dir/$(gen_bin_file)

    local cmd_line=(go build --ldflags="-s -w" "$@" -o "$output_file" "$MODULE_PATH")
    echo "[ OS: $os  Architecture: $arch ]"
    echo "${cmd_line[@]}"
    "${cmd_line[@]}"
    OUTPUT_FILE_ARRAY+=("$output_file")
}

upx_compress () {
    upx "${OUTPUT_FILE_ARRAY[@]}"
}

# $1 应用名称
# $2 模块路径
main () {
    APP_NAME=${1:-app}
    MODULE_PATH=${2:-'./cmd'}

    source "$BASE_DIR/supported_platforms.sh"

    shift 2
    OUTPUT_FILE_ARRAY=()
    for platform in "${SUPPORTED_PLATFORMS_LIST[@]}"; do
        parse_platform "$platform"
        do_build "$@"
    done
#    upx_compress
}

main "$@"
