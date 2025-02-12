import React from 'react';
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, YAxis} from 'recharts';
import {useQuery} from '@tanstack/react-query';
import {fetchTrafficStats} from '../../services/statsService';
import {toolTipContentStyle, XAxisUtil} from "./ChartUtils.jsx";

const TrafficLineChart = () => {
    const {data, isLoading, error} = useQuery({
        queryKey: ['trafficStats_count'],
        queryFn: () => fetchTrafficStats(true)
    });

    const transformedData = data ?? [];

    return (
        <div className="chart-bg py-4   pr-4">
            <h3 className="chart-title px-4">API Traffic Count</h3>

            {isLoading && <div className="p-10">
                <p>Loading...</p>
            </div>}

            {isLoading && <div className="p-10">
                <>{error?.message}</>
            </div>}

            <ResponsiveContainer width="100%" height={320}>
                <LineChart data={transformedData}>
                    {XAxisUtil()}
                    <YAxis className="text-xs  text-white font-medium"/>
                    <Tooltip
                        contentStyle={toolTipContentStyle}
                    />

                    <Legend

                        wrapperStyle={{
                            fontSize: '10px',
                            fontWeight: 'normal',
                            lineHeight: '1.5',
                        }}

                    />

                    <CartesianGrid stroke="#ccc" strokeDasharray="5 5" strokeOpacity={0.1}/>

                    <Line
                        stroke="#ff7300" // Change this to any color you want
                        strokeWidth={2}
                        radius={8}
                        key="request_count"
                        type="monotone"
                        dataKey="request_count"
                        name="Api Traffic Count"
                    />

                </LineChart>
            </ResponsiveContainer>
        </div>
    );
};

export default TrafficLineChart;
