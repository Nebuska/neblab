import {type ChangeEvent, useState} from "react";
import * as React from "react";
import {
    Alert,
    Box, Button, Checkbox,
    Container, Divider,
    FormControlLabel,
    IconButton,
    InputAdornment,
    Link,
    Paper,
    Stack,
    TextField,
    Typography
} from "@mui/material";
import {GitHub, Google, Visibility, VisibilityOff} from "@mui/icons-material";
import {login} from "../api/login.ts";
import {NewLoginFormData, ValidateLoginData} from "../types/LoginFormData.ts";
import {NavigationService} from "../components/NavigationService.tsx";

export default function LoginPage() {
    const [showPassword, setShowPassword] = useState(false);
    const [loading, setLoading] = useState(false);

    const [formData, setFormData] = useState(NewLoginFormData());

    const [error, setError] = useState('');

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value, checked, type } = e.target;

        setFormData((prev) => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : value,
        }));
    };

    const handleSubmit : React.SubmitEventHandler<HTMLFormElement> = async (e) => {
        e.preventDefault();

        setError('');

        const err = ValidateLoginData(formData)
        if (err != "") {
            setError(err)
            return;
        }

        try {
            setLoading(true);

            const response = await login(formData)

            console.log('LOGIN:', response.data);

            NavigationService.go("/")

        } catch (err) {
            setError(err as string);
        } finally {
            setLoading(false);
        }
    };
    return (
        <>
            <Box
                sx={{
                    minHeight: '100vh',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    background:
                        'linear-gradient(135deg, #f5f7fa 0%, #e4ecfb 100%)',
                    px: 2,
                }}
            >
                <Container maxWidth="sm">
                    <Paper
                        elevation={8}
                        sx={{
                            p: { xs: 3, sm: 5 },
                            borderRadius: 4,
                            backdropFilter: 'blur(10px)',
                        }}
                    >
                        {/* Header */}
                        <Stack spacing={1} sx={{ mb: 4 }}>
                            <Typography
                                variant="h4"
                                sx={{
                                    fontWeight: 700
                                }}
                            >
                                Welcome back
                            </Typography>

                            <Typography color="text.secondary">
                                Sign in to continue to your account
                            </Typography>
                        </Stack>

                        {/* Error */}
                        {error && (
                            <Alert severity="error" sx={{ mb: 3 }}>
                                {error}
                            </Alert>
                        )}

                        {/* Form */}
                        <Box
                            component="form"
                            onSubmit={handleSubmit}
                        >
                            <Stack spacing={2.5}>
                                <TextField
                                    label="Email Address"
                                    name="email"
                                    type="email"
                                    fullWidth
                                    autoComplete="email"
                                    value={formData.email}
                                    onChange={handleChange}
                                    autoFocus
                                />

                                <TextField
                                    label="Password"
                                    name="password"
                                    type={showPassword ? 'text' : 'password'}
                                    fullWidth
                                    autoComplete="current-password"
                                    value={formData.password}
                                    onChange={handleChange}
                                    slotProps={{
                                        input: {
                                            endAdornment: (
                                                <InputAdornment position="end">
                                                    <IconButton
                                                        onClick={() =>
                                                            setShowPassword((prev) => !prev)
                                                        }
                                                        edge="end"
                                                    >
                                                        {showPassword ? (
                                                            <VisibilityOff/>
                                                        ) : (
                                                            <Visibility/>
                                                        )}
                                                    </IconButton>
                                                </InputAdornment>
                                            ),
                                        }
                                    }}
                                />

                                {/* Remember + Forgot */}
                                <Box
                                    sx={{
                                        display: 'flex',
                                        justifyContent: 'space-between',
                                        alignItems: 'center',
                                    }}
                                >
                                    <FormControlLabel
                                        control={
                                            <Checkbox
                                                name="remember"
                                                checked={formData.remember}
                                                onChange={handleChange}
                                            />
                                        }
                                        label="Terms & Condition"
                                    />

                                    <Link
                                        href="/forgot-password"
                                        underline="hover"
                                    >
                                        Forgot password?
                                    </Link>
                                </Box>

                                {/* Submit */}
                                <Button
                                    type="submit"
                                    variant="contained"
                                    size="large"
                                    fullWidth
                                    disabled={loading}
                                    sx={{
                                        py: 1.4,
                                        borderRadius: 2,
                                        textTransform: 'none',
                                        fontSize: '1rem',
                                        fontWeight: 600,
                                    }}
                                >
                                    {loading ? 'Signing in...' : 'Sign In'}
                                </Button>
                            </Stack>
                        </Box>

                        {/* Divider */}
                        <Divider sx={{ my: 4 }}>
                            OR CONTINUE WITH
                        </Divider>

                        {/* Social Buttons */}
                        <Stack
                            direction={{ xs: 'column', sm: 'row' }}
                            spacing={2}
                        >
                            <Button
                                fullWidth
                                variant="outlined"
                                startIcon={<Google />}
                                sx={{
                                    py: 1.2,
                                    borderRadius: 2,
                                    textTransform: 'none',
                                }}
                            >
                                Google
                            </Button>

                            <Button
                                fullWidth
                                variant="outlined"
                                startIcon={<GitHub />}
                                sx={{
                                    py: 1.2,
                                    borderRadius: 2,
                                    textTransform: 'none',
                                }}
                            >
                                GitHub
                            </Button>
                        </Stack>

                        {/* Footer */}
                        <Typography
                            variant="body2"
                            color="text.secondary"
                            align="center"
                            sx={{ mt: 4 }}
                        >
                            Don&apos;t have an account?{' '}
                            <Link
                                href="/register"
                                underline="hover"
                                sx={{
                                    fontWeight: 600
                                }}
                            >
                                Sign up
                            </Link>
                        </Typography>
                    </Paper>
                </Container>
            </Box>
        </>
    )
}