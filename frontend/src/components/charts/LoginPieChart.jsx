import React from 'react';
import { Pie, PieChart, ResponsiveContainer } from 'recharts';

const data = [
    { name: 'Group A', value: 400, fill: '#8884d8' },
    { name: 'Group B', value: 300, fill: '#82ca9d' },
    { name: 'Group C', value: 300, fill: '#ff7300' },
    { name: 'Group D', value: 200, fill: '#ff7f50' },
    { name: 'Group E', value: 278, fill: '#00c49f' },
    { name: 'Group F', value: 189, fill: '#ffbb28' },
    { name: 'Group G', value: 189, fill: '#ff8042' },
    { name: 'Group H', value: 189, fill: '#d0ed57' },
];

const LoginPieChart = () => {
    const totalValue = data.reduce((sum, entry) => sum + entry.value, 0);

    return (
        <ResponsiveContainer width="100%" height="100%" className="chart-bg">
            <PieChart width={400} height={400}>
                <Pie
                    dataKey="value"
                    data={data}
                    startAngle={0}
                    cx="50%"
                    cy="50%"
                    outerRadius={80}
                    label={({ name, value }) => {
                        const percentage = ((value / totalValue) * 100).toFixed(2);
                        return `${name}: ${percentage}%`;
                    }}
                    fill={ entry => entry.fill}
                />
            </PieChart>
        </ResponsiveContainer>
    );
};

export default LoginPieChart;
