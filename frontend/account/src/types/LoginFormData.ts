export interface LoginFormData {
    email: string;
    password: string;
    remember: boolean;
}

export function NewLoginFormData(): LoginFormData {
    return {
        email: "",
        password: "",
        remember: false,
    }
}

export function ValidateLoginData(data: LoginFormData): string {

    if (!data.email || !data.password) {
        return "Please fill in all fields."
    }
    return ""
}
// Regexes:
//  .{8,} --- at least 8 character
//  (?=.*[A-Z])(?=.*[a-z]) -- at least 1 uppercase and 1 lowercase