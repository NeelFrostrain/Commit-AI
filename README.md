# Commit-AI 🤖

<div align="left">

**Professionalize your Git history with AI-generated Conventional Commits.**

[![Version](https://img.shields.io/badge/version-1.1.0-red?style=flat-square)](https://github.com/NeelFrostrain/Commit-Ai-Go)
[![Go](https://img.shields.io/badge/go-%3E%3D1.25.6-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org)
[![Groq](https://img.shields.io/badge/Groq-AI-cyan?style=flat-square)](https://groq.com)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)

Stop writing "fixed stuff" and start writing commits that tell a story.

</div>

---

## 📖 About Commit-AI

**Commit-AI** transforms raw, technical code changes into **human-readable, professional documentation**. By leveraging Groq's inference engine (Llama 3.1 8B), it acts as a bridge between your terminal and your project's history.

### 🧠 The Logic

The tool reads your **Git Diff** to understand:

- **Intent:** Adding a feature or fixing a regression?
- **Impact:** What specific logic changed within the functions?
- **Context:** Automatically filters out noise like `package-lock.json` or `node_modules`.

---

## ⚡ Quick Start

### 1️⃣ Installation

**For Go Users:**

```bash
go install github.com/NeelFrostrain/Commit-Ai-Go@latest

```

**For Windows (Non-Go Users):**

1. Download the latest `commit-ai-installer.exe` from [Releases](https://github.com/NeelFrostrain/Commit-Ai-Go/releases/latest/).
2. Run the installer. It will automatically move the binary to your `AppData` and update your `PATH`.

### 2️⃣ Configuration (One-Time Setup)

You don't need to manually create `.env` files.

1. Visit [Groq Console](https://console.groq.com/keys) to get your key.
2. Run `commit-ai` in any terminal.
3. Paste your key when prompted. **Commit-AI** will save it to `~/.commit-ai-key` and your System Registry for global access.

---

## ✨ Features

| Feature                   | Description                                                |
| ------------------------- | ---------------------------------------------------------- |
| **🧠 Deep Diff Analysis** | Understands code logic, not just file metadata.            |
| **📝 Conventional Style** | Strictly follows the `type: description` standard.         |
| **📊 Technical Reports**  | Generates a detailed bulleted summary for the commit body. |
| **🛡️ Global Config**      | Set your API key once, use it in any project folder.       |
| **🚀 Fast**               | Powered by Groq/Llama-3.1 for near-instant results.        |

---

## 📖 Usage

### Command Flags

| Flag        | Short | Description                                             |
| ----------- | ----- | ------------------------------------------------------- |
| `--commit`  | `-c`  | Performs the `git commit` after generating the message. |
| `--yes`     | `-y`  | Skips the confirmation prompt (Auto-pilot).             |
| `--version` | `-v`  | Displays version information.                           |

### Example Workflow

1. **Stage changes:** `git add .`
2. **Review AI suggestion:** `commit-ai`
3. **Commit with AI:** `commit-ai -c`

---

## ⚙️ Standards & Privacy

### Conventional Categories

`feat`, `fix`, `docs`, `style`, `refactor`, `chore`.

### 🛡️ Privacy

- **Local Keys:** Your API key is stored locally on your machine.
- **Diffs Only:** Only the code diff of your **staged** files is sent to the AI.

---

## 📄 License

MIT © [Neel Frostrain](https://github.com/NeelFrostrain)

---
