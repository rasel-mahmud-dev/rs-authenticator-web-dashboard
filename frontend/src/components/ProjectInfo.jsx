import React from "react";
import appCovers from "../assets/appCovers.json"


const ProjectInfo = () => {
    const frontendFeatures = [
        "Registration Trends",
        "Login Methods Breakdown (Password vs. Authenticator App)",
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
        <section className="py-10">
            <div className="container mx-auto p-4 md:p-6">
                <h2 className="text-3xl font-bold text-center text-primary mb-8">
                    About Project Web
                </h2>

                <div className="max-w-2xl mx-auto">
                    <p className="mb-6 text-gray-400">
                        The backend is built with <strong>Golang</strong> and uses <strong>PostgreSQL</strong> to store
                        user data,
                        authentication sessions, traffic logs, and MFA connection secrets.
                    </p>

                    <p className="mb-6 text-gray-400">
                        It implements the <strong>Chain of Responsibility</strong> design pattern to efficiently split
                        and reuse handlers
                        for specific tasks. Additionally, caching mechanisms are integrated into certain endpoints and
                        repositories
                        to reduce database load and enhance response times.
                    </p>

                    <h3 className="text-2xl font-semibold text-gray-300 mb-4">
                        Frontend Features
                    </h3>
                    <ul className="list-disc list-inside text-gray-400 space-y-2">
                        {frontendFeatures.map((feature, index) => (
                            <li key={index}>
                                <span>{feature}</span>
                            </li>
                        ))}
                    </ul>
                </div>

                <h2 className="text-3xl font-bold text-center text-primary mt-12 mb-8">
                    Android App: RsAuthenticator
                </h2>

                <div>
                    <div className=" ">
                        <div className="flex justify-center gap-6">
                            {appCovers?.map((project, index) => (
                                <div
                                    key={index}
                                    className="w-64 android-frame">

                                    {/* Screenshot inside the frame */}
                                    <figure className="w-full h-full overflow-hidden relative z-0">
                                        <img
                                            src={project}
                                            alt={`Project ${index + 1}`}
                                            className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
                                        />
                                    </figure>
                                </div>
                            ))}
                        </div>


                    </div>

                    <div className="max-w-2xl mx-auto">
                        <p className="mb-6 text-gray-400">
                            <strong>RsAuthenticator</strong> is a secure authenticator app developed using <strong>Kotlin
                            Jetpack Compose</strong>.
                            It ensures <strong>100% secure authentication</strong> without requiring internet access.
                        </p>

                        <h3 className="text-2xl font-semibold text-gray-300 mb-4">
                            Key Features
                        </h3>
                        <ul className="list-disc list-inside text-gray-400 space-y-2">
                            {rsAuthenticatorFeatures.map((feature, index) => (
                                <li key={index}>{feature}</li>
                            ))}
                        </ul>

                        <p className="mt-6 text-gray-400">
                            With <strong>RsAuthenticator</strong>, you gain complete control over your authentication
                            process while prioritizing
                            security and convenience.
                        </p>
                    </div>
                </div>
            </div>
        </section>
    );
}

export default ProjectInfo