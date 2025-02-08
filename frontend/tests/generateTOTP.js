import crypto from 'crypto';

function generateTOTP(secret, timeStep = 30, digits = 6) {
    const time = Math.floor(Date.now() / 1000);
    const timeCounter = Math.floor(time / timeStep);

    const secretBytes = decodeBase32(secret);
    const counterBytes = Buffer.alloc(8);

    for (let i = 0; i < 8; i++) {
        counterBytes[7 - i] = (timeCounter >> (i * 8)) & 0xff;
    }

    const hmac = hmacSha256(secretBytes, counterBytes);
    const offset = hmac[hmac.length - 1] & 0x0f;

    const otp = ((hmac[offset] & 0x7f) << 24 |
        (hmac[offset + 1] & 0xff) << 16 |
        (hmac[offset + 2] & 0xff) << 8  |
        (hmac[offset + 3] & 0xff)) % (10 ** digits);

    return otp.toString().padStart(digits, '0');
}

function decodeBase32(base32) {
    const base32Alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567';
    const cleanBase32 = base32.replace(/[^A-Z2-7]/g, '');
    const buffer = [];
    let bits = 0;
    let value = 0;

    for (let char of cleanBase32) {
        const index = base32Alphabet.indexOf(char);
        if (index === -1) throw new Error('Invalid Base32 character');

        value = (value << 5) | index;
        bits += 5;

        if (bits >= 8) {
            buffer.push((value >> (bits - 8)) & 0xff);
            bits -= 8;
        }
    }

    return Buffer.from(buffer);
}

function hmacSha256(key, data) {
    return crypto.createHmac('sha256', key).update(data).digest();
}

export default generateTOTP;

// // Example usage:
// const secret = 'JBSWY3DPEHPK3PXP'; // Example Base32 encoded secret
// console.log("Generated TOTP:", generateTOTP(secret));
