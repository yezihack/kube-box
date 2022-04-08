[![Docker Image CI](https://github.com/yezihack/kube-box/actions/workflows/docker-image.yml/badge.svg)](https://github.com/yezihack/kube-box/actions/workflows/docker-image.yml)

# checkbox

> 专于诊断网络使用的

## 接口

> 默认80端口

| 序列 | 接口地址 | 说明  |
| ---- | ----- | ----- |
| 1    | /   | 主页 |
| 2    | /ping    | 存活检测接口  |
| 3    | /healthz   |  健康接口|
|4  |  /check-ip | 检查IP是否通达
|4  |  /dry-check-ip | 检查IP是否通达，只返回失败的
|5   |/check-healthz| 检查健康接口是否通达
|5   |/dry-check-healthz| 检查健康接口是否通达，只返回失败的,

## 环境变量

| 名称 | 默认值 | 说明  |
| ---- | ----- | ----- |
|PORT | 80| 端口
|TARGET_PORT | 80 | 目标端口，即 healthz请求时使用的端口
|VERSION| v0.0.1 | 版本号
|DATA_PATH| ./data/ | 数据存储目录路径
|IP_DATA_FILENAME| ip.data| IP数据文件名|
|NETWORK_NAME|eth0 |网卡名称，用于获取本机IP地址
|GO_NUMBER | 10 | 并发数
|TIMEOUT | 5 | 请求超时，单位秒（s）
|HEALTHZ_PATH_NAME| healthz| 健康接口地址

## 运行测试

- windown

```bat
$env:DATA_PATH="./data/"
$env:NETWORK_NAME="WLAN"

go run .
```

- linux

```sh
export DATA_PATH="./data/"
export NETWORK_NAME="WLAN"
go run .
```
