import {expect, test} from '@playwright/test';
import generateTOTP from "./generateTOTP.js";

test('should register and login with authenticator successfully', async ({page}) => {
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


    await page.locator('a[href="/account/authenticator-apps"]').click();

    await page.locator('a[href="/account/authenticator-setup"]').click();

    await page.locator('a[href="/account/authenticator-setup/google"]').click();

    await page.waitForTimeout(500);
    const secretKey = await page.locator('span[data-test_id="secret_key"]');
    let a = await secretKey.innerText()

    const otp = generateTOTP(a);
    await page.locator('button[data-test_id="confirm_setup_button"]').click();

    await page.waitForTimeout(100);

    // await page.locator('button[data-test_id="auth_menu"]').click();
    // await page.locator('button[data-test_id="logout_button"]').first().click();

    // Go to login page
    await page.goto('http://localhost:5173/login/authenticator');

    for (let i = 0; i < otp.length; i++) {
        await page.locator(`input[type="text"]`).nth(i).fill(otp[i]);
    }

    await page.locator('button[type="submit"]').click();

    await page.waitForTimeout(5000);
});
