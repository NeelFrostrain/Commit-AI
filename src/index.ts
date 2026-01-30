#!/usr/bin/env bun

import "dotenv/config";
import { Command } from "commander";
import Groq from "groq-sdk";
import { simpleGit, type SimpleGit } from "simple-git";
import * as readline from "node:readline/promises";
import { readFile } from "node:fs/promises";
import { join } from "node:path";
import chalk from "chalk";

const program = new Command();
const git: SimpleGit = simpleGit();
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

/**
 * Custom Logger
 * Format: [Origin] [Type]: Message
 */
const origin = chalk.bold.magenta("[Commit-AI]");

const log = {
  info: (msg: string) =>
    console.log(`${origin} ${chalk.blue("[Info]")}: ${msg}`),
  success: (msg: string) =>
    console.log(`${origin} ${chalk.green("[Success]")}: ${msg}`),
  warn: (msg: string) =>
    console.log(`${origin} ${chalk.yellow("[Warn]")}: ${msg}`),
  error: (msg: string) =>
    console.error(`${origin} ${chalk.red("[Error]")}: ${msg}`),
  ai: (msg: string) => console.log(`${origin} ${chalk.cyan("[AI]")}: ${msg}`),
};

async function getIgnorePatterns(): Promise<string[]> {
  const defaultExcludes = [
    "package-lock.json",
    "bun.lockb",
    "yarn.lock",
    "pnpm-lock.yaml",
    "node_modules",
    "dist",
    "*.log",
  ];
  try {
    const gitignorePath = join(process.cwd(), ".gitignore");
    const content = await readFile(gitignorePath, "utf-8");
    const gitignoreLines = content
      .split(/\r?\n/)
      .map((line) => line.trim())
      .filter((line) => line && !line.startsWith("#"));
    return Array.from(new Set([...defaultExcludes, ...gitignoreLines])).map(
      (pattern) => `:(exclude)${pattern}`,
    );
  } catch (e) {
    return defaultExcludes.map((pattern) => `:(exclude)${pattern}`);
  }
}

program
  .name("commit-ai")
  .description("AI-powered git analysis and auto-committer")
  .version("1.2.3")
  .option("-c, --commit", "enable commit mode")
  .option("-y, --yes", "skip confirmation prompt");

program.action(async (options) => {
  const apiKey = process.env.GROQ_API_KEY;
  if (!apiKey) {
    log.error("GROQ_API_KEY is missing from your environment variables.");
    process.exit(1);
  }

  const groq = new Groq({ apiKey });

  try {
    const isRepo = await git.checkIsRepo();
    if (!isRepo) {
      log.error("Current directory is not a Git repository.");
      return;
    }

    log.info("Analyzing modified files...");
    await git.add(["--intent-to-add", "."]);

    const excludePatterns = await getIgnorePatterns();
    let diff: string = "";
    try {
      diff = await git.diff(["HEAD", "--", ".", ...excludePatterns]);
    } catch (e) {
      const EMPTY_TREE_HASH = "4b825dc642cb6eb9a060e54bf8d69288fbee4904";
      diff = await git.diff([EMPTY_TREE_HASH, "--", ".", ...excludePatterns]);
    }

    if (!diff || diff.trim() === "") {
      log.success("No changes detected.");
      return;
    }

    const MAX_CHAR = 5000;
    if (diff.length > MAX_CHAR) {
      diff = diff.substring(0, MAX_CHAR) + "\n\n...[TRUNCATED]...";
    }

    const prompt = `
      Analyze this Git diff.
      1. Provide a bulleted "REPORT" of changes.
      2. Provide a "COMMIT_MESSAGE" in [type]: description format.
      
      RULES:
      - Format: [type]: description (example: [feat]: add login)
      - Use imperative mood.
      - No period at the end.

      Diff:
      ${diff}
    `;

    log.ai("Generating commit suggestion...");

    const chatCompletion = await groq.chat.completions.create({
      messages: [
        {
          role: "system",
          content: "You are a professional Git workflow assistant.",
        },
        { role: "user", content: prompt },
      ],
      model: "llama-3.1-8b-instant",
      temperature: 0.2,
    });

    const response = chatCompletion.choices[0]?.message?.content || "";
    const reportPart =
      response
        .split(/COMMIT_MESSAGE:/i)[0]
        ?.replace(/REPORT:/i, "")
        .trim() || "";
    let titlePart = response.split(/COMMIT_MESSAGE:/i)[1]?.trim() || "";

    // Force [type]: message format
    titlePart = titlePart.replace(/^(\w+):/, "[$1]:").replace(/\.$/, "");

    console.log(`\n${chalk.bold.cyan("─── AI SUGGESTION ───")}`);
    console.log(chalk.white(response));
    console.log(`${chalk.bold.cyan("─────────────────────")}\n`);

    if (options.commit && titlePart) {
      let shouldCommit = false;
      if (options.yes) {
        shouldCommit = true;
      } else {
        const confirm = await rl.question(
          `${origin} ${chalk.yellow("[Prompt]")}: Use this commit message? (y/n): `,
        );
        if (confirm.toLowerCase() === "y") shouldCommit = true;
      }

      if (shouldCommit) {
        await git.add(".");
        await git.commit([titlePart, reportPart]);
        log.success("Changes committed to history.");
      } else {
        log.warn("Commit aborted.");
      }
    } else if (!options.commit) {
      log.info(
        `Tip: Use ${chalk.bold("'-c'")} to commit directly from this tool.`,
      );
    }
  } catch (error: any) {
    log.error(`Critical Failure: ${error.message}`);
  } finally {
    rl.close();
  }
});

program.parse(process.argv);
