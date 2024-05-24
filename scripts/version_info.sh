#!/usr/bin/env bash

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
readonly SCRIPT_DIR

PROJECT_ROOT=$(realpath "$SCRIPT_DIR/..")
readonly PROJECT_ROOT

brief_ver_info () {
    # 获取Git信息
    commit_hash=$(git rev-parse --short HEAD)
    branch_name=$(git rev-parse --abbrev-ref HEAD)
    tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "no-tag")
    build_time=$(date +%Y%m%d.%H%M%S)

    # 生成版本号
    if [ "$tag" != "no-tag" ]; then
      version="$tag"
    else
      version="1.0.0-$branch_name-$build_time+$commit_hash"
    fi

#    echo "Version file created with the following content:"

    # 创建版本文件
    cat <<EOL
Version: ${version}
Commit Hash: ${commit_hash}
Branch: ${branch_name}
Tag: ${tag}
Build Time: ${build_time}
EOL

}

detail_ver_info () {
    # 获取Git信息
    commit_hash=$(git rev-parse HEAD)
    branch_name=$(git rev-parse --abbrev-ref HEAD)
    tag=$(git describe --tags --abbrev=0 2>/dev/null)
    commit_message=$(git log -1 --pretty=%B)
    author_name=$(git log -1 --pretty=%an)
    commit_date=$(git log -1 --pretty=%ad)
    repository_url=$(git config --get remote.origin.url)

    # 检查工作目录状态
    dirty_state=$(git diff --quiet || echo "dirty")

    # 获取详细的更改信息
    staged_changes=$(git diff --cached --name-status)
    unstaged_changes=$(git diff --name-status)
#    untracked_files=$(git ls-files --others --exclude-standard)

    # 创建 build_info.txt 文件
    cat <<EOL

## Git Commit Information
Commit Hash: ${commit_hash}
Branch: ${branch_name}
Tag: ${tag}
Commit Message: ${commit_message}
Author: ${author_name}
Date: ${commit_date}
Repository URL: ${repository_url}


## Working Directory Status
Dirty State: ${dirty_state}


## Detailed Changes
### Staged Changes
${staged_changes}


### Unstaged Changes
${unstaged_changes}
EOL


### Untracked Files
#${untracked_files}

}

main () {
    local build_info_dir=$PROJECT_ROOT/internal/app/resource/build_info
    mkdir -p "$build_info_dir"

    local brief_ver_file=$build_info_dir/brief_ver_info.txt
    local detail_ver_file=$build_info_dir/detail_ver_info.txt

    brief_ver_info > "$brief_ver_file"
    detail_ver_info > "$detail_ver_file"
}

main
