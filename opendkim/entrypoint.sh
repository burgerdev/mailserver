#!/bin/bash

rsyslog -n &
/usr/sbin/opendkim -x /etc/opendkim.conf -f &

wait $!
