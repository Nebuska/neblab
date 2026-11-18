import {createContext, useContext, useEffect, useState} from "react";
import type {Credentials} from "../types/Credentials.ts";
import * as React from "react";

const AuthContext = createContext([null,null])

export function AuthProvider({ children }: { children: React.ReactNode }) {
    const [user, setUser] = useState<Credentials | null>(null)

    useEffect(() => {
        //todo: Get User information on load
    })

    function signIn(userData: Credentials) {
        setUser(userData)
    }

    function signOut() {
        setUser(null)
    }

    return (
        <AuthContext.Provider
            value={[
                user,
                signIn,
                signOut
            ]}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if (!context) {
        throw new Error("useAuth must be used within a AuthProvider")
    }

    return useContext(AuthContext)
}

interface AuthContextType {
    user: Credentials | null;
    signIn: (token: string) => void;
}