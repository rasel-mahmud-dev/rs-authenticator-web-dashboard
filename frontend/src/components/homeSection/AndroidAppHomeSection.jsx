import React from "react";
import appCovers from "../../assets/appCovers.json"
import {Link} from 'react-router-dom'
import AppPreviewGallery from "./AppPreviewGallery.jsx";

const AndroidAppHomeSection = () => {
    const frontendFeatures = [
        "Registration Trends",
        "Login Methods Breakdown (Password vs. Authenticator App)",
        "Account Recovery codes",
        "Authenticator Success and Failure Counts",
        "API Traffic Analytics",
        "API Latency Monitoring",
    ]

    const rsAuthenticatorFeatures = [
        "Offline Backup & Restore – Easily back up and restore your connected apps using \n JSON or CSV formats, without needing an internet connection.",
        "Data Security – Your authentication data remains safe and private.",
        "Simplicity – A seamless user experience with an intuitive interface.",
        "Elegant Design – A modern and visually appealing UI.",
    ]

    return (
        <section className="py-10 ">
            <div className="max-w-screen-xl mx-auto p-4 md:p-6">

                <h2 className="text-3xl font-bold text-center text-primary mt-20 mb-8">
                    Android App
                </h2>

                <div>

                    <div className="max-w-2xl mx-auto">


                        <div>
                            <img src="/ic_launcher.png" alt="app-logo" className="mx-auto w-16 h-16 mb-4"/>
                        </div>


                        <p className="mb-6 text-gray-400">
                            <strong>RsAuthenticator</strong> is a secure authenticator app developed using <strong>Kotlin
                            Jetpack Compose</strong>.
                            It ensures <strong>100% secure authentication</strong> without requiring internet access.
                        </p>
                    </div>

                    <div className="max-w-2xl mx-auto mt-10">
                        <h3 className="text-2xl font-semibold text-gray-300 mb-4">
                            Key Features
                        </h3>
                        <ul className="list-disc list-inside text-gray-400 space-y-2">
                            {rsAuthenticatorFeatures.map((feature, index) => (
                                <li key={index}>{feature}</li>
                            ))}
                        </ul>

                        <p className="mt-6 text-gray-400">
                            With <strong>RsAuthenticator</strong>, you gain complete control over your
                            authentication
                            process while prioritizing
                            security and convenience.
                        </p>
                    </div>

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