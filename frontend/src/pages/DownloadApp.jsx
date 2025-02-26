import React, {useState} from "react";
import appCovers from "../assets/appCovers.json"
import {api} from "../services/api.js";
import {toast} from "react-toastify";
import {getTotalDownload} from "../services/statsService.js";
import {useQuery} from "@tanstack/react-query";
import {DownloadCloud} from "lucide-react";


let loadingQueue = 0


const DownloadApp = () => {
    const appInfo = {
        name: "Download App",
        description: "Download the latest version of the app",
        images: appCovers
    };

    const [selectedIndex, setSelectedIndex] = useState(-1);

    const closePreview = () => {
        setSelectedIndex(-1);
    };

    const prevImage = (e) => {
        e.stopPropagation()
        setSelectedIndex((prev) => (prev > 0 ? prev - 1 : prev));
    };

    const nextImage = (e) => {
        e.stopPropagation()
        setSelectedIndex((prev) => (prev < appCovers.length - 1 ? prev + 1 : prev));
    };


    const versions = [
        {
            version: "1.0.0",
            link: "https://drive.google.com/uc?export=download&id=1ldE3Xx-GvRO0iUykeIw3g_xlFY2JPWIJ",
            releaseNotes: [
                "Initial release with core features.",
                "Basic user authentication added.",
                "Bug fixes and performance improvements."
            ],
        }
    ];

    const {data, refetch} = useQuery({
        queryKey: ["totalDownload"],
        queryFn: getTotalDownload
    })

    const totalDownload = data?.data?.totalDownload || 0


    async function handleDownload() {
        try {
            loadingQueue = toast.loading("Generating download link..");
            const {data} = await api.get(`/api/v1/download`);
            const downloadURL = data?.data?.downloadURL
            if (downloadURL) {
                refetch()
                window.open(downloadURL, "_blank")
            }
        } catch (error) {
            toast.error("Error downloading file: " + error?.message);
        } finally {
            toast.done(loadingQueue)
        }
    }

    return (
        <div className="mt-10 max-w-screen-lg mx-auto px-4 text-center ">

            <h1 className="text-gray-100 text-3xl font-bold text-center">{appInfo.name}</h1>
            <p className="text-gray-300 mt-2">{appInfo.description}</p>

            <div className="max-w-3xl mx-auto p-6  mt-4  text-white shadow-lg rounded-lg">
                <div className="block md:flex  text-start items-center justify-between gap-4">
                    <div className="flex items-center gap-4">
                        <img
                            src="/ic_launcher.png"
                            alt="App Logo"
                            className="w-28 h-28 "
                        />
                        <div>
                            <h1 className="text-xl font-bold">RsAuthenticator <span className="text-gray-400">Fast & Safe</span>
                            </h1>
                            <p className="text-sm text-gray-400">Secure</p>
                            <p className="text-sm text-gray-400">Version: 1.0.0</p>
                            <span className="badge badge-success text-xs">Trusted App</span>
                        </div>
                    </div>
                    <div className="mt-6 text-center">
                        <button onClick={handleDownload}
                                className="btn  text-white px-6 bg-primary-400 hover:bg-primary-600">⬇ Download
                            APK
                        </button>
                    </div>
                </div>


                <div className="flex justify-between items-center mt-10 text-gray-400">
                    <div className="text-yellow-400 gap-x-2 flex items-center font-semibold">
                        <DownloadCloud/>
                        <span className="text-yellow-400">{totalDownload} Downloads</span>
                    </div>
                    <div>Feb 20, 2025</div>
                    <div>Android 10.0+</div>
                </div>

                <div className="col-span-8 space-y-6">
                    {versions.map(({version, link, releaseNotes}) => (
                        <div
                            key={version}
                            className="  text-left"
                        >

                            <ul className="mt-3 text-gray-300 text-sm list-disc pl-5 space-y-1">
                                {releaseNotes.map((note, index) => (
                                    <li key={index}>{note}</li>
                                ))}
                            </ul>
                        </div>
                    ))}
                </div>

                <div className="mt-10">
                    <p className="text-start font-medium text-xl text-white">Rs Authenticator Screenshots</p>
                    <div className="flex flex-wrap gap-4 mt-4">
                        {appCovers?.map((project, index) => (
                            <div
                                key={index}
                                className="w-32  ">
                                <figure className="w-full h-full overflow-hidden relative z-0">
                                    <img onClick={() => setSelectedIndex(index)}
                                         src={project}
                                         alt={`Project ${index + 1}`}
                                         className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
                                    />
                                </figure>
                            </div>
                        ))}
                    </div>
                </div>

                <div
                    className={`fixed backdrop-blur inset-0  select-none  z-[1000]  bg-black bg-opacity-80 flex flex-col justify-center items-center  transition-opacity duration-300 ease-in-out  transform ${selectedIndex !== -1 ? 'opacity-100 scale-100 translate-y-0' : 'opacity-0 scale-95 translate-y-10 pointer-events-none'}`}>
                    <button onClick={closePreview} className="absolute top-4 right-4 text-white text-2xl">✖</button>
                    <img src={appCovers[selectedIndex]} alt="Preview"
                         className=" android-frame  max-w-full max-h-[90vh]"/>

                    <div className="flex md:block gap-x-4">
                        <div className="relative md:absolute top-1/2 bottom-0 md:left-1/4">
                            <button className="btn btn-outline btn-primary px-6" onClick={prevImage}
                                    disabled={selectedIndex === 0}>
                                Prev
                            </button>

                        </div>
                        <div className="relative md:absolute top-1/2 bottom-0 md:right-1/4    ">
                            <button className="btn btn-outline btn-primary px-6" onClick={nextImage}
                                    disabled={selectedIndex === appCovers.length - 1}>
                                Next
                            </button>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    )
};

export default DownloadApp;
