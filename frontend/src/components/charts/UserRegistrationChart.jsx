import React from 'react';
import {Area, AreaChart, ResponsiveContainer, Tooltip, XAxis, YAxis,} from 'recharts';

// Example data for the last month's user registration
const data = [
    {date: '2024-12-01', registrations: 10},
    {date: '2024-12-02', registrations: 20},
    {date: '2024-12-03', registrations: 15},
    {date: '2024-12-04', registrations: 25},
    {date: '2024-12-05', registrations: 30},
    {date: '2024-12-06', registrations: 22},
    {date: '2024-12-07', registrations: 18},
    {date: '2024-12-08', registrations: 28},
    {date: '2024-12-09', registrations: 35},
    {date: '2024-12-10', registrations: 40},
    {date: '2024-12-11', registrations: 50},
    {date: '2024-12-12', registrations: 42},
    {date: '2024-12-13', registrations: 30},
    {date: '2024-12-14', registrations: 38},
    {date: '2024-12-15', registrations: 45},
    {date: '2024-12-16', registrations: 55},
    {date: '2024-12-17', registrations: 60},
    {date: '2024-12-18', registrations: 65},
    {date: '2024-12-19', registrations: 70},
    {date: '2024-12-20', registrations: 80},
    {date: '2024-12-21', registrations: 85},
    {date: '2024-12-22', registrations: 75},
    {date: '2024-12-23', registrations: 65},
    {date: '2024-12-24', registrations: 55},
    {date: '2024-12-25', registrations: 50},
    {date: '2024-12-26', registrations: 60},
    {date: '2024-12-27', registrations: 70},
    {date: '2024-12-28', registrations: 80},
    {date: '2024-12-29', registrations: 90},
    {date: '2024-12-30', registrations: 100},
    {date: '2024-12-31', registrations: 110},
];

const RegistrationAreaChart = () => {
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
                    <XAxis dataKey="date"/>
                    <YAxis/>
                    <Tooltip/>
                    <Area
                        type="monotone"
                        dataKey="registrations"
                        stroke="#8884d8"
                        fill="url(#registrationsGradient)"
                    />
                </AreaChart>
            </ResponsiveContainer>
        </div>
    );
};

export default RegistrationAreaChart;
