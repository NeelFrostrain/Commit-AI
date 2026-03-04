# 🎬 Commit-AI Live Demo

## 🚀 Quick Demo

### Scenario: You just improved your codebase

```bash
# You made changes to multiple files
$ git status
modified:   cmd/root.go
modified:   internal/ai/parser.go
new file:   internal/config/config.go
new file:   internal/git/diff.go
```

### Step 1: Stage your changes
```bash
$ git add .
```

### Step 2: Run Commit-AI
```bash
$ commit-ai -v
```

### Step 3: See the magic happen ✨

```
[Info] Detected scope: internal
[Info] Diff size: 12036 characters
[Commit-AI] Analyzing changes with llama-3.1-8b-instant...
[Info] Generated 3 options

? Select commit title:
> refactor(internal): restructure modules for better maintainability
  feat(internal): add configuration and git operation modules
  chore(internal): reorganize internal package structure
  ✏️  Edit manually...
  ❌ Cancel
```

### Step 4: Choose your favorite
```
? Commit body:
> 📝 Keep AI report
  ✏️  Edit report
  ⊘  No report
```

### Step 5: Review and commit
```
─────────────────────────────────
Title: refactor(internal): restructure modules for better maintainability
Body:
- Created dedicated config module for centralized configuration
- Added git/diff.go for optimized git operations
- Enhanced AI parser with validation and scope detection
- Improved error handling and user feedback
─────────────────────────────────

[Hint] Use -c flag to commit automatically
```

---

## 🎯 Feature Demonstrations

### 1. Verbose Mode

**Command**: `commit-ai -v`

**What you see**:
```
[Info] Detected scope: api
[Info] Diff size: 2847 characters
[Commit-AI] Analyzing changes with llama-3.1-8b-instant...
[Info] Generated 3 options
```

**Benefits**:
- See what scope was detected
- Know how much context AI is analyzing
- Confirm which model is being used
- Track how many options were generated

---

### 2. Custom Model

**Command**: `commit-ai -m llama-3.1-70b-versatile`

**What happens**:
- Uses more powerful model
- Better for complex changes
- More accurate descriptions
- Slightly slower but worth it

**Use when**:
- Large refactoring
- Complex features
- Need extra accuracy

---

### 3. Auto-Commit Mode

**Command**: `commit-ai -c`

**Workflow**:
```
1. Generates options
2. You select one
3. You choose report option
4. Asks for confirmation
5. Commits automatically
```

**Even faster**: `commit-ai -cy`
- Skips all confirmations
- Commits immediately
- Use with caution!

---

### 4. Scope Detection

**Example 1: API Changes**
```bash
$ git add api/handler.go api/routes.go
$ commit-ai

# AI suggests:
feat(api): add user authentication endpoints
```

**Example 2: UI Changes**
```bash
$ git add ui/components/Login.tsx ui/styles/auth.css
$ commit-ai

# AI suggests:
feat(ui): implement login component with styling
```

**Example 3: Mixed Changes**
```bash
$ git add api/auth.go ui/Login.tsx database/users.sql
$ commit-ai -v

# Shows:
[Info] Detected scope: api
# AI suggests:
feat(api): add authentication system with UI and database
```

---

## 🎨 Real-World Examples

### Example 1: Bug Fix

**Changes**:
```diff
// cache/manager.go
- func (m *Manager) cleanup() {
-     // Memory leak here
-     m.items = make(map[string]Item)
+ func (m *Manager) cleanup() {
+     for k := range m.items {
+         delete(m.items, k)
+     }
```

**AI Generated**:
```
1. fix(cache): resolve memory leak in cleanup method
2. fix(cache): properly clear cache items to prevent memory leak
3. fix: fix memory management in cache cleanup
```

**Selected**: `fix(cache): resolve memory leak in cleanup method`

**Report**:
```
- Changed cleanup to delete items individually
- Prevents memory leak from map reallocation
- Improves long-running application stability
```

---

### Example 2: New Feature

**Changes**:
```diff
// api/auth.go
+ func LoginHandler(w http.ResponseWriter, r *http.Request) {
+     // JWT authentication logic
+ }

// api/middleware.go
+ func AuthMiddleware(next http.Handler) http.Handler {
+     // Token validation
+ }
```

**AI Generated**:
```
1. feat(api): add JWT authentication system
2. feat(auth): implement login handler and middleware
3. feat(api): add user authentication with JWT tokens
```

**Selected**: `feat(api): add JWT authentication system`

**Report**:
```
- Implemented JWT-based authentication
- Added login handler for user credentials
- Created middleware for protected routes
- Includes token validation and refresh logic
```

---

### Example 3: Refactoring

**Changes**:
```diff
// Before: Everything in one file
- // main.go (500 lines)

// After: Organized structure
+ // internal/config/config.go
+ // internal/git/diff.go
+ // internal/ai/parser.go
```

**AI Generated**:
```
1. refactor(internal): restructure codebase into modules
2. refactor: organize code into dedicated packages
3. chore(internal): split monolithic code into modules
```

**Selected**: `refactor(internal): restructure codebase into modules`

**Report**:
```
- Created config module for configuration management
- Added git module for git operations
- Enhanced AI module with validation
- Improved code organization and maintainability
```

---

## 🔄 Workflow Comparison

### Old Way (Manual)
```bash
$ git add .
$ git commit -m "update stuff"
# 😞 Vague, unhelpful commit message
```

### New Way (Commit-AI)
```bash
$ git add .
$ commit-ai -c
# 😊 Professional, descriptive commit message
# ✅ Follows Conventional Commits
# 📝 Includes detailed report
# ⚡ Takes 10 seconds
```

---

## 💡 Pro Tips Demo

### Tip 1: Selective Staging
```bash
# Only stage related changes
$ git add src/auth/*
$ commit-ai
# AI focuses on auth changes only
# Result: More accurate commit message
```

### Tip 2: Review Before Committing
```bash
# Generate message first
$ commit-ai

# Review the suggestion
# Then commit manually if you want to edit
$ git commit -m "your edited message"
```

### Tip 3: Use Verbose for Learning
```bash
# See what AI is analyzing
$ commit-ai -v

# Learn from the AI's reasoning
# Understand scope detection
# See diff size impact
```

### Tip 4: Different Models for Different Tasks
```bash
# Quick commits: default model
$ commit-ai -c

# Complex changes: powerful model
$ commit-ai -m llama-3.1-70b-versatile -c

# Experimentation: try different models
$ commit-ai -m mixtral-8x7b-32768
```

---

## 🎭 Before & After Examples

### Before Commit-AI
```
commit ab5d73f
Author: Developer
Date:   Mon Jan 1 12:00:00 2024

    fixed stuff

commit 9c8e2a1
Author: Developer
Date:   Mon Jan 1 11:00:00 2024

    updates

commit 7f3b9d2
Author: Developer
Date:   Mon Jan 1 10:00:00 2024

    wip
```

### After Commit-AI
```
commit ab5d73f
Author: Developer
Date:   Mon Jan 1 12:00:00 2024

    fix(cache): resolve memory leak in cleanup method
    
    - Changed cleanup to delete items individually
    - Prevents memory leak from map reallocation
    - Improves long-running application stability

commit 9c8e2a1
Author: Developer
Date:   Mon Jan 1 11:00:00 2024

    feat(api): add JWT authentication system
    
    - Implemented JWT-based authentication
    - Added login handler for user credentials
    - Created middleware for protected routes

commit 7f3b9d2
Author: Developer
Date:   Mon Jan 1 10:00:00 2024

    refactor(internal): restructure codebase into modules
    
    - Created config module for configuration management
    - Added git module for git operations
    - Enhanced AI module with validation
```

**Result**: Professional, searchable, meaningful git history! 🎉

---

## 🚦 Common Scenarios

### Scenario 1: "I forgot to stage files"
```bash
$ commit-ai
[!] No staged changes detected.
? Stage all files and continue? (y/N)
> y

[✓] All files staged
[Commit-AI] Analyzing changes...
```

### Scenario 2: "I don't like any of the suggestions"
```bash
? Select commit title:
  feat(api): add authentication
  feat(auth): implement login
  chore(security): add middleware
> ✏️  Edit manually...

? Enter custom commit title: feat(api): add OAuth2 authentication
```

### Scenario 3: "I want to skip the report"
```bash
? Commit body:
  📝 Keep AI report
  ✏️  Edit report
> ⊘  No report

─────────────────────────────────
Title: feat(api): add authentication
Body: (none)
─────────────────────────────────
```

### Scenario 4: "I need to see what changed"
```bash
$ commit-ai -v
[Info] Detected scope: api
[Info] Diff size: 2847 characters

# If you want more details:
$ git diff --cached --stat
```

---

## 🎓 Learning from AI

### Watch the AI Think
```bash
$ commit-ai -v

# AI analyzes:
# 1. File paths → Detects scope
# 2. Diff content → Understands changes
# 3. Change type → Determines feat/fix/refactor
# 4. Impact → Generates meaningful description
```

### Improve Your Commit Skills
```
Use Commit-AI regularly and you'll learn:
- How to write better commit messages
- What makes a good description
- How to structure commit bodies
- Conventional Commits format
```

---

## 🎬 Video Script (If You Make One)

```
[0:00] "Tired of writing vague commit messages?"
[0:05] "Meet Commit-AI - your AI-powered commit assistant"
[0:10] [Show: git add .]
[0:12] [Show: commit-ai -v]
[0:15] [Show: AI generating 3 options]
[0:20] [Show: Selecting an option]
[0:25] [Show: Beautiful commit message]
[0:30] "Professional commits in seconds"
[0:35] "Try Commit-AI today!"
```

---

## 📸 Screenshots (Recommended)

1. **Help Screen**: Show all flags
2. **Verbose Mode**: Show detailed output
3. **Option Selection**: Show 3 AI-generated options
4. **Final Commit**: Show formatted commit message
5. **Git Log**: Show before/after comparison

---

## 🎉 Try It Yourself!

```bash
# Install
go install github.com/NeelFrostrain/Commit-Ai-Go@latest

# Set up API key
commit-ai
# (Follow the prompts)

# Make some changes
echo "// New feature" >> main.go
git add .

# Generate commit
commit-ai -v

# Enjoy professional commits! 🚀
```

---

**Demo Status**: ✅ Ready to Show
**Recommended For**: README, Social Media, Presentations
**Duration**: 2-5 minutes
