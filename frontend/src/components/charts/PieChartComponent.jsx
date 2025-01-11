import React from 'react';
import {buildStyles, CircularProgressbar} from 'react-circular-progressbar';
import 'react-circular-progressbar/dist/styles.css';

const PieChartComponent = ({val, label}) => {
    return (
        <div className="flex items-center gap-x-4 w-full px-4 py-2 chart-bg ">
            <div style={{width: 90}}>
                <CircularProgressbar
                    value={val}
                    text={`${val}%`}
                    styles={buildStyles({
                        textSize: '16px',
                        pathTransitionDuration: 0.5,
                        pathColor: `#4e62e6`,
                        textColor: '#fff',
                        trailColor: 'rgba(78,98,230,0.18)',
                        backgroundColor: '#4e62e6',
                    })}
                />
            </div>
            <div className="pr-10">
                <h4 className="font-medium text-white text-base">$340/$450</h4>
                <p className="font-normal text-gray-500 text-sm">{label}</p>
            </div>
        </div>
    );
};

export default PieChartComponent;
