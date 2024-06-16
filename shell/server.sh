#!/bin/bash

function display_usage {
    echo "Usage: $0 {status|start|stop}"
    exit 1
}

if [ $# -ne 1 ]; then
    display_usage
fi

case "$1" in
    "status")
        echo "Server status: Running"
        ;;
    "start")
        echo "Starting server..."
        ;;
    "stop")
        echo "Stopping server..."
        ;;
    *)
        display_usage
        ;;
esac

exit 0

