import React from 'react';
import {Bar, BarChart, ResponsiveContainer, Tooltip, XAxis, YAxis} from 'recharts';

const data = [
    {name: '/home', count: 843},
    {name: '/documentation', count: 384},
    {name: '/blocks', count: 108},
    {name: '/imprint', count: 46},
    {name: '/cancellation', count: 3},
];

const PageVisitCountSlats = () => {
    return (
        <div className="chart-bg  flex  items-end p-4">
            <ResponsiveContainer width="100%" height={200} className=" ">
                <BarChart data={data} layout="vertical">
                    <XAxis type="number"/>
                    <YAxis dataKey="name" type="category"/>
                    <Tooltip/>
                    <Bar
                        dataKey="count"
                        fill="#8884d8"
                        barSize={15}
                        radius={[0, 10, 10, 0]}
                    />
                </BarChart>
            </ResponsiveContainer>
        </div>
    );
};

export default PageVisitCountSlats;