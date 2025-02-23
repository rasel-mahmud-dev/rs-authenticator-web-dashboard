import React, {useState} from "react";
import {Link, useNavigate, useSearchParams} from "react-router-dom";
import {useMutation} from "@tanstack/react-query";
import {login} from "../services/authSerivce.js";
import useAuthStore from "../store/authState.js";
import AuthenticationLogin from "../pages/AuthenticationLogin.jsx";
import {toast} from "react-toastify";

const LoginForm = () => {
    const [query] = useSearchParams()
    let email = query.get("email") || "test@gmail.com"
    const {setAuth} = useAuthStore()

    const [enabled2Fa, setEnabled2Fa] = useState("") // store userid

    const navigate = useNavigate()

    const mutation = useMutation({
        mutationFn: login,
        onSuccess: (data, variables) => {
            const responseData = data?.data

            if (!responseData) return;
            if (responseData?.enabled2Fa) {
                setEnabled2Fa(responseData?.userId)
                return;
            }

            setAuth(responseData)
            localStorage.setItem("auth-remember-me", variables.rememberMe);
            if (variables.rememberMe) {
                localStorage.setItem("token", data?.data?.token);
            } else {
                sessionStorage.setItem("token", data?.data?.token);
            }
            toast.success("Great! You have successfully logged in!");
            navigate('/account')

        },
        onError: (error) => {
            console.error("Oops! Authentication failed");
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
        <div className="flex items-center justify-center relative vh overflow-hidden">
            {enabled2Fa ? <AuthenticationLogin userId={enabled2Fa} /> : (
                <div
                    className="card w-96 bg-transparent absolute left-1/2 -translate-x-1/2 top-[10vh] md:bg-gray-800  shadow-none md:shadow-xl">
                    <div className="card-body">
                        <h2 className="text-2xl font-semibold text-center text-white">Welcome Back</h2>
                        <p className="text-center text-gray-400 mb-4">Sign in to your account</p>
                        {errorMessage && <p className="text-center text-red-500 mb-4">{errorMessage}</p>}
                        <form onSubmit={handleLogin} className="form-phone">
                            <div className="form-control mb-4">
                                <label className="label">
                                    <span className="label-text text-gray-300">Email</span>
                                </label>
                                <input
                                    type="email"
                                    name="email"
                                    defaultValue={email}
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
                            <button type="submit" className="btn btn-primary w-full">Sign In</button>
                        </form>

                        <p className="text-center text-gray-500 mt-4">
                            Don’t have an account?{" "}
                            <Link to="/registration" className="text-primary font-medium">
                                Sign Up
                            </Link>
                        </p>

                        <p className="text-center text-gray-500 mt-4">
                            <a href="/account-recovery" className="text-primary font-medium">
                                Recovery with backup code
                            </a>
                        </p>

                    </div>
                </div>
            )}
        </div>
    );
};

export default LoginForm;
