// src/components/Sidebar.js
import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import '../styles/Sidebar.css'; // Make sure this file exists for styling

const Sidebar = () => {
    const navigate = useNavigate();
    const location = useLocation();

    const handleNavigation = (path) => {
        if (location.pathname !== path) {
            navigate(path);
        }
    };

    return (
        <div className="sidebar">
            <h2>Dashboard</h2>
            <ul>
                <li
                    className={location.pathname === '/userdashboard/tables' ? 'active' : ''}
                    onClick={() => handleNavigation('/userdashboard/tables')}
                >
                    Tables
                </li>
                <li
                    className={location.pathname === '/userdashboard/menus' ? 'active' : ''}
                    onClick={() => handleNavigation('/userdashboard/menus')}
                >
                    Menus
                </li>
                <li
                    className={location.pathname === '/userdashboard/orders' ? 'active' : ''}
                    onClick={() => handleNavigation('/userdashboard/orders')}
                >
                    Your Order
                </li>
                <li
                    className={location.pathname === '/userdashboard/orderview' ? 'active' : ''}
                    onClick={() => handleNavigation('/userdashboard/orderview')}
                >
                    Placed Orders
                </li>
                
            </ul>
        </div>
    );
};

export default Sidebar;
