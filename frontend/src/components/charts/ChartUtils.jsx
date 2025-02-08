import dayjs from "dayjs";
import {XAxis} from "recharts";
import React from "react";


export function XAxisUtil(){
    return (
        <XAxis
            dataKey="date"
            className="text-xs  text-white font-medium"
            tickFormatter={(date) => dayjs(date).format("YYYY-MM-DD")} // Format date only
        />
    )
}

