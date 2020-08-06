#!/usr/bin/env bash

set -e

MAX_NUMBER=4294967296
MAX_COUNT=100000
FORMAT="int32_le"

# make fifo
rm -f datapipe
mkfifo datapipe

# start reader
dieharder -a -g 200 <datapipe &
# xxd -b <datapipe &

# open fifo for writing
exec 3>datapipe

# keep writing binary data into fifo -- will end if reader ends
while true; do
  curl -s "http://rng:8080?min=0&maxExclusive=$MAX_NUMBER&count=$MAX_COUNT&format=$FORMAT" >&3
done
# close fifo
exec 3>&-
