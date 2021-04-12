# dod
*dod = docker delete*
这是一个用于批量删除本地docker镜像和容器的小程序

# 安装
```go
go install github.com/kainhuck/dod
```

# 使用

如果要查看删除过程加上 `-v`

- 删除所有容器
```shell
$ dod delete container
```
```shell
$ dod delete con
```

- 删除指定容器
```shell
$ dod delete container id1 id2 ...
```

- 删除所有镜像
```shell
$ dod delete image
```
```shell
$ dod delete img
```

- 删除指定镜像
```shell
$ dod delete image id1 id2 ...
```

- 删除所有容器和镜像
```shell
$ dod delete
```