import React from 'react';
import {Link, useLocation} from 'react-router-dom';
import {FaCog, FaHome, FaUser} from 'react-icons/fa';
import {TbDashboard} from "react-icons/tb";
import {RiGoogleFill} from "react-icons/ri";
import {TfiMicrosoft} from "react-icons/tfi";

const sidebarItems = [
    {name: 'Home', route: '/', icon: <FaHome/>},
    {name: 'Dashboard', route: '/#', icon: <TbDashboard/>},
    {name: 'Settings', route: '/settings', icon: <FaCog/>},
    {name: 'Profile', route: '/login', icon: <FaUser/>},
    {name: 'Authenticator Apps', route: '/authenticator-apps', icon: <RiGoogleFill/>}, // Google Authenticator

];

const Sidebar = () => {
    const location = useLocation();  // Get the current location (route)

    return (
        <div className="sidebar h-screen bg-gray-800 text-white flex flex-col">
            <div className="flex justify-center items-center h-24 bg-gray-900">
                <img
                    src="https://via.placeholder.com/100"  // Replace with your dashboard image
                    alt="Dashboard"
                    className="rounded-full w-20 h-20"
                />
            </div>

            <nav className="flex-1 px-4 py-6">
                <ul className="space-y-4">
                    {sidebarItems.map((item, index) => (
                        <li key={index}>
                            <Link
                                to={item.route}
                                className={`flex items-center text-lg 
                                    ${location.pathname === item.route ? 'text-yellow-500' : 'text-white'} 
                                    hover:text-yellow-500`}
                            >
                                {item.icon}
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
