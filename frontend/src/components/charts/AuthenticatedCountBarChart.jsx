
import React from 'react';
import {
    BarChart,
    Bar,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer,
} from 'recharts';

// Example data for logins in the last month
const data = [
    { date: '2024-12-01', googleAuthenticator: 20, password: 50, google: 30 },
    { date: '2024-12-02', googleAuthenticator: 25, password: 60, google: 35 },
    { date: '2024-12-03', googleAuthenticator: 15, password: 40, google: 25 },
    { date: '2024-12-04', googleAuthenticator: 30, password: 70, google: 40 },
    { date: '2024-12-05', googleAuthenticator: 20, password: 60, google: 30 },
    { date: '2024-12-06', googleAuthenticator: 18, password: 50, google: 25 },
    { date: '2024-12-07', googleAuthenticator: 22, password: 55, google: 28 },
    // Add more data for the last month as required
];

const AuthenticatedCountBarChart = () => {
    return (
        <ResponsiveContainer width="100%" height={400} className="chart-bg py-4">
            <BarChart
                data={data}
                margin={{
                    top: 20,
                    right: 30,
                    left: 0,
                    bottom: 5,
                }}
            >
                <defs>
                    {/* Gradient for Google Authenticator */}
                    <linearGradient id="googleAuthGradient" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8} />
                        <stop offset="95%" stopColor="#8884d8" stopOpacity={0.2} />
                    </linearGradient>

                    {/* Gradient for Password Login */}
                    <linearGradient id="passwordGradient" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8} />
                        <stop offset="95%" stopColor="#82ca9d" stopOpacity={0.2} />
                    </linearGradient>

                    {/* Gradient for Google Login */}
                    <linearGradient id="googleGradient" x1="0" y1="0" x2="0" y2="1">
                        <stop offset="5%" stopColor="#ffc658" stopOpacity={0.8} />
                        <stop offset="95%" stopColor="#ffc658" stopOpacity={0.2} />
                    </linearGradient>
                </defs>

                <XAxis dataKey="date" />
                <YAxis />
                <Tooltip />
                <Legend />

                {/* Bars with gradient backgrounds */}
                <Bar
                    dataKey="googleAuthenticator"
                    fill="url(#googleAuthGradient)"
                    name="Google Authenticator"
                />
                <Bar
                    dataKey="password"
                    fill="url(#passwordGradient)"
                    name="Login with Password"
                />
                <Bar
                    dataKey="google"
                    fill="url(#googleGradient)"
                    name="Login with Google"
                />
            </BarChart>
        </ResponsiveContainer>
    );
};

export default AuthenticatedCountBarChart;
