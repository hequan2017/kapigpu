# AGENTS.md

## Cursor Cloud specific instructions

### Project Overview

This is **kapigpu**, a Gin-Vue-Admin (GVA) full-stack admin panel with a Go (Gin) backend and Vue 3 (Vite) frontend.

### Prerequisites

- **Go 1.24.0+** — the `go.mod` specifies `go 1.24.0` with `toolchain go1.24.2`. The system Go may be older; install Go 1.24+ to `/usr/local/go` and ensure `PATH=/usr/local/go/bin:$PATH`.
- **Node.js 20+** — available via nvm.
- **npm** — no lockfile is committed; use `npm install` in `web/`.

### Running Services

| Service | Directory | Command | Port |
|---------|-----------|---------|------|
| Backend | `server/` | `go run main.go` | 8888 |
| Frontend | `web/` | `npm run dev` | 8080 |

### Database Setup (SQLite — no external DB needed)

The default `config.yaml` points to an external MySQL server. For local dev without MySQL:

1. Clear the MySQL `db-name` in `server/config.yaml` (set to `""`) so the server starts without attempting a MySQL connection.
2. Start the backend (`cd server && go run main.go`).
3. Initialize the database via the API or web UI. API example:
   ```bash
   mkdir -p server/data
   curl -X POST http://127.0.0.1:8888/init/initdb \
     -H "Content-Type: application/json" \
     -d '{"dbType":"sqlite","adminPassword":"123456","dbName":"gva","dbPath":"./data"}'
   ```
4. After init, the server writes SQLite config back to `config.yaml` and operates with SQLite.

### Default Login

- Username: `admin`
- Password: `123456`
- Captcha is enabled by default (`open-captcha: 0` means always on).

### Lint & Build

- **Frontend lint**: `npx eslint .` in `web/` (32 pre-existing lint errors in the codebase)
- **Frontend build**: `npm run build` in `web/` (slow; for dev use `npm run dev`)
- **Backend build**: `go build -o server .` in `server/` (compiles successfully)
- **Backend `./...` build**: fails due to a pre-existing issue in `plugin/announcement/gen/gen.go` (misplaced compiler directive), but the main binary builds fine.

### Key Gotchas

- The frontend Vite dev server proxies `/api` requests to the backend at `http://127.0.0.1:8888`. Both services must be running for the app to work.
- When initializing SQLite via the API, the `dbPath` directory must exist beforehand or you'll get "unable to open database file" errors.
- The `server/config.yaml` is overwritten by the init process — it will change `system.db-type` to `sqlite` and update the `sqlite` section.
