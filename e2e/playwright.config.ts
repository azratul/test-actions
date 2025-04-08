import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  timeout: 5000,
  retries: 0,
  use: {
    baseURL: 'http://localhost:5000',
    headless: true,
  },
});
