[简体中文](README.md) | [English](README.en.md)

# 卡皮巴拉GPU 管理平台

轻量的 GPU 算力管理起步版：基于 Gin-Vue-Admin（GVA）全栈框架，先把多个 Docker 集群（TLS 凭证加密存储）统一纳管起来，再逐步长出算力调度能力。

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![Vue](https://img.shields.io/badge/Vue-3.5-4FC08D?logo=vue.js&logoColor=white)
![License](https://img.shields.io/badge/License-Apache--2.0-green)

## 📖 项目介绍

GPU 管理平台往往一上来就很重。本仓库选择从最小可用集出发：以 Gin-Vue-Admin 为底座，保留完整的用户/角色/菜单/RBAC 权限与日志审计能力，核心业务先做 **Docker 集群管理**——登记集群名称并保存 CA 证书、客户端证书与私钥（加密存储，TEXT 字段避免 MySQL 行大小限制），提供凭证查看与集群下拉选择接口，为后续的节点纳管、容器实例调度打好地基。

如果你需要开箱即用的完整 GPU 管理平台（实例调度、显存切分、SSH 跳板机、K8s 管理等），请使用同作者的 [docker-gpu-manage](https://github.com/hequan2017/docker-gpu-manage)；本仓库适合作为二开起点。

## ✨ 功能特性

- 🗂️ Docker 集群管理：新建、编辑、删除、批量删除、分页查询
- 🔐 集群 TLS 凭证：CA 证书 / 客户端证书 / 私钥加密存储（TEXT），支持凭证查看
- 📋 集群下拉接口：`getAllDockerClusters` 供其他模块选择集群
- 👥 GVA 内置能力：用户、角色、菜单、API、Casbin RBAC、JWT 认证、验证码、操作记录
- 📢 公告插件与 📧 邮件插件（GVA 插件机制，可按需扩展）
- 🧩 代码生成器等 GVA 系统工具，方便快速扩展业务模块

## 🛠 技术栈

| 层 | 技术 |
|---|---|
| 后端 | Go 1.24、Gin 1.10、GORM 1.25、Casbin v2、JWT v5、Viper、Zap |
| 前端 | Vue 3.5、Vite 6、Element Plus 2.10、Pinia 2、UnoCSS 66、Axios 1.8 |
| 数据库 | MySQL（默认），支持 SQLite / PostgreSQL 等经 GORM 接入 |

## 🚀 快速开始

### 环境要求

- Go 1.24+、Node.js 20+、npm

### 启动后端（默认 8888）

```bash
cd server
go run main.go
```

### 启动前端（默认 8080，/api 代理到 127.0.0.1:8888）

```bash
cd web
npm install
npm run dev
```

### 数据库初始化（推荐 SQLite，本地无需 MySQL）

默认 `server/config.yaml` 指向外部 MySQL。本地用 SQLite 时：

1. 将其中 MySQL 的 `db-name` 清空（`""`），避免启动时连接 MySQL
2. 启动后端后，执行：

```bash
mkdir -p server/data
curl -X POST http://127.0.0.1:8888/init/initdb \
  -H "Content-Type: application/json" \
  -d '{"dbType":"sqlite","adminPassword":"123456","dbName":"gva","dbPath":"./data"}'
```

> `dbPath` 目录需提前创建，否则会报 `unable to open database file`。初始化成功后配置会写回 `config.yaml` 并切换为 SQLite 运行。

### 默认账号

- 用户名 `admin` / 密码 `123456`；验证码默认开启（`open-captcha: 0`）

## 📁 目录结构

```text
.
├─ server/               # Go 后端（Gin + GORM）
│  └─ model|service|api|router/dockerCluster/   # Docker 集群管理模块
└─ web/                  # Vue 3 前端（Vite）
   └─ src/view/dockerCluster/
```

## 🔗 相关项目

同一 GPU 管理方向的演进版本：

- [docker-gpu-manage](https://github.com/hequan2017/docker-gpu-manage) — 天启算力管理平台（完整版：实例调度、HAMi 显存切分、SSH 跳板机、K8s 管理）
- [pcfarm-admin](https://github.com/hequan2017/pcfarm-admin) — 服务器资产 / IP / PXE / 远控管理系统
- [tianqi](https://github.com/hequan2017/tianqi) — TianQi GPU Manager
- [DockerGPU](https://github.com/hequan2017/DockerGPU) — 前代 GPU 算力租用管理系统

## 📄 License

本项目基于 [Apache-2.0](./LICENSE) 协议开源。
