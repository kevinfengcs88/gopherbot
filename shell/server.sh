#!/bin/bash

log_file="$(pwd)/logs/sotf.log"

usage() {
    echo "Usage: $0 {status|start|stop}"
    exit 1
}

log_check() {
    mkdir -p "$(pwd)/logs"
    touch "$log_file"
}

clean_log() {
    dos2unix "$log_file" > /dev/null 2>&1
    sed -i 's/\xEF\xBB\xBF//g' "$log_file"
}

status() {
    log_check
    status_check=$(tasklist.exe | grep "SonsOfTheForestDS.exe")
    if [ -n "$status_check" ]; then
        echo "Server status: UP"
        echo "$(date '+%Y-%m-%d %H:%M:%S') gopherbot Server status: UP" >> "$log_file"
    else
        echo "Server status: DOWN"
        echo "$(date '+%Y-%m-%d %H:%M:%S') gopherbot Server status: DOWN" >> "$log_file"
    fi
    clean_log
}

start() {
    # cd /mnt/c/Users/Kevin/Desktop/sotf/server && cmd.exe /c StartSOTFDedicated.bat

    status_check=$(tasklist.exe | grep "SonsOfTheForestDS.exe")
    if [ -n "$status_check" ]; then
        echo "The server is already up"
    else
        log_check
        fifo_file="/tmp/server_output_fifo"

        if [[ ! -p "$fifo_file" ]]; then
            mkfifo "$fifo_file"
        fi

        (cd /mnt/c/Users/Kevin/Desktop/sotf/server && cmd.exe /c StartSOTFDedicated.bat > "$fifo_file" 2>&1) &

        while IFS= read -r line; do
            echo "$(date '+%Y-%m-%d %H:%M:%S') $line" >> "$log_file"
        done < "$fifo_file"

        wait
        clean_log
    fi
}

stop() {
    status_check=$(tasklist.exe | grep "SonsOfTheForestDS.exe")
    if [ -n "$status_check" ]; then
        log_check
        stdout=$(taskkill.exe /IM SonsOfTheForestDS.exe /F)
        echo "$stdout"
        echo "$(date '+%Y-%m-%d %H:%M:%S') gopherbot (LOG) $stdout" >> "$log_file"
        echo "================================================================================================================================================================" >> "$log_file"
        clean_log
    else
        echo "The server is already down"
    fi
}

if [ $# -ne 1 ]; then
    usage
fi

case "$1" in
    "status")
        status
        ;;
    "start")
        start
        ;;
    "stop")
        stop
        ;;
    *)
        usage
        ;;
esac

exit 0

