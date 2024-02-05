#!/bin/bash

current_dir=`echo $PWD | awk -F '/' '{ print $NF }'`
project_name="sports-${current_dir}"

# 主程序
main() {
    clear
    if [[ "$1" == "upload" ]]; then
        upload_files
        exit
    fi

    echo "当前目录: $PWD - $current_dir"
    echo "进程名称: $project_name"
    echo

    if [[ ! -e $project_name ]]; then
        echo "缺少文件: $project_name"
        exit
    fi

    process_list=`ps aux | grep "\./$project_name" | grep -v grep | awk '{ print $2 }'`
    if [[ "$process_list" != "" ]]; then
        echo "准备停止进程列表: "
        echo "|_______________________________________________________________ ..."
        ps aux | grep "\./$project_name" | grep -v grep
        for p in $process_list; do
            echo "停止进程: $p ..."
            kill -9 $p
        done
        echo
    fi

    echo "启动程序: $project_name ..."
    if [[ "$1" == "dev" ]]; then
        ./$project_name -config=./setting.ini -ext=./setting_ext.ini
    else
        bin_log="../logs/${project_name}-bin.log"
        echo "程序日志: $bin_log"
        if [[ "`ulimit -n`" != "8192" ]]; then # 设置ulimit
            ulimit -n 8192
        fi
        echo
        if [[ ! -f ./setting_ext.ini ]]; then
            ./$project_name -config=./setting.ini > $bin_log 2>&1 &
        else
            ./$project_name -config=./setting.ini -ext=./setting_ext.ini > $bin_log 2>&1 &
        fi

        if [[ "$current_dir" == "game-cron" ]]; then
            consumer_log="../logs/${project_name}-consumer.log"
            echo "程序日志: $consumer_log"
            echo
            ./$project_name -consumerName=user_wager -config=./setting.ini -ext=./setting_ext.ini > $consumer_log  2>&1 &
        fi

        echo "重新启动进程列表: "
        echo "|_______________________________________________________________ ..."
        ps aux | grep "\./${project_name}" | grep -v grep
    fi
}

main $@
