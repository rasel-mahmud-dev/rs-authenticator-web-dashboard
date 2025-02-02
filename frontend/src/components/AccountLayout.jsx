import React, {useEffect, useState} from 'react';
import {Outlet, useNavigate} from 'react-router-dom';
import HeaderNavbar from "./HeaderNavbar.jsx";
import useAuthStore from "../store/authState.js";
import {verifyAuthentication} from "../services/authSerivce.js";
import Sidebar from "./Sidebar.jsx";

const AccountLayout = () => {
    const {user, authLoaded, setAuth} = useAuthStore()
    const [errorState, setErrorState] = useState({
        message: "",
        statusText: ""
    })

    const navigate = useNavigate()

    useEffect(() => {
        verifyAuthentication().then(data => {
            setAuth(data?.data ?? undefined)
        }).catch(ex => {
            setAuth(undefined)
        }).finally(() => {

        })
    }, []);

    useEffect(() => {
        if (authLoaded && !user?.id) {
            navigate("/login")
        }
    }, [authLoaded, user?.id]);

    return (
        <div className="">
            <HeaderNavbar/>
            {authLoaded && user?.id && (
                <div className="layout-content">
                    <Sidebar/>
                    <main className="pt-[75px]">
                        <Outlet/>
                    </main>
                </div>
            )}
        </div>
    );
};

export default AccountLayout;
