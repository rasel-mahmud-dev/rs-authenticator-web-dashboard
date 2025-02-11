import React, {useEffect, useState} from 'react';
import {Link, useLocation} from 'react-router-dom';
import {FaCog, FaHome, FaUser} from 'react-icons/fa';
import {TbDashboard} from "react-icons/tb";
import {RiGoogleFill} from "react-icons/ri";
import useAuthStore from "../store/authState.js";

const sidebarItems = [
    {name: 'Home', route: '/', icon: <FaHome/>},
    {name: 'Dashboard', route: '/account', icon: <TbDashboard/>},
    {name: 'Settings', route: '/account/settings', icon: <FaCog/>},
    {name: 'Profile', route: '/account/profile', icon: <FaUser/>},
    {name: 'Authenticator', route: '/account/authenticator-apps', icon: <RiGoogleFill/>}, // Google Authenticator
];

const Sidebar = () => {
    const location = useLocation();
    const {user} = useAuthStore()

    return (
            <div className={`sidebar vh bg-gray-900 text-white flex flex-col  `}>
                <div className="flex flex-col justify-center items-center  my-10 bg-gray-900">
                    <div className="avatar">
                        <div
                            className="w-8 sm:w-12 lg:w-28 rounded-full ">
                            <img src={user?.avatar}/>
                        </div>
                    </div>

                    <div className="text-center mt-4">
                        <h4 className="text-gray-100 font-bold"><span>{user?.username}</span></h4>
                        <h4 className="text-gray-400 font-normal "><span>{user?.email}</span></h4>
                    </div>
                </div>

                <nav className="flex-1 px-4 py-6">
                    <ul className="space-y-4">
                        {sidebarItems.map((item, index) => (
                            <li className="sidebar-item" key={index}>
                                <Link
                                    to={item.route}
                                    className={` flex items-center text-lg 
                                    ${location.pathname === item.route ? 'text-yellow-500' : 'text-white'} 
                                    hover:text-yellow-500`}
                                >
                                    <div className="sidebar-icon">{item.icon}</div>
                                    <span className="ml-3">{item.name}</span>
                                </Link>
                            </li>
                        ))}
                    </ul>
                </nav>
        </div>
    );
};

export default Sidebar;
