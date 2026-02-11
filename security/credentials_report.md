# Credentials / Secrets Report

Scanned repository for likely credentials and secret patterns. No obvious private key material (e.g. "-----BEGIN PRIVATE KEY-----") or AWS access keys were found. Below are relevant matches and recommendations.

- **.github/workflows/ci.yml**: contains CI-created values and a POSTGRES_PASSWORD entry.
  - Excerpt: "POSTGRES_PASSWORD: postgres" (redacted: postgres)
  - Recommendation: keep (value is a harmless placeholder for CI). Consider storing real secrets in GitHub Actions secrets and not in workflows.

- **.github/workflows/ci.yml**: workflow creates a `.env.ci` in CI with `POSTGRES_PASSWORD=postgres`.
  - Excerpt: "POSTGRES_PASSWORD=postgres"
  - Recommendation: keep as example in CI; ensure real credentials are provided via secrets.

- **docker-compose.yml**: references `env_file: .env`.
  - Excerpt: "env_file: .env"
  - Recommendation: keep; ensure `.env` is in `.gitignore` and does not contain real secrets in the repo.

- **config/config.go**: default fallback password in code.
  - Excerpt: `Password: getEnv("DB_PASSWORD", "inventory_pass")`
  - Recommendation: keep code; `inventory_pass` is a default placeholder. Do NOT rely on defaults for production; inject via environment or secret manager.

- **internal/repository/postgres_store_integration_test.go** and **internal/repository/postgres_integration_test.go**: tests search for `.env` files and environment variables for DB credentials.
  - Excerpt: envPaths includes ".env"
  - Recommendation: tests should use `.env.ci` for CI. I've added `.env.ci.example` and a helper script. Ensure CI populates `.env.ci` via secrets.

- **.gitignore**: contains `.env` and `.env.local` (good).
  - Excerpt: ".env" and ".env.local"
  - Recommendation: keep; ensure `.env.ci` is also ignored (added by this change).

Summary and follow-ups:
- No high-confidence secret strings (API keys, private keys, AWS keys) were found committed in the working tree.
- The only concrete credential-like values are placeholders (`postgres`, `inventory_pass`) used in CI/workflow or as code defaults — these are benign but should be replaced by secrets in production/CI.
- If you believe any of the placeholder values had been replaced by real credentials in past commits, rotate those credentials and remove them from history manually; do not rewrite history automatically in this script.

If you'd like, I can also:
- Add scanning to CI (e.g., truffleHog or git-secrets) to block accidental commits of secrets.
- Open a PR to update docs with secrets-handling guidance.
