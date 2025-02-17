import React, {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import useAuthStore from "../store/authState.js";
import {useQuery} from "@tanstack/react-query";
import {api} from "../services/api.js";
import {toast} from "react-toastify";

const UserProfile = () => {
    const query = useQuery({
        queryKey: ["profile"],
        queryFn: () => api.get("/api/v1/profile")
    })

    const {user, setAuth} = useAuthStore()

    const logout = () => {
        setAuth(null)
        localStorage.removeItem("token")
        sessionStorage.removeItem("token")
    };


    const profile = query?.data?.data ?? {}

    const {
        user_id,
        full_name,
        account_created_at,
        gender,
        phone,
        created_at,
        updated_at,
        about_me,
        avatar,
        cover,
        username
    } = profile

    const [preview, setPreview] = useState(avatar || "/boy.png");

    useEffect(() => {
        if (avatar) setPreview(avatar)
    }, [avatar]);

    const resizeAndCropImage = (file) => {
        return new Promise((resolve) => {
            const img = new Image();
            img.src = URL.createObjectURL(file);
            img.onload = () => {
                const canvas = document.createElement("canvas");
                const ctx = canvas.getContext("2d");

                const size = 500;
                canvas.width = size;
                canvas.height = size;

                const aspectRatio = img.width / img.height;
                let sx, sy, sWidth, sHeight;
                if (aspectRatio > 1) {
                    sWidth = img.height;
                    sHeight = img.height;
                    sx = (img.width - sWidth) / 2;
                    sy = 0;
                } else {
                    sWidth = img.width;
                    sHeight = img.width;
                    sx = 0;
                    sy = (img.height - sHeight) / 2;
                }

                ctx.drawImage(img, sx, sy, sWidth, sHeight, 0, 0, size, size);
                canvas.toBlob((blob) => resolve(blob), "image/jpeg", 0.8);
            };
        });
    };

    const handleFileChange = async (event) => {
        const file = event.target.files[0];
        if (!file) return;

        const resizedBlob = await resizeAndCropImage(file);

        const previewUrl = URL.createObjectURL(resizedBlob);
        setPreview(previewUrl);

        const formData = new FormData();
        formData.append("image", resizedBlob, "profile.jpg");

        try {
            toast.info("Uploading Profile Photo")
            await api.post("/api/v1/profile/avatar", formData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
            });

            toast.success("Successfully Update Profile Photo")
        } catch (error) {
            toast.error("Upload Error:");
        }
    };


    return (
        <div className="text-white vh
        ">
            <div className="relative w-full  bg-gray-800 flex justify-center items-center">

                <div className="profile-cover">
                    <img src="/magento-2-upload-product-images-placeholder.jpg" alt=""/>
                </div>

                <div className="absolute bottom-4 transform translate-y-1/2">
                    <div className="relative w-36 h-36">
                        <div className="avatar">
                            <div className="rounded-full w-36 h-36">
                                <img src={preview} alt="Profile" className="w-36 h-36 object-cover rounded-full"/>
                            </div>

                            {/* Hidden File Input */}
                            <input
                                type="file"
                                accept="image/*"
                                id="fileInput"
                                className="hidden"
                                onChange={handleFileChange}
                            />

                            {/* Button to trigger file input */}
                            <button
                                onClick={() => document.getElementById("fileInput").click()}
                                className="w-6 h-6 flex items-center justify-center text-[10px] absolute bottom-2 right-0 bg-gray-700 p-3 rounded-full hover:bg-gray-600"
                            >
                                ✏️
                            </button>
                        </div>


                    </div>
                </div>
            </div>

            {/* Profile Info */}
            <div className="text-center mt-24">
                <h1 className="text-3xl font-semibold">{user?.username || "John Doe"}</h1>
                <p className="text-gray-400">@{full_name || username}</p>
            </div>

            {/* Profile Details Section */}
            <div className="max-w-4xl mx-auto mt-8 px-4">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6 bg-gray-800 p-6 rounded-lg shadow-md">
                    <div>
                        <h2 className="text-gray-400 text-sm">Email</h2>
                        <p className="text-white">{profile?.email}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Phone</h2>
                        <p className="text-white">{profile?.phone}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Role</h2>
                        <p className="text-white">{profile.role || "User"}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Account Status</h2>
                        <p className="text-white">{profile.status || "Active"}</p>
                    </div>
                    <div className="col-span-2">
                        <h2 className="text-gray-400 text-sm">Bio</h2>
                        <p className="text-white">{profile.about_me || "No bio available."}</p>
                    </div>

                    <div>
                        <h2 className="text-gray-400 text-sm">Gender</h2>
                        <p className="text-white">{profile.gender || "Not specified"}</p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Date of Birth</h2>
                        <p className="text-white">
                            {profile.account_created_at ? new Date(profile.account_created_at).toLocaleDateString() : "N/A"}
                        </p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Joined</h2>
                        <p className="text-white">
                            {profile.account_created_at ? new Date(profile.account_created_at).toLocaleDateString() : "N/A"}
                        </p>
                    </div>
                    <div>
                        <h2 className="text-gray-400 text-sm">Last Updated</h2>
                        <p className="text-white">
                            {profile.updated_at ? new Date(profile.updated_at).toLocaleDateString() : "N/A"}
                        </p>
                    </div>
                </div>
            </div>

            {/* Action Buttons */}
            <div className="flex justify-center gap-4 mt-8 mb-10">
                <Link to="/account/profile/edit"
                      className="btn btn-primary px-6 py-2 rounded-lg bg-blue-500 hover:bg-blue-600">
                    Edit Profile
                </Link>
                <button
                    onClick={logout}
                    className="btn btn-secondary px-6 py-2 rounded-lg bg-red-500 hover:bg-red-600"
                >
                    Log Out
                </button>
            </div>

        </div>
    );
};

export default UserProfile;
