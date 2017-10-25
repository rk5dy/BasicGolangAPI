docker container stop $(docker container ls -q)
docker container prune -f
docker volume rm $(docker volume ls -qf dangling=true)
docker rmi $(docker images | select-string -pattern \<none\> | FOREACH-OBJECT {($_ -split '\s+')[2]})
