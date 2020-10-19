#!/bin/sh

if [[ $# -eq 0 ]]; then
    echo 'Please provide a title for the entry.'
    exit 0
fi


title=$@
filename="$(date +%s)_${title// /-}.md"

cp data/default.md entries/$filename
echo "Successfully created an entry for $title"
