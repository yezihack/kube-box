[![Docker Image CI](https://github.com/yezihack/kube-box/actions/workflows/docker-image.yml/badge.svg)](https://github.com/yezihack/kube-box/actions/workflows/docker-image.yml)
<!-- TOC -->

- [1. kube-box](#1-kube-box)
  - [1.1. 部署使用](#11-部署使用)
  - [1.2. 镜像选择](#12-镜像选择)
  - [1.3. 接口介绍](#13-接口介绍)
  - [1.4. 环境变量](#14-环境变量)
  - [1.5. 运行测试](#15-运行测试)
  - [1.6. 添加 metrics 接口](#16-添加-metrics-接口)

<!-- /TOC -->

# 1. kube-box

> 专于诊断网络使用的

## 1.1. 部署使用

两种使用方法，一种是使用 DeamonSet, 一种是 Deployment

- [DeamonSet 清单](mainfest/kube-box-ds.yaml)
  - 用于诊断整个集群的网络通达
- [Deployment 清单](mainfest/kube-box.yaml)
  - 用于单个应用的测试与调试使用

## 1.2. 镜像选择

<https://hub.docker.com/r/sgfoot/kube-box/tags>

```sh
# 最新版本
docker pull sgfoot/kube-box:latest
```

## 1.3. 接口介绍

> 默认80端口

| 序列 | 接口地址 | 说明  |
| ---- | ----- | ----- |
| 1    | /   | 主页 |
| 2    | /ping    | 存活检测接口  |
| 3    | /healthz   |  健康接口|
|4  |  /check-ip | 检查IP是否通达
|5  |  /dry-check-ip | 检查IP是否通达，只返回失败的
|6   |/check-healthz| 检查健康接口是否通达
|7   |/dry-check-healthz| 检查健康接口是否通达，只返回失败的
|8   |/metrics| prometheus metrics infomation

## 1.4. 环境变量

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

## 1.5. 运行测试

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

## 1.6. 添加 metrics 接口

> 用于接入 prometheus 监控使用。

访问方法：<http://localhost:80/metrics>

```sh
# 统计接口访问状态次数。如 200 表示成功次数为 10 次，404表示未找到页面 7 次。
# HELP kubebox_requests_total Number of the http requests received since the server started
# TYPE kubebox_requests_total counter
kubebox_requests_total{status="200"} 10
kubebox_requests_total{status="404"} 7
```
