import React, {useEffect, useState} from "react";
import {useMutation, useQuery} from "@tanstack/react-query";
import {api} from "../services/api.js";
import {toast} from "react-toastify";
import {Link} from "react-router-dom"

const EditProfilePage = () => {
    const [profile, setProfile] = useState({
        fullName: "",
        birthDate: "",
        gender: "",
        phone: "",
        location: "",
        aboutMe: "",
        website: "",
        facebook: "",
        twitter: "",
        linkedin: "",
        instagram: "",
        github: "",
        youtube: "",
        tiktok: "",
    });

    const query = useQuery({
        queryKey: ["profile"],
        queryFn: () => api.get("/api/v1/profile")
    })
    const userProfile = query?.data?.data
    const isLoading = query?.isLoading


    useEffect(() => {
        if (userProfile) {
            const bd = userProfile?.birth_date?.split("T")?.[0]
            setProfile({
                fullName: userProfile.full_name || "",
                birthDate: bd || "",
                gender: userProfile.gender || "",
                phone: userProfile.phone || "",
                location: userProfile.location || "",
                aboutMe: userProfile.about_me || "",
                website: userProfile.website || "",
                facebook: userProfile.facebook || "",
                twitter: userProfile.twitter || "",
                linkedin: userProfile.linkedin || "",
                instagram: userProfile.instagram || "",
                github: userProfile.github || "",
                youtube: userProfile.youtube || "",
                tiktok: userProfile.tiktok || "",
            });
        }
    }, [userProfile]);

    const mutation = useMutation({
        mutationFn: (profileData) => api.put("/api/v1/profile", profileData),
        onSuccess: () => {
            // handle success (e.g., navigate, show success message)
            toast("Successfully updated profile.")
        },
        onError: () => {
            // handle error (e.g., show error message)
            toast("Failed to updated profile.")
        },
    });

    const handleChange = (e) => {
        setProfile({
            ...profile,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        mutation.mutate(profile);
    };

    if (isLoading) return <div>Loading...</div>;

    return (
        <div className="mt-10 max-w-screen-2xl mx-auto px-4">
            <div className="flex justify-center relative ">
                <Link to="/account/profile">
                    <button className="btn btn-outline btn-primary absolute left-0">Back to profile</button>
                </Link>
                <h2 className="text-3xl  text-gray-100 font-bold text-center">Edit Profile</h2>
            </div>

            <div className="max-w-3xl mx-auto ">
                <form onSubmit={handleSubmit} className="space-y-4 my-8">

                    <div role="tablist" className="tabs tabs-bordered profile-edit-tablist">
                        <input defaultChecked type="radio" name="my_tabs_1" role="tab" className="tab"
                               aria-label="General Information"/>
                        <div role="tabpanel" className="tab-content py-10">

                            <div>
                                <div>
                                    <label className="label">
                                        <span className="label-text">Full Name</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="fullName"
                                        value={profile.fullName}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Birth Date</span>
                                    </label>
                                    <input
                                        type="date"
                                        name="birthDate"
                                        value={profile.birthDate}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Gender</span>
                                    </label>
                                    <select
                                        name="gender"
                                        value={profile.gender}
                                        onChange={handleChange}
                                        className="select select-bordered w-full"
                                    >
                                        <option value="">Select Gender</option>
                                        <option value="male">Male</option>
                                        <option value="female">Female</option>
                                        <option value="other">Other</option>
                                    </select>
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Phone</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="phone"
                                        value={profile.phone}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Location</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="location"
                                        value={profile.location}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">About Me</span>
                                    </label>
                                    <textarea
                                        name="aboutMe"
                                        value={profile.aboutMe}
                                        onChange={handleChange}
                                        className="textarea textarea-bordered w-full"
                                        rows="4"
                                    />
                                </div>
                            </div>

                        </div>


                        <input type="radio" name="my_tabs_1" role="tab" className="tab" aria-label="Social Links"/>
                        <div role="tabpanel" className="tab-content py-10">

                            <div>
                                <div>
                                    <label className="label">
                                        <span className="label-text">Website</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="website"
                                        value={profile.website}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Facebook</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="facebook"
                                        value={profile.facebook}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Twitter</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="twitter"
                                        value={profile.twitter}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">LinkedIn</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="linkedin"
                                        value={profile.linkedin}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">Instagram</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="instagram"
                                        value={profile.instagram}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">GitHub</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="github"
                                        value={profile.github}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">YouTube</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="youtube"
                                        value={profile.youtube}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>

                                <div>
                                    <label className="label">
                                        <span className="label-text">TikTok</span>
                                    </label>
                                    <input
                                        type="text"
                                        name="tiktok"
                                        value={profile.tiktok}
                                        onChange={handleChange}
                                        className="input input-bordered w-full"
                                    />
                                </div>
                            </div>


                        </div>
                    </div>


                    <button type="submit" className="btn btn-primary w-full mt-6">
                        Save Changes
                    </button>


                </form>

            </div>
        </div>
    )
};

export default EditProfilePage;
