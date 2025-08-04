#!/bin/sh
/usr/local/bin/webserver & exec "$@" & /home/tfc-agent/bin/tfc-agent