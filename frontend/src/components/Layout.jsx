import React, {useEffect, useState} from 'react';
import {Outlet} from 'react-router-dom';
import HeaderNavbar from "./HeaderNavbar.jsx";
import useAuthStore from "../store/authState.js";
import {verifyAuthentication} from "../services/authSerivce.js";

const Layout = () => {
    const {user, authLoaded, setAuth} = useAuthStore()
    const [errorState, setErrorState] = useState({
        message: "",
        statusText: ""
    })

    useEffect(() => {
        verifyAuthentication().then(data => {
            setAuth(data?.data ?? undefined)
        }).catch(ex => {
            setAuth(undefined)
        }).finally(() => {

        })
    }, []);

    // if (!authLoaded) {
    //     return <div>Auth loading...</div>;
    // }

    return (
        <div className="">
            <HeaderNavbar/>
            <main>
                <Outlet/>
            </main>
        </div>
    );
};

export default Layout;
