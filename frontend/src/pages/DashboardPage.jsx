import React from 'react';
import RegistrationAreaChart from "../components/charts/UserRegistrationChart.jsx";
import AuthenticatedCountBarChart from "../components/charts/AuthenticatedCountBarChart.jsx";
import UserLoginSlatsBarChart from "../components/charts/UserLoginSlatsBarChart.jsx";
import PieChartComponent from "../components/charts/PieChartComponent.jsx";
import {useQuery} from "@tanstack/react-query";
import {api} from "../services/api.js";
import TrafficLineChart from "../components/charts/TrafficLineChart.jsx";
import TrafficCountLineChart from "../components/charts/TrafficCountLineChart.jsx";

const DashboardPage = () => {

    const query = useQuery({
        queryKey: ["attempts"],
        queryFn: () => api.get("/api/v1/slats/auth-attempts"),

    })

    const data = query?.data?.data ?? {}

    return (
        <div className="px-4">
            <div className="grid grid-cols-1 md:grid-cols-3  gap-4 py-4">
                <UserLoginSlatsBarChart/>
                <PieChartComponent label="Success" total={data?.total ?? 0} val={data?.success ?? 0}/>
                <PieChartComponent isError={true} label="Failed" total={data?.total ?? 0} val={data?.failed}/>
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-2 gap-4 ">
                <RegistrationAreaChart/>
                <AuthenticatedCountBarChart/>
            </div>

            <div className="grid grid-cols-1  gap-4 pt-4 ">
                <TrafficLineChart/>
            </div>
            <div className="grid grid-cols-1  gap-4 py-4 ">
                <TrafficCountLineChart/>
            </div>

        </div>
    );
};

export default DashboardPage;