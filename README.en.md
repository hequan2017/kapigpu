[简体中文](README.md) | [English](README.en.md)

# KaPi GPU Management Platform

A lightweight starting point for GPU compute management: built on the Gin-Vue-Admin (GVA) full-stack framework, it first brings multiple Docker clusters (with encrypted TLS credentials) under unified management, and grows scheduling capabilities from there.

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![Vue](https://img.shields.io/badge/Vue-3.5-4FC08D?logo=vue.js&logoColor=white)
![License](https://img.shields.io/badge/License-Apache--2.0-green)

## 📖 Introduction

GPU management platforms tend to be heavy out of the gate. This repo starts from a minimal viable set: it uses Gin-Vue-Admin as the foundation — keeping the complete user/role/menu/RBAC permission system with audit logging — and focuses its core business on **Docker cluster management**: register clusters by name and store their CA certificate, client certificate, and private key (encrypted, in TEXT columns to avoid MySQL row size limits), with credential viewing and a cluster dropdown API that lay the groundwork for future node onboarding and container scheduling.

If you want a full-featured, ready-to-use GPU management platform (instance scheduling, VRAM splitting, SSH jump box, Kubernetes management), use the same author's [docker-gpu-manage](https://github.com/hequan2017/docker-gpu-manage). This repo works best as a fork-friendly starting point.

## ✨ Features

- 🗂️ Docker cluster management: create, edit, delete, batch delete, and paginated queries
- 🔐 Cluster TLS credentials: CA cert / client cert / private key stored encrypted (TEXT), with credential viewing
- 📋 Cluster dropdown API: `getAllDockerClusters` for other modules to pick a cluster
- 👥 Built-in GVA capabilities: users, roles, menus, APIs, Casbin RBAC, JWT auth, captcha, operation records
- 📢 Announcement plugin and 📧 email plugin (GVA plugin mechanism, extensible as needed)
- 🧩 GVA system tools such as the code generator for fast module development

## 🛠 Tech Stack

| Layer | Technologies |
|---|---|
| Backend | Go 1.24, Gin 1.10, GORM 1.25, Casbin v2, JWT v5, Viper, Zap |
| Frontend | Vue 3.5, Vite 6, Element Plus 2.10, Pinia 2, UnoCSS 66, Axios 1.8 |
| Database | MySQL (default); SQLite / PostgreSQL and others available via GORM |

## 🚀 Quick Start

### Requirements

- Go 1.24+, Node.js 20+, npm

### Start the backend (port 8888 by default)

```bash
cd server
go run main.go
```

### Start the frontend (port 8080 by default, /api proxied to 127.0.0.1:8888)

```bash
cd web
npm install
npm run dev
```

### Database initialization (SQLite recommended — no local MySQL needed)

The default `server/config.yaml` points to an external MySQL server. To use SQLite locally:

1. Clear the MySQL `db-name` in it (set to `""`) so the server does not try to connect to MySQL on startup
2. Start the backend, then run:

```bash
mkdir -p server/data
curl -X POST http://127.0.0.1:8888/init/initdb \
  -H "Content-Type: application/json" \
  -d '{"dbType":"sqlite","adminPassword":"123456","dbName":"gva","dbPath":"./data"}'
```

> Create the `dbPath` directory first, otherwise you may get `unable to open database file`. After initialization, the config is written back to `config.yaml` and the server switches to SQLite.

### Default Account

- Username `admin` / password `123456`; captcha is always on by default (`open-captcha: 0`)

## 📁 Directory Structure

```text
.
├─ server/               # Go backend (Gin + GORM)
│  └─ model|service|api|router/dockerCluster/   # Docker cluster module
└─ web/                  # Vue 3 frontend (Vite)
   └─ src/view/dockerCluster/
```

## 🔗 Related Projects

Evolutions in the same GPU management direction:

- [docker-gpu-manage](https://github.com/hequan2017/docker-gpu-manage) — TianQi compute management platform (full edition: instance scheduling, HAMi VRAM splitting, SSH jump box, K8s management)
- [pcfarm-admin](https://github.com/hequan2017/pcfarm-admin) — server assets / IP pools / PXE / remote power control
- [tianqi](https://github.com/hequan2017/tianqi) — TianQi GPU Manager
- [DockerGPU](https://github.com/hequan2017/DockerGPU) — the previous-generation GPU rental management system

## 📄 License

Released under the [Apache-2.0](./LICENSE) license.
