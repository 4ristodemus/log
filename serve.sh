#!/bin/sh

sigint_handler()
{
  kill $PID
  exit
}

trap sigint_handler SIGINT

while true; do
  go build log.go
  ./log &
  PID=$!
  inotifywait -r -e modify .
  kill $PID
done
