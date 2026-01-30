# Commit-AI 🤖

<div align="center">

[![Version](https://img.shields.io/badge/version-1.2.4-blue?style=flat-square)](https://github.com/NeelFrostrain/Commit-AI)
[![Node.js](https://img.shields.io/badge/node-%3E%3D18.0.0-brightgreen?style=flat-square)](https://nodejs.org)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)
[![TypeScript](https://img.shields.io/badge/typescript-%3E%3D5.0-blue?style=flat-square)](https://www.typescriptlang.org)

AI-assisted, CLI-first commit message generator and auto-committer

</div>

---

## 📋 Table of Contents

- [Overview](#-overview)
- [Quick Start](#-quick-start)
- [Features](#-features)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration & .env setup](#-configuration--env-setup)
- [Usage & Examples](#-usage--examples)
- [Prompt & Parsing Details](#-prompt--parsing-details)
- [Project Structure](#-project-structure)
- [Scripts & Commands](#-scripts--commands)
- [Troubleshooting](#-troubleshooting)
- [Security & Privacy](#-security--privacy)
- [Contributing & Docs](#-contributing--docs)
- [License](#-license)

---

## 🎯 Overview

**Commit-AI** analyzes your repository diff and uses an LLM (via Groq) to generate a concise, conventional commit-style title (`type: description`) and a short commit body. It sanitizes AI output and optionally performs the commit using `simple-git`.

This helps keep commit history meaningful while reducing friction when writing commit messages.

---

## ⚡ Quick Start

### 1) Prerequisites

- Bun (recommended) or Node.js (>=18)
- A `GROQ_API_KEY` (Groq chat completion API key)
- Git repo (run commands from a repository root)

### 2) Install

```bash
bun install
# or (npm)
npm install
```

### 3) Create `.env`

Create a `.env` file in the repository root:

```
GROQ_API_KEY=your_groq_api_key_here
```

PowerShell (session):

```powershell
$env:GROQ_API_KEY = "your_groq_api_key_here"
```

macOS/Linux/Bash:

```bash
export GROQ_API_KEY=your_groq_api_key_here
```

> Tip: add `.env` to `.gitignore` and use `.env.example` as a template (already included in repository).

### 4) Run

Dry run (generate suggestion only):

```bash
bun .
```

Generate and commit (interactive):

```bash
bun . -c
# Confirm with 'y' to commit
```

Auto-confirm and commit:

```bash
bun . -c -y
```

---

## ✨ Features

- AI-generated `COMMIT_MESSAGE` following `type: description` (e.g., `feat: add parser`)
- Short `COMMIT_BODY` (1–3 sentences) plus a `REPORT` of changes
- Sanitization: strips Markdown, normalizes bullets, enforces formatting
- Safe defaults & fallbacks for malformed AI output
- Uses `simple-git` to stage and commit using `[title, body]`
- Configurable truncation (limits diff size sent to the LLM)

---

## 🛠 Installation (Global / CLI)

If you want to use `commit-ai` as a global CLI tool:

- Using Bun linking (development):

```bash
bun link && bun link commit-ai
```

- Global npm install (after packaging/build):

```bash
# Build first
npm run build
# Then install globally
npm install -g .
```

- Build an executable (Windows example):

```bash
bun run build-exe
# Then run
./bin/commit-ai.exe -c -y
```

---

## ⚙️ Configuration & .env setup

- `GROQ_API_KEY` — required. Place in `.env` or set environment variable.
- `MAX_CHAR` — truncation limit (implemented as `5000` in `src/index.ts`); change in source if needed.

Example `.env.example` included in repository.

CI: Set `GROQ_API_KEY` as a repository secret in your CI provider; do not commit sensitive keys.

---

## 🔍 Prompt & Parsing Details

- The tool sends a focused prompt to Groq requesting three sections: `REPORT`, `COMMIT_MESSAGE` (single-line), and `COMMIT_BODY` (1-3 sentences).
- The diff is truncated to `MAX_CHAR` to limit tokens and cost.
- Parsing is done using regex; sanitization removes `**` and normalizes bullet points.
- Title validation enforces `type: description`. If missing, a fallback (`chore: update project files`) is used. If the title is too terse, a phrase from the body is appended.

---

## ✅ Usage Examples

Dry run:

```bash
bun .
```

Interactive commit with prompt:

```bash
bun . -c
# You'll see AI suggestion and be asked: Use this commit message? (y/n):
```

Auto-commit (skip prompt):

```bash
bun . -c -y
```

What gets committed:

- The commit subject is the `COMMIT_MESSAGE` (title)
- The commit body is `COMMIT_BODY` (or the top bullets from the REPORT as a fallback)

---

## 🗂 Project Structure

```
Commit-AI/
├── src/
│   └── index.ts       # Main CLI logic
├── bin/
│   └── commit-ai.exe  # Optional built Windows binary
├── package.json
├── tsconfig.json
├── README.md
├── .env.example
└── LICENSE
```

---

## 🧪 Scripts & Commands

|                           Command | What it does                             |
| --------------------------------: | :--------------------------------------- |
|                           `bun .` | Run CLI from source (suggest only)       |
|                        `bun . -c` | Generate suggestion and prompt to commit |
|                     `bun . -c -y` | Auto-confirm and commit                  |
| `bun run build` / `npm run build` | TypeScript compile to `dist/`            |
|               `bun run build-exe` | Build Windows executable                 |

---

## ⚠️ Troubleshooting

- `GROQ_API_KEY is missing.` — Create `.env` or export the env var.
- `git.commit: requires the commit message to be supplied as a string|string[]` — Ensure you are running the current source (run `bun .`) or rebuild `dist/` if running from compiled JS. The app now calls `git.commit([title, body])`.
- Commit fails — Check `git status`, pre-commit hooks, and file permissions. Errors are logged as `Git Commit Failed: <reason>`.

If the commit message looks too terse, use the `COMMIT_BODY` output (printed after suggestion) — the tool constructs a more descriptive title if the title is under 15 chars using body content.

---

## 🔒 Security & Privacy

- This tool sends diffs (code) to Groq — **do not** run against repositories with secrets or sensitive data unless you accept the risk.
- Add patterns to `.gitignore` or extend `getIgnorePatterns()` in `src/index.ts` to exclude sensitive paths from the diff.

---

## 🤝 Contributing & Docs

- See `CONTRIBUTING.md` for PR workflow and guidelines.
- `SECURITY.md` explains responsible disclosure.
- `CODE_OF_CONDUCT.md` outlines expected community behavior.

---

## 📝 License

MIT © Neel Frostrain — see `LICENSE` for details.

---

## 👤 Author

Neel Frostrain — https://github.com/NeelFrostrain

---

If you'd like, I can also:

- Add a `CHANGELOG.md` and example GitHub Actions CI workflow ✅
- Add unit + integration tests for parsing logic ✅
- Add additional sample outputs and a `docs/` folder for longer guides ✅

Tell me which extras you want next and I'll add them.

## Table of contents

- [Overview](#overview) ✅
- [Features](#features) ✨
- [Quickstart](#quickstart) 🔧
- [Detailed usage](#detailed-usage) 🧭
- [Prompt & parsing details](#prompt--parsing-details) 💬
- [Architecture & flow](#architecture--flow) 🏗️
- [Configuration & constants](#configuration--constants) ⚙️
- [Development & build notes](#development--build-notes) 🛠️
- [Project layout](#project-layout) 🗂️
- [Testing & recommended checks](#testing--recommended-checks) ✅
- [Troubleshooting](#troubleshooting) ⚠️
- [Security & privacy](#security--privacy) 🔒
- [Contributing](#contributing) 🤝
- [License](#license) 📄

---

## Overview

commit-ai helps maintain better commit hygiene by leveraging an LLM to craft conventional commit-style titles and short commit bodies that summarize the change. The tool is built to be safe for day-to-day development: it sanitizes AI output, enforces formatting rules, and provides fallbacks when the AI returns unexpected content.

## Features

- Structured AI output: `REPORT`, `COMMIT_MESSAGE`, and `COMMIT_BODY`.
- Sanitization: remove Markdown, normalize bullets, enforce `type: description` for the title.
- Commit workflow: stages changes and commits with `[title, body]` (string array) using `simple-git`.
- Configurable truncation of diffs to control prompt size and cost.
- Error handling: commit failures are reported clearly without crashing the tool.

## Quickstart

### Prerequisites

- Bun (recommended) or Node.js (if you change scripts)
- A Groq API key (set `GROQ_API_KEY` in your environment or `.env` file)
- Git (a local repo to run against)

### Install dependencies

```bash
bun install
# or (if you prefer npm)
npm install
```

### Setup your API key (.env)

Create a `.env` file at the project root with this content:

```
GROQ_API_KEY=your_groq_api_key_here
```

- For PowerShell (temporary, current session):

```powershell
$env:GROQ_API_KEY = "your_groq_api_key_here"
```

- For Command Prompt (temporary):

```cmd
set GROQ_API_KEY=your_groq_api_key_here
```

- For macOS/Linux or Git Bash:

```bash
export GROQ_API_KEY=your_groq_api_key_here
```

To set the variable permanently on Windows, add it via **System Properties → Environment Variables**.

> Tip: Keep `.env` out of version control. The project already respects `.gitignore`, but confirm you don't commit secrets.

### Run the CLI (dry run — suggestions only)

```bash
bun .
```

### Run and commit (interactive)

```bash
bun . -c
# Answer 'y' to confirm commit
```

### Run and commit (auto-confirm)

```bash
bun . -c -y
```

### Optional: global install (useful if you want `commit-ai` as a CLI)

- Using Bun (link the local package):

```bash
bun link && bun link commit-ai
```

- Using npm (publish or local link):

```bash
# after building (npm package-ready)
npm install -g .
# OR (development linking)
npm link
```

### Optional: build an executable (Windows example)

```bash
bun run build-exe
# or run the bundled exe if present
./bin/commit-ai.exe -c -y
```

If you need a `node`-based run (without Bun):

```bash
# transpile first
npm run build
# run built JS
node dist/index.js
```

## Detailed usage

- `-c, --commit` — enable commit mode (stages + commits changes if you approve)
- `-y, --yes` — bypass the confirmation prompt and auto-commit

Example output (formatted):

```
[Commit-AI] [AI]: Generating commit suggestion...

─── AI SUGGESTION ───
REPORT:
- Fixed parsing of markdown in AI responses
- Updated commit message generation

COMMIT_MESSAGE: feat: improve commit generation

COMMIT_BODY:
Generate a cleaner commit title and add a short description for context.
─────────────────────
[Commit-AI] [Prompt]: Use this commit message? (y/n):
```

When confirmed, the tool calls `git.commit([<title>, <body>])` so the title becomes the subject and the body becomes the commit body.

## Prompt & parsing details

- Prompt template: A focused prompt is sent to Groq's chat completion that requests `REPORT`, `COMMIT_MESSAGE`, and a short `COMMIT_BODY`.
- Truncation: the diff is truncated to the first 5000 characters (`diff.substring(0, 5000)`) to limit token usage and avoid overlong prompts.
- Parsing: responses are parsed via regex into three sections. The code sanitizes bold/markdown, normalizes bullets, and trims whitespace.
- Title validation: The title is forced to the `type: description` format. If missing or malformed, the tool sets a fallback (e.g., `chore: update project files`). If the description part is too short, a phrase from the body is appended to make the title informative.

### Example of expected AI response format

```
REPORT:
- Refactor index.ts parsing logic
- Update package.json scripts to point to source

COMMIT_MESSAGE: refactor: improve parsing and scripts

COMMIT_BODY:
Improve parsing logic to handle markdown and enforce title format. Update `package.json` for dev runs.
```

## Architecture & flow

1. Stage files (intent-to-add) and compute a diff.
2. Truncate diff and send to Groq chat completion with a strict system prompt.
3. Parse response into `report`, `title`, `commitBody`.
4. Sanitize and validate text.
5. Show suggestion to user; if confirmed, stage and commit using `simple-git`.

Key behaviors

- Non-fatal commit errors are caught and logged as `Git Commit Failed: <reason>`.
- The tool passes commit parameters as a string array to `simple-git` to satisfy the library's accepted signatures.

## Configuration & constants

- `GROQ_API_KEY` — **required** (set via environment or `.env`). See Quickstart for `.env` and platform-specific examples.
- `MAX_CHAR` (truncation length) — currently implemented as `5000` in `src/index.ts`.

Example `.env` file (project root):

```
# .env (do not commit this file)
GROQ_API_KEY=your_groq_api_key_here
```

To change behavior, edit `src/index.ts` and adjust:

- Prompt wording (how much detail you ask the model to include)
- Truncation length (`MAX_CHAR` / how much of the diff to include)
- Parsing/sanitization rules (how `REPORT`, `COMMIT_MESSAGE`, and `COMMIT_BODY` are extracted)

Environment variable notes

- You can export `GROQ_API_KEY` per-session (PowerShell/CMD/Bash) or set it permanently via system environment variables (Windows) or your shell profile (macOS/Linux).
- For CI, set `GROQ_API_KEY` as a repository/organization secret rather than committing it to the repo.

Security reminder: Do not include secrets in diffs you send to the LLM. Add patterns to `.gitignore` or extend `getIgnorePatterns()` to exclude sensitive files.

## Development & build notes

Scripts (from `package.json`):

- `bun run index.ts` — run source directly using Bun (recommended for development)
- `build` / `tsc` — compile to `dist/` for production
- `build-exe` — build a platform executable via Bun
- `full-build` — build + package + link

Notes:

- The project is TypeScript (`tsconfig.json` present). `@types/node` is in devDependencies.
- For local CLI testing, run: `bun link && bun link commit-ai`.

## Project layout

- `src/index.ts` — main CLI code
- `package.json` — metadata and scripts
- `tsconfig.json` — TypeScript configuration
- `bin/` — output build artifacts (e.g., `commit-ai.exe`)

## Testing & recommended checks

- Add unit tests for parsing logic (validate `REPORT`, `COMMIT_MESSAGE`, `COMMIT_BODY` extraction).
- Add integration tests that mock Groq responses and verify commit behavior with a temporary git repo.
- Add a linter and formatter (ESLint + Prettier) for consistency.

## Troubleshooting

- `GROQ_API_KEY is missing.` — Ensure `GROQ_API_KEY` is set in environment or `.env`.
- `git.commit: requires the commit message to be supplied as a string|string[]` — Rebuild `dist` if running from `dist/`, or run `bun .` to use `src` during development. The tool calls `git.commit([title, body])`.
- Commit fails — Check `git status`, hooks, and file permissions. Commit errors are logged as `Git Commit Failed: <reason>`.

## Security & privacy

- This tool sends diffs to an external AI provider (Groq). Do **not** use on repositories with sensitive secrets or proprietary data unless you accept the risk.
- If required, add local exclusion rules to avoid sending files with secrets (update `getIgnorePatterns()` in `src/index.ts`).

## Contributing

- Open an issue describing your change or file a PR.
- Recommended contributions:
  - Add unit & integration tests for parsing and commit logic
  - Improve prompt engineering for more concise or detailed commit bodies
  - Add CI configuration and release automation

## Repository docs

- `CONTRIBUTING.md` — Contribution workflow and PR guidance
- `SECURITY.md` — How to report security issues sensitively
- `CODE_OF_CONDUCT.md` — Community expectations
- `.env.example` — Template for local environment variables (do not commit secrets)

## License

This repository now includes an `LICENSE` file (MIT). See `LICENSE` for full terms.

---

## Appendix — Helpful tips

- Want longer commit bodies? Change the prompt to request more detail and/or use the full `REPORT` as the body.
- Want different commit types (e.g., `fix`, `feat`, `chore`)? Change the system prompt rules or add a mapping step in `src/index.ts`.

If you'd like, I can:

- Add `CONTRIBUTING.md`, `SECURITY.md`, and a `LICENSE` (MIT) file ✅
- Add unit tests for the parsing logic and a GitHub Actions CI workflow ✅
- Add sample AI responses as fixtures for tests ✅

Tell me which extras you'd like and I will add them. 🚀
