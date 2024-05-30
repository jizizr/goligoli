#!/bin/bash
current_dir_name=${PWD##*/}
if [[ "$current_dir_name" == "service" ]]; then
  echo "Current directory is: $current_dir_name"
  # 列出当前目录的所有文件夹
  for dir in */; do
    # 检查是否为目录
    if [ -d "$dir" ]; then
      echo "Entering directory: $dir"
      cd "$dir/etcd"
      go run . &
      cd -
    fi
  done
else
  echo "Current directory is not service. Script will not run."
fi
