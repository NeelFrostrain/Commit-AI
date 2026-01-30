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
  .description(
    "AI-powered git change analysis and conventional commit generator",
  )
  .version("1.1.1");

program.action(async () => {
  const apiKey = process.env.GROQ_API_KEY;

  if (!apiKey) {
    console.error("❌ Error: GROQ_API_KEY is not set in .env");
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

    await git.add(["--intent-to-add", "."]);

    // Added filters: Ignore lockfiles and large assets to save tokens
    let diff: string = "";
    try {
      diff = await git.diff([
        "HEAD",
        "--",
        ".",
        ":!package-lock.json",
        ":!bun.lockb",
        ":!node_modules",
      ]);
    } catch (e) {
      const EMPTY_TREE_HASH = "4b825dc642cb6eb9a060e54bf8d69288fbee4904";
      diff = await git.diff([
        EMPTY_TREE_HASH,
        "--",
        ".",
        ":!package-lock.json",
        ":!bun.lockb",
        ":!node_modules",
      ]);
    }

    if (!diff || diff.trim() === "") {
      console.log("✅ No changes found. Your working tree is clean.");
      return;
    }

    // Lowered character limit to stay safe under Groq's 6000 TPM limit
    const MAX_CHAR = 5000;
    if (diff.length > MAX_CHAR) {
      diff =
        diff.substring(0, MAX_CHAR) +
        "\n\n...[DIFF TRUNCATED TO SAVE TOKENS]...";
    }

    const prompt = `
      You are an expert software engineer. Analyze this Git diff and provide:
      1. A bulleted "REPORT" of technical changes.
      2. A "COMMIT_MESSAGE" (Conventional Commits style).
      
      Structure your response exactly like this:
      REPORT:
      - change details
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
          content: "You are commit-ai, a professional Git assistant.",
        },
        { role: "user", content: prompt },
      ],
      model: "llama-3.1-8b-instant",
      temperature: 0.2,
    });

    const response = chatCompletion.choices[0]?.message?.content || "";

    // Extract commit message using split
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
        console.log("👋 Commit aborted.");
      }
    }
  } catch (error: any) {
    console.error("❌ commit-ai Error:", error.message);
  } finally {
    rl.close();
  }
});

program.parse(process.argv);
