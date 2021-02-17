#!/bin/bash

docker run -p 8080:8080   \
  -e PPS_DBCONFIG_HOST=$(ifconfig en0  | grep 'inet ' |awk '{print $2}') \
  qiujiahong/page-ss:qiujiahong/page-ss:latest
