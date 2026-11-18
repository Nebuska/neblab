import {BrowserRouter, Route, Routes, useNavigate} from "react-router-dom";
import LoginPage from "./pages/LoginPage.tsx";
import ErrorPage from "./pages/ErrorPage.tsx";
import HealthCheck from "./components/HealthCheck.tsx";
import IndexPage from "./pages/IndexPage.tsx";
import RegisterPage from "./pages/RegisterPage.tsx";
import {ThemeProvider} from "@mui/material/styles";
import theme from "./theme.tsx";
import TestPage from "./pages/TestPage.tsx";
import {useEffect} from "react";
import {NavigationService} from "./components/NavigationService.tsx";
import MainLayout from "./layouts/MainLayout.tsx";
import {AuthProvider} from "./contexts/AuthContext.tsx";

function App() {
  return (
    <>
        <BrowserRouter>
            <ThemeProvider theme={theme}>
                <NavigationBridge />
                <HealthCheck />
                <AuthProvider>
                <Routes>
                    <Route element={<MainLayout />}>

                        <Route
                            path="/error"
                            element={<ErrorPage />}
                        />
                        <Route
                            path="/"
                            element={<IndexPage />}
                        />
                        <Route
                            path="/login"
                            element={<LoginPage />}
                        />
                        <Route
                            path="/register"
                            element={<RegisterPage/>}
                        />
                        <Route
                            path="/test"
                            element={<TestPage/>}
                        />

                    </Route>
                </Routes>
                </AuthProvider>
            </ThemeProvider>
        </BrowserRouter>
    </>
  )
}

function NavigationBridge() {
    const navigate = useNavigate();

    useEffect(() => {
        NavigationService.setNavigate(navigate);
    }, [navigate]);

    return null;
}

export default App
