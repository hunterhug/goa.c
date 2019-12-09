# 文档部署

参考以下步骤可以部署该文档。

`Linux/Mac` 安装 `docker` 后，请按以下部署文档：

```
./docker_build.sh
./docker_run.sh
```

提示没有权限，请加上 `sudo`。

打开 [http://127.0.0.1:9999](http://127.0.0.1:9999)

# 其他部署方法（建议不用）

如果不想使用 `docker`，可以：

```
# npm install -g  n
# n stable
# n
npm install -g gitbook-cli --registry=https://registry.npm.taobao.org
gitbook install
gitbook serve
# gitbook build . --log=debug --debug
```

提示没有权限，请加上 `sudo`。