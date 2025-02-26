import React from "react";
import ProjectInfo from "../components/homeSection/ProjectInfo.jsx";
import {Link} from "react-router-dom"
import ProjectVideo from "../components/homeSection/ProjectVideo.jsx";
import AndroidAppHomeSection from "../components/homeSection/AndroidAppHomeSection.jsx";

const covers = [

    "/preview/png-home.png",
    "/preview/png (10)33.png",
    "/preview/api_latency.png",
    "/preview/png (1).png",
    "/preview/png (2).png",
    "/preview/png (3).png",
    "/preview/png (4).png",
    "/preview/png (5).png",
    "/preview/png (6).png",
    "/preview/png (7).png",
    "/preview/png (8).png",
    "/preview/png (9).png",
]


const BlogHomePage = () => {

    return (
        <div className="">
            <ProjectInfo/>
            <ProjectVideo/>
            <AndroidAppHomeSection/>

            <section className="py-10  home-section  ">
                <div className="container mx-auto p-4 md:p-6">
                    <h2 className="text-3xl font-bold text-center text-primary mb-8">
                        Web UI/Screens
                    </h2>
                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                        {covers.map((project, index) => (
                            <div
                                key={index}
                                className="group relative bg-neutral shadow-lg overflow-hidden project-image"
                            >
                                <figure className="w-full  overflow-hidden">
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
            </section>

            <section className="container  text-center mx-auto py-10">
                <div className=" max-w-6xl mx-auto shadow-lg rounded-lg p-4 md:p-6">
                    <h2 className="text-3xl font-bold text-primary mb-4">
                        How Setup MFA Authenticator
                    </h2>
                    <p className=" text-center mb-6 text-gray-400">
                        Multi-Factor Authentication (MFA) is a security mechanism that
                        requires users to verify their identity using multiple steps.
                    </p>

                    <div className="grid  text-start grid-cols-1 lg:grid-cols-2 gap-8">
                        {/* Step 1 */}
                        <div className="card bg-neutral text-neutral-content shadow-lg">
                            <div className="card-body">
                                <h3 className="card-title text-xl">Step 1: Setting Up the App</h3>
                                <p>
                                    - First, <Link to="/download-app" className="text-blue-500 hover:underline">download
                                    RsAuth</Link> from our <strong>Download Page</strong>, or alternatively, you can
                                    use <strong>Google Authenticator</strong> to test the setup.
                                    <br/>
                                    - Next, log in to the application or service where you want to enable Multi-Factor
                                    Authentication (MFA).
                                    <br/>
                                    - Once logged in, navigate to the <strong>"Enable MFA"</strong> section. This is
                                    usually where the system will prompt you to scan a QR code for setting up the
                                    authentication.
                                </p>
                            </div>

                        </div>

                        {/* Step 2 */}
                        <div className="card bg-neutral text-neutral-content shadow-lg">
                            <div className="card-body">
                                <h3 className="card-title text-xl">Step 2: Scanning the QR Code</h3>
                                <p>
                                    Open the Authenticator app and use the "Scan a QR code"
                                    option to link your account. The app will then generate a unique
                                    six-digit code for your account.
                                </p>
                            </div>
                        </div>

                        {/* Step 3 */}
                        <div className="card bg-neutral text-neutral-content shadow-lg">
                            <div className="card-body">
                                <h3 className="card-title text-xl">Step 3: Using the Generated Code</h3>
                                <p>
                                    Whenever you log in, the system will ask for the six-digit code
                                    from Authenticator. This code changes every 30 seconds
                                    and ensures only authorized users can access the account.
                                </p>
                            </div>
                        </div>

                        {/* Step 4 */}
                        <div className="card bg-neutral text-neutral-content shadow-lg">
                            <div className="card-body">
                                <h3 className="card-title text-xl">Step 4: Validating the Code</h3>
                                <p>
                                    The service verifies the code using a shared secret key and the
                                    current time. If the code matches, you're granted access.
                                </p>
                            </div>
                        </div>
                    </div>


                </div>
            </section>

            {/* Footer */}
            <footer className="text-primary-content py-4">
                <div className="container mx-auto text-center">
                    <p>&copy; {new Date().getFullYear()} Rasel Mahmud. All rights reserved.</p>
                    <p>Made by <a href="https://rasel-portfolio.vercel.app/#" className="text-primary">Rasel Mahmud</a>
                    </p>
                </div>
            </footer>

        </div>
    );
};

export default BlogHomePage;
