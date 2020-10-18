#!/bin/sh

sigint_handler()
{
  kill $PID
  exit
}

trap sigint_handler SIGINT

go build

while true; do
  go build
  ./log &
  PID=$!
  inotifywait -r -e modify .
  kill $PID
done
