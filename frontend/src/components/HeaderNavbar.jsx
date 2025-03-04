import React from "react";
import useAuthStore from "../store/authState.js";
import {Link, NavLink, useNavigate} from "react-router-dom";
import {FaGithub} from "react-icons/fa";


const HeaderNavbar = () => {
    const {user, setAuth} = useAuthStore()

    const logout = () => {
        setAuth(null)
        localStorage.removeItem("token")
        sessionStorage.removeItem("token")
    };

    const links = (
        <>

            <li>
                <NavLink to="/" className="text-lg font-medium hover:text-indigo-300">
                    Home
                </NavLink>

            </li>
            <li>
                <NavLink to="/account" className="text-lg font-medium hover:text-indigo-300">
                    Dashboard
                </NavLink>
            </li>
            <li>
                <NavLink to="/download-app" className="text-lg font-medium hover:text-indigo-300">
                    Download App
                </NavLink>
            </li>
            <li>
                <NavLink to="/about-me" className="text-lg font-medium hover:text-indigo-300">
                    Me
                </NavLink>
            </li>
        </>
    )

    return (
        <div className="header  bg-blue-950 w-full ">
            <div className="navbar container mx-auto text-white shadow-lg">
                <div className="navbar-start">
                    <div className="dropdown">
                        <div tabIndex={0} role="button" className="btn btn-ghost lg:hidden">
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                className="h-5 w-5"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor">
                                <path
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth="2"
                                    d="M4 6h16M4 12h8m-8 6h16"/>
                            </svg>
                        </div>
                        <ul
                            tabIndex={0}
                            className="menu menu-sm dropdown-content bg-gray-800 rounded-box z-[1] mt-3 w-52 p-2 shadow">
                            {links}
                        </ul>
                    </div>
                    <Link className="px-0 btn btn-ghost text-2xl " style={{fontFamily: "'Pacifico', cursive"}}
                          data-tour="step-1" to="/">
                        Rs Authenticator
                    </Link>
                </div>
                <div className="navbar-center hidden lg:flex">
                    <ul className="menu menu-horizontal px-1">
                        {links}
                    </ul>
                </div>
                <div className="navbar-end relative">

                    <div className="flex flex-1 items-center justify-end px-2">

                        <a target="_blank" className="btn btn-ghost"
                           href="https://github.com/rasel-mahmud-dev/rs-authenticator-web-dashboard">
                            <FaGithub className="text-2xl"/>
                        </a>

                        {user ? (
                            <div className="flex items-stretch">

                                <div className="dropdown  dropdown-end">
                                    <div data-test_id="auth-menu" tabIndex={0} role="button"
                                         className="btn btn-ghost rounded-btn">

                                        <div className="avatar">
                                            <div
                                                className="w-8 rounded-full ">
                                                <img src={user?.avatar || "/boy.png"}/>
                                            </div>
                                        </div>

                                        <span className="hidden md:block">{user?.username}</span>
                                    </div>

                                    <ul
                                        tabIndex={0}
                                        className="menu dropdown-content bg-primary-100 rounded-box relative z-[100] mt-4 w-52 p-2 shadow">

                                        <li>
                                            <NavLink to="/account/profile" className="hover:bg-indigo-200">
                                                Profile
                                            </NavLink>
                                        </li>
                                        <li>
                                            <NavLink to="/account/settings" className="hover:bg-indigo-200">
                                                Settings
                                            </NavLink>
                                        </li>
                                        <li>
                                            <button data-test_id="logout_button" onClick={logout}
                                                    className="hover:bg-indigo-200 w-full text-left">
                                                Logout
                                            </button>
                                        </li>


                                    </ul>
                                </div>
                            </div>
                        ) : (
                            <div>
                                <Link to="/login">
                                    <button className="btn btn-outline btn-light text-white hover:bg-indigo-400">
                                        Sign In
                                    </button>
                                </Link>

                            </div>
                        )}
                    </div>


                </div>
            </div>
        </div>

    );
};

export default HeaderNavbar;
