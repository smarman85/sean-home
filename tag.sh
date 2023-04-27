#!/bin/bash

APP=seanhom

docker rm -f $APP

docker rmi -f $APP

docker build -t $APP .

containerCommit=$(docker ps -a | grep $APP | awk '{print $1}')
VERSION=$1

if [[ -z $VERSION  ]];
then
  echo "Missing Version"
  exit 1
fi

docker commit $containerCommit $APP

imageID=$(docker images | grep $APP | head -1 | awk '{print $3}')

docker tag $imageID smarman/${APP}:${VERSION}

docker push smarman/${APP}:${VERSION}
