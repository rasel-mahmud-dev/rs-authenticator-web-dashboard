import React from "react";
import {Link} from 'react-router-dom'

const rsAuthenticatorFeatures = [
    {
        title: "App Lock",
        description: "Enhance your app's security with a built-in lock, preventing unauthorized access to your sensitive data. Keep your information safe, even if your device is lost or stolen."
    },
    {
        title: "Offline Operation",
        description: "Stay productive with offline support. Use the app’s full functionality even without an internet connection, ensuring you’re always in control."
    },
    {
        title: "Offline Data Backup",
        description: "Back up your data offline without the need for an internet connection. Restore your information securely at any time, giving you peace of mind."
    },
    {
        title: "Secure & No Tracking",
        description: "Your privacy is our priority. We ensure your data remains secure, stored only on your device, with no tracking or sharing. Enjoy complete control over your information."
    }
];


const AndroidAppHomeSection = () => {
    return (
        <section className="py-10 ">
            <div className="max-w-screen-xl mx-auto p-4 md:p-6">

                <h2 className="text-3xl font-bold text-center text-primary mt-20 mb-8">
                    Android App
                </h2>

                <div className="max-w-4xl mx-auto">
                    <div>
                        <img src="/ic_launcher.png" alt="app-logo" className="mx-auto w-16 h-16 mb-4"/>
                    </div>

                    <p className="mb-6 text-gray-400">
                        <strong>RsAuthenticator</strong> is a secure authenticator app developed using <strong>Kotlin
                        Jetpack Compose</strong>.
                        It ensures <strong>100% secure authentication</strong> without requiring internet access.
                    </p>
                </div>

                <h3 className="text-2xl font-semibold text-center text-gray-300 mb-4 ">
                    Key Features
                </h3>

                <ul className="grid grid-cols-4 gap-x-4 text-gray-400 space-y-2">
                    {rsAuthenticatorFeatures.map((feature, index) => (
                        <div key={feature.title}
                             className="bg-gray-800 text-white rounded-lg shadow-lg p-6 hover:scale-105 transform transition-all">
                            <h4 className="text-xl font-medium text-primary mb-4">{feature.title}</h4>
                            <p className="text-gray-400">{feature.description}</p>
                        </div>
                    ))}
                </ul>

                <div className="max-w-4xl mx-auto">
                    <p className="mt-6 text-gray-400">
                        With <strong>RsAuthenticator</strong>, you gain complete control over your
                        authentication
                        process while prioritizing
                        security and convenience.
                    </p>
                    <div className="flex justify-center mt-10">
                        <Link to="/download-app">
                            <button className="btn btn-primary  ">Download</button>
                        </Link>
                    </div>
                </div>
            </div>
        </section>
    );
}

export default AndroidAppHomeSection