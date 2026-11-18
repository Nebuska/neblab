
import {SnackbarProvider} from "notistack";
import {Container} from "@mui/material";
import {Outlet} from "react-router-dom";
import Navbar from "../components/Navbar.tsx";

export default function MainLayout() {
    return (
        <>
            <SnackbarProvider/>
            <Container>
                <Navbar/>
                <Outlet/>
            </Container>
        </>
    )
}