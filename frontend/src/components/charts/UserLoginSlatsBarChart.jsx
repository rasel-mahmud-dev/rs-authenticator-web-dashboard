import React from 'react';
import {Bar, BarChart, Cell, ResponsiveContainer, Tooltip,} from 'recharts';
import {useQuery} from "@tanstack/react-query";
import {api} from "../../services/api.js";
import {toolTipContentStyle} from "./ChartUtils.jsx";



const UserLoginSlatsBarChart = () => {

    const query = useQuery({
        queryKey: ["attempts_detail"],
        queryFn: () => api.get("/api/v1/slats/auth-attempts?t=detail"),
    });

    const data = query?.data?.data ?? [];

    console.log(data, "dsfsds")

    const chartData = data?.map(entry => ({
        date: entry.date,
        failed: entry.failed,
        success: entry.success
    }));


    return (
        <ResponsiveContainer  className="chart-bg p-4 py-1">
            <BarChart

                data={data}
                margin={{
                    top: 20,
                    right: 30,
                    left: 0,
                    bottom: 5,
                }}
            >

                <Tooltip
                    contentStyle={toolTipContentStyle}
                />
                <Bar dataKey="success" radius={[10, 10, 0, 0]} barSize={10}>
                    {chartData?.map((entry, index) => (
                        <Cell
                            key={`success-${index}`}
                            fill="#4d88ff"
                        />
                    ))}
                </Bar>

                <Bar dataKey="failed" radius={[10, 10, 0, 0]} barSize={10}>
                    {chartData?.map((entry, index) => (
                        <Cell
                            key={`failed-${index}`}
                            fill="#ff4d4d"
                        />
                    ))}
                </Bar>


            </BarChart>
        </ResponsiveContainer>
    );
};


export default UserLoginSlatsBarChart;
