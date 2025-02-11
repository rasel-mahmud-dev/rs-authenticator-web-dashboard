import dayjs from "dayjs";
import {XAxis} from "recharts";
import React from "react";


export function XAxisUtil() {
    return (
        <XAxis
            dataKey="date"
            className="text-xs  text-white font-medium"
            tickFormatter={(date) => dayjs(date).format("YYYY-MM-DD")} // Format date only
        />
    )
}


export const toolTipContentStyle = {
    backgroundColor: '#131626',
    borderColor: "#293583",
    borderRadius: '5px',
    padding: '10px',
    color: '#fff',
    fontSize: "12px",
    boxShadow: '0px 4px 10px rgba(0, 0, 0, 0.3)'
}