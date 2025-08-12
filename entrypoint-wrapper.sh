#!/bin/sh
/usr/local/bin/agentfunctionserver & exec "$@" & /home/tfc-agent/bin/tfc-agent