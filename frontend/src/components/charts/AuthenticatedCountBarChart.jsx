import React from 'react';
import {Bar, BarChart, Legend, ResponsiveContainer, Tooltip,  YAxis,} from 'recharts';
import {useQuery} from "@tanstack/react-query";
import {authenticatorSlats} from "../../services/statsService.js";
import {XAxisUtil} from "./ChartUtils.jsx";

const data = [
    {date: '2024-12-01', authenticator: 20, password: 50},
];

const AuthenticatedCountBarChart = () => {

    const query = useQuery({
        queryKey: ["authenticationSlats"],
        queryFn: authenticatorSlats
    })

    const data = query?.data?.data ?? []

    return (
        <div className="chart-bg pt-4 pb-14 pr-4">
            <ResponsiveContainer width="100%" height={350}>
                <h3 className="chart-title px-4">Authenticated</h3>
                <BarChart data={data}>
                    <defs>
                        {/* Gradient for Google Authenticator */}
                        <linearGradient id="googleAuthGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8}/>
                            <stop offset="95%" stopColor="#8884d8" stopOpacity={0.2}/>
                        </linearGradient>

                        {/* Gradient for Password Login */}
                        <linearGradient id="passwordGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8}/>
                            <stop offset="95%" stopColor="#82ca9d" stopOpacity={0.2}/>
                        </linearGradient>

                    </defs>

                    {XAxisUtil()}

                    <YAxis/>
                    <Tooltip/>
                    <Legend/>

                    {/* Bars with gradient backgrounds */}
                    <Bar
                        dataKey="authenticator"
                        fill="url(#googleAuthGradient)"
                        name="Authenticator"
                    />
                    <Bar
                        dataKey="password"
                        fill="url(#passwordGradient)"
                        name="Password"
                    />

                </BarChart>
            </ResponsiveContainer>
        </div>
    );
};

export default AuthenticatedCountBarChart;
