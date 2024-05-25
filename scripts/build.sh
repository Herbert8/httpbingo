#!/bin/bash

BASE_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly BASE_DIR

# 清理前面输出的指定行数
clear_scroll_lines() {
    local lines_count=${1:-}
    [[ -z "$lines_count" ]] && return
    echo -ne "\033[${lines_count}A"
    for ((i = 0; i < lines_count; i++)); do
        echo -e "\033[K"
    done
    echo -ne "\033[${lines_count}A"
}

# 在指定范围内滚动
# 参考：https://zyxin.xyz/blog/2020-05/TerminalControlCharacters/
print_scroll_in_range() {
    # 默认最多显示滚动行数，默认为 8
    local scroll_lines=${1:-8}
    # 每行字符数，避免折行，默认 120
    local chars_per_line=${2:-120}
    local txt=''
    local last_line_count=0
    while read -r line; do
        line=${line:0:$chars_per_line}
        [[ "${last_line_count}" -gt "0" ]] && echo -ne "\033[${last_line_count}A"
        if [[ -z "$txt" ]]; then
            txt=$(echo -e "\033[2m$line\033[K" | tail -n"$scroll_lines")
        else
            txt=$(echo -e "$txt\n$line\033[K" | tail -n"$scroll_lines")
        fi
        last_line_count=$(($(wc -l <<<"$txt")))
        echo "$txt"
    done
    echo -ne "\033[0m"
    if [[ "$last_line_count" -gt "0" ]]; then
        clear_scroll_lines "$last_line_count"
    fi
}

parse_platform() {
    local platform

    platform=${1:-$(go version | awk '{ print $4 }')}

    os=$(echo "${platform}" | awk -F '/' '{ print $1 }')
    arch=$(echo "${platform}" | awk -F '/' '{ print $2 }')

    #    echo "$os" "$arch"

    export CGO_ENABLED=0
    export GOOS=$os
    export GOARCH=$arch
}

gen_bin_file() {
    echo "${APP_NAME}_${GOOS}_${GOARCH}"
}

do_build() {
    # 输出目录
    local output_dir=$BASE_DIR/../build
    # 由于 macOS 版的 realpath 必须针对存在的路径操作，所以先创建，再取绝对路径
    mkdir -p "$output_dir" && output_dir=$(realpath "$output_dir")

    # 输出文件
    local output_file
    output_file=$output_dir/$(gen_bin_file)

    local cmd_line=(go build --ldflags="-s -w" "$@" -o "$output_file" "$MODULE_PATH")
    echo -e "\033[1m[ OS: $os  Architecture: $arch ]\033[0m"
    echo -e "  \033[32m${cmd_line[@]}\033[0m"
    "${cmd_line[@]}" 2>&1 | print_scroll_in_range 8 120
    OUTPUT_FILE_ARRAY+=("$output_file")
}

upx_compress() {
    upx "${OUTPUT_FILE_ARRAY[@]}"
}

# $1 应用名称
# $2 模块路径
main() {
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
