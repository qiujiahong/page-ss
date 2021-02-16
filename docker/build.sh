#!/bin/bash

version=$(cat conf/conf.yaml| grep version: | awk '{print $2}')
echo $version
sh build.sh
rm -rf docker/apps
cp -rf ./dist/latest/page-ss_linux docker/apps
cd docker

docker build -t qiujiahong/page-ss:$version .
docker push qiujiahong/page-ss:$version


# docker run -p 8080:8080  qiujiahong/page-ss:v1.0