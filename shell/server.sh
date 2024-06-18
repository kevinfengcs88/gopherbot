#!/bin/bash

usage() {
    echo "Usage: $0 {status|start|stop}"
    exit 1
}

status() {
    stdout=$(tasklist.exe | grep "SonsOfTheForestDS.exe")
    if [ -n "$stdout" ]; then
        echo "Server status: UP"
    else
        echo "Server status: DOWN"
    fi
}

start() {
    # implement a check to see if the server is already running
    # cd /mnt/c/Users/Kevin/Desktop/sotf/server && cmd.exe /c StartSOTFDedicated.bat
    # also check after this command that the server truly started

    # new stuff

    # make sure that log file and named pipe exist
    mkdir -p "$(pwd)/logs"
    log_file="$(pwd)/logs/sotf.log"
    fifo_file="/tmp/server_output_fifo"

    # create named pipe
    if [[ ! -p "$fifo_file" ]]; then
        mkfifo "$fifo_file"
    fi

    # run the command in Windows mnt and redirect output to the named pipe
    (cd /mnt/c/Users/Kevin/Desktop/sotf/server && cmd.exe /c StartSOTFDedicated.bat > "$fifo_file" 2>&1) &

    while IFS= read -r line; do
        echo "$line" >> "$log_file"
    done < "$fifo_file"

    wait
    # get rid of carriage returns (^M)
    dos2unix "$log_file" > /dev/null 2>&1
    # get rid of zero width characters (<feff>)
    sed -i 's/\xEF\xBB\xBF//g' "$log_file"
}

stop() {
    # implement a check to see if the server is already down
    stdout=$(taskkill.exe /IM SonsOfTheForestDS.exe /F)
    echo "$stdout"
    # also check after this command that the server truly terminated
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

