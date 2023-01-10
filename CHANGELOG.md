# CHANGELOG

## v0.7.0

- Dockerfile 基本 busybox 包版本改为: latest
- 新增 tcpdump 抓包工具

## v0.6.0

- 修改 busybox 版本 v0.1.2(新增 AB 压测工具)

## v0.5.0

- 添加 [nali](https://github.com/zu1k/nali) 一个查询IP地理信息和CDN提供商的离线终端工具.

## v0.4.0

- 添加 [httpstat](https://github.com/davecheney/httpstat) 分析URL请求不同阶段的耗时

## v0.3.0

- 添加 /check-mysql 检查 MySQL 连接状态
- 添加 /create-mysql-db 创建 MySQL 新的数据库

## v0.2.0

- 添加 metrics 接口，用于接入 prometheus 监控使用。

## v0.1.0

- 基础镜像替换为 sgfoot/busybox:v0.1.1

## v0.0.2

- 修复 check-healthz 无法解析 heathz 返回的结果

## v0.0.1

- 新版本发布
