import {useEffect, useState} from "react";
import {api} from "../api/client.ts";
import {useNavigate} from "react-router-dom";

export default function HealthCheck() {
    const [status, setStatus] = useState("loading")
    const navigate = useNavigate()

    useEffect(() => {
        api.get("/health")
            .then((res) => {
                setStatus(res.data.status)
                console.log(res.data)
            })
            .catch(() => setStatus("error"))
    }, []);

    useEffect(() => {
        if (status === "error") {
            navigate("/error")
        }
    }, [status, navigate])

    return null
}