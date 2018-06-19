#!/bin/sh

DESCRIBE=$(git describe --tags 2>/dev/null)
if [ -z "$(git status --porcelain)" ]; then
    echo $DESCRIBE
else
    echo $DESCRIBE-dirty
fi
