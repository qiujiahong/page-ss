#!/bin/bash


#sh build.sh
#rm -rf docker/apps
#cp -rf ./dist/latest/page-ss_linux docker/apps
cd docker

docker build -t qiujiahong/page-ss:v1.0 .

