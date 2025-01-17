import React from "react";
import useAuthStore from "../store/authState.js";
import {NavLink} from "react-router-dom";


const HeaderNavbar = () => {
    const {user} = useAuthStore()

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
                            <a href="/login" className="text-lg font-medium hover:text-indigo-300">
                                About
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
                        <div>
                            {user?.username}
                        </div>
                    ) : (
                        <div>
                            <button className="btn btn-outline btn-light text-white hover:bg-indigo-400">
                                Sign In
                            </button>

                        </div>
                    )}
                </div>
            </div>
        </header>
    );
};

export default HeaderNavbar;
