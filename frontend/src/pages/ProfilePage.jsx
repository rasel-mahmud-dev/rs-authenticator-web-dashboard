import React from "react";
import {Link} from "react-router-dom";
import useAuthStore from "../store/authState.js";
import {useQuery} from "@tanstack/react-query";
import {api} from "../services/api.js";

const UserProfile = () => {
    const {user} = useAuthStore(); // Get authenticated user details

    const query = useQuery({
        queryKey: ["profile"],
        queryFn: () => api.get("/api/v1/profile")
    })

    const data = query?.data?.data ?? {}

    const {
        user_id,
        full_name,
        account_created_at,
        gender,
        phone,
        created_at,
        updated_at,
    } = data


    return (
        <div className="text-white">
            {/* Profile Header */}
            <div className="relative w-full  bg-gray-800 flex justify-center items-center">

                <div className="profile-cover">
                    <img src="/magento-2-upload-product-images-placeholder.jpg" alt=""/>
                </div>

                <div className="absolute bottom-4 transform translate-y-1/2">
                    <div className="relative w-36 h-36">
                        <div className="avatar  ">
                            <div className="  rounded-full">
                                <img src="/boy.png"/>
                            </div>
                            <button
                                className="w-5 h-5 flex items-center justify-center  text-[10px] absolute bottom-2 right-0 bg-gray-700 p-3 rounded-full hover:bg-gray-600">
                                ✏️
                            </button>
                        </div>


                    </div>
                </div>
            </div>

            {/* Profile Info */}
            <div className="text-center mt-24">
                <h1 className="text-3xl font-semibold">{user?.username || "John Doe"}</h1>
                <p className="text-gray-400">@{full_name || "johndoe"}</p>
            </div>

            {/* Profile Details Section */}
            <div className="max-w-4xl mx-auto mt-8 px-4">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6 bg-gray-800 p-6 rounded-lg shadow-md">
                    <div>
                        <h2 className="text-gray-400 text-sm">Email</h2>
                        <p className="text-white">{user?.email || "johndoe@example.com"}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Phone</h2>
                        <p className="text-white">{phone || "+123 456 7890"}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Role</h2>
                        <p className="text-white">{user?.role || "User"}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Account Status</h2>
                        <p className="text-white">{user?.status || "Active"}</p>
                    </div>
                    <div className="col-span-2">
                        <h2 className="text-gray-400 text-sm">Bio</h2>
                        <p className="text-white">
                            {user?.bio || "No bio available. Add a short bio about yourself."}
                        </p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Location</h2>
                        <p className="text-white">{user?.location || "Unknown"}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Joined</h2>
                        <p className="text-white">
                            {user?.joined ? new Date(auth.joined).toLocaleDateString() : "N/A"}
                        </p>
                    </div>
                </div>
            </div>

            {/* Action Buttons */}
            <div className="flex justify-center gap-4 mt-8 mb-10">
                <Link to="/profile/edit" className="btn btn-primary px-6 py-2 rounded-lg bg-blue-500 hover:bg-blue-600">
                    Edit Profile
                </Link>
                <button
                    onClick={() => {
                        // Logout logic here
                    }}
                    className="btn btn-secondary px-6 py-2 rounded-lg bg-red-500 hover:bg-red-600"
                >
                    Log Out
                </button>
            </div>

        </div>
    );
};

export default UserProfile;
