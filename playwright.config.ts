import { defineConfig } from "@playwright/test";

export default defineConfig({
  testDir: "./tests/e2e",
  timeout: 30_000,
  use: {
    baseURL: "http://127.0.0.1:8080",
    headless: true,
  },
  webServer: {
    command: 'sh -lc "npm run build:css && go run ./examples/showcase"',
    port: 8080,
    reuseExistingServer: !process.env.CI,
    timeout: 120_000,
  },
});

