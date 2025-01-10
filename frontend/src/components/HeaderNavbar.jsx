import React from "react";
import useAuthStore from "../store/authState.js";


const HeaderNavbar = () => {
    const {user} = useAuthStore()

    console.log("user", user)

    return (
        <header className="bg-gradient-to-r from-blue-500 to-indigo-600 text-white shadow-lg">
            <div className="container mx-auto p-5 flex justify-between items-center">
                <h1 className="text-3xl font-bold">YourLogo</h1>
                <nav>
                    <ul className="flex space-x-6">
                        <li>
                            <a href="#" className="text-lg font-medium hover:text-indigo-300">
                                Home
                            </a>
                        </li>
                        <li>
                            <a href="#" className="text-lg font-medium hover:text-indigo-300">
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
