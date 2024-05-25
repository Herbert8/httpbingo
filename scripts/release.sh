#!/usr/bin/env bash

set -eu

download_code() {
    local work_dir=$1
    (
        local latest_tag
        cd "$work_dir" &&
            # 克隆仓库
            git clone 'git@github.com:Herbert8/httpbingo.git' &&
            cd httpbingo &&
            # 获取最新的代码和标签
            #            git fetch --all --tags &&
            # 获取最新的 release 标签
            latest_tag=$(git describe --tags "$(git rev-list --tags --max-count=1)") &&
            echo "Latest release tag: $latest_tag" >&2 &&
            # 切换到最新的 release 标签
            git checkout "$latest_tag" &&
            # 输出当前目录
            pwd
    )
}

list() {
    cmd_exists() {
        command -v "$1" >/dev/null
    }
    exa_list() {
        exa -Fghl --time-style=long-iso --group-directories-first --color-scale "$@"
    }
    gls_list() {
        gls -lp --time-style=long-iso --group-directories-first --color=auto "$@"
    }
    ls_list() {
        ls -lpG "$@"
    }
    local LL
    if cmd_exists exa; then
        LL=exa_list
    elif cmd_exists gls; then
        LL=gls_list
    else
        LL=ls_list
    fi
    "$LL" "$@"
}

# $1 发布目录
# $2 是否需要下载代码
main() {
    # 处理目标目录
    local target_dir=${1:-dist}
    local need_download_code=${2:-n}
    SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
    readonly SCRIPT_DIR
    PROJECT_ROOT=$(realpath "$SCRIPT_DIR"/..)

    # 创建目标目录
    mkdir -p "$target_dir"
    # 获取目标目录绝对路径
    target_dir=$(realpath "$target_dir")

    # 工作目录
    local work_dir=''
    # 代码目录
    local code_dir=$PROJECT_ROOT

    if [[ "$need_download_code" == 'y' ]]; then
        # 创建临时文件夹作为工作目录
        work_dir=$(mktemp -d)
        # 下载代码到工作目录
        code_dir=$(download_code "$work_dir")
    fi

    # 构建
    gmake -C "$code_dir" release_rebuild &&
        # 复制构建结果
        cp "$code_dir/build"/* "$target_dir/" &&
        echo && list "$target_dir"

    # 清理目录
    if [[ "$need_download_code" == 'y' && -n "$work_dir" && -d "$work_dir" ]]; then
        rm -rf "$work_dir"
    fi
}

main "$@"
