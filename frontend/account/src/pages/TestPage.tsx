import SetToken, {api} from "../api/client.ts";
import {Button} from "@mui/material";
import Navbar from "../components/Navbar.tsx";
import {enqueueSnackbar} from "notistack";

export default function TestPage() {
    async function load() {
        const response = await api.get("/test");
        console.log(response.data);
        if (response.status === 200) {
            enqueueSnackbar("Test Response:" + JSON.stringify(response.data), { variant: 'success' });
        } else {
            enqueueSnackbar("Test Failed", { variant: 'error' });
        }
    }
    async function testjwt() {
        const response = await api.get("/jwt");
        SetToken(response.data.token);
        console.log(response.data);
        if (response.status === 200) {
            enqueueSnackbar("Jwt Response:" + JSON.stringify(response.data), { variant: 'success' });
        } else {
            enqueueSnackbar("Jwt Failed", { variant: 'error' });
        }
    }
    async function removeToken() {
        SetToken(null);
        enqueueSnackbar("Token Removed.", { variant: 'info'});
    }




    return (
        <>
            <Button variant="contained" color="primary" onClick={load}>
                Test /test
            </Button>
            <Button variant="contained" color="primary" onClick={testjwt}>
                Test /jwt
            </Button>
            <Button variant="contained" color="primary" onClick={removeToken}>
                Remove JWT token
            </Button>
        </>
    )
}