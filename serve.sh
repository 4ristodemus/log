#!/bin/sh

sigint_handler()
{
  kill $PID
  exit
}

trap sigint_handler SIGINT

go build

while true; do
  clear
  ./compile_assets.sh
  go build
  ./log serve &
  PID=$!
  inotifywait -r -e modify .
  kill $PID
done
