import React, {useEffect, useState} from 'react';
import {Link, useLocation} from 'react-router-dom';
import {FaCog, FaHome, FaUser} from 'react-icons/fa';
import {TbDashboard} from "react-icons/tb";
import {RiCollapseDiagonalFill, RiGoogleFill} from "react-icons/ri";

const sidebarItems = [
    {name: 'Home', route: '/', icon: <FaHome/>},
    {name: 'Dashboard', route: '/account', icon: <TbDashboard/>},
    {name: 'Settings', route: '/account/settings', icon: <FaCog/>},
    {name: 'Profile', route: '/login', icon: <FaUser/>},
    {name: 'Authenticator Apps', route: '/account/authenticator-apps', icon: <RiGoogleFill/>}, // Google Authenticator
];

const Sidebar = () => {
    const location = useLocation();
    const [isOpen, setOpen] = useState(true)

    function handleResize() {
        if (window.innerWidth > 1025) {
            setOpen(true)
        } else {
            setOpen(false)
        }
    }

    useEffect(() => {
        window.addEventListener("resize", handleResize)
        return () => window.removeEventListener("resize", handleResize)
    }, []);

    return (
        <div className={`${isOpen ? "" : "compact-sidebar"}`}>
            <div className="sidebar-w-mask"></div>
            <div className={`sidebar vh bg-gray-900 text-white flex flex-col  `}>
                <div className="flex justify-center items-center h-24 bg-gray-900">
                    <img
                        src="https://via.placeholder.com/100"
                        alt="Dashboard"
                        className="rounded-full w-20 h-20"
                    />
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
        </div>
    );
};

export default Sidebar;
