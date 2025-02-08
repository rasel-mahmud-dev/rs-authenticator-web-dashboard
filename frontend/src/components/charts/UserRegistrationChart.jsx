import React from 'react';
import {Area, AreaChart, ResponsiveContainer, Tooltip, XAxis, YAxis,} from 'recharts';
import {useQuery} from "@tanstack/react-query";
import {registrationSlats} from "../../services/statsService.js";
import dayjs from "dayjs";
import {XAxisUtil} from "./ChartUtils.jsx";

const RegistrationAreaChart = () => {

    const query = useQuery({
        queryKey: ["registrationSlats"],
        queryFn: registrationSlats,

    })

    const data = query?.data?.data ?? []


    return (
        <div className="chart-bg pt-4 pb-14 pr-4">
            <ResponsiveContainer width="100%" height={350} className="">
                <h3 className="chart-title px-4">User Onboarding</h3>

                <AreaChart data={data} margin={{top: 0, right: 0, left: 0, bottom: 0}}>
                    <defs>
                        <linearGradient id="registrationsGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="5%" stopColor="#fff000" stopOpacity={0.8}/>
                            <stop offset="95%" stopColor="#1184d8" stopOpacity={0}/>
                        </linearGradient>
                    </defs>

                    {XAxisUtil()}

                    <YAxis className="text-xs  text-white font-medium" />
                    <Tooltip/>
                    <Area
                        type="monotone"
                        dataKey="count"
                        stroke="#8884d8"
                        fill="url(#registrationsGradient)"
                    />
                </AreaChart>
            </ResponsiveContainer>
        </div>
    );
};

export default RegistrationAreaChart;
