import React from "react";
import appCovers from "../assets/appCovers.json"

const DownloadApp = () => {
    const appInfo = {
        name: "Download App",
        description: "Download the latest version of the app",
        logo: "https://rasel-portfolio.vercel.app/logo.png",
        images: appCovers
    };

    const versions = [
        {
            version: "0.0.1",
            link: "https://rasel-portfolio.vercel.app/#contact-section",
            releaseNotes: [
                "Initial release with core features.",
                "Basic user authentication added.",
                "Bug fixes and performance improvements."
            ],
        },
        {
            version: "0.0.2",
            link: "https://rasel-portfolio.vercel.app/#contact-section",
            releaseNotes: [
                "Improved UI/UX design.",
                "Added dark mode support.",
                "Fixed login issue for some users."
            ],
        },
        {
            version: "0.0.3",
            link: "https://rasel-portfolio.vercel.app/#contact-section",
            releaseNotes: [
                "Enhanced security features.",
                "Added multi-language support.",
                "Fixed minor bugs in dashboard."
            ],
        },
        {
            version: "0.0.4",
            link: "https://rasel-portfolio.vercel.app/#contact-section",
            releaseNotes: [
                "Major performance improvements.",
                "New file upload functionality added.",
                "Improved API response time."
            ],
        },
    ];

    return (
        <div className="mt-10 max-w-screen-lg mx-auto px-4 text-center">
            <img src={appInfo.logo} alt="App Logo" className="w-20 h-20 mx-auto mb-4"/>

            <h1 className="text-white text-3xl font-bold">{appInfo.name}</h1>
            <p className="text-gray-300 mt-2">{appInfo.description}</p>

            <div className="grid grid-cols-12 mt-6 gap-x-2 py-10">

                <div className="col-span-4">
                    <div className="flex flex-wrap gap-4">
                        {appCovers?.map((project, index) => (
                            <div
                                key={index}
                                className="w-32 android-frame">

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

                <div className="col-span-8 space-y-6">
                    {versions.map(({version, link, releaseNotes}) => (
                        <div
                            key={version}
                            className=" border border-gray-700  p-6 rounded-lg shadow-lg text-left"
                        >
                            <div className="flex justify-between items-center">
                                <h2 className="text-white text-xl font-semibold">Version {version}</h2>

                            </div>

                            <ul className="mt-3 text-gray-300 text-sm list-disc pl-5 space-y-1">
                                {releaseNotes.map((note, index) => (
                                    <li key={index}>{note}</li>
                                ))}
                            </ul>

                            <div className="mt-6 ">
                                <a
                                    href={link}
                                    download
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    className="btn btn-outline   text-white font-bold   px-4 rounded-lg transition-all"
                                >
                                    Download
                                </a>
                            </div>
                        </div>
                    ))}
                </div>

            </div>


        </div>
    );
};

export default DownloadApp;
