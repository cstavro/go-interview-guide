# AGENTS.md

## Run & Verify

- **Server**: `go run server.go` (port 8080)
- **Docker**: `docker compose up --build`
- **Test a template**: `cd templates/<category>/<qNN-name> && go test ./...`
- No root-level `go.mod`, test suite, linter, or CI. `server.go` compiles standalone.

## Architecture

- **server.go** — single-file Go HTTP server (stdlib only). Routes: `/` (static files), `/api/template`, `/api/generate`, `/api/workspaces`.
- Existing workspace directories are surfaced in the UI via `/api/workspaces` as a green "Workspace" badge on each question, a per-section "Collapse all with workspaces" button, and started/total counts on the home page category cards.
- **go-interview-guide/** — vanilla HTML/CSS/JS frontend, no build step, no dependencies.
- **templates/** — 32 problem boilerplate dirs across 4 categories: `fundamentals/`, `concurrency/`, `algorithms/`, `systems-design/`.
- **workspaces/** — gitignored; generated at runtime by `/api/generate` copying a template dir.

## Template Conventions

Each template is a **standalone Go module** (`go 1.22` in its own `go.mod`):
- `main.go` with TODO stubs, `main_test.go` with tests, `README.md`
- **systems-design** templates have only `go.mod` + `README.md` (design exercises, no code)
- Module names are the short problem name (e.g. `module slice-append`), not a full path

## Gotchas

- No root `go.mod` — do not run `go build ./...` or `go test ./...` from the repo root.
- `workspaces/` is gitignored and ephemeral; don't put source-of-truth files there.
- Frontend has no bundler — edit HTML/CSS/JS directly in `go-interview-guide/`.
