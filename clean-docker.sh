#!/bin/sh
echo "remove containers"
docker rm -vf $(docker ps -a -q)
wait

echo "remove images"
docker rmi -f $(docker images -a -q)
wait

echo "remove volumes"
docker volume rm $(docker volume ls -q)
wait

echo "done"

