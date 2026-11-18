import {api} from "./client.ts";
import type {LoginFormData} from "../types/LoginFormData.ts";

export async function login(data: LoginFormData) {
    const response = await api.post("/login", {
        username: data.email,
        password: data.password
    })
    return response.data;
}