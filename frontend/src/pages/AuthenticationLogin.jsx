import React, {useRef, useState} from "react";
import {useMutation} from "@tanstack/react-query";
import {loginWithAuthenticator} from "../services/authSerivce.js";
import useAuthStore from "../store/authState.js";
import {Link, useNavigate} from "react-router-dom";
import {toast} from "react-toastify";

const AuthenticationLogin = ({userId}) => {
    const {setAuth} = useAuthStore();
    const [otp, setOtp] = useState(["", "", "", "", "", ""]);
    const inputRefs = useRef([]);

    const navigate = useNavigate()

    const mutation = useMutation({
        mutationFn: loginWithAuthenticator,
        onSuccess: (data, variables) => {
            setAuth(data?.data);
            localStorage.setItem("auth-remember-me", true);
            localStorage.setItem("token", data?.data?.token);
            toast.success("Great! You have successfully logged in by otp code!");
            navigate('/account')
        },
        onError: (error) => {
            toast.error("Oops! Authentication failed.");
            console.error("Authentication failed:", error);
        }
    })

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
        mutation.mutate({otpCode, userId});
    }

    const errorMessage = mutation?.error?.response?.data?.message

    return (
        <div className="vh  mt-10 md:mt-0 overflow-hidden">
            <div
                className="relative top-0 md:top-[20vh] mx-auto  card  max-w-md  bg-transparent md:bg-gray-800  shadow-none md:shadow-xl">
                <div className="card-body">
                    <h2 className="text-2xl font-semibold text-center text-white">Authenticator</h2>
                    <p className="text-center text-gray-400 mb-4">Enter the OTP from your authenticator app.</p>
                    {errorMessage && <p className="text-center text-red-500 mb-4">{errorMessage}</p>}

                    <form onSubmit={handleLogin} className="form-phone">
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

                    <div>

                        <p className="text-center text-gray-600 mt-4 text-sm">
                            Lost your phone?{" "}
                            <Link to="/account-recovery" className="text-primary font-semibold hover:underline">
                                Recover using a backup code
                            </Link>
                        </p>
                        
                        <p className="text-center text-gray-500 mt-4">
                            <a href="/login" className="text-primary font-medium">
                                Back
                            </a>
                        </p>

                    </div>
                </div>
            </div>
        </div>
    );
};

export default AuthenticationLogin;
