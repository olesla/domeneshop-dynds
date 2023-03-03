# olesla/domeneshop-dynds

Uses [domeneshop.no api](https://api.domeneshop.no/docs/#tag/ddns/paths/~1dyndns~1update/get) to dynamically update your IP-address.  

## DockerHub

[https://hub.docker.com/r/olesla/domeneshop-dynds](https://hub.docker.com/r/olesla/domeneshop-dynds)


## docker-compose

```
version: "2.1"
services:
  domeneshop-dyndns:
    image: olesla/domeneshop-dynds
    container_name: domeneshop-dynds
    restart: unless-stopped
    environment:
      - HOSTNAME=your.hostname.com
      - TOKEN=your-domeneshop-token
      - SECRET=your-domeneshop-secret
```
