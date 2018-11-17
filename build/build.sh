#!/bin/bash
pushd ..
export $(egrep -v '^#' .env | xargs)
docker rmi -f $GRAMMAR_BOT_IMAGE
docker build -t $GRAMMAR_BOT_IMAGE -f build/dockerfile .
popd
