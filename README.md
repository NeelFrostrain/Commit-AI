# Commit-AI 🤖

<div align="center">

[![Version](https://img.shields.io/badge/version-1.2.4-blue?style=flat-square)](https://github.com/NeelFrostrain/Commit-AI)
[![Node.js](https://img.shields.io/badge/node-%3E%3D18.0.0-brightgreen?style=flat-square)](https://nodejs.org)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)
[![TypeScript](https://img.shields.io/badge/typescript-%3E%3D5.0-blue?style=flat-square)](https://www.typescriptlang.org)

**Professionalize your Git history with AI-generated Conventional Commits**

</div>

---

## 📋 Table of contents

- [Overview](#-overview)
- [Quick start](#-quick-start)
- [Features](#-features)
- [Installation](#-installation)
- [Configuration (.env)](#-configuration-env)
- [CLI usage & examples](#-cli-usage--examples)
- [Prompt & parsing details](#-prompt--parsing-details)
- [Advanced configuration](#-advanced-configuration)
- [Testing & CI](#-testing--ci)
- [Publishing to npm](#-publishing-to-npm)
- [Troubleshooting](#-troubleshooting)
- [Security & privacy](#-security--privacy)
- [Contributing & docs](#-contributing--docs)
- [License & author](#-license--author)

---

## 🎯 Overview

Commit-AI reads your staged changes (via `git diff`), sends a focused prompt to a chat model (Groq/Llama), and returns three structured outputs: **REPORT**, **COMMIT_MESSAGE**, and **COMMIT_BODY**. It sanitizes the output and optionally commits using `simple-git`.

Use it to consistently produce clear, conventional commit messages across your repo.

---

## ⚡ Quick start

1. Install dependencies

```bash
bun install
# or
npm install
```

2. Create a `.env` file (see configuration)

3. Try a dry run

```bash
bun .
```

4. Commit interactively

```bash
bun . -c
# confirm with 'y' to commit
```

5. Auto-commit (no prompt)

```bash
bun . -c -y
```

---

## ✨ Features

- Generates `COMMIT_MESSAGE` in `type: description` format (conventional commits)
- Provides a concise `COMMIT_BODY` (1–3 sentences) and a bulleted `REPORT`
- Sanitizes AI output (removes Markdown, normalizes bullets)
- Safe fallbacks for malformed responses
- Uses `simple-git` for staging and committing with a title + body

---

## 🧩 Installation

- Local development: `bun install` (or `npm install`)
- Build for release: `npm run build` (emits `dist/`)
- Development link (global CLI):

```bash
bun link && bun link commit-ai
# or after build
npm install -g .
```

---

## ⚙️ Configuration (.env)

Create a `.env` at the project root:

```env
GROQ_API_KEY=your_groq_api_key_here
```

Platform examples:

- PowerShell (session):

```powershell
$env:GROQ_API_KEY = "your_key_here"
```

- macOS / Linux / Git Bash:

```bash
export GROQ_API_KEY=your_key_here
```

> Tip: add `.env` to `.gitignore` and keep `.env.example` as a template.

---

## 💻 CLI usage & examples

- Dry run (preview only):

```bash
bun .
```

- Interactive commit (prompt):

```bash
bun . -c
```

- Auto-confirm and commit:

```bash
bun . -c -y
```

Sample output (what you'll see):

```
─── AI SUGGESTION ───
REPORT:
- Fix markdown parsing in AI responses
- Add commit body generation

COMMIT_MESSAGE: feat: improve commit generation
COMMIT_BODY:
Improve parsing of AI responses to avoid artifacts and add brief context.
─────────────────────
Use this commit message? (y/n):
```

---

## 🔍 Prompt & parsing details

- The model is asked to return **REPORT**, **COMMIT_MESSAGE** (single line), and **COMMIT_BODY** (1–3 sentences).
- We truncate the diff to `MAX_CHAR` (default 5000) to limit token usage.
- Parsing uses regex and sanitization steps to ensure stable output.
- If the title is malformed, fallback is used: `chore: update project files`.

---

## ⚙️ Advanced configuration

- Edit `MAX_CHAR` in `src/index.ts` to increase/decrease diff size.
- Modify the `prompt` template in `src/index.ts` to change the model's output style.
- To include more detail in `COMMIT_BODY`, request it explicitly in the prompt.

---

## 🧪 Testing & CI (suggested)

- Add unit tests for parsing (`REPORT`, `COMMIT_MESSAGE`, `COMMIT_BODY`) using fixtures
- Example GitHub Actions workflow:

```yaml
name: CI
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 18
      - run: npm ci
      - run: npm run build
      - run: npm test
```

I can scaffold tests + workflow if you'd like.

---

## 📦 Publishing to npm

Checklist before publishing:

- `npm run build` (verify `dist/` contains `index.js`)
- Ensure `package.json` `name` & `version` are correct
- `npm whoami` → confirm authentication
- `npm publish --access public` (if using a scoped public package)

Common issues:

- `401/404` → token expired or not logged in (run `npm login`)
- Name conflicts → choose a different package name or scope

---

## ⚠️ Troubleshooting

- **Missing API key:** set `GROQ_API_KEY` or use `.env`
- **Commit errors:** check `git status`, hooks, and permissions
- **AI output looks wrong:** tweak the prompt or check the truncated diff

---

## 🔒 Security & privacy

- Diffs (not your entire repo) are sent to Groq. Do not include secrets in staged diffs.
- Use `.gitignore` or `getIgnorePatterns()` to exclude sensitive paths.

---

## 🤝 Contributing & docs

- `CONTRIBUTING.md`, `SECURITY.md`, and `CODE_OF_CONDUCT.md` are included; please follow them.
- Open issues or PRs for improvements and test additions.

---

## 📝 License & author

MIT © Neel Frostrain — see `LICENSE` for full text.

---

If you'd like, I can also:

- Add a `CHANGELOG.md`, unit tests for parsing, and a GitHub Actions workflow ✅
- Add `docs/` with example AI responses and expected parser outputs ✅

Tell me which extras you want next and I'll add them.
Run this command (replace `your_key_here` with your actual key):

```cmd
setx GROQ_API_KEY "your_key_here"

```

#### **Option B: PowerShell**

Run this command:

```powershell
[System.Environment]::SetEnvironmentVariable('GROQ_API_KEY', 'your_key_here', 'User')

```

> **⚠️ Important:** You **must restart** your terminal (CMD, PowerShell, or VS Code) after running these commands for the changes to take effect.

---

## ✨ Features

| Feature                   | Description                                                      |
| ------------------------- | ---------------------------------------------------------------- |
| **🧠 Deep Diff Analysis** | Understands code logic, not just file metadata.                  |
| **📝 Conventional Style** | Strictly follows the `type: description` standard.               |
| **📊 Technical Reports**  | Generates a detailed bulleted summary for the commit body.       |
| **🛡️ Smart Filtering**    | Respects `.gitignore` and ignores heavy lockfiles automatically. |
| **🚀 Sub-second Speed**   | Powered by Groq for nearly instant commit generation.            |

---

## 📖 Usage

### Command Flags

| Flag        | Short | Description                                             |
| ----------- | ----- | ------------------------------------------------------- |
| `--commit`  | `-c`  | Performs the `git commit` after generating the message. |
| `--yes`     | `-y`  | Skips the confirmation prompt (Auto-pilot).             |
| `--version` | `-v`  | Displays the current version.                           |
| `--help`    | `-h`  | Displays the help menu.                                 |

### Example Workflow

1. **Stage your changes:**

```bash
git add .

```

2. **Run Commit-AI:**

```bash
commit-ai -c

```

3. **Review & Confirm:** The AI will show you a report and the suggested message. Type `y` to finalize!

---

## ⚙️ Standards & Security

### Conventional Commit Types

Commit-AI automatically categorizes your work into:

- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation updates
- `style`: Formatting/Linting
- `refactor`: Code restructuring
- `chore`: Build tasks/dependencies

### 🛡️ Privacy

- **Local Keys:** Your API key stays on your machine.
- **Diffs Only:** Only the `git diff` of your **staged** files is sent to the AI for processing. No other system data is accessed.

---

## 📄 License

MIT © [Neel Frostrain](https://www.google.com/search?q=https://github.com/NeelFrostrain)

---
