import {expect, test} from '@playwright/test';

test('should register and login successfully', async ({page}) => {
    const email = `test-${Date.now()}@gmail.com`;
    const username = `test-${Date.now()}`;
    const password = "123456";

    // Go to registration page
    await page.goto('http://localhost:5173/registration');

    // Fill registration form
    await page.fill('input[name="username"]', username);
    await page.fill('input[name="email"]', email);
    await page.fill('input[name="password"]', password);
    await page.click('button[type="submit"]');


    // Go to login page
    await page.goto('http://localhost:5173/login');


    await page.fill('input[name="email"]', email);
    await page.fill('input[name="password"]', password);
    await page.click('button[type="submit"]');

    await expect(page).toHaveURL('http://localhost:5173/account');
});
