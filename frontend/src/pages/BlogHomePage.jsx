import React from "react";

const covers = [
    // "/preview/ca236c3c-ff20-4e68-a87b-6bf97d047f80.webp",
    // "/preview/ca236c3c-ff20-4e68-a87b-6bf97d047f80.png",

    "/preview/png-home.png",
    "/preview/png.png",
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

            <section className="py-10 home-section  ">
                <div className="container mx-auto p-4 md:p-6">
                    <h2 className="text-3xl font-bold text-center text-primary mb-8">
                        Featured Projects
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

            <section className="container  mx-auto py-10">
                <div className=" shadow-lg rounded-lg p-4 md:p-6">
                    <h2 className="text-3xl font-bold text-primary mb-4">
                        How Setup MFA Authenticator
                    </h2>
                    <p className="  mb-6 text-gray-400">
                        Multi-Factor Authentication (MFA) is a security mechanism that
                        requires users to verify their identity using multiple steps. Google
                        Authenticator is a popular tool for implementing MFA. Here' ..
                    </p>

                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
                        {/* Step 1 */}
                        <div className="card bg-neutral text-neutral-content shadow-lg">
                            <div className="card-body">
                                <h3 className="card-title text-xl">Step 1: Setting Up the App</h3>
                                <p>
                                    - Download and install Google Authenticator from your app store.
                                    <br/>
                                    - Log in to the application or service where you want to enable
                                    MFA.
                                    <br/>
                                    - Look for the "Enable MFA" option, which often generates a QR
                                    code.
                                </p>
                            </div>
                        </div>

                        {/* Step 2 */}
                        <div className="card bg-neutral text-neutral-content shadow-lg">
                            <div className="card-body">
                                <h3 className="card-title text-xl">Step 2: Scanning the QR Code</h3>
                                <p>
                                    Open the Google Authenticator app and use the "Scan a QR code"
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
                                    from Google Authenticator. This code changes every 30 seconds
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

                    <div className="text-center mt-8">
                        <a
                            href="#"
                            className="btn btn-primary btn-lg"
                        >
                            Read More Articles
                        </a>
                    </div>
                </div>
            </section>

            {/* Footer */}
            <footer className="bg-primary text-primary-content py-4">
                <div className="container mx-auto text-center">
                    <p>&copy; 2025 MFA Blog. All rights reserved.</p>
                </div>
            </footer>
        </div>
    );
};

export default BlogHomePage;
