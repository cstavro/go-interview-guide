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
- **templates/** — 52 problem boilerplate dirs across 6 categories: `fundamentals/`, `concurrency/`, `data-structures/`, `algorithms/`, `systems-design/`, `projects/`.
- **workspaces/** — gitignored; generated at runtime by `/api/generate` copying a template dir.

## Template Conventions

Each template is a **standalone Go module** (`go 1.22` in its own `go.mod`):
- `main.go` with TODO stubs, `main_test.go` with tests, `README.md`
- **systems-design** templates have only `go.mod` + `README.md` (design exercises, no code)
- Module names are the short problem name (e.g. `module slice-append`), not a full path

## Creating New Questions

When adding a new problem, create a directory under `templates/<category>/qNN-short-name/` (for example, `templates/concurrency/q09-new-pattern`). Follow the existing layout:

- `go.mod` — standalone module using the short problem name (e.g., `module new-pattern`)
- `main.go` — starter code with clear TODO stubs and exported function signatures
- `main_test.go` — behavioral tests (see below)
- `README.md` — problem statement, hints, follow-ups, and relevant doc links
- `AGENTS.md` — per-question interviewer-agent instructions

### Per-question `AGENTS.md`

Every question must include its own `AGENTS.md`. It configures the interviewer agent that reviews the candidate's solution. Copy an existing one as a starting point and customize:

- **Problem title and overview** — state the goal in one or two sentences.
- **Hints** — three progressively more specific hints. Keep them conceptual; never give away the implementation.
- **Follow-up questions** — a few depth/extension prompts tailored to the problem.
- **Solution review criteria** — what to evaluate (correctness, efficiency, clarity, testing, completeness).
- **Review tone** — honest, constructive, specific praise and criticism.

The agent instructions must keep the agent in interviewer mode: ask before telling, be vague by default, and do not solve the problem for the candidate.

### Writing the Test Suite

Tests should validate behavior without dictating implementation or giving away the solution.

- **Test the contract, not the algorithm.** Assert inputs, outputs, invariants, and observable side effects. Do not require a specific data structure, helper function name, or step-by-step approach.
- **Avoid implementation hints.** Test names, variable names, and comments should not telegraph the intended solution.
- **Cover meaningful edge cases** such as empty inputs, single elements, duplicates, negative values, nil slices, timeouts, and concurrent callers — but only when they are part of the stated problem.
- **Prefer table-driven tests** for multiple similar cases.
- **Keep tests minimal enough** that a reasonable solution passes, but strict enough that an incomplete or buggy solution fails.
- **Do not import internal test helpers** from outside the template; every template must be self-contained.

If a question is design-only or debugging-focused, it may omit `main.go` and `main_test.go`, matching the systems-design convention.

## Gotchas

- No root `go.mod` — do not run `go build ./...` or `go test ./...` from the repo root.
- `workspaces/` is gitignored and ephemeral; don't put source-of-truth files there.
- Frontend has no bundler — edit HTML/CSS/JS directly in `go-interview-guide/`.
