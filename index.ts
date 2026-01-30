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
  .version("1.2.0")
  .option("-c, --commit", "enable commit mode (prompts to commit changes)")
  .option("-y, --yes", "skip confirmation prompt (requires -c)");

program.action(async (options) => {
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

    // Stage changes for analysis
    await git.add(["--intent-to-add", "."]);

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
      diff = await git.diff(["HEAD", "--", ".", ...excludePatterns]);
    } catch (e) {
      const EMPTY_TREE_HASH = "4b825dc642cb6eb9a060e54bf8d69288fbee4904";
      diff = await git.diff([EMPTY_TREE_HASH, "--", ".", ...excludePatterns]);
    }

    if (!diff || diff.trim() === "") {
      console.log("✅ No changes found. Your working tree is clean.");
      return;
    }

    const MAX_CHAR = 5000;
    if (diff.length > MAX_CHAR) {
      diff =
        diff.substring(0, MAX_CHAR) +
        "\n\n...[DIFF TRUNCATED TO SAVE TOKENS]...";
    }

    const prompt = `
      You are an expert engineer. Analyze this Git diff.
      1. Provide a bulleted "REPORT" of technical changes.
      2. Provide a "COMMIT_MESSAGE" (Conventional Commits style).
      
      Response Format:
      REPORT:
      - change details
      COMMIT_MESSAGE:
      type(scope): description

      Diff:
      ${diff}
    `;

    console.log("🚀 AI is drafting your professional report...");

    const chatCompletion = await groq.chat.completions.create({
      messages: [
        {
          role: "system",
          content:
            "You are commit-ai, a specialized assistant for professional Git documentation.",
        },
        { role: "user", content: prompt },
      ],
      model: "llama-3.1-8b-instant",
      temperature: 0.2,
    });

    const response = chatCompletion.choices[0]?.message?.content || "";
    const commitMsg = response.split(/COMMIT_MESSAGE:/i)[1]?.trim() || "";

    console.log("\n--- 📝 commit-ai: PROFESSIONAL REPORT ---");
    console.log(response);
    console.log("------------------------------------------\n");

    // Logic for committing
    if (options.commit && commitMsg) {
      let shouldCommit = false;

      if (options.yes) {
        console.log("⏩ Auto-commit flag detected (-y).");
        shouldCommit = true;
      } else {
        const confirm = await rl.question(
          `🤔 Commit with message: "${commitMsg}"? (y/n): `,
        );
        if (confirm.toLowerCase() === "y") shouldCommit = true;
      }

      if (shouldCommit) {
        await git.add(".");
        await git.commit(commitMsg);
        console.log("✅ Changes committed successfully!");
      } else {
        console.log("👋 Commit aborted.");
      }
    } else if (!options.commit) {
      console.log("💡 Note: Run with '-c' to enable commit mode.");
    }
  } catch (error: any) {
    if (error.status === 413) {
      console.error("❌ Error: Diff too large for AI. Stage smaller chunks.");
    } else {
      console.error("❌ commit-ai Error:", error.message);
    }
  } finally {
    rl.close();
  }
});

program.parse(process.argv);
