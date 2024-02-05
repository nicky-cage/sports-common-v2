#!/usr/bin/env bash

#set -x
trap 'cd ../;' SIGINT

# 应用名称
app_name=''
# 包命名规则 - 开发阶段: sports-admin-dev-201222_1326
project_dir='/Users/ai/work/sports'
project_name='sports-'                  # 编译生成的应用程序名称
remote_dir='/data/sports/'              # 远程部署目录
remote_user='ubuntu'                    # 远程用户
remote_host='sports.test'               # 远程主机
remote_args=''                          # 远程需要执行的额外参数
app_version="2.0"						# 版本
if [[ "$SPORT_VERSION" != "" ]]; then
    app_version="${SPORT_VERSION}"
fi

# 本地运行 - 位于src目录执行
run_local() {
    if [[ -e ~/go/bin/realize ]]; then
        ~/go/bin/realize start
    elif [[ -e ~/.go/bin/realize ]]; then
        ~/.go/bin/realize start
    else
        echo "缺少文件: realize"
    fi
}

# 上传到运端服务器文件
upload_file() {
    scp $1 $2
}

# 远程部署 - 位于src目录执行
run_deploy() {
    echo "|->>> 当前目录: $PWD"
    echo "|->>> 远程主机: ${remote_user}@${remote_host} (${remote_key_file})"

    project_name="sports-${app_name}-${app_version}" # 程序名称
    deploy_dir="${remote_dir}/${app_name}-${app_version}" # 程序部署目录 - 远程服务器上
    build_version="`date '+%y%m%d_%H%M'`"
    ###########################################################################
    # 1. 创建目录
    if [[ ! -d target ]]; then
        echo "|-> 创建目录: target"
        mkdir target
    fi
    temp_dir="target/${project_name}-compiled-${build_version}"
    if [[ -d $temp_dir ]]; then
        echo "|-> 删除之前存在的目录: $temp_dir ..."
        rm -rf $temp_dir
    fi
    echo "|-> 创建临时目录: $temp_dir ..."
    mkdir $temp_dir

    echo "|-> 进入目录: ./src"
    current_project_dir="${project_dir}/${app_name}"
    cd $current_project_dir/src

    echo "|-> 检查三方依赖 ..."
    go mod tidy
    if [[ "$?" != "0" ]]; then
        exit
    fi

    ###########################################################################
    # 2. 编译程序 - 上传程序
    echo "|-> 编译程序 ..."
    temp_bin="${project_name}-dev-${build_version}"
    # 判断是否是Mac系统/android, 此时需要交叉编译
    if [[ "`uname`" == "Darwin" || "`uname -a | grep -v grep | grep Android`" != "" ]]; then
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../$temp_dir/$temp_bin  main.go
    else
        go build -o ../$temp_dir/$temp_bin  main.go # 编译程序
    fi
    if [[ "$?" != "0" ]]; then
        echo "|-> 编译程序失败"
        exit
    fi
    echo "|-> 上传二进制程序至远程服务器 ..."
    cd ../${temp_dir}
    remote_bin="${project_name}-dev-${build_version}"
    upload_file $temp_bin ${remote_user}@${remote_host}:${deploy_dir}/${remote_bin}

    ###########################################################################
    # 3. 打包模板文件 - 上传模板文件
    echo "|-> 从远程服务器重启程序 ..."
    if [[ "$remote_host" == "prod" || "$remote_host" == "prod2" || "$remote_host" == "sports.test" ]]; then
        echo "|-> 远程服务器: ${remote_host} - ${app_name}"
        bin_log="${project_name}-bin.log" # 日志文件名称
        if [[ "$app_name" == "admin" ]]; then # 如果管理后台,需要上传模板
            template_dir="templates"
            temp_tpl="templates-${build_version}.tar.gz"
            cd $current_project_dir
            if [[ ! -d ${template_dir} ]]; then # 如果存在模板目录
                echo "错误: 无法找到模板目录"
                exit
            fi
#
            echo "|-> 打包管理后台模板文件: ${template_dir} ..."
            tar czf $temp_tpl ${template_dir}
            echo "|-> 上传管理后台模板文件 ..."
            upload_file $temp_tpl ${remote_user}@${remote_host}:${deploy_dir}/
            echo "|-> 删除本地临时文件: $temp_tpl ..."
            rm -rf $temp_tpl
            cd $temp_dir
            ssh ${remote_user}@${remote_host} 2>&1 << COMMAND
                cd ${deploy_dir}
                rm -rf ${template_dir}
                tar zxf ${temp_tpl}
                rm -rf ${temp_tpl}
                chmod +x ${temp_bin}
                rm -rf ${project_name}
                ln -sf ${temp_bin} ${project_name}
                ./restart.sh
                exit
COMMAND
        else
            ssh ${remote_user}@${remote_host} 2>1 << COMMAND
                cd ${deploy_dir}
                chmod +x ${temp_bin}
                rm -rf ${project_name}
                ln -sf ${temp_bin} ${project_name}
                ./restart.sh
                exit
COMMAND
        fi
    fi

    # 关于生产环境三/测试环境的处理
    if [[ "$remote_host" == "prod3" || "remote_host" == "TJ_TEST" ]]; then
        echo "|-> 远程服务器: ${remote_host} - ${app_name}"
        if [[ "$remote_host" == "prod3" ]]; then
            if [[ "$app_name" == "game-api" || "$app_name" == "payment-cron" ]]; then
                ssh ${remote_user}@${remote_host} 2>1 << COMMAND
                    cd ${deploy_dir}
                    chmod +x ${temp_bin}
                    rm -rf ${project_name}
                    ln -sf ${temp_bin} ${project_name}
                    ./restart.sh
                    exit
COMMAND
            fi
        fi
        bin_log="${project_name}-bin.log" # 日志文件名称
        if [[ "$app_name" == "game-cron" || "$app_name" == "game-cron-consumer" ]]; then
            ssh -i $remote_key_file ${remote_user}@${remote_host} 2>1 << COMMAND
                cd ${deploy_dir}
                chmod +x ${temp_bin}
                rm -rf ${project_name}
                ln -sf ${temp_bin} ${project_name}
                ./restart.sh
                exit
COMMAND
        fi
    fi
    echo "|-> 远程服务器重启成功,  请删除临时目录: $temp_dir ..."
    cd ..
}

# 停止服务
stop_service() {
    project_name="sports-${app_name}" # 程序名称
    deploy_dir="${remote_dir}/${app_name}" # 程序部署目录 - 远程服务器上
    ssh -i $remote_key_file ${remote_user}@${remote_host} 2>1 << COMMAND
        cd ${deploy_dir}
        ps aux | grep -v grep | grep ${project_name} | awk '{ print $2 }' | xargs kill -9
        exit
COMMAND
}

# 重启服务 - 等同于启动服务
restart_service() {
    project_name="sports-${app_name}" # 程序名称
    deploy_dir="${remote_dir}/${app_name}" # 程序部署目录 - 远程服务器上
    bin_log="${project_name}-bin.log" # 日志文件名称
    ssh -i $remote_key_file ${remote_user}@${remote_host} 2>1 << COMMAND
        cd ${deploy_dir}
        ./restart.sh
        exit
COMMAND
}

# 清理所有的文件
clean_target() {
    if [[ ! -d ./target ]]; then
        echo "缺少目录： ./target"
        return
    fi
    count=0
    echo "删除 target 目录下生成文件"
    for d in `ls ./target`; do
        full_path="./target/${d}"
        echo "cleaning directory: $full_path"
        rm -rf $full_path
        ((count=$count+1))
    done
    echo "删除 src/sports-*相关文件"
    if [[ -d ./src ]]; then
        for d in `ls ./src/sports-*`; do
            if [[ -f $d ]]; then
                echo "删除可执行文件: $d"
                rm -rf $d
                ((count=$count+1))
            fi
        done
    fi
    echo "deleted total directories: $count"
}

# 启动服务
start_service() {
    restart_service
}

# 显示帮助信息
show_help() {
    echo "用法: ./run.sh dev|deploy|start|stop|restart|clean [prod] [admin|member-api|game-api|oss|game-cron|game-cron-consumer|quantum-bank]"
    echo "参数: "
    echo "      dev:        开发模式"
    echo "      deploy:     部署到远程开发服务器"
    echo "      start:      启动服务"
    echo "      stop:       停止服务"
    echo "      restart:    重启服务"
    echo
    echo "      admin:                  管理后台"
    echo "      member-api:             前端后台接口服务"
    echo "      game-api:               前端游戏接口服务"
    echo "      game-cron:              游戏拉单服务"
    echo "      game-cron-consumer:     游戏拉单消费者服务"
    echo "      quantum-bank:           量子银行"
    exit
}

# 主函数
main() {

    # 操作说明 - 仅限已经指定命令
    if [[ "$1" != "dev" && "$1" != "deploy" && "$1" != "restart" && "$1" != "start" && "$1" != "clean" ]]; then
        echo "错误: 参数有误! 参数仅限于 dev|deploy|restart|start|clean"
        show_help
    fi
    # 如果参数仅仅是清理
    if [[ "$1" == "clean" ]]; then
        clean_target
        exit
    fi
    # 需要指定应用名称
    if [[ "$3" == "" ]]; then
        app_name=`echo $PWD | awk -F '/' '{ print $NF }'`
    else
        app_name="$3"
    fi

    # 命令参数 - 仅限已经指定参数
    if [[ "$app_name" != "admin" && "$app_name" != "member-api" && "$app_name" != "game-api" \
        && "$app_name" != "game-cron" && "$app_name" != "game-cron-consumer" && "$app_name" != "oss" ]]; then
        echo "错误: 必须指定有效应用名称!"
        show_help
    fi

    # 检测环境、进入运行目录
    echo "*** 当前目录: $app_name ($PWD)"
    echo "*** 当前时间: `date '+%Y-%m-%d_%H:%M:%S'` ***"

    case "$1" in
        "dev")
            cd src && run_local && exit ;;
        "deploy")
            run_deploy ;;
        "start")
            start_service ;;
        "restart")
            restart_service ;;
        *)
        echo "-*- 错误 -*-" ;;
    esac
}

# 用法: ./start.sh dev|deploy|release
main $@
