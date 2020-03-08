#!/bin/bash

/usr/sbin/opendkim -x /etc/opendkim.conf -f &

wait
