import React from "react";
import {useQuery} from "@tanstack/react-query";
import {api} from "../services/api.js";
import {Link} from "react-router-dom"

const ConnectedAuthenticators = () => {
    const fetchAuthenticatorsQuery = useQuery({
        queryKey: ["connected-authenticators"],
        queryFn: () => api.get("/api/v1/mfa/authenticated-apps"),
    });

    const authenticators = fetchAuthenticatorsQuery?.data?.data?.data || [];


    const pageTitle = (
        <div className="flex items-center justify-between">
            <h2 className="mb-4 text-gray-100 text-3xl font-bold text-center ">Connected Authenticator Apps</h2>
            <Link to="/account/authenticator-setup" className="btn btn-primary">Add Provider</Link>
        </div>
    )

    if (fetchAuthenticatorsQuery.isLoading) {
        return (
            <div className="mt-10 vh max-w-screen-xl mx-auto px-4">
                {pageTitle}
                <h2 className="text-lg font-semibold text-gray-100">Loading connected authenticators...</h2>
            </div>
        );
    }

    if (fetchAuthenticatorsQuery.isError) {
        return (
            <div className="mt-10 vh max-w-screen-xl mx-auto px-4">
                {pageTitle}
                <h2 className="text-lg font-semibold text-red-500">
                    Failed to load connected authenticators. Please try again later.
                </h2>
            </div>
        );
    }

    return (
        <div className="mt-10 vh max-w-screen-xl mx-auto px-4">

            {pageTitle}

            <div className="mt-4">

                <ul className="space-y-4">
                    {authenticators?.map((authenticator, index) => (
                        <li
                            key={index}
                            className="flex justify-between items-center p-4 bg-gray-800 rounded-lg"
                        >
                            <div>
                                <h3 className="text-lg font-semibold text-gray-100">
                                    {authenticator.code_name}
                                </h3>

                                <p className="text-sm text-gray-400">
                                    Linked on: {new Date(authenticator.created_at).toLocaleString()}
                                </p>
                            </div>
                            <button
                                className="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-lg"
                                onClick={() => handleUnlinkAuthenticator(authenticator.id)}
                            >
                                Unlink
                            </button>
                        </li>
                    ))}
                </ul>

            </div>

        </div>
    );
};

const handleUnlinkAuthenticator = async (authenticatorId) => {
    try {
        await api.post(`/api/v1/mfa/unlink-authenticator/${authenticatorId}`);
        alert("Authenticator unlinked successfully!");
        window.location.reload();
    } catch (error) {
        console.error("Failed to unlink authenticator:", error);
        alert("Failed to unlink the authenticator. Please try again.");
    }
};

export default ConnectedAuthenticators;
