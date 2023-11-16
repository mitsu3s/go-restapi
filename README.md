## go-restapi

```shell
### Docker を起動
$ docker compose up -d

### 確認
$ docker ps
CONTAINER ID   IMAGE               COMMAND   CREATED          STATUS         PORTS                     NAMES
21d3e9f13fb7   go-restapi-client   "bash"    18 seconds ago   Up 7 seconds   0.0.0.0:58414->8000/tcp   client
6e4d0b1f0698   go-restapi-server   "bash"    18 seconds ago   Up 7 seconds                             server

### Server コンテナに入る
$ docker compose exec server bash

### Client コンテナに入る
$ docker compose exec client bash
```
