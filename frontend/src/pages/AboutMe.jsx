import React from "react";
import {FaFacebook, FaGithub, FaLinkedin} from "react-icons/fa";
import {TbWorldWww} from "react-icons/tb";

const AboutMe = () => {
    return (
        <div className="flex mt-0 container mx-auto px-4">
            <div className="top-0 md:top-[20vh] relative  mx-auto ">
                <h1 className="text-gray-100 text-3xl font-bold text-center">About Me</h1>

                <div className="grid grid-cols-12 gap-x-10 mt-10">
                    <div className="col-span-4 ml-auto max-w-[300px]">
                        <img className="rounded-xl" src="/preview/1740327521513.webp"
                             alt="profile-avatar"/>
                    </div>

                    <div className="col-span-8">
                        <p className="text-lg text-gray-300">
                            Hello! I'm Rasel, a passionate developer specializing in <strong
                            className="text-orange-500"> JavaScript/Typescript, Node.js,
                            React, and
                            PostgreSQL</strong>.
                            I love building scalable applications and optimizing performance.
                        </p>
                        <p className="text-lg text-gray-300 mt-4">
                            With experience in backend and frontend technologies, I focus on creating efficient and
                            user-friendly software solutions.
                        </p>
                        <p className="text-lg text-gray-300 mt-4">
                            I specialize in building high-performance applications using
                            <strong className="text-orange-500"> Go (Golang)</strong>. While I have some experience with
                            mobile development,
                            I’m still growing my skills in <strong className="text-orange-500">Jetpack Compose
                            (Kotlin Android)</strong>.
                        </p>

                        <p className="mt-6 text-xl font-medium text-gray-100">
                            Contact
                        </p>
                        <div className="flex space-x-4 mt-4">
                            <a href="https://github.com/rasel-mahmud-dev" target="_blank" rel="noopener noreferrer"
                               className="w-12 h-12 border border-gray-500 hover:bg-primary flex items-center justify-center rounded-lg      text-gray-300 hover:text-white text-2xl">
                                <FaGithub/>
                            </a>
                            <a href="https://www.facebook.com/rasel.mahmud.dev" target="_blank"
                               rel="noopener noreferrer"
                               className=" w-12 h-12 border border-gray-500 hover:bg-primary  flex items-center justify-center rounded-lg  text-gray-300 hover:text-white text-2xl">
                                <FaFacebook/>
                            </a>
                            <a href="https://www.linkedin.com/in/rasel-mahmud-dev/" target="_blank"
                               rel="noopener noreferrer"
                               className=" w-12 h-12 border border-gray-500 hover:bg-primary  flex items-center justify-center rounded-lg  text-gray-300 hover:text-white text-2xl">
                                <FaLinkedin/>
                            </a>

                            <a href="https://rasel-portfolio.vercel.app/" target="_blank"
                               rel="noopener noreferrer"
                               className=" w-12 h-12 border border-gray-500 hover:bg-primary  flex items-center justify-center rounded-lg  text-gray-300 hover:text-white text-2xl">
                                <TbWorldWww/>
                            </a>
                        </div>


                    </div>

                </div>

            </div>
        </div>
    );
};

export default AboutMe;
