#!/usr/bin/env bash

docker build -t bentsolheim/weather-station-data-receiver .
docker run \
 --rm \
 -p 9010:9010 \
 --name weather-station-data-receiver \
 bentsolheim/weather-station-data-receiver:latest
