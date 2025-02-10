import React from 'react';
import {buildStyles, CircularProgressbar} from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';

const PieChartComponent = ({val= 0, total=0, label, isError = false}) => {
    const percentage = ((val / total) * 100) || 0
    const formattedAmount = `${val}/${total}`; // Format the amount

    const pathColor = isError ? '#e74c3c' : '#4e62e6'; // Red for error, blue otherwise
    const trailColor = isError ? 'rgba(231, 76, 60, 0.18)' : 'rgba(78,98,230,0.18)'; // Light red for error, light blue otherwise

    return (
        <div className="flex items-center gap-x-4 w-full px-4 py-2 chart-bg">
            <div style={{width: 90}}>
                <CircularProgressbar
                    value={percentage} // Set the percentage value for the progress circle
                    text={`${Math.round(percentage)}%`} // Display the rounded percentage
                    styles={buildStyles({
                        textSize: '16px',
                        pathTransitionDuration: 0.5,
                        pathColor, // Use dynamic pathColor based on isError
                        textColor: '#fff',
                        trailColor, // Use dynamic trailColor based on isError
                        backgroundColor: '#4e62e6',
                    })}
                />
            </div>
            <div className="pr-10">
                <h4 className="font-medium text-white text-base">{formattedAmount}</h4>
                <p className="font-normal text-gray-500 text-sm">{label}</p>
            </div>
        </div>
    );
};

export default PieChartComponent;
