import React, {useEffect} from "react";
import {useNavigate, useSearchParams} from "react-router-dom";
import {useMutation} from "@tanstack/react-query";
import {api} from "../services/api.js";
import {toast} from "react-toastify";

const SetupGoogleAuthenticator = () => {

    const navigate = useNavigate()
    const generateSecret = useMutation({
        mutationFn: ({isNew}) => api.post("/api/v1/mfa/generate-2fa-secret", {isNew}),
    })

    const completeAuthSetup = useMutation({
        mutationFn: ({id, isCompleted, provider}) => api.post("/api/v1/mfa/generate-2fa-secret-complete", {
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
                toast.success("Successfully connected.")
            }

        } catch (ex) {
            toast.error("Failed to complete setup")
        }
    }

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
            <div className="max-w-3xl mx-auto">


                <h3 className="text-gray-100 text-3xl font-bold text-center">Authenticator Setup</h3>
                <p className="text-gray-300 my-4">
                    Scan the QR code below with your Google Authenticator app or enter the secret key manually.
                </p>
                <img
                    src={data?.qr_code_url}
                    alt="QR Code"
                    className="  mb-4"
                />
                <p className="font-mono text-gray-200">Code Name: {data?.code_name}</p>
                <p className="font-mono text-gray-200">Secret Key: <span data-test_id="secret_key">{data?.secret}</span>
                </p>
                <button
                    onClick={() => generateQR(true)}
                    className="bg-gray-600 hover:bg-gray-700 text-white px-4 mt-2 py-2 rounded-lg"
                >
                    ReGenerate
                </button>



                <div className="mt-6 flex space-x-4">

                    <button data-test_id="confirm_setup_button" onClick={handleCancel} onClick={handleCompleteSetup}
                            className="btn mx-auto btn-primary text-white px-32 py-2 rounded-lg">
                        Confirm Setup
                    </button>
                </div>

            </div>
        </div>
    );
};

export default SetupGoogleAuthenticator;
