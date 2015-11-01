#!/usr/bin/env bash

docker build -t cm-voting .

docker run -i -t -p 9000:9000 --link mongo_cmvoting:mongo cm-voting env