const NotFound = () => {
    return (
        <div className="flex items-center justify-center vh bg-gray-100">
            <div className="bg-white rounded-lg   max-w-sm w-full p-8">
                <h2 className="text-2xl font-semibold text-center text-gray-800 mb-6">Not Found.</h2>
                <p className="text-sm font-medium text-center text-gray-700 mb-6">Seems like you are lost the right
                    path!</p>
                <div className="space-y-4 flex justify-center  text-center">
                    <a
                        href="/"
                        className="px-6 py-2 bg-blue-500 text-white rounded-lg font-medium text-base hover:bg-blue-600 transition duration-300"
                    >
                        Back to Home</a>
                </div>
            </div>
        </div>
    );
};

export default NotFound;
