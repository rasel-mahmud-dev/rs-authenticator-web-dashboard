import {api} from "./api.js";

export async function registrationSlats() {
    const response = await api.get("/api/v1/slats/registration");
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("registration status code");
    }
}

export async function authenticatorSlats() {
    const response = await api.get("/api/v1/slats/authentication");
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("Authenticator slats: Unexpected status code");
    }
}

export async function fetchTrafficStats(isCount = false) {
    const response = await api.get(`/api/v1/slats/traffic?t=${isCount ? "" : "detail"}`);
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("traffic slats: Unexpected status code");
    }
}
export async function fetchApiLatencySlats() {
    const response = await api.get(`/api/v1/slats/api-latency`);
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("traffic slats: Unexpected status code");
    }
}

export async function fetchUsers(currentPage, itemsPerPage) {
    const response = await api.get("/api/v1/slats/users?page=" + currentPage + "&limit=" + itemsPerPage);
    if (response.status === 200) {
        return response.data;
    } else {
        throw Error("traffic slats: Unexpected status code");
    }
}