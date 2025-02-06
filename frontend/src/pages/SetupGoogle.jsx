import React, {useEffect} from "react";
import {useNavigate, useSearchParams} from "react-router-dom";
import {useMutation} from "@tanstack/react-query";
import {api} from "../services/api.js";


const SetupGoogleAuthenticator = () => {


    const navigate = useNavigate()
    const generateSecret = useMutation({
        mutationFn: ({isNew}) => api.post("/api/v1/generate-2fa-secret", {isNew}),
    })

    const completeAuthSetup = useMutation({
        mutationFn: ({id, isCompleted, provider}) => api.post("/api/v1/generate-2fa-secret-complete", {
            provider,
            id,
            isCompleted
        }),
    })

    const data = generateSecret?.data?.data?.data

    async function handleCompleteSetup() {
        try {
            const response = await completeAuthSetup.mutateAsync({provider: "google", id: data.id, isCompleted: true})
            if (response?.status === 200) {
                navigate("/account/authenticator-apps")
            }

        } catch (ex) {
            console.log(ex, "hisdfjskd")
            // toast.error("Failed to complete setup")
        }
    }

    const [getParams] = useSearchParams()

    const provider = getParams.get("provider")

    console.log(provider, "providerproviderprovider")

    useEffect(() => {
        generateQR()
    }, [])

    function generateQR(isNew = false) {
        generateSecret.mutateAsync({isNew})
            .then(d => {
                console.log(d)
            }).catch(e => {
            console.log(e)
        })
    }


    function handleCancel() {

    }

    return (
        <div className="p-4 vh bg-gray-800 rounded-lg">
            <h3 className="text-xl font-bold text-gray-100">Google Authenticator Setup</h3>
            <p className="text-gray-300 my-4">
                Scan the QR code below with your Google Authenticator app or enter the secret key manually.
            </p>
            <img
                src={data?.qr_code_url}
                alt="QR Code"
                className="  mb-4"
            />
            <p className="font-mono text-gray-200">Code Name: {data?.code_name}</p>
            <p className="font-mono text-gray-200">Secret Key: {data?.secret}</p>
            <button
                onClick={() => generateQR(true)}
                className="bg-gray-600 hover:bg-gray-700 text-white px-4 mt-2 py-2 rounded-lg"
            >
                ReGenerate
            </button>

            <div className="mt-6">
                {Array.isArray(data?.recovery_codes) && (
                    <>
                        <div className="flex items-center justify-between">
                            <h4 className="text-lg font-semibold text-gray-200">Backup Codes</h4>
                            <button className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg">
                                Download
                            </button>
                        </div>
                        <ul className="mt-4 flex flex-wrap gap-2">
                            {data?.recovery_codes?.map((code, index) => (
                                <li
                                    key={index}
                                    className="bg-gray-700  p-2 rounded-lg font-mono text-gray-100"
                                >
                                    {code}
                                </li>
                            ))}
                        </ul>
                    </>
                )}
            </div>

            <div className="mt-6 flex space-x-4">

                <button onClick={handleCompleteSetup}
                        className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg">
                    Confirm Setup
                </button>
            </div>
        </div>
    );
};

export default SetupGoogleAuthenticator;
