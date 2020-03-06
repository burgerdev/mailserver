#!/bin/bash

function _shutdown() {
    if 
        /usr/sbin/postfix status
    then
        /usr/sbin/postfix stop
    fi
}

trap _shutdown SIGTERM SIGINT SIGQUIT EXIT

/usr/sbin/postfix start-fg &

wait
