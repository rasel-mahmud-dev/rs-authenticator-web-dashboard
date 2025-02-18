import {api} from "./api.js";

export async function login(credentials) {
    const response = await api.post("/api/v1/auth/login", credentials);
    if (response.status === 200) {
        console.log("Login successful:", response.data);
        return response.data;
    } else {
        throw Error("Login failed: Unexpected status code");
    }
}

export async function register(credentials) {
    const response = await api.post("/api/v1/auth/registration", credentials);
    return response.data;
}

export async function loginWithAuthenticator({otpCode}) {
    const response = await api.post("/api/v1/auth/login-with-authenticator", {otpCode});
    if (response.status === 200) {
        console.log("Login successful:", response.data);
        return response.data;
    } else {
        throw Error("Login failed: Unexpected status code");
    }
}

export async function verifyAuthentication() {
    const response = await api.get("/api/v1/auth/verify");
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("")
    }
}



export async function accountRecoveryWithBackupCode({code}) {
    const response = await api.post("/api/v1/mfa/account-recovery", {code});
    if (response.status === 200) {
        console.log("account-recovery successful:", response.data);
        return response.data;
    } else {
        throw Error("account-recovery failed: Unexpected status code");
    }
}
