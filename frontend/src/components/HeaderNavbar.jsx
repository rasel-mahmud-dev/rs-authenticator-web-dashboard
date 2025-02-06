import React, {useState} from "react";
import useAuthStore from "../store/authState.js";
import {Link, NavLink} from "react-router-dom";


const HeaderNavbar = () => {
    const {user, setAuth} = useAuthStore()

    const [isDropdownOpen, setIsDropdownOpen] = useState(false);

    const toggleDropdown = () => {
        setIsDropdownOpen(!isDropdownOpen);
    };

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
             <a href="/account" className="text-lg font-medium hover:text-indigo-300">
             Dashboard
             </a>
             </li>
             <li>
             <a href="#" className="text-lg font-medium hover:text-indigo-300">
             Services
             </a>
             </li>
             <li>
             <a href="#" className="text-lg font-medium hover:text-indigo-300">
             Contact
             </a>
             </li>
        </>
    )

    return (
        <div className="navbar header bg-blue-950  text-white shadow-lg">
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
                        className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow">
                        {links}
                    </ul>
                </div>
                <a className="btn btn-ghost text-xl">daisyUI</a>
            </div>
            <div className="navbar-center hidden lg:flex">
                <ul className="menu menu-horizontal px-1">
                    {links}
                </ul>
            </div>
            <div className="navbar-end relative">
                {/*<a className="btn">Button</a>*/}

                { user ? (
                    <div className="flex flex-1 justify-end px-2">
                        <div className="flex items-stretch">

                            <div className="dropdown  dropdown-end">
                                <div tabIndex={0} role="button" className="btn btn-ghost rounded-btn">{user?.username}</div>
                                <ul
                                    tabIndex={0}
                                    className="menu dropdown-content bg-primary-100 rounded-box relative z-[100] mt-4 w-52 p-2 shadow">

                                    <li>
                                        <NavLink to="/profile" className="hover:bg-indigo-200">
                                            Profile
                                        </NavLink>
                                    </li>
                                    <li>
                                        <NavLink to="/settings" className="hover:bg-indigo-200">
                                            Settings
                                        </NavLink>
                                    </li>
                                    <li>
                                        <button onClick={logout} className="hover:bg-indigo-200 w-full text-left">
                                            Logout
                                        </button>
                                    </li>


                                </ul>
                            </div>
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

    );
};

export default HeaderNavbar;
