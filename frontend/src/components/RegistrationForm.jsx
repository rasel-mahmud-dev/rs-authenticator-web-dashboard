import React, {useEffect} from "react";
import {Link, useNavigate} from "react-router-dom";
import {useMutation} from "@tanstack/react-query";
import {register} from "../services/authSerivce.js";
import useAuthStore from "../store/authState.js";
import {toast} from "react-toastify";
import { faker } from '@faker-js/faker';

const RegistrationForm = () => {
    const navigate = useNavigate();

    const fakeEmail = faker.internet.email().toLowerCase();
    const randomName = faker.internet.username().toLowerCase();



    const mutation = useMutation({
        mutationFn: register,
        onSuccess: (_data, variables) => {
            toast.success("Great! You have successfully registered!");
            navigate(`/login?email=${variables.email}`);
        },
        onError: (error) => {
            toast.error("Oops! Registration failed.");
            console.error("Registration failed:", error);
        },
    });

    function handleRegister(e) {
        e.preventDefault();
        const username = e.target.username.value;
        const email = e.target.email.value;
        const password = e.target.password.value;
        mutation.mutate({username, email, password});
    }

    const errorMessage = mutation?.error?.response?.data?.message;

    return (
        <div className="vh  mt-10 md:mt-0 overflow-hidden">
            <div className="relative top-0 md:top-[20vh] w-96 rounded-xl mx-auto md:bg-gray-800  shadow-none md:shadow-xl">
                <div className="card-body">
                    <h2 className="text-2xl font-semibold text-center text-white">Create an Account</h2>
                    <p className="text-center text-gray-400 mb-4">Sign up to get started</p>
                    {errorMessage && <p className="text-center text-red-500 mb-4">{errorMessage}</p>}
                    <form onSubmit={handleRegister} className="form-phone">
                        <div className="form-control mb-4">
                            <label className="label">
                                <span className="label-text text-gray-300">Username</span>
                            </label>
                            <input
                                type="text"
                                defaultValue={randomName}
                                name="username"
                                placeholder="Enter username"
                                className="input input-bordered w-full"
                            />
                        </div>
                        <div className="form-control mb-4">
                            <label className="label">
                                <span className="label-text text-gray-300">Email</span>
                            </label>
                            <input
                                type="email"
                                defaultValue={fakeEmail}
                                name="email"
                                placeholder="Enter your email"
                                className="input input-bordered w-full"
                            />
                        </div>
                        <div className="form-control mb-4">
                            <label className="label">
                                <span className="label-text text-gray-300">Password</span>
                            </label>
                            <input
                                type="password"
                                name="password"
                                placeholder="Create a password"
                                className="input input-bordered w-full"
                            />
                        </div>
                        <button type="submit" className="btn btn-primary w-full">Sign Up</button>
                    </form>
                    <p className="text-center text-gray-500 mt-4">
                        Already have an account?{" "}
                        <Link to="/login" className="text-primary font-medium">
                            Sign In
                        </Link>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default RegistrationForm;
