let jwtToken: string | null = null;

export function SetAccessToken(accessToken: string) {
    jwtToken = accessToken;
}

export function GetAccessToken() {
    return jwtToken;
}