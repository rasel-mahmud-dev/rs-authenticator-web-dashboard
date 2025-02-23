import React from "react";

const ProjectVideo = () => {
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
        <section className="py-10 home-section">
            <div className="max-w-screen-xl mx-auto p-4 md:p-6">
                <h2 className="text-3xl font-bold text-center text-primary mb-8">
                    About Project App/Web
                </h2>

                <div className="max-w-8xl mx-auto">
                    <video src="/guide-rasel-mahmud.mp4" controls/>
                </div>

            </div>
        </section>
    );
}

export default ProjectVideo