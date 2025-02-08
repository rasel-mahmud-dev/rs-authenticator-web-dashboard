import React from "react";
import {Link, useNavigate} from "react-router-dom";
import {useMutation} from "@tanstack/react-query";
import {login} from "../services/authSerivce.js";
import useAuthStore from "../store/authState.js";

const LoginForm = () => {

    const {setAuth} = useAuthStore()

    const navigate = useNavigate()

    const mutation = useMutation({
        mutationFn: login,
        onSuccess: (data, variables) => {
            setAuth(data?.data)
            localStorage.setItem("auth-remember-me", variables.rememberMe);
            if (variables.rememberMe) {
                localStorage.setItem("token", data?.data?.token);
            } else {
                sessionStorage.setItem("token", data?.data?.token);
            }
            navigate('/account')
        },
        onError: (error) => {
            console.error("Login failed:", error);
        }
    });

    function handleLogin(e) {
        e.preventDefault();
        const email = e.target.email.value;
        const password = e.target.password.value;
        const rememberMe = e.target.rememberMe.checked;
        mutation.mutate({email, password, rememberMe});
    }

    const errorMessage = mutation?.error?.response?.data?.message


    return (
        <div className="flex items-center justify-center vh  overflow-hidden">
            <div className="card w-96 bg-gray-800  shadow-xl">
                <div className="card-body">
                    <h2 className="text-2xl font-semibold text-center text-white">Welcome Back</h2>
                    <p className="text-center text-gray-400 mb-4">Sign in to your account</p>
                    {errorMessage && <p className="text-center text-red-500 mb-4">{errorMessage}</p>}
                    <form onSubmit={handleLogin}>
                        <div className="form-control mb-4">
                            <label className="label">
                                <span className="label-text text-gray-300">Email</span>
                            </label>
                            <input
                                type="email"
                                name="email"
                                defaultValue="rasel@gmail.com"
                                placeholder="Enter your email"
                                className="input input-bordered w-full"
                            />
                        </div>
                        <div className="form-control mb-4">
                            <label className="label">
                                <span className="label-text text-gray-300">Password</span>
                            </label>
                            <input
                                defaultValue="123456"
                                name="password"
                                type="password"
                                placeholder="Enter your password"
                                className="input input-bordered w-full"
                            />
                        </div>
                        <div className="form-control mb-4">
                            <label className="label cursor-pointer justify-start gap-2">
                                <input name="rememberMe" type="checkbox" className="checkbox checkbox-primary"/>
                                <span className="label-text text-gray-300">Remember me</span>
                            </label>
                        </div>
                        <button className="btn btn-primary w-full">Sign In</button>
                    </form>
                    <p className="text-center text-gray-500 mt-4">
                        <Link to="/login/authenticator" className="text-primary font-medium">
                            Login with authenticator app
                        </Link>
                    </p>
                    <p className="text-center text-gray-500 mt-4">
                        Don’t have an account?{" "}
                        <Link to="/registration" className="text-primary font-medium">
                            Sign Up
                        </Link>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default LoginForm;
