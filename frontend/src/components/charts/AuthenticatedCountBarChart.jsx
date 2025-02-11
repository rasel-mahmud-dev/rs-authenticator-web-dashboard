import React from 'react';
import {Bar, BarChart, Legend, ResponsiveContainer, Tooltip,  YAxis,} from 'recharts';
import {useQuery} from "@tanstack/react-query";
import {authenticatorSlats} from "../../services/statsService.js";
import {toolTipContentStyle, XAxisUtil} from "./ChartUtils.jsx";


const AuthenticatedCountBarChart = () => {

    const query = useQuery({
        queryKey: ["authenticationSlats"],
        queryFn: authenticatorSlats
    })

    const data = query?.data?.data ?? []

    return (
        <div className="chart-bg pt-4 pb-14 pr-4">
            <ResponsiveContainer width="100%" height={320}>
                <h3 className="chart-title px-4">Authenticator</h3>
                <BarChart data={data}>
                    <defs>
                        <linearGradient id="googleAuthGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8}/>
                            <stop offset="95%" stopColor="#8884d8" stopOpacity={0.2}/>
                        </linearGradient>

                        <linearGradient id="passwordGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8}/>
                            <stop offset="95%" stopColor="#82ca9d" stopOpacity={0.2}/>
                        </linearGradient>
                    </defs>

                    {XAxisUtil()}

                    <YAxis className="text-xs  text-white font-medium" />
                    <Tooltip
                        contentStyle={toolTipContentStyle}
                    />
                    <Legend/>

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
