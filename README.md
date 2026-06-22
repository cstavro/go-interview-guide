# Go Interview Guide

A hands-on interview preparation guide for senior and staff-level Go engineering positions. Features 52 practical questions with hints, follow-ups, and workspace generation.

## Quick Start

### Option 1: Run with Docker Compose (Recommended)

```bash
docker compose up --build
```

Then open http://localhost:8080 in your browser.

### Option 2: Run with Docker

```bash
docker build -t go-interview-guide .
docker run -p 8080:8080 -v $(pwd)/workspaces:/app/workspaces go-interview-guide
```

### Option 3: Run locally

```bash
go run server.go
```

Then open http://localhost:8080 in your browser.

## Project Structure

```
.
├── server.go                        # Go HTTP server
├── Dockerfile                       # Docker image definition
├── docker-compose.yml               # Docker Compose configuration
├── go-interview-guide/              # Static HTML/CSS/JS files
│   ├── index.html
│   ├── go-fundamentals.html
│   ├── concurrency.html
│   ├── data-structures.html
│   ├── algorithms.html
│   ├── systems-design.html
│   ├── projects.html
│   ├── style.css
│   └── script.js
├── templates/                       # 52 problem boilerplate directories
│   ├── fundamentals/
│   ├── concurrency/
│   ├── data-structures/
│   ├── algorithms/
│   ├── systems-design/
│   └── projects/
└── workspaces/                      # Generated on demand (see below)
```

## Volume Mount

The `docker-compose.yml` mounts the `workspaces/` directory as a **bind mount**:

```yaml
volumes:
  - ./workspaces:/app/workspaces
```

### Why this matters

When you click **Generate** on any question, the server creates a new workspace directory under `workspaces/` containing the problem's boilerplate code, tests, and README. By mounting this directory as a volume:

- **Generated workspaces persist** between container restarts
- **You can access them from the host** via the `workspaces/` directory in the project root
- **You can open them directly in your IDE** — each workspace is a standalone Go module
- **No data is lost** if you rebuild or remove the container

### Without the volume mount

If the volume were not mounted, every time you stopped the container, all generated workspaces would be lost. You'd have to regenerate them every time you restart.

### With the volume mount

1. Start the container: `docker-compose up`
2. Open http://localhost:8080 and click **Generate** on a problem
3. The workspace appears in `workspaces/` on your host machine
4. Open `workspaces/fundamentals-q01/` in your IDE
5. Stop the container with `Ctrl+C`
6. Restart with `docker-compose up` — your workspaces are still there

## Features

- **52 questions** across 6 categories: Go Fundamentals, Concurrency, Data Structures, Algorithms, Systems Design, Projects
- **Interactive hints** that reveal progressively
- **Follow-up questions** that simulate interviewer depth
- **Code boilerplate** loaded dynamically from templates
- **One-click workspace generation** — creates a ready-to-go directory with `go.mod`, code, tests, and README
- **Official documentation links** on every question
- **Copy buttons** for all code blocks
- **Overwrite protection** — asks before replacing an existing workspace

## Technology Stack

- **Backend**: Go 1.23+ standard library (net/http)
- **Frontend**: Vanilla HTML/CSS/JS (no build step, no external dependencies)
- **Deployment**: Docker / Docker Compose
- **Templates**: Pure Go files served as static text

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Static file server (HTML, CSS, JS) |
| `/api/template`    | GET  | Load a template file by path            |
| `/api/generate`    | POST | Generate a workspace from a template    |
| `/api/workspaces`  | GET  | List existing generated workspace dirs  |

## AI-Powered Interview Practice (`go-interviewer` skill)

The repository ships with an OpenCode skill (`.opencode/skills/go-interviewer/`) that turns the AI assistant into a technical interviewer. When the skill is active it:

- **Acts as an interviewer** — asks clarifying questions, reveals hints only on request, and reviews your solution without giving away the answer.
- **Loads per-question instructions** automatically from each template's `AGENTS.md` (hints, follow-ups, review criteria).
- **Generates missing scaffolds** on demand — if you start working in a new problem directory, it creates `go.mod`, `main.go`, `main_test.go`, and `AGENTS.md` automatically.
- **Reviews for correctness, efficiency, clarity, and testing** when you ask for feedback.

### How to use it

1. **Generate a workspace** from the web UI or by copying a template manually.
2. **Open that workspace directory** in your terminal / IDE.
3. **Start the conversation** with phrases like:
   - `"Interview me on this problem"`
   - `"Quiz me"`
   - `"Give me a hint"`
   - `"Review my solution"`
4. The skill detects the problem context, loads `README.md` and `AGENTS.md`, and adopts the interviewer persona.

### LeetCode Support

The skill also supports **any LeetCode problem URL**. To practice a LeetCode question:

1. **Provide the URL** to the assistant, e.g.:
   ```
   https://leetcode.com/problems/two-sum
   ```
2. The skill runs the fetch script to download the problem description, examples, constraints, and any image assets:
   ```bash
   python3 .opencode/skills/go-interviewer/fetch_leetcode.py <url>
   ```
3. It creates a new workspace under `workspaces/leetcode-<number>-<slug>/` containing:
   - `README.md` — the full LeetCode description with difficulty tag
   - `main.go` — LeetCode's exact Go boilerplate (compiles with zero-value stubs)
   - `main_test.go` — table-driven behavioural tests derived from the LeetCode examples
   - `go.mod` — standalone module named after the problem slug
   - `assets/` — any images referenced in the problem statement
4. The skill then engages you in an interview on that problem, following the same boundaries (ask before telling, no full solutions, etc.).

This lets you practice real LeetCode problems with the same interactive interview experience as the built-in template questions.

## Templates

Each problem lives under `templates/<category>/qNN-short-name/` as a **standalone Go module** with its own `go.mod`. Most templates contain:

- `README.md` — problem statement, hints, follow-ups, and doc links
- `main.go` — starter code with TODO stubs
- `main_test.go` — behavioral tests
- `AGENTS.md` — interviewer-agent instructions used when reviewing a candidate's solution

Systems-design and some project-style questions may have only `go.mod` + `README.md` + `AGENTS.md` if they are discussion or debugging exercises rather than coding problems.

## Development

To rebuild after making changes:

```bash
docker-compose up --build
```

To stop and remove the container:

```bash
docker-compose down
```

To remove the container and the volume (deletes all generated workspaces):

```bash
docker-compose down -v
```

## Contributing / Adding Questions

See `AGENTS.md` for the conventions used when adding new question templates, including how to write the per-question `AGENTS.md` and the test suite.

## License

MIT
