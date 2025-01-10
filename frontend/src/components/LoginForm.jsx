import React from "react";

const LoginForm = () => {
    return (
        <div className="flex items-center justify-center min-h-screen  ">
            <div className="card w-96 bg-gray-800  shadow-xl">
                <div className="card-body">
                    <h2 className="text-2xl font-semibold text-center text-white">Welcome Back</h2>
                    <p className="text-center text-gray-400 mb-4">Sign in to your account</p>
                    <form>
                        <div className="form-control mb-4">
                            <label className="label">
                                <span className="label-text text-gray-300">Email</span>
                            </label>
                            <input
                                type="email"
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
                                placeholder="Enter your password"
                                className="input input-bordered w-full"
                            />
                        </div>
                        <div className="form-control mb-4">
                            <label className="label cursor-pointer justify-start gap-2">
                                <input type="checkbox" className="checkbox checkbox-primary"/>
                                <span className="label-text text-gray-300">Remember me</span>
                            </label>
                        </div>
                        <button className="btn btn-primary w-full">Sign In</button>
                    </form>
                    <p className="text-center text-gray-500 mt-4">
                        Don’t have an account?{" "}
                        <a href="#" className="text-primary font-medium">
                            Sign Up
                        </a>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default LoginForm;
