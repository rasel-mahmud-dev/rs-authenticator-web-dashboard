import React from 'react';
import {Bar, BarChart, Cell, ResponsiveContainer, Tooltip,} from 'recharts';

const data = Array.from({length: 30}, (_, i) => ({
    date: i + 1,
    logins: Math.floor(Math.random() * 100) + 1
}));

const UserLoginSlatsBarChart = () => {
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
                {/*<XAxis dataKey="date"/>*/}
                {/*<YAxis/>*/}

                <Tooltip/>

                <Bar dataKey="logins" radius={[10, 10, 0, 0]} barSize={10}>
                    {data.map((entry, index) => (
                        <Cell
                            key={index}
                            fill={index % 2 === 0 ? "#ff4d4d" : "#4d88ff"}
                        />
                    ))}
                </Bar>
            </BarChart>
        </ResponsiveContainer>
    );
};


export default UserLoginSlatsBarChart;
