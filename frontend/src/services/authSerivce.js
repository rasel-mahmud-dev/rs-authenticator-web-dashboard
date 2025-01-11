import {api} from "./api.js";

export async function login(credentials) {
    const response = await api.post("/api/v1/login", credentials);
    if (response.status === 200) {
        console.log("Login successful:", response.data);
        return response.data;
    } else {
        throw Error("Login failed: Unexpected status code");
    }
}

export async function verifyAuthentication() {
    const response = await api.get("/api/v1/verify");
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("")
    }
}
