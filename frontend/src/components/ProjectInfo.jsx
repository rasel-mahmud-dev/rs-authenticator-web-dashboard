import React from "react";
import appCovers from "../assets/appCovers.json"
import {Link } from 'react-router-dom'

const ProjectInfo = () => {
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
        <section className="py-10">
            <div className="container mx-auto p-4 md:p-6">
                <h2 className="text-3xl font-bold text-center text-primary mb-8">
                    About Project Web
                </h2>

                <div className="max-w-2xl mx-auto">

                    <p className="mb-6">
                        The backend is built with <strong>Golang</strong> and
                        utilizes <strong>PostgreSQL</strong> to manage
                        user data, authentication sessions, traffic logs, <strong>recovery codes</strong> and <strong>MFA
                        connection secrets</strong>. This
                        ensures a robust
                        and scalable authentication system with secure data storage.
                    </p>

                    <p className="mb-6">
                        It implements <strong>image upload functionality using Cloudinary</strong> on the
                        server side,
                        allowing users to update their profile photos efficiently. This integration ensures
                        optimized
                        image handling with minimal server load.
                    </p>

                    <p className="mb-6">
                        To enhance code reusability and maintainability, the system follows the
                        <strong> Chain of Responsibility design pattern</strong>. This approach enables the
                        efficient splitting
                        and reuse of handlers for specific tasks, improving modularity and reducing
                        redundancy.
                    </p>

                    <p className="mb-6">
                        Additionally, caching mechanisms are incorporated into selected endpoints and
                        repositories. This
                        significantly reduces database load, enhances system performance, and ensures faster
                        response
                        times for a smoother user experience.
                    </p>


                    <h3 className="text-2xl font-semibold text-gray-300 mb-4">
                        Frontend Features
                    </h3>
                    <ul className="list-disc list-inside text-gray-400 space-y-2">
                        {frontendFeatures.map((feature, index) => (
                            <li className="text-orange-500" key={index}>
                                <span className="">{feature}</span>
                            </li>
                        ))}
                    </ul>
                </div>

                <h2 className="text-3xl font-bold text-center text-primary mt-12 mb-8">
                    Android App: RsAuthenticator
                </h2>

                <div>

                    <div className="max-w-2xl mx-auto">


                        <Link to="/download-app">
                            <img src="/ic_launcher.png" alt="app-logo"
                                 className="mx-auto border rounded-3xl p-2 border-gray-500 w-16 h-16 mb-4"/>
                        </Link>

                        <p className="mb-6 text-gray-400">
                            <strong>RsAuthenticator</strong> is a secure authenticator app developed using <strong>Kotlin
                            Jetpack Compose</strong>.
                            It ensures <strong>100% secure authentication</strong> without requiring internet access.
                        </p>
                    </div>

                    <div className="flex justify-center gap-6">
                        {appCovers?.map((project, index) => (
                            <div
                                key={index}
                                className="w-36 android-frame">

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
                </div>
            </div>
        </section>
    );
}

export default ProjectInfo