import React from 'react';
import RegistrationAreaChart from "../components/charts/UserRegistrationChart.jsx";
import AuthenticatedCountBarChart from "../components/charts/AuthenticatedCountBarChart.jsx";
import UserLoginSlatsBarChart from "../components/charts/UserLoginSlatsBarChart.jsx";
import PieChartComponent from "../components/charts/PieChartComponent.jsx";

const HomePage = () => {
    return (
        <div className="px-4">
            <div className="grid grid-cols-3  gap-4 py-4">
                <UserLoginSlatsBarChart/>
                <PieChartComponent label="React " val={29}/>
                <PieChartComponent label="React " val={29}/>
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-2 gap-4 ">
                <RegistrationAreaChart/>
                <AuthenticatedCountBarChart/>
            </div>
        </div>
    );
};

export default HomePage;