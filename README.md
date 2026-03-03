# kapigpu

基于 Gin-Vue-Admin（GVA）的全栈后台管理系统，后端使用 Go + Gin，前端使用 Vue 3 + Vite。

## 技术栈

### 后端（server/）

- Go 1.24+（仓库 go.mod：go 1.24.0，toolchain go1.24.2）
- Gin 1.10.0
- GORM 1.25.x
- Casbin 2.x（RBAC 权限）
- Viper 配置、Zap 日志

### 前端（web/）

- Vue 3.5.x + Vite 6.x
- Pinia 2.x
- Element Plus 2.x
- UnoCSS 66.x
- Axios 1.8.x

## 快速开始（本地开发）

### 环境要求

- Go 1.24+
- Node.js 20+
- npm

### 启动后端

```bash
cd server
go run main.go
```

默认端口：8888

### 启动前端

```bash
cd web
npm install
npm run dev
```

默认端口：8080  
前端开发服务器会将 `/api` 代理到 `http://127.0.0.1:8888`，因此需要同时启动前后端。

## 数据库初始化（推荐 SQLite，本地无需 MySQL）

项目默认的 `server/config.yaml` 可能指向外部 MySQL。若本地不使用 MySQL，可按以下方式用 SQLite 初始化：

1. 将 `server/config.yaml` 中 MySQL 的 `db-name` 清空（设置为 `""`），避免启动时尝试连接 MySQL。
2. 启动后端（见上文）。
3. 通过接口初始化数据库：

```bash
mkdir -p server/data
curl -X POST http://127.0.0.1:8888/init/initdb \
  -H "Content-Type: application/json" \
  -d '{"dbType":"sqlite","adminPassword":"123456","dbName":"gva","dbPath":"./data"}'
```

注意：`dbPath` 目录需要提前存在，否则可能出现 “unable to open database file”。

初始化成功后，服务会把 SQLite 配置写回 `server/config.yaml`，并切换为 SQLite 运行。

## 默认账号

- 用户名：`admin`
- 密码：`123456`
- 验证码：默认开启（`open-captcha: 0` 表示始终开启）

## 常用命令

### 前端

```bash
cd web
npm run dev
npm run build
npx eslint .
```

### 后端

```bash
cd server
go build -o server .
```

## 目录结构

```text
.
├─ server/   # Go 后端（Gin + GORM）
└─ web/      # Vue 3 前端（Vite）
```
