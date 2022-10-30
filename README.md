## TODO

- [ ] buildkitd 如何查看拉下来的镜像

## 说明

builder，基于 k8s & buildkitd 的部署工具，基于 Github Repo build 景象，并推送到 docker.io 上。

## 部署

### 配置代理

如果你本地有安装 http 代理，可以取消注释 `k8s-config/buildkitd_deploy.yaml` 中的 `env` 字段，打开 http 代理

### 配置认证信息

`REGISTRY_USER` 是登陆 docker 所用的用户
`REGISTRY_PASSWORD` 是登陆 docker 的密码

```shell
echo -n "REGISTRY_USER:REGISTRY_PASSWORD" | base64
```

修改 k8s-config/docker_config.yaml 文件，将上述命令生成的字符串，替换掉 `BASE64_AUTH`

### 部署 buildkitd

```shell
k apply -f k8s-config/docker_config.yaml
k apply -f k8s-config/buildkitd_deploy.yaml
```

## 笔记

### buildkitd 设置代理

设置了 `HTTP_PROXY` 和 `HTTPS_PROXY` 环境变量后，拉取镜像就会自动走代理

### buildkitd build 镜像并自动推送

```shell
buildctl -addr tcp://192.168.56.13:1234 build --frontend dockerfile.v0 --local context=`pwd` --local dockerfile=`pwd` --output type=image,name=docker.io/bwangel/kubia:ee9f3c5,push=true
```

### 设置 registry 的认证

```shell
BASE64_AUTH=`echo -n "$CI_REGISTRY_USER:$CI_REGISTRY_PASSWORD" | base64`
mkdir -p /.docker
echo "{\"auths\": {\"https://index.docker.io/v1/\": {\"auth\": \"$BASE64_AUTH\"}}}" > ~/.docker/config.json
```

使用 `docker login` 登陆到 docker registry 之后，也能够在 `~/.docker/config.json` 中看到认证信息