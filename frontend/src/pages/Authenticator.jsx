import React, {useState} from "react";
import {useQuery} from '@tanstack/react-query';
import SetupGoogleAuthenticator from "./SetupGoogle.jsx";
import {api} from "../services/api.js";


const Authenticator = ({source}) => {
    const [isGenNew, setIsGenNewCode] = useState(false);
    const [loadingPrevious, setLoadingPrevious] = useState(true); // Loader for fetching previous setup

    const fetch2faQuery = useQuery({
        queryKey: ["previous-2fa"],
        queryFn: () => api.get("/previous-2fa"),
    })
    const twoFactor = fetch2faQuery?.data?.data?.data
    console.log(twoFactor)

    const handleGenerateNew = () => {
        setIsGenNewCode(true);
    };

    if (fetch2faQuery.isFetching) {
        return <div>
            <h2>Loading previous 2FA setup</h2>
        </div>
    }

    return (
        <div className="p-4">
            <div className="">

                {
                    isGenNew ? (
                        <>
                            <SetupGoogleAuthenticator/>
                        </>
                    ) : (
                        <>
                            {!fetch2faQuery.isFetching && twoFactor ? (
                                <>
                                    <h2 className="text-2xl font-bold mb-4 text-gray-100">{source === "google" ? "Google " : "RsAuth "}
                                        Authenticator</h2>
                                    <p className="mb-4 text-gray-300">
                                        You already have a 2FA setup. Scan the QR code or use the secret below.
                                    </p>

                                    <img src={twoFactor?.qrCode} alt="QR Code"
                                         className="mb-4 w-40 h-40 mx-auto"/>
                                    <div className="mb-6 text-gray-300">
                                        <span className="font-bold">Secret Key: </span>
                                        <span className="font-mono">{twoFactor?.secret?.base32}</span>
                                    </div>

                                    <div className="flex items-center space-x-2">
                                        <button
                                            theme="orange"
                                            className="rounded-lg w-full mb-4"
                                            onClick={handleGenerateNew}
                                        >
                                            Unlink Authenticator
                                        </button>
                                        <button
                                            theme="orange"
                                            className="rounded-lg w-full mb-4"
                                            onClick={handleGenerateNew}
                                        >
                                            Generate New QR Code
                                        </button>
                                    </div>
                                </>
                            ) : !fetch2faQuery.isFetching && (
                                <SetupGoogleAuthenticator/>
                            )}
                        </>
                    )
                }

            </div>
        </div>
    );
};

export default Authenticator;