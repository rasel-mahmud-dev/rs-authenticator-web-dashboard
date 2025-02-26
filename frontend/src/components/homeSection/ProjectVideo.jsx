import React from "react";

const ProjectVideo = () => {
    return (
        <section className="py-10 home-section">
            <div className="max-w-screen-xl mx-auto p-4 md:p-6">
                <h2 className="text-3xl font-bold text-center text-primary mb-8">
                    See It In Action: Video Overview
                </h2>

                <div
                    className="transition-all duration-500 ease-in-out hover:scale-105    hover:shadow-lg  max-w-7xl border border-primary rounded-xl overflow-hidden mx-auto   ">



                    <iframe width="100%" height="600px" src="https://www.youtube.com/embed/0WMP7-_NKi0?vq=hd1080p"  ></iframe>


                    {/*<video*/}
                    {/*    poster="/rs-authenticator-thumb.webp"*/}
                    {/*    src="https://drive.google.com/uc?export=download&id=1ZhKG1f7cSAc8x7B1L1w1yLw_J7OUJX7p"*/}
                    {/*    className="w-full "*/}
                    {/*    controls>*/}
                    {/*</video>*/}


                </div>
            </div>
        </section>
    )
        ;
}

export default ProjectVideo