import React, {useRef, useState} from "react";
import {useMutation} from "@tanstack/react-query";
import {loginWithAuthenticator} from "../services/authSerivce.js";
import useAuthStore from "../store/authState.js";

const TwoFactorAuthentication = () => {
    const {setAuth} = useAuthStore();
    const [otp, setOtp] = useState(["", "", "", "", "", ""]);
    const inputRefs = useRef([]);

    const mutation = useMutation({
        mutationFn: loginWithAuthenticator,
        onSuccess: (data, variables) => {
            setAuth(data?.data);
            localStorage.setItem("auth-remember-me", variables.rememberMe);
            if (variables.rememberMe) {
                localStorage.setItem("token", data?.data?.token);
            } else {
                sessionStorage.setItem("token", data?.data?.token);
            }
        },
        onError: (error) => {
            console.error("Login failed:", error);
        },
    });

    function handleOtpChange(value, index) {
        const newOtp = [...otp];
        newOtp[index] = value;
        setOtp(newOtp);

        if (value && index < otp.length - 1) {
            inputRefs.current[index + 1]?.focus();
        }
    }

    function handleKeyDown(e, index) {
        if (e.key === "ArrowRight" && index < otp.length - 1) {
            inputRefs.current[index + 1]?.focus();
        } else if (e.key === "ArrowLeft" && index > 0) {
            inputRefs.current[index - 1]?.focus();
        } else if (e.key === "Backspace" && !otp[index] && index > 0) {
            inputRefs.current[index - 1]?.focus();
        } else if (e.key === "Enter" && index === otp.length - 1 && otp.every((char) => char)) {
            handleLogin(e);
        }
    }

    function handleLogin(e) {
        e.preventDefault();
        const otpCode = otp.join("");
        mutation.mutate({otpCode});
    }

    return (
        <div className="flex items-center justify-center h-screen overflow-hidden">
            <div className="card  bg-gray-800 shadow-xl">
                <div className="card-body">
                    <h2 className="text-2xl font-semibold text-center text-white">Authenticator</h2>
                    <p className="text-center text-gray-400 mb-4">Enter the OTP sent to your email</p>
                    <form onSubmit={handleLogin}>
                        <div className="flex justify-center mb-4 space-x-2">
                            {otp.map((_, index) => (
                                <input
                                    key={index}
                                    type="text"
                                    maxLength={1}
                                    value={otp[index]}
                                    onChange={(e) => handleOtpChange(e.target.value, index)}
                                    onKeyDown={(e) => handleKeyDown(e, index)}
                                    ref={(el) => (inputRefs.current[index] = el)}
                                    className="input input-bordered w-12 h-12 text-center text-xl"
                                />
                            ))}
                        </div>
                        <button className="btn btn-primary w-full" type="submit">Verify</button>
                    </form>
                    <p className="text-center text-gray-500 mt-4">
                        <a href="#" className="text-primary font-medium">
                            Login with password
                        </a>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default TwoFactorAuthentication;
