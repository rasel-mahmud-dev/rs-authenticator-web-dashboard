import React, {useState} from 'react';
import {toast} from 'react-toastify';
import {useMutation} from '@tanstack/react-query';
import {accountRecoveryWithBackupCode} from '../services/authSerivce.js';
import {useNavigate} from "react-router-dom";
import useAuthStore from "../store/authState.js";

const RecoveryAccount = () => {
    const [backupCode, setBackupCode] = useState('');

    const navigate = useNavigate()
    const {setAuth} = useAuthStore()

    const mutation = useMutation({
        mutationFn: accountRecoveryWithBackupCode,
        
        onSuccess: (data, variables) => {
            setAuth(data?.data);
            localStorage.setItem("auth-remember-me", true);
            localStorage.setItem("token", data?.data?.token);
            toast.success('Backup code accepted! You have successfully logged in.');
            navigate('/account')
        },

        onError: () => {
            toast.error('Invalid backup code. Please try again.');
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        mutation.mutate({code: backupCode});
    };

    return (
        <div className="container mx-auto">
            <div className="p-6">
                <h2 className="text-2xl font-semibold text-gray-200">Lost Access to Authenticator?</h2>
                <p className="mt-2 text-gray-400">
                    If you lost access to your authenticator app, you can use one of your backup codes to log in.
                </p>
                <form onSubmit={handleSubmit} className="mt-6 space-y-4">
                    <div>
                        <label htmlFor="backupCode" className="block text-gray-200">Enter Backup Code</label>
                        <input
                            type="text"
                            id="backupCode"
                            value={backupCode}
                            onChange={(e) => setBackupCode(e.target.value)}
                            className="mt-2 p-2 w-full bg-gray-700 text-gray-100 rounded-lg"
                            placeholder="Enter your backup code"
                            required
                        />
                    </div>

                    {mutation.isError && (
                        <div className="text-red-500 text-sm mt-2">Invalid backup code. Please try again.</div>
                    )}

                    <button
                        type="submit"
                        className="btn btn-primary text-white px-4 py-2 rounded-lg"
                        disabled={mutation.isLoading}
                    >
                        {mutation.isLoading ? 'Submitting...' : 'Submit'}
                    </button>
                </form>
            </div>
        </div>
    );
};

export default RecoveryAccount;
