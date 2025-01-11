import React, {useEffect} from "react";

const SetupGoogleAuthenticator = ({data, onCompleteSetup, onCancel, generateSecret}) => {

    useEffect(() => {
        generateQR()
    }, [])

    function generateQR() {
        generateSecret.mutateAsync()
            .then(d => {
                console.log(d)
            }).catch(e => {
            console.log(e)
        })
    }

    function handleCompleteSetup() {
        onCompleteSetup(data)
    }

    return (
        <div className="p-4 bg-gray-800 rounded-lg">
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
                onClick={generateQR}
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
                <button
                    onClick={onCancel}
                    className="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-lg"
                >
                Cancel
                </button>
                <button onClick={handleCompleteSetup} className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg">
                    Confirm Setup
                </button>
            </div>
        </div>
    );
};

export default SetupGoogleAuthenticator;
