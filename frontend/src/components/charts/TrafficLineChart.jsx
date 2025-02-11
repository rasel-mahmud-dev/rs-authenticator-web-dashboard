import React from 'react';
import {CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, YAxis} from 'recharts';
import {useQuery} from '@tanstack/react-query';
import {fetchTrafficStats} from '../../services/statsService';
import {toolTipContentStyle, XAxisUtil} from "./ChartUtils.jsx";

const TrafficLineChart = () => {
    const {data, isLoading, error} = useQuery({
        queryKey: ['trafficStats'],
        queryFn: fetchTrafficStats
    });

    const transformedData = transformData(data ?? []);

    const uniqueRoutes = [...new Set(data?.map((d) => d.route_path))];

    function renderRoutePath(route) {
        return route?.replace("/api/v1", "");
    }

    return (
        <div className="chart-bg py-4   pr-4">
            <h3 className="chart-title px-4">API Traffic</h3>

            {isLoading && <div className="p-10">
                <p>Loading...</p>
            </div>}

            {isLoading && <div className="p-10">
                <>{error?.message}</>
            </div>}

            <ResponsiveContainer width="100%" height={400}>
                <LineChart data={transformedData}>
                    {XAxisUtil()}
                    <YAxis className="text-xs  text-white font-medium" />
                    <Tooltip
                        contentStyle={toolTipContentStyle}
                    />


                    <Legend
                        wrapperStyle={{
                            fontSize: '10px',
                            fontWeight: 'normal',
                            lineHeight: '1.5'
                        }}
                        payload={uniqueRoutes.map((route, index) => ({
                            value: renderRoutePath(route),
                            type: 'line',
                            color: ["#8884d8", "#82ca9d", "#ff7300"][index % 3],
                        }))}
                    />

                    <CartesianGrid stroke="#ccc" strokeDasharray="5 5" strokeOpacity={0.1}/>

                    {uniqueRoutes.map((route, index) => (
                        <Line
                            strokeWidth={2}
                            radius={8}
                            key={route}
                            type="monotone"
                            dataKey={route}
                            name={renderRoutePath(route)}
                            stroke={["#8884d8", "#82ca9d", "#ff7300"][index % 3]}
                        />
                    ))}
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
};

// Helper function to transform the raw data
const transformData = (rawData) => {
    const groupedData = {};

    rawData.forEach(({date, route_path, request_count}) => {
        if (!groupedData[date]) {
            groupedData[date] = {date};
        }
        groupedData[date][route_path] = request_count;
    });

    return Object.values(groupedData);
};

export default TrafficLineChart;
