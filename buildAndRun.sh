#!/usr/bin/env bash

docker pull mongo
docker run -i -t --name mongo_cmvoting -p 27017:27017 mongo
docker build -t karesti/cm_voting .
docker run -i -t -p 9000:9000 --link mongo_cmvoting:mongo karesti/cm_voting