# 删除所有名字中带 “provider”
docker rmi $(docker images | grep "provider" | awk '{print $3}')
# 查看容器ip
docker inspect d7f29df68dd4 | grep IPAddress
# 删除所有容器
docker rm -f `docker ps -a -q`