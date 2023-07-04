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

## Contributors ✨

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://github.com/stianhb"><img src="https://avatars.githubusercontent.com/u/2493466?v=4" width="100px;" alt=""/><br /><sub><b>Stian Hoel</b></sub></a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
