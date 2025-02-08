import { defineConfig } from '@playwright/test';

export default defineConfig({
    use: {
        baseURL: 'http://localhost:5173',
        headless: false, // Run tests in visible mode
        viewport: { width: 980, height: 620 },
        screenshot: 'off',
        video: 'off',
        browserName: 'chromium', // Ensure it's set to Chromium
        channel: 'chrome', // Use the system-installed Chrome
    },
    webServer: {
        command: 'npm run dev',
        port: 5173,
        timeout: 3000,
    },
});
