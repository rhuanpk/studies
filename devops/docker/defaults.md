## POSTGRES
```
docker volume create postgres
docker run -d [--rm] --name postgres -e POSTGRES_PASSWORD=postgres [-v [/path/]postgres:/var/lib/postgresql/data] [-p 5432:5432] postgres
```

## ADMINER
```
docker run -d [--rm] --name adminer [--link database:db] -p 80:8080 adminer
```

## WINDOWS
```
docker run -it [--rm] --name windows -e "VERSION={11|11l|11e|10|10l|10e|8e|7u|vu|xp|2k|2025|2022|2019|2016|2012|2008|2003}" [-e "RAM_SIZE=8G"] [-e "CPU_CORES=8"] -p 8006:8006 --device=/dev/kvm --device=/dev/net/tun --cap-add NET_ADMIN [-v "./windows:/storage[:rw]"] --stop-timeout 120 dockurr/windows
```
