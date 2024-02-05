#!/bin/bash

#project_dir="~/web/"
project_dir="/Users/ai/work/sports"
projects="common models admin member-api game-api game-cron game-cron-consumer"

# pull
sports_pull() { 
  for p in $projects; do
    cd $p
    echo "当前路径: $PWD"
    pull
    cd ../
  done
}

# push
sports_push() { 
  for p in $projects; do
    cd $p
    echo "项目名称: $p"
    echo "当前路径: $PWD"
    echo ""
    push
    cd ../
  done
}

# checkout
sports_checkout() {
  branch="$1"
  for p in $projects; do
    cd $p
    echo "当前路径: $PWD"
    git checkout $branch
    cd ../
    echo "+-------------------------------------------------------------+"
    echo
  done
}

# status
sports_status() { 
  for p in $projects; do
	echo "目录: $p"
    cd $p
    echo "当前路径: $PWD"
    git status
    cd ../
    echo "+-------------------------------------------------------------+"
    echo
  done
}

# show
sports_show() { 
  echo "Projects:"
  echo "+-------------------------------------------------------------+"
  for p in $projects; do
      echo "project: $p"
  done
  echo ""
  echo "Command:"
  echo "+-------------------------------------------------------------+"
  echo "pull | push | status | show | st"
}

# main
sports_main() {
  case "$1" in
    "push")
      sports_push;;
    "pull")
      sports_pull;;
    "status")
      sports_status;;
    "st")
     sports_status;;
      "tidy")
      for f in $projects; do
        if [ -d $f/src ]; then
            cd $f/src
            echo $PWD
            go mod tidy
            go vet
            cd -
        fi
      done;;
    "co")
      if [[ $# -lt 2 ]]; then
        echo "缺少参数: 分支名称"
        exit
      fi
      sports_checkout $2;;
    *)
      sports_show;;
  esac
}

# main
sports_main $@
