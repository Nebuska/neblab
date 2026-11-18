import axios from "axios";
import {NavigationService} from "../components/NavigationService.tsx";
import type {Credentials} from "../types/Credentials.ts";

export const api = axios.create({
    baseURL: "http://localhost:8080",
    withCredentials: true,
    timeout: 5000,
    maxBodyLength: 1000,
    maxContentLength: 1000,
    allowAbsoluteUrls: false,
    responseType: "json",
})

export const access: Credentials = {
    token: null,
    isRefreshing: false,
    refreshQueue: [],
}

export default function SetToken(token: string | null) {
    access.token = token;
}

api.interceptors.request.use((config) => {
    const token = access.token;

    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
})

api.interceptors.response.use(
    (res) => res,
    async (err) => {
        const originalRequest = err.config;

        if (err.response?.status !== 401 || originalRequest._retry) {
            return Promise.reject(err);
        }

        if (access.isRefreshing) {
            return new Promise((resolve, reject) => {
                access.refreshQueue.push({resolve,reject});
            }).then((token) => {
                originalRequest.headers.Authorization = `Bearer ${token}`;
                return api(originalRequest);
            });
        }

        originalRequest._retry = true;
        access.isRefreshing = true;

        try {
            const res = await api.get(
                "/jwt",
            )
            SetToken(res.data.token);
            processQueue(access.token || "",null);
            originalRequest.headers.Authorization = `Bearer ${access.token}`;
            return api(originalRequest);
        } catch (refreshError) {
            console.log(refreshError);
            processQueue("",refreshError);
            //todo: redirect to login
            NavigationService.go("/login");
            return Promise.reject(err);
        } finally {
            access.isRefreshing = false;
        }

    }
)

const processQueue = (token: string, error: unknown) => {
    access.refreshQueue.forEach((item) => {
        if (error) {
            item.reject(error);
        } else {
            item.resolve(token);
        }
    });
    access.refreshQueue = [];
}