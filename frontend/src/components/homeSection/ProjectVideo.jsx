import React from "react";

const ProjectVideo = () => {
    return (
        <section className="py-10 home-section">
            <div className="max-w-screen-xl mx-auto p-4 md:p-6">
                <h2 className="text-3xl font-bold text-center text-primary mb-8">
                    See It In Action: Video Overview
                </h2>
                <div
                    className="transition-all duration-500 ease-in-out hover:scale-105 hover:shadow-lg max-w-7xl border border-primary rounded-xl overflow-hidden mx-auto">
                    <video
                        poster="/rs-authenticator-thumb.webp"
                        src="/rs-authenticator-2.mp4"
                        controls
                        className="w-full "
                    />
                </div>
            </div>
        </section>
    );
}

export default ProjectVideo