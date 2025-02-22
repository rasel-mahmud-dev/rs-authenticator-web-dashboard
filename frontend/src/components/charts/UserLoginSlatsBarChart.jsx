import React from 'react';
import {Bar, BarChart, Cell, ResponsiveContainer, Tooltip, XAxis,} from 'recharts';
import {useQuery} from "@tanstack/react-query";
import {api} from "../../services/api.js";
import {toolTipContentStyle, XAxisUtil} from "./ChartUtils.jsx";
import dayjs from "dayjs";


const UserLoginSlatsBarChart = () => {

    const query = useQuery({
        queryKey: ["attempts_detail"],
        queryFn: () => api.get("/api/v1/slats/auth-attempts?t=detail"),
    });

    const data = query?.data?.data ?? [];


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
                    top: 0,
                    right: 0,
                    left: 0,
                    bottom: -10,
                }}
            >

                <XAxis
                    dataKey="date"
                    className="text-[10px]  text-white"
                    tickFormatter={(date) => dayjs(date).format("DD/MM")}
                />

                <Tooltip
                    cursor={{ radius: [10, 10, 0, 0] }}
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
