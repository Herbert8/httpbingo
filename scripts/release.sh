#!/usr/bin/env bash

run_release() {
    local target_dir=$1
    gmake release_rebuild
    cp build/* "$target_dir/"
}

download_code() {
    (
        # 克隆仓库
        git clone 'git@github.com:Herbert8/httpbingo.git'

        cd httpbingo || return 2

        # 获取最新的代码和标签
        git fetch --all --tags

        # 获取最新的 release 标签
        local latest_tag
        latest_tag=$(git describe --tags "$(git rev-list --tags --max-count=1)")

        echo "Latest release tag: $latest_tag"

        # 切换到最新的 release 标签
        git checkout "$latest_tag"
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

main() {
    local target_dir=${1:-dist}

    mkdir -p "$target_dir"
    target_dir=$(realpath "$target_dir")

    if ! [[ -d "$target_dir" ]]; then
        echo 'Target path does not exist.' >&2
        exit 2
    fi

    local work_dir
    work_dir=$(mktemp -d)
    (
        cd "$work_dir" &&
            download_code &&
            cd "$work_dir/httpbingo" &&
            run_release "$target_dir"
    ) >/dev/null
    rm -rf "$work_dir"
    echo
    list "$target_dir"
}

main "$@"
