export interface Credentials {
    token: string | null;
    isRefreshing: boolean;
    refreshQueue: {
        resolve: (token: string) => void;
        reject: (error: unknown) => void;
    }[];
}