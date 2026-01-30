#!/usr/bin/env bun

import * as dotenv from "dotenv";
import { Command } from "commander";
import Groq from "groq-sdk";
import { simpleGit, type SimpleGit } from "simple-git";
import * as readline from "node:readline/promises";

dotenv.config();

const program = new Command();
const git: SimpleGit = simpleGit();
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

program
  .name("commit-ai")
  .description("AI-powered git analysis and auto-committer")
  .version("1.1.2");

program.action(async () => {
  const apiKey = process.env.GROQ_API_KEY;

  if (!apiKey) {
    console.error("❌ Error: GROQ_API_KEY is not set in your .env file.");
    process.exit(1);
  }

  const groq = new Groq({ apiKey });

  try {
    const isRepo = await git.checkIsRepo();
    if (!isRepo) {
      console.error("❌ Not a Git repo. Run 'git init' first.");
      return;
    }

    console.log("🔍 commit-ai is scanning your changes...");

    // 1. Stage changes (respects .gitignore naturally)
    await git.add(["--intent-to-add", "."]);

    // 2. Prepare diff with filters to save tokens (avoids 413 error)
    const excludePatterns = [
      ":!package-lock.json",
      ":!bun.lockb",
      ":!yarn.lock",
      ":!pnpm-lock.yaml",
      ":!node_modules",
      ":!dist",
      ":!*.log",
    ];

    let diff: string = "";
    try {
      // Normal diff against HEAD
      diff = await git.diff(["HEAD", "--", ".", ...excludePatterns]);
    } catch (e) {
      // Fallback for initial commit (diff against empty tree)
      const EMPTY_TREE_HASH = "4b825dc642cb6eb9a060e54bf8d69288fbee4904";
      diff = await git.diff([EMPTY_TREE_HASH, "--", ".", ...excludePatterns]);
    }

    if (!diff || diff.trim() === "") {
      console.log("✅ No changes found. Your working tree is clean.");
      return;
    }

    // 3. Strict token management (Stay under Groq's 6000 TPM limit)
    const MAX_CHAR = 5000;
    if (diff.length > MAX_CHAR) {
      diff =
        diff.substring(0, MAX_CHAR) +
        "\n\n...[DIFF TRUNCATED TO SAVE TOKENS]...";
    }

    const prompt = `
      You are an expert software engineer. Analyze this Git diff and provide:
      1. A bulleted "REPORT" of technical changes.
      2. A "COMMIT_MESSAGE" following Conventional Commits.
      
      Response Format:
      REPORT:
      - detail 1
      COMMIT_MESSAGE:
      type(scope): description

      Diff:
      ${diff}
    `;

    console.log("🚀 AI is drafting your professional commit...");

    const chatCompletion = await groq.chat.completions.create({
      messages: [
        {
          role: "system",
          content:
            "You are commit-ai, a specialized assistant that only generates professional Git documentation.",
        },
        { role: "user", content: prompt },
      ],
      model: "llama-3.1-8b-instant",
      temperature: 0.2,
    });

    const response = chatCompletion.choices[0]?.message?.content || "";

    // Extract the commit message part
    const commitMsg = response.split(/COMMIT_MESSAGE:/i)[1]?.trim() || "";

    console.log("\n--- 📝 commit-ai: PROFESSIONAL REPORT ---");
    console.log(response);
    console.log("------------------------------------------\n");

    if (commitMsg) {
      const confirm = await rl.question(
        `🤔 Commit with message: "${commitMsg}"? (y/n): `,
      );

      if (confirm.toLowerCase() === "y") {
        await git.add("."); // Final stage
        await git.commit(commitMsg);
        console.log("✅ Changes committed successfully!");
      } else {
        console.log("👋 Commit aborted by user.");
      }
    }
  } catch (error: any) {
    if (error.status === 413) {
      console.error(
        "❌ Error: The diff is too large for the AI to process. Try staging smaller chunks.",
      );
    } else {
      console.error("❌ commit-ai Error:", error.message);
    }
  } finally {
    rl.close();
  }
});

program.parse(process.argv);
