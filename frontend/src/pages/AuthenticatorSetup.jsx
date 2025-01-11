import React, {useState} from "react";
import SetupGoogleAuthenticator from "./SetupGoogle.jsx";
// import SetupMicrosoftAuthenticator from "./SetupMicrosoft.jsx";
import {FiKey, FiShield} from "react-icons/fi";
import {useMutation} from "@tanstack/react-query";
import {api} from "../services/api.js"; // Icons for UI

const AuthenticatorSetup = () => {
    const [selectedAuthenticator, setSelectedAuthenticator] = useState(null);

    const authenticators = [
        {id: "google", name: "Google Authenticator", icon: <FiShield className="text-blue-500 text-4xl"/>},
        {id: "microsoft", name: "Microsoft Authenticator", icon: <FiKey className="text-green-500 text-4xl"/>},
    ];

    const generateSecret = useMutation({
        mutationFn: () => api.get("/api/v1/generate-2fa-secret"),
    })

    const completeAuthSetup = useMutation({
        mutationFn: (id) => api.post("/api/v1/generate-2fa-secret", {id}),
    })

    const handleSetup = (authenticatorId) => {
        setSelectedAuthenticator(authenticatorId);
    };

    const handleResetSelection = () => {
        setSelectedAuthenticator(null);
    };

    const data = generateSecret?.data?.data?.data

    async function handleCompleteSetup(data) {
       try{
           const response = await completeAuthSetup.mutateAsync(data.id)
           console.log(response)

       } catch (ex){
           console.log(ex, "hisdfjskd")
       }
    }

    return (
        <div className="p-6 bg-gray-900 text-white full-viewport">
            <h2 className="text-3xl font-bold mb-6 ">
                Setup Multi-Factor Authentication
            </h2>

            {!selectedAuthenticator ? (
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
                                <button
                                    onClick={() => handleSetup(authenticator.id)}
                                    className="mt-4 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-lg w-full text-center"
                                >
                                    Setup
                                </button>
                            </div>
                        ))}
                    </div>
                </div>
            ) : (
                <div>
                    {selectedAuthenticator === "google" && (
                        <SetupGoogleAuthenticator
                            onCompleteSetup={handleCompleteSetup}
                            data={data}
                            generateSecret={generateSecret}
                            onCancel={handleResetSelection}
                        />
                    )}
                    {/* Uncomment when Microsoft Authenticator is ready */}
                    {/* {selectedAuthenticator === "microsoft" && (
                        <SetupMicrosoftAuthenticator onCancel={handleResetSelection} />
                    )} */}
                </div>
            )}
        </div>
    );
};

export default AuthenticatorSetup;
