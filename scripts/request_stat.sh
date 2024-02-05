#!/bin/bash

# 请求统计
request_stat() {
    log_file="$1"
    awk '
        function get_ms(ms_str) { # 转换为毫秒计数
            if (index(ms_str, "ms") > 0) { # 毫秒
                gsub("ms", "", ms_str); # 去掉ms尾部
                return strtonum(ms_str);
            }
            if (index(ms_str, "µs") > 0) { # 纳秒
                gsub("µs", "", ms_str); # 转换为ms
                return strtonum(ms_str) / 1000.0;
            }
            if (index(ms_str, "d") > 0 || index(ms_str, "h") > 0) { # 超过小时极端情况
                return 999999.9;
            }
            if (match(ms_str, "[0-9.]+m[0-9.]+s") > 0) {  # 几分几秒
                len = split(ms_str, result, "m")
                if (len == 0) {
                    return 0.0;
                }
                minutes = strtonum(result[1]);
                seconds_str = result[2];
                gsub("s", "", seconds_str);
                seconds = strtonum(seconds_str);
                return minutes * 60 * 1000.0 + seconds * 1000.0;
            }
            return 0.0;
        }

        BEGIN {
            request_begin = "";         # 开始
            request_end = "";           # 结束
            request_ms_total = 0.0;     # 总计请求时间
            request_total = 0;          # 总计请求数量
            api_slow = "";              # 最慢接口
            api_slow_request_time = ""; # 最慢请求时间
            api_slow_request_ms = 0.0;  # 最慢请求ms
            api_fast = "";              # 最快接口
            api_fast_request_time = ""; # 最快请求时间
            api_fast_request_msg = 0.0; # 最快请求ms
            #array_request_times = []; # 请求次数
            #array_request_total = []; # 请求时间总计
        }
        {
            if ($12 != "GET" && $12 != "POST") { # 跳过非get/post的行
                next;
            }
            if (match($0, "favicon") > 0) { # 跳过favicon
                next;
            }
            if (match($13, "/ws") > 0) { # 跳过ws接口
                next;
            }
            if (request_begin == "") { # 如果还没有记录开始时间则记录
                request_begin = sprintf("%s %s", $2, $4);
            }
            if (FNR == NR) { # 是否最后一行 - 记录最后时间
                request_end = sprintf("%s %s", $2, $4);
            }

            request_ms = get_ms($8); # 转换为ms
            if (request_ms > api_slow_request_ms) { # 计算最慢请求
                api_slow_request_ms = request_ms;
                api_slow = $13;
                api_slow_request_time = sprintf("%s %s", $2, $4);
            }
            if (api_fast_request_time == "") { # 如果还没有记录最快时间
                api_fast_request_time = sprintf("%s %s", $2, $4);
                api_fast = $13;
                api_fast_request_ms = request_ms;
            } else if (request_ms < api_fast_request_ms) { # 计算最快请求 - 进行对比
                api_fast_request_time = sprintf("%s %s", $2, $4);
                api_fast = $13;
                api_fast_request_ms = request_ms;
            }

            request_total = request_total + 1; # 对于请求总计进行累加
            request_ms_total = request_ms_total + request_ms; # 对于请求计时进行累加

            # 对数组相关进行统计
            split($13, urls, "?") # 只取?前部分
            api_url = urls[1];
            if (api_url in array_request_times) {
                array_request_times[api_url] = array_request_times[api_url] + 1;
            } else {
                array_request_times[api_url] = 1;
            }
            if (api_url in array_request_total) {
                array_request_total[api_url] = array_request_total[api_url] + request_ms;
            } else {
                array_request_total[api_url] = request_ms;
            }

            if ($6 != "200") {  # 如果是错误接口
                if (api_url in array_request_errors) {
                    array_request_errors[api_url] = array_request_errors[api_url] + 1;
                } else {
                    array_request_errors[api_url] = 1;
                }
            }
        }
        END {
            printf("统计区间: %s - %s\n", request_begin, request_end);
            printf("总计时间: %.4fms - Req: %d - Time/Req: %.4fms\n", request_ms_total, request_total, request_ms_total / request_total);
            printf("最快请求: %-20s - %-20s - %s\n", sprintf("%.4fms", api_fast_request_ms), api_fast_request_time, api_fast);
            printf("最慢请求: %-20s - %-20s - %s\n", sprintf("%.4fms", api_slow_request_ms), api_slow_request_time, api_slow);
            printf("\n");
            printf("*** 详细接口访问统计 ***\n");
            printf("----------------------------------------------------------------------------------------------------\n");
            for (k in array_request_times) {
                printf("%-40s - Req: %-4d - Time/Req: %.4fms (Total: %.4fms)\n",
                    k,
                    array_request_times[k],
                    array_request_total[k] / array_request_times[k],
                    array_request_total[k]);
            }
            if (length(array_request_errors) > 0) {
                printf("*** 访问接口错误统计 ***\n");
                printf("----------------------------------------------------------------------------------------------------\n");
                for (k in array_request_errors) {
                    printf("%-40s - Total: %d\n", k, array_request_errors[k]);
                }
            }
        }
    ' $log_file
}

main() {
    log_file=""
    case "$1" in
        "admin")
            log_file="logs/sports-admin-bin.log"
            request_stat $log_file;;
        "member-api")
            log_file="logs/sports-member-api-bin.log"
            request_stat $log_file;;
        "oss")
            log_file="logs/sports-admin-bin.log"
            request_stat $log_file;;
        "game-api")
            log_file="logs/sports-admin-bin.log"
            request_stat $log_file;;
        *)
            echo "参数错误: $1"
            echo "可用参数: admin|member-api|game-api|oss"
            exit;;
    esac
}

main $@