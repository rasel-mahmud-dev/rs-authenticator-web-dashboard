import { api } from "./api.js";

export async function login(credentials) {
    try {
        const response = await api.post("/api/v1/login", credentials);
        if (response.status === 200) {
            console.log("Login successful:", response.data);
            return response.data; // Return the data (e.g., token or user info)
        } else {
            console.error("Login failed: Unexpected status code", response.status);
            return null;
        }
    } catch (error) {
        if (error.response) {
            console.error("Error response from server:", error.response.data);
            return null;
        } else if (error.request) {
            console.error("No response received from server:", error.request);
        } else {
            console.error("Error in setting up request:", error.message);
        }
        return null;
    }
}
