import axios from "axios";


export const api = axios.create({
    baseURL: "http://192.168.0.148:8080",
    withCredentials: true
})