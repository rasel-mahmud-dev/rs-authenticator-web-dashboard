import React, {useEffect, useState} from 'react';
import {useMutation, useQuery} from '@tanstack/react-query';
import BackupCodes from "../components/BackupCodes.jsx";
import {useNavigate} from "react-router-dom";
import {api} from "../services/api.js";

const RecoveryCodes = () => {

    const navigate = useNavigate()

    const generateCodes = useMutation({
        mutationKey: ["generate-recovery-code"],
        mutationFn: ({isNew = false}) => api.post(`/api/v1/mfa/generate-recovery-code?isNew=${isNew}`),
    })

    useEffect(() => {
        generateCodes.mutate({isNew: false})
    }, []);

    function handleGenNew(){
        generateCodes.mutate({isNew: true})
    }

    const data = generateCodes?.data?.data?.data

    return (
        <div className="container mx-auto">
            <div className="p-6">
                <h2 className="text-2xl font-semibold text-gray-200">Recovery Codes</h2>
                <p className="mt-2 text-gray-400">
                    You can use each recovery code one time. Save your recovery codes in a safe place.
                    Without these, you may not be able to log into your account if you lose access to your phone or can’t log in using your security method.
                </p>


                <div className="mt-6">
                    <BackupCodes onGenerateNewCode={handleGenNew} recovery_codes={data}/>
                </div>

            </div>
        </div>
    );
};

export default RecoveryCodes;
