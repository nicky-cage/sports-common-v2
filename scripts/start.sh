#!/bin/bash


# 捕捉 ctrl+c
trap "cd ../" SIGINT

app_name=`pwd | awk -F '/' '{ print $NF }'`
if [[ "$app_name" != "admin" && \
    "$app_name" != "member-api" ]]; then
    app_name=`echo $0 | sed 's/\/start\.sh//g' | awk -F '/' '{ print $NF }'`
fi

project_name="sports-${app_name}"
project_dir="/data/sports/${app_name}"

if [[ ! -d $project_dir ]]; then
    echo "缺少目录: $project_dir"
    exit
fi

# 源码目录 ./src
if [[ ! -e ./src ]]; then
    echo "缺少目录: ./src"
    exit
fi

process_id=`ps aux | grep -v grep | grep $project_name | awk '{ print $2 }'`
if [[ "$process_id" != "" ]]; then
    echo "停止现有进程: $process_id"
    kill -9 $process_id
fi

cd ./src

echo "目录名称: $project_name"

echo "删除旧的编译程序: $project_name"
rm -rf $project_name

echo "编译代码 ..."
go build -o $project_name main.go
echo "编译完成."

if [[ "$?" != "0" ]]; then
    echo "编译失败 "
    exit
fi

echo "运行程序 ..."
echo "-----------------------------------------------------------------------------------------------------------"
./$project_name -config=./setting.ini -ext=./setting_ext.ini
