// src/components/Dashboard.js

import React from 'react';
import '../styles/Dashboard.css';

const Dashboard = () => {
    return (
        <div className="dashboard-container">
            <h1>Welcome to Your Dashboard</h1>
            <p>You have successfully logged in!</p>
            <div className="dashboard-content">
                <p>This is your personalized dashboard where you can manage your account, view data, and explore features.</p>
                <button className="dashboard-button">Explore Features</button>
            </div>
        </div>
    );
};

export default Dashboard;
