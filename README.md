[![Docker Image CI](https://github.com/yezihack/kube-box/actions/workflows/docker-image.yml/badge.svg)](https://github.com/yezihack/kube-box/actions/workflows/docker-image.yml)
<!-- TOC -->

- [1. kube-box](#1-kube-box)
  - [1.1. 特色](#11-特色)
  - [1.2. 部署使用](#12-部署使用)
    - [1.2.1. Kubenetes 部署](#121-kubenetes-部署)
    - [1.2.2. Docker 部署](#122-docker-部署)
  - [1.3. 镜像选择](#13-镜像选择)
  - [1.4. 接口介绍](#14-接口介绍)
  - [1.5. 环境变量](#15-环境变量)
  - [1.6. 运行测试](#16-运行测试)
  - [1.7. 添加 metrics 接口](#17-添加-metrics-接口)
  - [1.8. 检查数据库](#18-检查数据库)
    - [1.8.1. 参数说明](#181-参数说明)
  - [1.9. 创建数据库](#19-创建数据库)
    - [1.9.1. 参数说明](#191-参数说明)

<!-- /TOC -->
# 1. kube-box

> 专于诊断网络使用的

## 1.1. 特色

- ping 存活检查
- healthz 健康检查
- 检查IP是否通达
- 检查健康接口是否通达
- 检查 MySQL 连接状态
- 创建 MySQL 新的数据库
- 提供 Prometheus metrics 数据
- 添加 [httpstat](https://github.com/davecheney/httpstat) 分析URL请求不同阶段的耗时

## 1.2. 部署使用

### 1.2.1. Kubenetes 部署

两种使用方法，一种是使用 DeamonSet, 一种是 Deployment

- [DeamonSet 清单](mainfest/kube-box-ds.yaml)
  - 用于诊断整个集群的网络通达
- [Deployment 清单](mainfest/kube-box.yaml)
  - 用于单个应用的测试与调试使用

### 1.2.2. Docker 部署

```sh
# 使用默认端口 80
docker run -itd --name kube-box -p 9110:80 sgfoot/kube-box:latest
```

## 1.3. 镜像选择

<https://hub.docker.com/r/sgfoot/kube-box/tags>

```sh
# 最新版本
docker pull sgfoot/kube-box:latest
```

## 1.4. 接口介绍

> 默认80端口

| 序列 | 接口地址           | 说明                               |
| ---- | ------------------ | ---------------------------------- |
| 1    | /                  | 主页                               |
| 2    | /ping              | 存活检测接口                       |
| 3    | /healthz           | 健康接口                           |
| 4    | /check-ip          | 检查IP是否通达                     |
| 5    | /dry-check-ip      | 检查IP是否通达，只返回失败的       |
| 6    | /check-healthz     | 检查健康接口是否通达               |
| 7    | /dry-check-healthz | 检查健康接口是否通达，只返回失败的 |
| 8    | /metrics           | prometheus metrics infomation      |
| 9    | /check-mysql       | 检查 MySQL 连接状态                |
| 10   | /create-mysql-db   | 创建 MySQL 新的数据库              |

## 1.5. 环境变量

| 名称              | 默认值  | 说明                                 |
| ----------------- | ------- | ------------------------------------ |
| PORT              | 80      | 端口                                 |
| TARGET_PORT       | 80      | 目标端口，即 healthz请求时使用的端口 |
| VERSION           | v0.0.1  | 版本号                               |
| DATA_PATH         | ./data/ | 数据存储目录路径                     |
| IP_DATA_FILENAME  | ip.data | IP数据文件名                         |
| NETWORK_NAME      | eth0    | 网卡名称，用于获取本机IP地址         |
| GO_NUMBER         | 10      | 并发数                               |
| TIMEOUT           | 5       | 请求超时，单位秒（s）                |
| HEALTHZ_PATH_NAME | healthz | 健康接口地址                         |

## 1.6. 运行测试

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

## 1.7. 添加 metrics 接口

> 用于接入 prometheus 监控使用。

访问方法：<http://localhost:80/metrics>

```sh
# 统计接口访问状态次数。如 200 表示成功次数为 10 次，404表示未找到页面 7 次。
# HELP kubebox_requests_total Number of the http requests received since the server started
# TYPE kubebox_requests_total counter
kubebox_requests_total{status="200"} 10
kubebox_requests_total{status="404"} 7
```

## 1.8. 检查数据库

### 1.8.1. 参数说明

- host 连接地址
- port 端口号，必须整型
- user 用户名称
- pass 用户密码

例：

```sh
curl localhost/check-mysql?host=127.0.0.1&port=3306&user=root&pass=123456
```

## 1.9. 创建数据库

- 创建数据库，默认采用 utf8mb4

### 1.9.1. 参数说明

- host 连接地址
- port 端口号，必须整型
- user 用户名称
- pass 用户密码
- dbname 需要创建的数据库名称

例：

```sh
curl localhost/create-mysql-db?host=127.0.0.1&port=3306&user=root&pass=123456&dbname=test
```

