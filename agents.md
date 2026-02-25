# Agents

## Ralph Loop

The Ralph Loop is an autonomous, iterative agent methodology.

### How it works

Run a loop. Each iteration:

1. **Pick** — read the plan, find the next uncompleted task
2. **Implement** — do the work for that task
3. **Verify** — confirm the implementation is correct
4. **Commit** — commit the changes to git
5. **Update progress** — mark the task as done in the plan
6. **Repeat** — start the next iteration

### Key principle

Each iteration begins with a fresh context. All state lives on disk and in git history — never in memory.
