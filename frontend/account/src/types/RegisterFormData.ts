export interface RegisterFormData {
    email: string;
    verificationCode: string;
    password: string;
    password_confirmation: string;
    termAndCondition: boolean;
}

export function NewRegisterFormData(): RegisterFormData {
    return {
        email: "",
        verificationCode: "",
        password: "",
        password_confirmation: "",
        termAndCondition: false,
    }
}

export function ValidateRegisterData(data: RegisterFormData): string {

    if (!data.email || !data.password) {
        return "Please fill in all fields."
    }
    if (!data.password || data.password.length < 6) {
        return "Please enter at least 6 characters long."
    }
    if (data.password !== data.password_confirmation) {
        return "Passwords don't match."
    }
    if (!data.termAndCondition) {
        return "Terms & Conditions must be accepted."
    }

    return ""
}