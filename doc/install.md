# 文档部署

在本地阅读速度会比较快，请安装 `Docker` 后，参考以下步骤可以部署该文档。

## Gitbook 容器部署

可以使用这种方式部署，但我觉得样式有点丑，可以选择 `Docsify` 部署，会漂亮很多。

```
git clone https://github.com/hunterhug/goa.c.git
cd goa.c

./docker_build.sh
./docker_run.sh
```

如果提示没有权限，请加上 `sudo`。

打开 [http://127.0.0.1:12346](http://127.0.0.1:12346)

## Docsify 容器部署（强烈推荐）

支持使用 [Docsify](https://docsify.js.org/#/zh-cn/quickstart) 部署：

```
git clone https://github.com/hunterhug/goa.c.git
cd goa.c

./docker_build_docsify.sh
./docker_run_docsify.sh
```

如果你想在本地调试，可以执行：

```
./docker_debug_docsify.sh
```

打开 [http://127.0.0.1:12346](http://127.0.0.1:12346) 。

## <del>Gitbook裸机部署（建议不用）</del>

如果不想使用 `Docker`，可以：

```
git clone https://github.com/hunterhug/goa.c.git
cd goa.c

# npm install -g  n
# n stable
# n

npm install -g gitbook-cli --registry=https://registry.npm.taobao.org
gitbook install
gitbook serve
# gitbook build . --log=debug --debug
```

提示没有权限，请加上 `sudo`。