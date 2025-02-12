import axios from "axios";


export const api = axios.create({
    // baseURL: "http://192.168.0.148:8080",
    // baseURL: "https://go-netlify-hujj4oazq-raselmahmuddevs-projects.vercel.app",
    baseURL: import.meta.env.DEV ? "http://localhost:8080" : "https://go-netlify-app.vercel.app",
    withCredentials: true
})

api.interceptors.request.use(
    (config) => {
        let token = ""
        const rememberMe = localStorage.getItem("auth-remember-me");
        if (rememberMe == "false") {
            token = sessionStorage.getItem("token")
        } else {
            token = localStorage.getItem("token")
        }
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);