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


    return (
        <header className="fixed top-0 z-50 left-0 w-full bg-blue-950  text-white shadow-lg">
            <div className="container mx-auto p-5 flex justify-between items-center">
                <h1 className="text-3xl font-bold">YourLogo</h1>
                <nav>
                    <ul className="flex space-x-6">
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
                    </ul>
                </nav>
                <div>
                    {user ? (
                        <div className="relative">
                            <button
                                onClick={toggleDropdown}
                                className="btn btn-outline btn-light text-white hover:bg-indigo-400"
                            >
                                {user?.username}
                            </button>

                            {isDropdownOpen && (
                                <ul className="bg-primary dropdown-content menu shadow bg-base-100 text-black absolute right-0 mt-2 p-2 w-48 rounded-lg">
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
                            )}

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
        </header>
    );
};

export default HeaderNavbar;
