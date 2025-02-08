import React, {useState} from "react";
import {FiKey, FiShield} from "react-icons/fi";
import {Link} from "react-router-dom";


const AuthenticatorSetup = () => {

    const authenticators = [
        {id: "google", name: "Google Authenticator", icon: <FiShield className="text-blue-500 text-4xl"/>},
        {id: "microsoft", name: "Microsoft Authenticator", icon: <FiKey className="text-green-500 text-4xl"/>},
    ];

    return (
        <div className="p-6 bg-gray-900 text-white full-viewport">
            <h2 className="text-3xl font-bold mb-6 ">
                Setup Multi-Factor Authentication
            </h2>

                <div>
                    <p className="mb-4 text-gray-300  ">
                        Choose an authenticator app to set up for enhanced account security:
                    </p>
                    <div className="grid grid-cols-1 sm:grid-cols-2 gap-6 max-w-lg  ">
                        {authenticators.map((authenticator) => (
                            <div
                                key={authenticator.id}
                                className="bg-gray-800 hover:bg-gray-700 p-6 rounded-lg flex flex-col items-center shadow-md transition transform hover:scale-105"
                            >
                                {authenticator.icon}
                                <h3 className="mt-4 text-lg font-semibold text-gray-100">
                                    {authenticator.name}
                                </h3>
                                <Link to={`/account/authenticator-setup/${authenticator.id}`}
                                      className="mt-4 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg w-full text-center"
                                >
                                    Setup
                                </Link>
                            </div>
                        ))}
                    </div>
                </div>

        </div>
    );
};


const name = "asdf"
export default AuthenticatorSetup;
