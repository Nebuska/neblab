import {
    Alert, Paper,
    Box,
    Button, Checkbox,
    Container, Divider,
    FormControlLabel,
    IconButton,
    InputAdornment,
    Stack,
    TextField,
    Typography,
    Link
} from "@mui/material";
import Navbar from "../components/Navbar.tsx";
import {GitHub, Google, Visibility, VisibilityOff} from "@mui/icons-material";
import {type ChangeEvent, useState} from "react";
import {NewRegisterFormData, ValidateRegisterData} from "../types/RegisterFormData.ts";

export default function RegisterPage(){
    const [showPassword, setShowPassword] = useState(false);
    const [loading, setLoading] = useState(false);

    const [formData, setFormData] = useState(NewRegisterFormData());

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

        const err = ValidateRegisterData(formData)
        if (err != "") {
            setError(err)
            return;
        }

        try {
            setLoading(true);

            // Fake API delay
            await new Promise((resolve) => setTimeout(resolve, 1500));

            console.log('LOGIN:', formData);

            // Navigate after login
            // navigate('/dashboard');

        } catch (err) {
            setError(err);
        } finally {
            setLoading(false);
        }
    };

    return (
        <>
            <Navbar />
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
                                Welcome
                            </Typography>

                            <Typography color="text.secondary">
                                Sign up to create your account
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

                                <TextField
                                    label="Password Confirmation"
                                    name="password_confirmation"
                                    type={showPassword ? 'text' : 'password'}
                                    fullWidth
                                    value={formData.password_confirmation}
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
                                                name="termAndCondition"
                                                checked={formData.termAndCondition}
                                                onChange={handleChange}
                                            />
                                        }
                                        label="Terms & Condition"
                                    />

                                    <Link
                                        href=""
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
                            Already Have an account?{' '}
                            <Link
                                href="/login"
                                underline="hover"
                                sx={{
                                    fontWeight: 600
                                }}
                            >
                                Sign in
                            </Link>
                        </Typography>
                    </Paper>
                </Container>
            </Box>
        </>
    )
}