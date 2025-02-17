import React from 'react';
import {Bar, BarChart, Cell, Legend, ResponsiveContainer, Tooltip, YAxis} from 'recharts';
import {useQuery} from '@tanstack/react-query';
import {fetchApiLatencySlats} from '../../services/statsService';
import {COLORS, toolTipContentStyle} from "./ChartUtils.jsx";


const ApiLatencyChart = () => {
    const {data, isLoading, error} = useQuery({
        queryKey: ['api-latencyss'],
        queryFn: fetchApiLatencySlats
    });



    function renderRoutePath(route) {
        return route?.replace("/api/v1", "");
    }

    const CustomTooltip = ({ active, payload, label, ...a }) => {
        if (active && payload && payload.length) {
            const item = payload?.[0]
            return (
                <div className="p-4 flex flex-col gap-2" style={toolTipContentStyle}>
                    <h4 className="text-sm  font-semibold text-white">{`${renderRoutePath(item?.payload?.route_path)}`}</h4>
                    <span className="text-xs   text-white">{`AVG Response time: ${item?.value?.toFixed(2)}ms`}</span>
                    <p className="text-xs   text-white">{`Request Count: ${item?.payload?.request_count}`}</p>
                </div>
            );
        }

        return null;
    };
    return (
        <div className="chart-bg py-4 pr-4">
            <h3 className="chart-title px-4">API Traffic (Latency)</h3>

            {isLoading && <div className="p-10"><p>Loading...</p></div>}
            {error && <div className="p-10"><p>{error.message}</p></div>}


            <ResponsiveContainer width="100%" height={320} className="api-latency-chart">
                <BarChart data={data?.data} >
                    <YAxis className="text-xs text-white font-medium"
                           label={{value: 'Response Time (ms)', angle: -90, position: 'insideLeft'}}/>
                    <Tooltip
                        cursor={{ radius: [28, 28, 0, 0] }}
                        content={<CustomTooltip />}
                        contentStyle={toolTipContentStyle}
                    />

                    <Legend
                        wrapperStyle={{
                            marginTop: "30px",
                            fontSize: '10px',
                            fontWeight: 'normal',
                            lineHeight: '1.5'
                        }}
                        payload={data?.data?.map((item, index) => ({
                            value: renderRoutePath(item.route_path),
                            type: 'line',
                            color: COLORS[index % 3],
                        }))}
                    />

                    <Bar  barSize={20} radius={[8, 8, 0, 0]} dataKey="avg_response_time" name="Response Time (ms)" className="flex justify-center">
                        {data?.data?.map((entry, index) => (
                            <Cell
                                va
                                key={`cell-${index}`}
                                fill={COLORS[index % COLORS.length]}
                            />
                        ))}
                    </Bar>

                </BarChart>
            </ResponsiveContainer>

        </div>
    );
};


export default ApiLatencyChart;
