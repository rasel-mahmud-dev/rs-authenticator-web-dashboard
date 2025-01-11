import React, {useEffect} from "react";

const SetupGoogleAuthenticator = ({data, onCancel, generateSecret}) => {

    useEffect(() => {
        generateSecret.mutateAsync().then(d => {
            console.log(d)
        }).catch(e => {
            console.log(e)
        })
    }, [])

    return (
        <div className="p-4 bg-gray-800 rounded-lg">
            <h3 className="text-xl font-bold text-gray-100">Google Authenticator Setup</h3>
            <p className="text-gray-300 my-4">
                Scan the QR code below with your Google Authenticator app or enter the secret key manually.
            </p>
            <img
                src={data?.qrCode} // Replace with actual QR code from API
                alt="QR Code"
                className="  mb-4"
            />
            <p className="font-mono text-gray-200">Secret Key: {data?.secret}</p>
            <div className="mt-6 flex space-x-4">
                <button
                    onClick={onCancel}
                    className="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-lg"
                >
                    Cancel
                </button>
                <button className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-lg">
                    Confirm Setup
                </button>
            </div>
        </div>
    );
};

export default SetupGoogleAuthenticator;
