import type { NavigateFunction } from "react-router-dom";

let navigate: NavigateFunction | null = null;

export const NavigationService = {
    setNavigate(fn: NavigateFunction) {
        navigate = fn;
    },

    go(path: string) {
        navigate?.(path);
    },
};