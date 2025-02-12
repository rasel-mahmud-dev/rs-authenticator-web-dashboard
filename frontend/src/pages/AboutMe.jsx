import React from "react";

const AboutMe = () => {
    return (

        <div className="mt-10 max-w-screen-2xl mx-auto px-4">
            <div className="max-w-3xl mx-auto ">
                <h1 className="text-gray-100 text-3xl font-bold text-center">About Me</h1>

                <p className="text-lg text-gray-300 mt-8">
                    Hello! I'm Rasel, a passionate developer specializing in JavaScript, Node.js, React, and PostgreSQL.
                    I love building scalable applications and optimizing performance.
                </p>
                <p className="text-lg text-gray-300 mt-4">
                    With experience in backend and frontend technologies, I focus on creating efficient and
                    user-friendly software solutions. Let's build something amazing together!
                </p>
                <div className="mt-6 flex justify-center">
                    <a href="https://rasel-portfolio.vercel.app/#contact-section" target="_blank"
                       className="btn btn-primary">Contact Me</a>
                </div>
            </div>
        </div>
    );
};

export default AboutMe;
